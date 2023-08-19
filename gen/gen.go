package gen

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jackc/pgx/v5"
	"github.com/michaelzx/pgx-z/gen/gen_tpl"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Gen struct {
	Config
	conn   *pgx.Conn
	Tables []table
}

func New(cfg Config) {
	g := &Gen{}
	g.Config = cfg
	var err error
	g.conn, err = pgx.Connect(context.TODO(), cfg.PgDSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		g.conn.Close(context.TODO())
		fmt.Println("success")
	}()
	g.getTables()
	g.getColumns()
	g.genModels()
	g.genCols()

}

func (g *Gen) getTables() {
	rows, _ := g.conn.Query(context.TODO(), `--
select table_name 
from information_schema.tables 
where table_schema='public'`)
	tableRows, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[table])
	if err != nil {
		panic(err)
	}
	for i, row := range tableRows {
		tableRows[i].NameForGo = strcase.ToCamel(row.NameForDb)
	}
	g.Tables = tableRows
	// fmt.Println("tables", g.Tables)
}

func (g *Gen) getColumns() {
	for tIdx, t := range g.Tables {
		rows, _ := g.conn.Query(context.TODO(), `--sql
SELECT column_name,udt_name,is_nullable::bool
,col_description((table_schema||'.'||table_name)::regclass::oid, ordinal_position) as comment
FROM information_schema.columns
WHERE table_name=$1 order by ordinal_position`, t.NameForDb)
		cols, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[column])
		if err != nil {
			panic(err)
		}
		for colIdx, col := range cols {
			cols[colIdx].NameForGo = strcase.ToCamel(col.NameForDb)
			cols[colIdx].NameForJson = strcase.ToSnake(col.NameForDb)
			goType := pg2go(strings.TrimPrefix(col.TypeForDb, "_"))
			if strings.HasPrefix(col.TypeForDb, "_") {
				goType = "[]" + goType
			}
			if col.IsNullable {
				goType = "*" + goType
			}
			if strings.Contains(goType, "pgtype") {
				g.Tables[tIdx].Imports = append(g.Tables[tIdx].Imports, "github.com/jackc/pgx/v5/pgtype")
			}
			cols[colIdx].TypeForGo = goType
		}
		g.Tables[tIdx].Columns = cols
		// fmt.Println("tables", cols)
	}
}

func (g *Gen) genModels() {
	for _, tbl := range g.Tables {
		fmt.Println("gen", "model."+tbl.NameForGo, "...")
		tmpl := template.Must(
			template.New("test").Parse(gen_tpl.Model),
		)

		tbl.NowTime = time.Now()

		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		err := tmpl.ExecuteTemplate(w, "test", tbl)
		if err != nil {
			panic(err)
		}
		w.Flush()
		code, err := format.Source(b.Bytes())
		if err != nil {
			fmt.Println(b.String())
			panic(err)
		}
		destDir := filepath.Join(g.OutputDir, "model")
		filename := strcase.ToSnake(tbl.NameForDb) + ".go"
		os.MkdirAll(destDir, 0755)
		if err := os.WriteFile(filepath.Join(destDir, filename), code, 0644); err != nil {
			panic(err)
		}
	}
}
func (g *Gen) genCols() {
	for _, tbl := range g.Tables {
		fmt.Println("gen", "col."+tbl.NameForGo, "...")
		tmpl := template.Must(
			template.New("test").Parse(gen_tpl.Col),
		)

		modelImport := filepath.Join(g.PkgName, g.OutputDir, "model")
		modelImport = strings.ReplaceAll(modelImport, "\\", "/")
		tbl.Imports = append(tbl.Imports, modelImport)
		tbl.Imports = append(tbl.Imports, "github.com/michaelzx/pgx-z/pgxz")
		tbl.NowTime = time.Now()

		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		err := tmpl.ExecuteTemplate(w, "test", tbl)
		if err != nil {
			panic(err)
		}
		w.Flush()
		code, err := format.Source(b.Bytes())
		if err != nil {
			fmt.Println(b.String())
			panic(err)
		}
		destDir := filepath.Join(g.OutputDir, "col")
		filename := strcase.ToSnake(tbl.NameForDb) + ".go"
		os.MkdirAll(destDir, 0755)
		if err := os.WriteFile(filepath.Join(destDir, filename), code, 0644); err != nil {
			panic(err)
		}
	}
}

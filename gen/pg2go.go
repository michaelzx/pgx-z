package gen

import (
	"log"
)

func pg2go(udt string) string {
	switch udt {
	case "serial", "serial4", "pg_catalog.serial4":
		return "int32"
	case "bigserial", "serial8", "pg_catalog.serial8":
		return "int64"
	case "smallserial", "serial2", "pg_catalog.serial2":
		return "int16"
	case "integer", "int", "int4", "pg_catalog.int4":
		return "int32"
	case "bigint", "int8", "pg_catalog.int8":
		return "int64"
	case "smallint", "int2", "pg_catalog.int2":
		return "int16"
	case "float", "double precision", "float8", "pg_catalog.float8":
		return "float64"
	case "real", "float4", "pg_catalog.float4":
		return "float32"
	case "numeric", "pg_catalog.numeric", "money":
		// return "pgtype.Numeric"
		return "string"

	case "boolean", "bool", "pg_catalog.bool":
		return "bool"
	case "json":
		return "[]byte"
		// return "pgtype.JSON"
		// return "json.RawMessage"
		// return "pqtype.NullRawMessage"
		// return "any"

	case "jsonb":
		return "[]byte"
		// return "pgtype.JSONB"
		// return "json.RawMessage"
		// return "pqtype.NullRawMessage"
		// return "any"
	case "bytea", "blob", "pg_catalog.bytea":
		return "[]byte"
	case "date":
		return "pgtype.Date"
		// return "time.Time"
	case "pg_catalog.time":
		return "pgtype.Time"
		// return "time.Time"
	case "pg_catalog.timetz":
		return "time.Time"
	case "pg_catalog.timestamp":
		return "pgtype.Timestamp"
		// return "time.Time"
	case "pg_catalog.timestamptz", "timestamptz":
		return "pgtype.Timestamptz"
		// return "time.Time"
	case "text", "varchar", "pg_catalog.varchar", "pg_catalog.bpchar", "string", "citext":
		return "string"
	case "uuid":
		return "pgtype.UUID"
		// return "uuid.UUID"
	}
	log.Fatalf("unknown PostgreSQL type: %s\n", udt)
	return "any"
}

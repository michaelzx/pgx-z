package gen_tpl

import _ "embed"

//go:embed model.gotpl
var Model string

//go:embed col.gotpl
var Col string

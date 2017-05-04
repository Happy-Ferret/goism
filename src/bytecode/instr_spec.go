package bytecode

import (
	"bytecode/ir"
	"emacs"
)

type instrSpec struct {
	argc   int
	output bool
	fn     emacs.Symbol
}

var instrSpecs = [...]instrSpec{
	ir.OpStackSet: {argc: 1},
	ir.OpReturn:   {argc: 1},

	ir.OpNumAdd: {argc: 2, output: true, fn: "+"},
	ir.OpNumGt:  {argc: 2, output: true, fn: ">"},
	ir.OpNumEq:  {argc: 2, output: true, fn: "="},
}

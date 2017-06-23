package compiler

import (
	"elapc/instr"
	"exn"
	"sexp"
)

func compileStmt(cl *Compiler, form sexp.Form) {
	switch form := form.(type) {
	case *sexp.Return:
		compileReturn(cl, form)
	case *sexp.If:
		compileIf(cl, form)
	case *sexp.Block:
		compileBlock(cl, form)
	case *sexp.FormList:
		compileStmtList(cl, form.Forms)
	case *sexp.Bind:
		compileBind(cl, form)
	case *sexp.Rebind:
		compileRebind(cl, form.Name, form.Expr)
	case *sexp.VarUpdate:
		compileVarUpdate(cl, form.Name, form.Expr)
	case *sexp.ExprStmt:
		compileExprStmt(cl, form)
	case *sexp.Repeat:
		compileRepeat(cl, form)
	case *sexp.While:
		compileWhile(cl, form)
	case *sexp.ArrayUpdate:
		compileArrayUpdate(cl, form)
	case *sexp.StructUpdate:
		compileStructUpdate(cl, form)
	case *sexp.Goto:
		compileGoto(cl, form)
	case *sexp.Label:
		compileLabel(cl, form)

	case *sexp.Let:
		compileLetStmt(cl, form)

	default:
		panic(exn.Logic("unexpected stmt: %#v", form))
	}
}

func compileExpr(cl *Compiler, form sexp.Form) {
	switch form := form.(type) {
	case sexp.Int:
		emit(cl, instr.ConstRef(cl.cvec.InsertInt(int64(form))))
	case sexp.Float:
		emit(cl, instr.ConstRef(cl.cvec.InsertFloat(float64(form))))
	case sexp.Str:
		emit(cl, instr.ConstRef(cl.cvec.InsertString(string(form))))
	case sexp.Symbol:
		emit(cl, instr.ConstRef(cl.cvec.InsertSym(form.Val)))
	case sexp.Bool:
		compileBool(cl, form)
	case sexp.Var:
		compileVar(cl, form)

	case *sexp.SparseArrayLit:
		compileSparseArrayLit(cl, form)

	case *sexp.ArrayIndex:
		compileArrayIndex(cl, form)

	case *sexp.LispCall:
		compileCall(cl, form.Fn.Sym, form.Args)
	case *sexp.Call:
		compileCall(cl, form.Fn.Name, form.Args)
	case *sexp.InstrCall:
		compileInstrCall(cl, form)

	case *sexp.StructLit:
		compileStructLit(cl, form)
	case *sexp.StructIndex:
		compileStructIndex(cl, form)

	case *sexp.Let:
		compileLetExpr(cl, form)

	case *sexp.And:
		compileAnd(cl, form)
	case *sexp.Or:
		compileOr(cl, form)

	case nil: // #REFS: 65
		emit(cl, instr.ConstRef(cl.cvec.InsertSym("nil")))

	default:
		panic(exn.Logic("unexpected expr: %#v", form))
	}
}

package opt

import (
	"sexp"
)

// RemoveDeadCode removes unreachable statements and expressions
// from given sexp form.
func RemoveDeadCode(form sexp.Form) sexp.Form {
	cr := codeRemover{}
	return cr.rewrite(form)
}

type codeRemover struct{}

func (cr codeRemover) rewrite(form sexp.Form) sexp.Form {
	return sexp.Rewrite(form, cr.walkForm)
}

func (cr codeRemover) walkForm(form sexp.Form) sexp.Form {
	switch form := form.(type) {
	case *sexp.If:
		if b, ok := form.Cond.(sexp.Bool); ok && !bool(b) {
			return sexp.EmptyStmt
		}
		form.Then.Forms = cr.walkBody(form.Then.Forms)
		form.Else = cr.rewrite(form.Else)
		return form

	case *sexp.Block:
		form.Forms = cr.walkBody(form.Forms)
		return form

	case *sexp.FormList:
		form.Forms = cr.walkBody(form.Forms)
		return form

	case *sexp.While:
		if b, ok := form.Cond.(sexp.Bool); ok && !bool(b) {
			return sexp.EmptyStmt
		}
		form.Body.Forms = cr.walkBody(form.Body.Forms)
		return form
	}

	return nil
}

func (cr codeRemover) walkBody(forms []sexp.Form) []sexp.Form {
	for i, form := range forms {
		forms[i] = cr.rewrite(form)
		if sexp.IsReturning(form) {
			return forms[:i+1]
		}
	}
	return forms
}

package sexp

func (atom Bool) Copy() Form   { return Bool(atom) }
func (atom Int) Copy() Form    { return Int(atom) }
func (atom Float) Copy() Form  { return Float(atom) }
func (atom Str) Copy() Form    { return Str(atom) }
func (atom Symbol) Copy() Form { return Symbol(atom) }
func (atom Var) Copy() Form {
	return Var{Name: atom.Name, Typ: atom.Typ}
}
func (atom Local) Copy() Form {
	return Local{Name: atom.Name, Typ: atom.Typ}
}

func (lit *ArrayLit) Copy() Form {
	return &ArrayLit{Vals: CopyList(lit.Vals), Typ: lit.Typ}
}
func (lit *SparseArrayLit) Copy() Form {
	return &SparseArrayLit{
		Ctor:    lit.Ctor.Copy(),
		Vals:    CopyList(lit.Vals),
		Indexes: append([]int(nil), lit.Indexes...),
		Typ:     lit.Typ,
	}
}
func (lit *SliceLit) Copy() Form {
	return &SliceLit{Vals: CopyList(lit.Vals), Typ: lit.Typ}
}
func (lit *StructLit) Copy() Form {
	return &StructLit{Vals: CopyList(lit.Vals), Typ: lit.Typ}
}

func (form *ArrayUpdate) Copy() Form {
	return &ArrayUpdate{
		Array: form.Array.Copy(),
		Index: form.Index.Copy(),
		Expr:  form.Expr.Copy(),
	}
}
func (form *SliceUpdate) Copy() Form {
	return &SliceUpdate{
		Slice: form.Slice.Copy(),
		Index: form.Index.Copy(),
		Expr:  form.Expr.Copy(),
	}
}
func (form *StructUpdate) Copy() Form {
	return &StructUpdate{
		Struct: form.Struct.Copy(),
		Index:  form.Index,
		Expr:   form.Expr.Copy(),
		Typ:    form.Typ,
	}
}
func (form *Bind) Copy() Form {
	return &Bind{Name: form.Name, Init: form.Init.Copy()}
}
func (form *Rebind) Copy() Form {
	return &Rebind{Name: form.Name, Expr: form.Expr.Copy()}
}
func (form *VarUpdate) Copy() Form {
	return &VarUpdate{Expr: form.Expr.Copy()}
}
func (form FormList) Copy() Form {
	return FormList(CopyList(form))
}
func (form Block) Copy() Form {
	return Block(CopyList(form))
}
func (form *If) Copy() Form {
	return &If{
		Cond: form.Cond.Copy(),
		Then: form.Then.Copy().(Block),
		Else: form.Else.Copy(),
	}
}
func (form *Switch) Copy() Form {
	return &Switch{
		Expr:       form.Expr.Copy(),
		SwitchBody: copySwitchBody(form.SwitchBody),
	}
}
func (form *SwitchTrue) Copy() Form {
	return &SwitchTrue{SwitchBody: copySwitchBody(form.SwitchBody)}
}
func (form *Return) Copy() Form {
	return &Return{Results: CopyList(form.Results)}
}
func (form *ExprStmt) Copy() Form {
	return &ExprStmt{Expr: form.Expr.Copy()}
}
func (form *Goto) Copy() Form  { return &Goto{LabelName: form.LabelName} }
func (form *Label) Copy() Form { return &Label{Name: form.Name} }

func (form *Repeat) Copy() Form {
	return &Repeat{
		N:    form.N,
		Body: form.Body.Copy().(Block),
	}
}
func (form *DoTimes) Copy() Form {
	return &DoTimes{
		N:    form.N.Copy(),
		Iter: form.Iter,
		Step: form.Step.Copy(),
		Body: form.Body.Copy().(Block),
	}
}
func (form *Loop) Copy() Form {
	return &Loop{
		Init: form.Init.Copy(),
		Post: form.Post.Copy(),
		Body: form.Body.Copy().(Block),
	}
}
func (form *While) Copy() Form {
	return &While{
		Init: form.Init.Copy(),
		Cond: form.Cond.Copy(),
		Post: form.Post.Copy(),
		Body: form.Body.Copy().(Block),
	}
}

func (form *ArrayIndex) Copy() Form {
	return &ArrayIndex{
		Array: form.Array.Copy(),
		Index: form.Index.Copy(),
	}
}
func (form *SliceIndex) Copy() Form {
	return &SliceIndex{
		Slice: form.Slice.Copy(),
		Index: form.Index.Copy(),
	}
}
func (form *StructIndex) Copy() Form {
	return &StructIndex{
		Struct: form.Struct.Copy(),
		Index:  form.Index,
		Typ:    form.Typ,
	}
}

func (form *ArraySlice) Copy() Form {
	return &ArraySlice{
		Array: form.Array.Copy(),
		Typ:   form.Typ,
		Span:  copySpan(form.Span),
	}
}
func (form *SliceSlice) Copy() Form {
	return &SliceSlice{
		Slice: form.Slice.Copy(),
		Span:  copySpan(form.Span),
	}
}

func (form *TypeAssert) Copy() Form {
	return &TypeAssert{Expr: form.Expr.Copy(), Typ: form.Typ}
}

func (call *Call) Copy() Form {
	return &Call{Fn: call.Fn, Args: CopyList(call.Args)}
}
func (call *LispCall) Copy() Form {
	return &LispCall{Fn: call.Fn, Args: CopyList(call.Args)}
}
func (call *LambdaCall) Copy() Form {
	return &LambdaCall{
		Args: copyBindList(call.Args),
		Body: call.Body.Copy().(Block),
		Typ:  call.Typ,
	}
}
func (call *DynCall) Copy() Form {
	return &DynCall{
		Callable: call.Callable.Copy(),
		Args:     CopyList(call.Args),
		Typ:      call.Typ,
	}
}

func (form *Let) Copy() Form {
	binds := copyBindList(form.Bindings)
	if form.Expr != nil {
		return &Let{Bindings: binds, Expr: form.Expr.Copy()}
	}
	return &Let{Bindings: binds, Stmt: form.Stmt.Copy()}
}
func (form *TypeCast) Copy() Form {
	return &TypeCast{Form: form.Form.Copy(), Typ: form.Typ}
}

func (form *And) Copy() Form {
	return &And{X: form.X.Copy(), Y: form.Y.Copy()}
}
func (form *Or) Copy() Form {
	return &Or{X: form.X.Copy(), Y: form.Y.Copy()}
}

func (form *emptyForm) Copy() Form { return form }

func CopyList(forms []Form) []Form {
	if forms == nil {
		return nil
	}
	res := make([]Form, len(forms))
	for i, form := range forms {
		res[i] = form.Copy()
	}
	return res
}

func copySpan(span Span) Span {
	return Span{
		Low:  span.Low.Copy(),
		High: span.High.Copy(),
	}
}

func copySwitchBody(b SwitchBody) SwitchBody {
	var clauses []CaseClause
	if b.Clauses != nil {
		clauses = make([]CaseClause, len(b.Clauses))
		for i, cc := range b.Clauses {
			clauses[i] = CaseClause{
				Expr: cc.Expr.Copy(),
				Body: cc.Body.Copy().(Block),
			}
		}
	}
	return SwitchBody{
		Clauses:     clauses,
		DefaultBody: b.DefaultBody.Copy().(Block),
	}
}

func copyCaseClauseList(clauses []CaseClause) []CaseClause {
	if clauses == nil {
		return nil
	}
	res := make([]CaseClause, len(clauses))
	for i, cc := range clauses {
		res[i] = CaseClause{
			Expr: cc.Expr.Copy(),
			Body: cc.Body.Copy().(Block),
		}
	}
	return res
}

func copyBindList(binds []*Bind) []*Bind {
	res := make([]*Bind, len(binds))
	for i, bind := range binds {
		res[i] = &Bind{
			Name: bind.Name,
			Init: bind.Init.Copy(),
		}
	}
	return res
}

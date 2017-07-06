package rt

import (
	"emacs/lisp"
)

// Panic triggers run-time panic.
func Panic(errorData lisp.Object) {
	{
	} // Prevent inlining (#REFS: 37)
	lisp.Call("signal", lisp.Intern("error"), lisp.Call("list", errorData))
}

// Print prints all arguments;
// formatting of arguments is implementation-specific.
func Print(args lisp.Object) {
	lisp.Call("princ", lisp.MapConcat(lisp.Prin1ToString, args, ""))
}

// Println is like print but prints spaces between arguments
// and a newline at the end.
func Println(args lisp.Object) {
	lisp.Call("princ", lisp.MapConcat(lisp.Prin1ToString, args, " "))
	lisp.Call("terpri") // Better than princ("\n") due to spared "\n"
}

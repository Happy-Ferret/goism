# Translation spec

This document describes design decisions and 
implementation limitations that affect
either **Go spec confirmity** or **Emacs Lisp experience**.

If something breaks compatibility with Go and is not
described here, it is a bug. 
That incompatibility should be fixed or at least 
appended to this document.
There is also `unimplemented.md` 
which supplements `translation_spec.md`.

## Abbreviations and conventions

`Go`->`Emacs Lisp` translated code is called `GE` inside 
this document for brevity.

`Emacs Lisp` here spelled as `Elisp` for the same reason 
even if it is innacurate generally.

* The most significant information presented in this way (UL)

## Implementation details

### (1) Constants

Go constants are inlined at the compilation time.
They do not produce `defvar` or `defconst`.

* You can not use Go constants inside Elisp

### (2) Numeric types

On type system level, everything is the same as in normal Go,
but during execution all types are implemented in terms
of Elisp numbers (int and float). 

All signed types are implemented without special treatment.
This means that `int8`, ..., `int64` are simple Elisp integer.

* You can not rely on `intX` types overflow

Unsigned types are emulated. 
There is no overhead on arithmetics. 
The most significant bits are cleared only when not doing 
so will affect *visible results*.

* `uint64` type behaves like `uint32`
* `float32` type behaves like `float64`

`float64` depends on the Elisp float,
which implemented in terms of C `double`. 

* Most conforming types are: `uint8`/`byte`, `uint16`, `uint32`, `int`, `float64`

If `int16` is boxed into `interface{}`, it can be type-matched
with `int16` only; this also applies to floating point types.

* Type assertions distinguish all numeric types

### (3) Functions

Void-result GE functions return value is unspecified and should not be assigned
inside Elisp. 

### (4) Symbol type

New symbols can be created by `lisp.Intern`.
If argument is constant, symbol literal is inserted inplace
instead of `intern` call.

* `lisp.Symbol` default value is `nil`
* `lisp.Symbol` has method-based API

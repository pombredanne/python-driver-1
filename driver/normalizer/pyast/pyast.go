// Package pyast defines constants from Python 2 and 3 AST.
package pyast

// GENERATED BY python-driver/native/gogen/gogen.py
// DO NOT EDIT

// Python 2+3 AST node types.
// FIXME: ColumnKey? (col_offset in Python)
// This includes all classes extending from _ast.AST for both Python 2 and 3.
// See:
// https://docs.python.org/3.6/library/ast.html#abstract-grammar
// https://docs.python.org/2.7/library/ast.html#abstract-grammar

const (
	Add                   = "Add"
	Alias                 = "alias"
	And                   = "And"
	AnnAssign             = "AnnAssign"
	Arg                   = "arg"
	Arguments             = "arguments"
	Assert                = "Assert"
	Assign                = "Assign"
	AsyncFor              = "AsyncFor"
	AsyncFunctionDef      = "AsyncFunctionDef"
	AsyncWith             = "AsyncWith"
	Attribute             = "Attribute"
	AugAssign             = "AugAssign"
	AugLoad               = "AugLoad"
	AugStore              = "AugStore"
	Await                 = "Await"
	BinOp                 = "BinOp"
	BitAnd                = "BitAnd"
	BitOr                 = "BitOr"
	BitXor                = "BitXor"
	BoolLiteral           = "BoolLiteral"
	BoolOp                = "BoolOp"
	BoolopInternal        = "boolop"
	Break                 = "Break"
	ByteLiteral           = "ByteLiteral"
	Bytes                 = "Bytes"
	Call                  = "Call"
	ClassDef              = "ClassDef"
	Cmpop                 = "cmpop"
	Compare               = "Compare"
	Comprehension         = "comprehension"
	Constant              = "Constant"
	Continue              = "Continue"
	Del                   = "Del"
	Delete                = "Delete"
	Dict                  = "Dict"
	DictComp              = "DictComp"
	Div                   = "Div"
	Ellipsis              = "Ellipsis"
	Eq                    = "Eq"
	ExceptHandler         = "ExceptHandler"
	ExcepthandlerInternal = "excepthandler"
	Exec                  = "Exec"
	Expr                  = "Expr"
	ExprInternal          = "expr"
	Expr_context          = "expr_context"
	Expression            = "Expression"
	ExtSlice              = "ExtSlice"
	FloorDiv              = "FloorDiv"
	For                   = "For"
	FormattedValue        = "FormattedValue"
	FunctionDef           = "FunctionDef"
	GeneratorExp          = "GeneratorExp"
	Global                = "Global"
	Gt                    = "Gt"
	GtE                   = "GtE"
	If                    = "If"
	IfExp                 = "IfExp"
	Import                = "Import"
	ImportFrom            = "ImportFrom"
	In                    = "In"
	Index                 = "Index"
	Interactive           = "Interactive"
	Invert                = "Invert"
	Is                    = "Is"
	IsNot                 = "IsNot"
	JoinedStr             = "JoinedStr"
	Keyword               = "keyword"
	LShift                = "LShift"
	Lambda                = "Lambda"
	List                  = "List"
	ListComp              = "ListComp"
	Load                  = "Load"
	Lt                    = "Lt"
	LtE                   = "LtE"
	MatMult               = "MatMult"
	Mod                   = "Mod"
	ModInternal           = "mod"
	Module                = "Module"
	Mult                  = "Mult"
	Name                  = "Name"
	NameConstant          = "NameConstant"
	NoneLiteral           = "NoneLiteral"
	Nonlocal              = "Nonlocal"
	NoopLine              = "NoopLine"
	Noop_lineInternal     = "noop_line"
	Not                   = "Not"
	NotEq                 = "NotEq"
	NotIn                 = "NotIn"
	Num                   = "Num"
	NumLiteral            = "NumLiteral"
	Operator              = "operator"
	Or                    = "Or"
	Param                 = "Param"
	Pass                  = "Pass"
	Pow                   = "Pow"
	PreviousNoops         = "PreviousNoops"
	Print                 = "Print"
	RShift                = "RShift"
	Raise                 = "Raise"
	RemainderNoops        = "RemainderNoops"
	Repr                  = "Repr"
	Return                = "Return"
	SameLineNoops         = "SameLineNoops"
	Set                   = "Set"
	SetComp               = "SetComp"
	Slice                 = "Slice"
	SliceInternal         = "slice"
	Starred               = "Starred"
	Stmt                  = "stmt"
	Store                 = "Store"
	Str                   = "Str"
	StringLiteral         = "StringLiteral"
	Sub                   = "Sub"
	Subscript             = "Subscript"
	Suite                 = "Suite"
	Try                   = "Try"
	TryExcept             = "TryExcept"
	TryFinally            = "TryFinally"
	Tuple                 = "Tuple"
	UAdd                  = "UAdd"
	USub                  = "USub"
	UnaryOp               = "UnaryOp"
	UnaryopInternal       = "unaryop"
	While                 = "While"
	With                  = "With"
	Withitem              = "withitem"
	Yield                 = "Yield"
	YieldFrom             = "YieldFrom"
)

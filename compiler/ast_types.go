package compiler

type AstDataType int

const (
	AstDataTypeUnknoen AstDataType = 0
	AstDataTypeNum     AstDataType = 1
	AstDataTypeStr     AstDataType = 2
)

type Ast struct {
	Functions []AstFunction
}

type AstFunction struct {
	Name       string
	ReturnType string
	Body       []AstStatement
}

type AstStatementType int

const (
	AstStatementTypeAssignment  = 1
	AstStatementTypeReturn      = 2
	AstStatementTypeConditional = 3
)

type AstStatement struct {
	Type AstStatementType
	val  interface{}
}

type AstStatementAssignment struct {
	VariableName string
	ErrorHandled bool
	Expr         AstExpr
}
type AstStatementReturn struct {
	VariableName string
}
type AstStatementConditional struct {
	Expr AstExpr
	Body []AstStatement
}

type AstBinOpType int

const (
	AstBinOpTypePlus  = 1
	AstBinOpTypeMinus = 2
	AstBinOpTypeMul   = 3
	AstBinOpTypeDiv   = 4
)

type AstExpr struct {
	ConstNumber *AstExprConstNumber
	ConstStr    *AstExprConstStr
	BinOp       *AstExprBinOp
	Indentifier *AstExprIdentifier
}

type AstExprConstNumber struct {
	Value int64
}

type AstExprConstStr struct {
	Value string
}

type AstExprBinOp struct {
	Op        AstBinOpType
	LeftNode  AstExpr
	RightNode AstExpr
}

type AstExprIdentifier struct {
	Name string
	Type AstDataType
}

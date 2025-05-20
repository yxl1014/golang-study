package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 未知的
	ILLEGAL = "ILLEGAL"
	// 终止符
	EOF = "EOF"

	/// 标识符 + 字面量
	// 变量
	IDENT = "IDENT"
	// 整形
	INT = "INT"

	/// 运算符
	// 复制
	ASSIGN = "="
	// 加
	PLUS = "+"

	/// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	/// 关键字
	// 方法
	FUNCTION = "FUNCTION"
	// 定义
	LET = "LET"
)

package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

/*
To disable `exported method should have comment or be unexported`
https://github.com/golang/lint/issues/186

> For command line:
> gometalinter --exclude="\bexported \w+ (\S*['.]*)([a-zA-Z'.*]*) should have comment or be unexported\b"
>
> For VSCode settings.json:
>
> "go.lintTool": "gometalinter",
> "go.lintFlags": [
> 	"--exclude=\"\bexported \\w+ (\\S*['.]*)([a-zA-Z'.*]*) should have comment or be unexported\b\""
> ],
*/
const (
	/*
		use string for now for simplicy,
		may change to number btyes fpr performance
	*/
	// special
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	// 1343456
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

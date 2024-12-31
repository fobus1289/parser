package parser

import (
	"fmt"
	"strings"
)

type TokenType int

const (
	EOF               TokenType = iota // End of file
	PLACEHOLDER                        // Placeholder
	INVALID                            // Invalid
	NEWLINE                            // Newline
	STRING                             // String
	NUMBER                             // Number
	BOOL                               // Boolean
	NULL                               // Null
	COMMENT                            // Comment
	TEXT                               // Text
	VAR                                // Variable
	LPAREN                             // Left parenthesis (
	RPAREN                             // Right parenthesis )
	LBRACE                             // Left brace {
	RBRACE                             // Right brace }
	IDENT                              // Identifier
	ASSIGN                             // Assignment operator =
	COLON                              // Colon :
	COMMA                              // Comma ,
	SEMICOLON                          // Semicolon ;
	PIPE                               // Pipe |
	QUOTE                              // Quote "
	QUOTE_END                          // Quote end "
	QUOTE_ESCAPE                       // Quote escape \
	QUOTE_ESCAPE_CHAR                  // Quote escape character
)

var tokenTypeNames = map[TokenType]string{
	EOF:               "EOF",
	NEWLINE:           "NEWLINE",
	STRING:            "STRING",
	NUMBER:            "NUMBER",
	BOOL:              "BOOL",
	NULL:              "NULL",
	COMMENT:           "COMMENT",
	TEXT:              "TEXT",
	VAR:               "VAR",
	LPAREN:            "LPAREN",
	RPAREN:            "RPAREN",
	LBRACE:            "LBRACE",
	RBRACE:            "RBRACE",
	IDENT:             "IDENT",
	ASSIGN:            "ASSIGN",
	COLON:             "COLON",
	COMMA:             "COMMA",
	SEMICOLON:         "SEMICOLON",
	PIPE:              "PIPE",
	QUOTE:             "QUOTE",
	QUOTE_END:         "QUOTE_END",
	QUOTE_ESCAPE:      "QUOTE_ESCAPE",
	QUOTE_ESCAPE_CHAR: "QUOTE_ESCAPE_CHAR",
	PLACEHOLDER:       "PLACEHOLDER",
	INVALID:           "INVALID",
}

type Token struct {
	Type  TokenType
	Value string
	Start int
	End   int
}

func (t Token) String() string {
	if name, ok := tokenTypeNames[t.Type]; ok {
		return fmt.Sprintf("%s: %s [%d:%d]", name, t.Value, t.Start, t.End)
	}
	return ""
}

func (t Token) Trim() string {

	value := strings.TrimSpace(t.Value)
	{
		if len(value) < 2 {
			return ""
		}
		value = value[1 : len(value)-1]
	}

	return value
}

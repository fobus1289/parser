package parser

import "sort"

type Parser struct {
	lexer *Lexer
}

func NewParser(input string) *Parser {
	return NewParserWithLexer(NewLexer(input))
}

func NewParserWithLexer(lexer *Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) ParsePlaceholder() (Token, bool) {

	for {
		token, ok := p.lexer.Peek()
		{
			if !ok {
				return Token{Type: EOF, Value: "PLACEHOLDER not found"}, false
			}
		}

		if token == '{' {
			break
		}

		p.lexer.NextPos()
	}

	token, ok := p.lexer.Read('{', '}', PLACEHOLDER)
	{
		if !ok {
			return Token{Type: EOF, Value: "PLACEHOLDER not found"}, false
		}
	}

	return token, true
}

func (p *Parser) ParsePlaceholders() []Token {

	var tokens []Token

	for {
		token, ok := p.ParsePlaceholder()
		if !ok {
			break
		}
		tokens = append(tokens, token)
	}

	return tokens
}

func Filter[T any](input []T, fn func(T) bool) []T {
	var output []T
	for _, v := range input {
		if fn(v) {
			output = append(output, v)
		}
	}
	return output
}

type Replacer interface {
	map[string]string | func(string) string
}

func ReplaceWithTokens[T Replacer](input string, tokens []Token, replacer T) string {

	tokens = Filter(tokens, func(token Token) bool {
		return token.Type == PLACEHOLDER
	})

	if len(tokens) == 0 {
		return input
	}

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i].Start < tokens[j].Start
	})

	var repl func(string) string
	{
		switch v := any(replacer).(type) {
		case map[string]string:
			repl = func(key string) string {
				return v[key]
			}
		case func(string) string:
			repl = v
		}
	}

	var prevEnd int

	buff := make([]byte, 0, len(input))

	for _, token := range tokens {

		segment := input[prevEnd:token.Start]

		value := repl(token.Trim())

		combined := segment + value

		buff = append(buff, combined...)

		prevEnd = token.End
	}

	buff = append(buff, input[prevEnd:]...)

	return string(buff)
}

package parser

import "strings"

func CamelCase(s string) string {
	p := NewParser(s)

	var result []byte

	tokens := Filter(p.Parse(), func(t Token) bool {
		return t.Type == IDENT
	})

	for i, token := range tokens {
		if token.Type == IDENT {
			res := token.Join(token.SplitUpper, func(s string) string {
				if len(s) == 0 {
					return s
				}
				if i == 0 {
					return strings.ToLower(s[0:1]) + s[1:]
				}
				return strings.ToUpper(s[0:1]) + s[1:]
			})

			result = append(result, res...)
		}
	}

	return string(result)
}

// func PascalCase(s string) string {
// 	p := NewParser(s)

// 	tokens := p.Parse()

// 	for _, token := range tokens {
// 		token.UpperFirst()
// 	}

// 	return strings.Join(tokens, "")
// }

func SnakeCase(s string) string {
	p := NewParser(s)

	tokens := Filter(p.Parse(), func(t Token) bool {
		return t.Type == IDENT
	})

	var result []byte

	for i, token := range tokens {
		if token.Type == IDENT {

			res := token.Join(token.SplitUpper, func(s string) string {
				if i == 0 {
					return strings.ToLower(s)
				}
				return "_" + strings.ToLower(s)
			})

			result = append(result, res...)
		}
	}

	return string(result)
}

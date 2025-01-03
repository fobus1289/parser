package parser

func CamelCase(s string) string {
	p := NewParser(s)

	var result []byte

	tokens := p.Parse()

	for _, token := range tokens {
		if token.Type == IDENT {
			result = append(result, token.UpperFirst()...)
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

	tokens := p.Parse()

	var result []byte

	for i, token := range tokens {
		if token.Type == IDENT {

			res := token.Join(token.SplitUpper, func(s string) string {
				if i == 0 {
					return s
				}
				return "_" + s
			})

			result = append(result, res...)
		}
	}

	return string(result)
}

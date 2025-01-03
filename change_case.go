package parser

import (
	"reflect"
	"strings"
	"unicode"
	"unsafe"
)

func CamelCaseOptimized(input string) string {
	var result strings.Builder
	result.Grow(len(input))

	capitalizeNext := false
	for i, r := range input {
		if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
			capitalizeNext = true
			continue
		}

		if capitalizeNext || i == 0 {
			result.WriteRune(unicode.ToUpper(r))
			capitalizeNext = false
		} else {
			result.WriteRune(unicode.ToLower(r))
		}
	}

	res := []rune(result.String())
	res[0] = unicode.ToLower(res[0])
	return string(res)
}

func CamelCaseUnsafe(input string) string {
	// Преобразуем строку в байтовый срез без копирования
	inBytes := *(*[]byte)(unsafe.Pointer(&input))
	// outBytes := make([]byte, 0, len(inBytes)) // Создаём выходной массив с избыточной длиной

	outSliceHeader := &reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&inBytes[0])),
		Len:  0,
		Cap:  len(inBytes),
	}

	outBytes := *(*[]byte)(unsafe.Pointer(outSliceHeader))

	capitalizeNext := true
	for _, b := range inBytes {
		if b == '_' || b == '-' || b == ' ' || unicode.IsPunct(rune(b)) { // Проверяем разделители
			capitalizeNext = true
			continue
		}
		if capitalizeNext {
			// Преобразуем к верхнему регистру
			if b >= 'a' && b <= 'z' {
				b -= 'a' - 'A'
			}
			capitalizeNext = false
		} else {
			// Преобразуем к нижнему регистру
			if b >= 'A' && b <= 'Z' {
				b += 'a' - 'A'
			}
		}
		outBytes = append(outBytes, b)
	}

	// Преобразуем байты обратно в строку
	return *(*string)(unsafe.Pointer(&outBytes))
}

func CamelCase(s string) string {
	p := NewParser(s)

	var result = make([]byte, 0, len(s))

	// tokens := Filter(p.Parse(), func(t Token) bool {
	// 	return t.Type == IDENT
	// })

	for i, token := range p.Parse() {
		if token.Type == IDENT {
			res := token.Join(token.SplitUpper, func(s string) string {
				if len(s) == 0 {
					return s
				}

				if i == 0 {
					q := []rune(s)
					if unicode.IsLower(q[0]) {
						return s
					}
					q[0] = unicode.ToLower(q[0])
					return string(q)
				}

				q := []rune(s)

				if unicode.IsUpper(q[0]) {
					return s
				}

				q[0] = unicode.ToUpper(q[0])
				return string(q)
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

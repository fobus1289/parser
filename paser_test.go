package parser

import (
	"fmt"
	"testing"
)

func TestFindTexts(t *testing.T) {

	const EXP = " Testing_@@@testing _______testing"

	fmt.Println(SnakeCase(EXP))
	fmt.Println(CamelCase(EXP))
}

func BenchmarkSnakeCase(b *testing.B) {
	const EXP = " Testing_@@@testing _______testing"
	for i := 0; i < b.N; i++ {
		_ = SnakeCase(EXP)
	}
}

func TestParsePlaceholders_ValidInput(t *testing.T) {
	const EXP = "id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}"
	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()

	expected := []string{"{ID}", "{NAME}", "{AGE}", "{SALARY}"}
	for _, exp := range expected {
		found := false
		for _, ph := range placeholders {
			if ph.Value == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected placeholder %s not found", exp)
		}
	}
}

func TestParsePlaceholders_EmptyInput(t *testing.T) {
	const EXP = ""
	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()

	if len(placeholders) != 0 {
		t.Errorf("Expected no placeholders, got %v", placeholders)
	}
}

func TestParsePlaceholders_NoPlaceholders(t *testing.T) {
	const EXP = "No placeholders here"
	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()

	if len(placeholders) != 0 {
		t.Errorf("Expected no placeholders, got %v", placeholders)
	}
}

func TestParsePlaceholders_MultiplePlaceholders(t *testing.T) {
	const EXP = "id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}/city:{CITY}"
	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()
	expected := []string{"{ID}", "{NAME}", "{AGE}", "{SALARY}", "{CITY}"}

	for _, exp := range expected {
		found := false
		for _, ph := range placeholders {
			if ph.Value == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected placeholder %s not found", exp)
		}
	}
}

func TestReplaceWithTokens_ReplaceWithMap(t *testing.T) {

	const EXP = "id:{ID}/name:{NAME}/age:{AGE}/salary:{SALARY}"

	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()

	replacer := map[string]string{
		"ID":     "1",
		"NAME":   "John",
		"SALARY": "50000",
		"AGE":    "25",
	}

	result := ReplaceWithTokens(EXP, placeholders, replacer)
	t.Log(result)
	if result != "id:1/name:John/age:25/salary:50000" {
		t.Errorf("Expected result to be 'id:1/name:John/age:25/salary:50000', got %s", result)
	}
}

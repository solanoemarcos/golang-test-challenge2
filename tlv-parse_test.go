package main

import "testing"

func TestTlvParseEmptyArray(t *testing.T) {
	input := []byte{}
	m, err := TlvParse(input)

	if m != nil {
		t.Error("Return map should be empty")
	}
	if err == nil {
		t.Error("Parse with empty array should return error")
	}
	emp, ok := err.(*EmptyArrayError)
	if !ok {
		t.Errorf("Expected EmptyArrayError but got: %v", emp)
	}
}

func TestTlvParseSuccessCase(t *testing.T) {
	//11AB398765UJ1A05
	caso1 := []byte{'1', '1', 'A', 'B', '3', '9', '8', '7', '6', '5', 'U', 'J', '1', 'A', '0', '5'}
	m, err := TlvParse(caso1)
	if err != nil {
		t.Errorf("Unexpected Error but got: %v", err)
	}
	if m != nil {
		if m["largo"] != "11" {
			t.Errorf("Wrong length expected 11, got: %s", m["largo"])
		}
		if m["valor"] != "AB398765UJ1" {
			t.Errorf("Wrong value, expected \"AB398765UJ1\", got: %s", m["valor"])
		}
		if m["tipo"] != "A05" {
			t.Errorf("Wrong type, expected \"A05\", got: %s", m["tipo"])
		}
	}
}

package test

import (
	"testing"

	"github.com/solanoemarcos/golang-test-challenge2/parser"
)

func TestTlvParseEmptyArray(t *testing.T) {
	input := []byte{}
	m, err := parser.TlvParse(input)

	if m != nil {
		t.Error("Return map should be empty")
	}
	if err == nil {
		t.Error("Parse with empty array should return error")
	}
	emp, ok := err.(*parser.InvalidFieldSizeError)
	if !ok {
		t.Errorf("Expected EmptyArrayError but got: %v", emp)
	}
}

func TestTlvParseSuccessCase(t *testing.T) {
	//test cases
	case1 := []byte("11AB398765UJ1A05")
	result1 := map[string]string{"length": "11", "value": "AB398765UJ1", "type": "A05"}
	case2 := []byte("0255N23")
	result2 := map[string]string{"length": "02", "value": "55", "type": "N23"}
	testcases := [][]byte{case1, case2}
	testresult := []map[string]string{result1, result2}
	for i, testcase := range testcases {
		m, err := parser.TlvParse(testcase)
		if err != nil {
			t.Errorf("Unexpected Error for case %d: %v", i, err)
		}
		if m != nil {
			if m["length"] != testresult[i]["length"] {
				t.Errorf("Wrong length expected \"%s\", got: \"%s\"", testresult[i]["length"], m["length"])
			}
			if m["value"] != testresult[i]["value"] {
				t.Errorf("Wrong value, expected \"%s\", got: %s", testresult[i]["value"], m["value"])
			}
			if m["type"] != testresult[i]["type"] {
				t.Errorf("Wrong type, expected \"%s\", got: %s", testresult[i]["type"], m["type"])
			}
		}
	}

}

func TestTlvParseInvalidCase(t *testing.T) {
	case1 := []byte("11AB3985UJ1A05")
	case2 := []byte("025AN23")
	case3 := []byte("0200X23")
	case4 := []byte("0A00N23")
	case5 := []byte("02**A23")
	case6 := []byte("1133398500AA05")
	cases := [][]byte{case1, case2, case3, case4, case5, case6}
	for i, testcase := range cases {
		_, err := parser.TlvParse(testcase)
		if err == nil {
			t.Errorf("Expected Error but got nothing, for case %d", i+1)
		}
	}

}

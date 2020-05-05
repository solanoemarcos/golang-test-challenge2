package parser

import (
	"regexp"
	"strconv"
)

//InvalidTlvError If value field does not match the specified type
type InvalidTlvError struct {
}

//InvalidFieldSizeError If any field does not match specified size
type InvalidFieldSizeError struct {
	field string
}

func (e *InvalidTlvError) Error() string       { return "Value does not match data type." }
func (e *InvalidFieldSizeError) Error() string { return e.field + " incomplete." }

//TlvParse parse the byte array
func TlvParse(input []byte) (map[string]string, error) {
	//range
	length := len(input)
	i := 0
	j := i + 2
	//TLV length
	if length < j {
		return nil, &InvalidFieldSizeError{field: "Length"}
	}
	lengths := string(input[i:j])
	lengthi, err := strconv.Atoi(lengths)
	if err != nil {
		return nil, err
	}
	//validate size
	k := lengthi + j
	if length < k+3 {
		return nil, &InvalidFieldSizeError{field: "Value or Type"}
	}
	//TLV value
	values := string(input[j:k])
	//TLV type
	typeb := input[k : k+3]
	types := string(typeb)
	//validation
	validate := validation(typeb[0])
	if !validate(values) {
		return nil, &InvalidTlvError{}
	}
	//build result
	result := map[string]string{
		"length": lengths,
		"type":   types,
		"value":  values,
	}
	return result, nil
}

type strategy func(string) bool

func validation(tipo byte) strategy {
	switch tipo {
	case 'A': //alphanumeric
		return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	case 'N': //numeric
		return regexp.MustCompile(`^[0-9]*$`).MatchString
	default:
		return func(value string) bool { return false }
	}
}

package parser

import (
	"regexp"
	"strconv"
)

//EmptyArrayError Si la entrada esta vacia.
type EmptyArrayError struct {
}

//InvalidTlvError Si algún campo no cumple con el tipo de dato especificado.
type InvalidTlvError struct {
}

//InvalidFieldSizeError Si algún campo no cumple con el tamaño especificado.
type InvalidFieldSizeError struct {
	field string
}

func (e *EmptyArrayError) Error() string       { return "Array is empty." }
func (e *InvalidTlvError) Error() string       { return "Value does not match data type." }
func (e *InvalidFieldSizeError) Error() string { return e.field + " incomplete." }

//TlvParse ejecuta el parseo del arreglo de bytes
func TlvParse(input []byte) (map[string]string, error) {
	//range
	length := len(input)
	i := 0
	j := i + 2
	//TLV length
	if length < j {
		return nil, &InvalidFieldSizeError{field: "Length"}
	}
	largos := string(input[i:j])
	largoi, err := strconv.Atoi(largos)
	if err != nil {
		return nil, err
	}
	//validate size
	k := largoi + j
	if length < k+3 {
		return nil, &InvalidFieldSizeError{field: "Value or Type"}
	}
	//TLV value
	valorb := input[j:k]
	valors := string(valorb)
	//TLV type
	tipob := input[k : k+3]
	tipos := string(tipob)
	//validation
	if !tlvValidate(tipob, valorb, largoi) {
		return nil, &InvalidTlvError{}
	}
	//build result
	result := make(map[string]string)
	result["largo"] = largos
	result["tipo"] = tipos
	result["valor"] = valors
	return result, nil
}

func tlvValidate(tipo []byte, valor []byte, length int) bool {
	valid := validation(tipo[0])
	b := true
	for i := 0; b && i < length; i++ {
		b = valid(valor[i])
	}
	return b
}

type strategy func(byte) bool

func validation(tipo byte) strategy {
	switch tipo {
	case 'A':
		return alphanumericValidate
	case 'N':
		return numericValidate
	default:
		return func(value byte) bool { return false }
	}
}

func alphanumericValidate(value byte) bool {
	isStringAlphaNumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	return isStringAlphaNumeric(string(value))
}

func numericValidate(value byte) bool {
	_, err := strconv.Atoi(string(value))
	if err != nil {
		return false
	}
	return true
}

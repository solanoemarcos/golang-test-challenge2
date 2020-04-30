package main

import "strconv"

//EmptyArrayError Si la entrada esta vacia.
type EmptyArrayError struct {
}

//InvalidDataTypeError Si algún campo no cumple con el tipo de dato especificado.
type InvalidDataTypeError struct {
	dataType string
}

func (e *EmptyArrayError) Error() string      { return "Array is empty." }
func (e *InvalidDataTypeError) Error() string { return "Character is not of type " + e.dataType }

//TlvParse ejecuta el parseo del arreglo de bytes
func TlvParse(input []byte) (map[string]string, error) {
	length := len(input)
	if length < 5 {
		return nil, &EmptyArrayError{}
	}
	res := make(map[string]string)
	//----------- atributo largo TLV
	largos := string(input[0:2])
	largoi, _ := strconv.Atoi(largos)
	res["largo"] = largos
	//----------- atributo valor TLV
	valors := string(input[2 : 2+largoi])
	res["valor"] = valors
	//----------- atributo tipo TLV
	tipos := string(input[largoi+2 : largoi+5])
	res["tipo"] = tipos
	return res, nil
}

func main() {

}

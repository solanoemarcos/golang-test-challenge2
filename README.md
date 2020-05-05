# TLV problem solution

Field | Position | Size
------------ | ------------ | ------------
LENGTH | 0 | 2
VALUE | 2 | LENGTH
TYPO | LENGTH+2 | 3

For this problem the first step is to find the ranges and validate the size of the input.
```golang
i := 0
j := i + 2
length := len(input)
if length < j {
  return nil, &InvalidFieldSizeError{field: "Length"}
}
```

If that's ok, then validate type of the value.
```golang
if !validate(values) {
  return nil, &InvalidTlvError{}
}
```

I decide to use a regexp pattern for this validation, and the pattern will be selected by the "type" in the TLV:
```golang
func validation(tipo byte) strategy {
	switch tipo {
	case 'A':
		return regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString
	case 'N':
		return regexp.MustCompile(`^[0-9]*$`).MatchString
	default:
		return func(value string) bool { return false }
	}
}
```

# Development

For testing use:
```
go test ..\test -coverpkg ./...
```
The coverage is 92.6% of statements, only left untested the return from the errors:
![Uncovered](./assets/Coverage.PNG)

Build/run:
```
go build
go run .
```

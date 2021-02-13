package typeguard

import (
	"reflect"
	"strconv"
	"strings"
)

var (
	arrStringsType = reflect.TypeOf([]string{}).Name()
	arrIntsType    = reflect.TypeOf([]int{}).Name()
	intType        = reflect.TypeOf(1).Name()
)

// Output ..
type Output struct {
	Value   string
	Default interface{}
}

// WantArrString returns string for type []string
func WantArrString() string {
	return arrStringsType
}

// ToArrString returns value of type []string
func (o Output) ToArrString() (output []string, err error) {
	if o.Value == "" {
		return o.Default.([]string), nil
	}

	output = strings.Split(o.Value, ",")

	return output, err
}

// WantInt returns string for type int
func WantInt() string {
	return intType
}

// ToInt returns value of type int
func (o Output) ToInt() (output int, err error) {
	if o.Value == "" {
		return o.Default.(int), nil
	}

	output, err = strconv.Atoi(o.Value)

	return output, err
}

// WantArrInt returns string for type []int
func WantArrInt() string {
	return arrIntsType
}

// ToArrInt returns value of type []int
func (o Output) ToArrInt() (output []int, err error) {
	if o.Value == "" {
		return o.Default.([]int), nil
	}

	arrStrings := strings.Split(o.Value, ",")

	for i := 0; i < len(arrStrings); i++ {
		n, err := strconv.Atoi(arrStrings[i])
		if err != nil {
			return output, err
		}

		output = append(output, n)
	}

	return output, err
}

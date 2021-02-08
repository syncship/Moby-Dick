package typeguard

import (
	"reflect"
	"strconv"
	"strings"
)

var (
	arrStringTyp = reflect.TypeOf([]string{}).Name()
	intTyp       = reflect.TypeOf(1).Name()
	arrIntTyp    = reflect.TypeOf([]int{}).Name()
)

// Output ..
type Output struct {
	Value   string
	Default interface{}
}

// WantArrString returns string for type []string
// TODO: dava para cachear o tipo
func WantArrString() string {
	return arrStringTyp
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
	return intTyp
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
	return arrIntTyp
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

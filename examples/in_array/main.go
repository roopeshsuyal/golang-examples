package main

import "fmt"
import "reflect"

func in_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func main() {

	names := []string{"Rohan", "Sohan", "Mohit", "Komal", "Raju"}
	fmt.Println(in_array("Rahul", names))
	fmt.Println(in_array("Deepak", names))

	ints := []int{7, 5, 2, 4, 8}
	fmt.Println(in_array(4, ints))
	fmt.Println(in_array(6, ints))
}

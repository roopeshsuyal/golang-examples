package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Category struct {
	ID       uint
	Name     string
	ParentID uint
}

var categories = []Category{
	Category{1, "Root", 1},
	Category{2, "Grooming", 1},
	Category{3, "Health", 1},
	Category{4, "Technology", 1},
	Category{5, "Relationships", 1},
	Category{6, "Smart Phones", 4},
	Category{7, "Web & Social", 4},
	Category{8, "PCs", 4},
	Category{9, "Internet", 4},
	Category{10, "Games", 4},
	Category{11, "Better Partner", 5},
	Category{12, "Impress a Girl", 16},
	Category{13, "Learn to Handle Cheating", 5},
	Category{14, "Long Distance Relationship", 5},
	Category{15, "Break Ups", 5},
	Category{16, "Understanding Women", 5},
	Category{17, "Nutrition", 3},
	Category{18, "Weight Loss", 3},
	Category{19, "Bodybuilding", 3},
	Category{20, "Celebrity Fitness", 3},
	Category{21, "Mental Health", 3},
	Category{22, "Hair Care Tips", 2},
	Category{23, "Skin Care", 2},
	Category{24, "Manscaping", 2},
	Category{25, "Hairfall", 28},
	Category{26, "Shaving", 2},
	Category{28, "Hair Care", 2},
	Category{29, "Hair Removal", 2},
	Category{30, "Personal Hygiene", 2},
	Category{31, "Beards and Shaving", 2},
	Category{32, "Hairstyle", 28},
	Category{33, "Skincare", 2},
	Category{34, "Wellness", 2},
	Category{35, "Celebrity Grooming", 2},
	Category{36, "Reviews", 4},
	Category{37, "Desktop PCs", 8},
	Category{38, "XBox Games", 10},
	Category{39, "3D Games", 10},
	Category{40, "Social Media", 4},
	Category{41, "News", 4},
	Category{42, "Science And Future", 4},
	Category{43, "What to Wear", 6},
	Category{44, "Style Guide", 6},
	Category{45, "Fashion Tips", 43},
	Category{46, "Style Trends", 43},
	Category{47, "Men's Watches", 6},
	Category{48, "Trends", 6},
	Category{49, "Celebrity Style", 6},
	Category{50, "Interviews", 6},
	Category{51, "Watches", 48},
	Category{52, "Shoes", 48},
	Category{53, "Laptop PCs", 8},
}

//okay tested
func main() {

	fmt.Println("Input categories=>", categories)

	for _, cat := range categories {
		urlValues := structToMap(&cat)
		fmt.Printf("\n%#v", urlValues)
	}

}

// strt should be a pointer (*)
func structToMap(strt interface{}) (values map[string]interface{}) {

	values = map[string]interface{}{}
	iVal := reflect.ValueOf(strt).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values[typ.Field(i).Name] = v
	}
	return
}

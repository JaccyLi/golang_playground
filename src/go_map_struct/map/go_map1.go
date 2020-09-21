package main

import (
	"fmt"
)

func main() {
	var map1 = map[string]map[string]string{"stu1": map[string]string{"name": "jack", "age": "23", "home": "haidian"},
		"stu2": map[string]string{"name": "john", "age": "18", "home": "xichen"},
		"stu3": map[string]string{"name": "jeessca", "age": "18", "home": "fengtai"}}
	fmt.Printf("type is %T\nvalue is %v\nlen is %v\n", map1, map1, len(map1))
	v, ok := map1["stu2"]
	fmt.Printf("v.name= %v\nv.age= %v\nv.home= %v\nok? %v", v["name"], v["age"], v["home"], ok)

}

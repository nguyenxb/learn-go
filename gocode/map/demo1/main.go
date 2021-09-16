package main

import "fmt"

func main() {

	// for range 遍历结构复杂的map
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街"

	studentMap["stu02"] = make(map[string]string, 2)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1, "v1=", v1)
		for k2, v2 := range v1 {
			fmt.Printf("k2=%v,v2=%v\n", k2, v2)
		}
		fmt.Println("\n")
	}
	fmt.Println("map size=", len(studentMap))
	fmt.Println("map size=", len(studentMap["stu01"]))
	fmt.Println("map size=", len(studentMap["stu02"]))
	fmt.Println("map size=", len(studentMap["stu02"]["name"]))
	fmt.Println("map size=", len(studentMap["stu01"]["name"]))
}

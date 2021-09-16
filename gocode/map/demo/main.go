package main

import "fmt"

func main() {
	/*
			var map的变量名 map[keyType]valueType
		key的类型：
			可以是bool,数字,string,指针,channel,还可以是只包含
		前面几个类型的 接口，结构体，数组
		通常是int、string

		注意：slice,map 还有function 不可以，因为这几个没法用 == 来判断

		value的类型：
			与key的类型基本一样，
			通常为数字（int,float)string,map,struct
	*/
	// 注意：map声明是不会分配内存的,初始化需要make,
	//分配内存后才能赋值和使用
	// define
	var map1 map[string]string
	var map2 map[string]map[string]string
	var map3 map[int]string
	// map1["123"] = "sdasd" // compile error
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	// map 的key-value是无序的，key不能重复，重复就会覆盖，value可以重复

	map1 = make(map[string]string, 10)
	map1["123"] = "asdcsd"
	fmt.Println(map1)

	// define and use map
	//	 第一种方式
	var a map[string]string
	// 在使用map前，需要先make，给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "松江"
	a["no2"] = "武松"
	a["no3"] = "哇塞"
	a["no1"] = "阿迪斯"
	fmt.Println(a)

	// 2
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "天津"
	fmt.Println(cities)

	// 3
	heroes := map[string]string{
		"heros1": "松江",
		"hero2":  "武松",
	}
	heroes["123"] = "使得"
	fmt.Println(heroes)

	// 4
	var aa = map[string]string{
		"ads":  "aadsd",
		"adsa": "asdadsadss",
	}
	fmt.Println(aa)

	// 5
	var aaa map[string]string = map[string]string{
		"123": "asdfs",
		"456": "adssda",
	}
	fmt.Println(aaa)

	// 存放3个学生的姓名和性别
	studenMap := make(map[string]map[string]string)

	studenMap["stu01"] = make(map[string]string, 3)
	studenMap["stu01"]["name"] = "tom"
	studenMap["stu01"]["sex"] = "男"
	studenMap["stu01"]["address"] = "北京长安街"

	studenMap["stu02"] = make(map[string]string, 3)
	studenMap["stu02"]["name"] = "mary"
	studenMap["stu02"]["sex"] = "女"
	studenMap["stu02"]["address"] = "上海黄浦江"

	fmt.Println(studenMap)
	fmt.Println(studenMap["stu01"])
	fmt.Println()

	/////////////////////////////////////////////
	// curd 操作
	myMap := make(map[string]string)
	fmt.Println(myMap)
	// 增加
	myMap["no1"] = "张三"
	myMap["no2"] = "李四"
	myMap["no3"] = "王五"
	fmt.Println("增加:", myMap)

	// 修改
	myMap["no1"] = "李达"
	fmt.Println("修改:", myMap)

	// 删除 func delete(m map[Type]Type1, key Type)
	//内建函数delete按照指定的键将元素从映射中删除。
	// 若m为nil或无此元素，delete不进行操作。
	delete(myMap, "no2")
	delete(myMap, "no7") // 不进行操作
	fmt.Println("删除:", myMap)

	// 删除map 中所有的key
	// 1 ,遍历删除
	// for key, value := range myMap {
	// 	fmt.Println("key=", key, "value=", value)
	// 	delete(myMap, key)
	// }
	// fmt.Println(myMap)

	// 2

	// myMap = nil // 指针置为空
	// 或者
	//myMap = make(map[string]string) // 重新指向另一个map
	// fmt.Println("删除所有2：", myMap)

	// 查询

	fmt.Println()

	// 查询
	// 如果 找到那么 ok == true ,没有找到 , ok == false
	key, ok := myMap["no5"]
	if ok {
		fmt.Println("找到了 key=", key)
	} else {
		fmt.Println("没有找到key = no1")
	}

}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 1统计字符串长度 按字节len(str)
	str := "hello北" //中文有3个字节
	fmt.Println("str len=", len(str))

	// 2遍历字符串
	str2 := "hello北京"
	for i := 0; i < len(str2); i++ {
		fmt.Printf("str2 = %c\n", str2[i])
	}
	fmt.Println()
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("str2 = %c\n", r[i])
	}

	// 3字符串转整数 n,err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误：", err)
	} else {
		fmt.Printf("转换成功：%v,n=%T\n", n, n)
	}

	// 4整数转字符串 str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v, str=%T\n", str, str)

	// 5字符串转 []byte   var  bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Println(bytes)

	// 6 []byte 转字符串 str = string("[]byte{97,98,100}")
	str = string([]byte{97, 98, 100})
	fmt.Printf("str=%v\n", str)

	// 7 十进制转 2，8，16进制 str = strconv.FormatInt(132,2)
	// 转2进制
	str = strconv.FormatInt(123, 2)
	fmt.Printf("str=%v\n", str)
	// 转8进制
	str = strconv.FormatInt(123, 8)
	fmt.Printf("str=%v\n", str)
	// 转16进制
	str = strconv.FormatInt(123, 16)
	fmt.Printf("str=%v\n", str)

	// 8 查找子串是否在指定字符串中：strings.Contains("seafood","foo")// true
	b := strings.Contains("seafood", "foo") //true
	fmt.Printf("b=%v\n", b)

	b = strings.Contains("seafood", "asd") // false
	fmt.Printf("b=%v\n", b)

	// 9 统计一个字符串有几个指定的子串，strings.Count("cheese","e")// 4
	num := strings.Count("cheese", "e")
	fmt.Println("num=", num)

	num = strings.Count("cheese", "se")
	fmt.Println("num=", num)

	num = strings.Count("cheese", "sad")
	fmt.Println("num=", num)

	//10 不区分大小写的字符串比较：strings.EqualFold("abc","Abc")// true
	// == 是区分大小写的
	b = strings.EqualFold("abc", "Abc")
	fmt.Println("b=", b)

	b = "abc" == "Abc"
	fmt.Println("b=", b)

	// 11 返回子串在字符串第一次出现的index值，如果没有返回-1：
	// strings.Index("NLT_abc","abc")//4
	index := strings.Index("NLT_abc", "abc")
	fmt.Println("11index=", index)

	index = strings.Index("NLT_abcabc", "abc")
	fmt.Println("11index=", index)

	index = strings.Index("NLT_mbcsdbc", "abc")
	fmt.Println("11index=", index)

	// 12 返回子串在字符串最后一次出现的位置
	// 如没有则返回-1，strings.LastIndex("go golang","go")

	index = strings.LastIndex("go golang", "go")
	fmt.Println("12index=", index) //3

	// 13 将指定的子串替换成 另一个子串：strings.replace("go golang","go","语言",n)
	// n 可以指定替换几个，如果n = -1表示全部替换
	str2 = strings.Replace("go golang", "go", "语言", 1)
	fmt.Println("str2=", str2)

	str2 = strings.Replace("go golang", "go", "语言", -1)
	fmt.Println("str2=", str2)

	// 14 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组：
	// strings.Split("hello,wrold,ok",",")
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	// 15 将字符串的字母进行大小写转换：
	// strings.ToLower("Go") go
	// strings.ToUpper("Go") GO
	str = "goLang Hello"
	str = strings.ToLower(str)
	fmt.Println("str=", str)

	str = "goLang Hello"
	str = strings.ToUpper(str)
	fmt.Println("str=", str)

	// 16 将字符串左右两边的空格去掉
	// strings.trimSpace("  tn a lone gopher ntrn  ")
	str = strings.TrimSpace("  tn a lone gopher ntrn  ")
	fmt.Printf("str=%q\n", str)

	// 17 将字符串左右两边指定的字符去掉：
	// strings.Trim("!_hello!__!","_!")
	// ["hello"]// 将左右两边的"_"和"!"去掉
	str = strings.Trim("!_hello!__!", "_!")
	fmt.Printf("str=%q\n", str) // str="hello"

	// 18 将字符串左边的指定字符去掉：
	// strings.TrimLeft("!_hello!__!","_!")
	// ["hello!__!"]// 将左边的"_"和"!"去掉
	str = strings.TrimLeft("!_hello!__!", "_!")
	fmt.Printf("str=%q\n", str)

	// 19 将字符串右边的指定字符去掉：
	// strings.TrimRight("!_hello!__!","_!")
	// ["!_hello"]// 将右边的"_"和"!"去掉
	str = strings.TrimRight("!_hello!__!", "_!")
	fmt.Printf("str=%q\n", str)

	// 20 判断字符串是否以指定的字符串开头：
	// strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Println("b=", b) // true

	// 21 判断字符串是否以指定的字符串结尾：
	// strings.Hassuffix("NLT_abs.jpg","jpg")
	b = strings.HasSuffix("NLT_abs.jpg", "jpg")
	fmt.Println("b=", b) // true

}

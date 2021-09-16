package main

import (
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("链接redis失败")
		return
	}
	defer conn.Close()
	fmt.Println("链接redis成功")
	var count int

	for {
		var monster Monster
		fmt.Println("请输入姓名：")
		fmt.Scanln(&monster.Name)
		fmt.Println("请输入年龄：")
		fmt.Scanln(&monster.Age)
		fmt.Println("请输入技能：")
		fmt.Scanln(&monster.Skill)
		chioce := ""
		for {
			fmt.Println("是否确认保存数据（y/n）")
			fmt.Scanln(&chioce)
			if chioce == "y" || chioce == "Y" || chioce == "N" || chioce == "n" {
				break
			}
		}
		if chioce == "N" || chioce == "n" {
			for {
				fmt.Println("是否继续录入数据（y/n）")
				fmt.Scanln(&chioce)
				if chioce == "y" || chioce == "Y" || chioce == "N" || chioce == "n" {
					break
				}
			}
			if chioce == "N" || chioce == "n" {
				break
			}
			continue
		}
		fmt.Printf("monster[].Name = %v\n", monster.Name)
		fmt.Printf("monster[].Age = %v\n", monster.Age)
		fmt.Printf("monster[].Skill = %v\n", monster.Skill)
		// 将数据保存到redis中
		_, err := conn.Do("HMSet", "Monster"+strconv.Itoa(count), "name", monster.Name,
			"age", monster.Age, "skill", monster.Skill)
		if err != nil {
			fmt.Println("保存数据到redis 出错 err=", err)
			return
		}
		fmt.Println("Monster" + strconv.Itoa(count))
		fmt.Printf("monster[].Name = %v\n", monster.Name)
		fmt.Printf("monster[].Age = %v\n", monster.Age)
		fmt.Printf("monster[].Skill = %v\n", monster.Skill)
		fmt.Println("保存数据成功")
		fmt.Println()
		fmt.Println("是否继续录入数据（y/n）")
		fmt.Scanln(&chioce)
		if chioce == "N" || chioce == "n" {
			break
		}
		count++
	}
	var monsters []Monster
	fmt.Println("从redis 中读取所有monster的数据")
	for i := 0; i <= count; i++ {
		var monster Monster
		data, err := redis.Strings(conn.Do("HMGet", "monster"+strconv.Itoa(i),
			"name", "age", "skill"))
		if err != nil {
			fmt.Println("数据从redis读取失败")
			return
		}
		monster.Name = data[0]
		monster.Age, _ = strconv.Atoi(data[1])
		monster.Skill = data[2]
		monsters = append(monsters, monster)
	}
	for i, v := range monsters {
		fmt.Printf("monster[%v].Name = %v\n", i, v.Name)
		fmt.Printf("monster[%v].Age = %v\n", i, v.Age)
		fmt.Printf("monster[%v].Skill = %v\n", i, v.Skill)
	}

}

package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 将对象序列化，并保存到文件中
func (this *Monster) Store() bool {
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}
	filePath := "C:\\Users\\asd\\Desktop.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write err =", err)
		return false
	}

	return true
}
func (this *Monster) ReStore() bool {
	filePath := "C:\\Users\\asd\\Desktop.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read err=", err)
		return false
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("unmarshal err=", err)
		return false
	}
	return true

}

package monster

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "夏明",
		Age:   19,
		Skill: "吃饭",
	}
	b := monster.Store()
	if !b {
		t.Fatalf("monster.Store()错误，希望为%v , 实际为%v", true, b)
	}
	t.Logf("测试成功...")
}
func TestReStore(t *testing.T) {
	monster := &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误，希望为%v , 实际为%v", true, res)
	}
	t.Logf("测试成功")
}

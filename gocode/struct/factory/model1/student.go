package model1

type student struct {
	Name  string
	score float64
}

// 工厂模式：
func NewStudent(name string, score1 float64) *student {
	return &student{
		Name:  name,
		score: score1,
	}
}
func (s *student) GetScore() float64 {
	return s.score
}

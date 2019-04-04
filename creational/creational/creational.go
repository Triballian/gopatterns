package creational

type singleton struct {
	count int
}

var instance *singleton

//GetInstance GetIntance method
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

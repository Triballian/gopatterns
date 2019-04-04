package singleton

// // Singleton interface
// type Singleton interface {
// 	AddOne() int
// }

type singleton struct {
	count int
}

var instance *singleton

// // GetInstance singleton function
// func GetInstance() Singleton {
// 	return nil
// }

//GetInstance GetIntance method
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// func (s *singleton) AddOne() int {
// 	return 0
// }

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

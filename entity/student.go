package entity

import "fmt"

type Student struct {
	Age int
	Name string
	Score float32
}

func (s *Student)Show(){
	fmt.Printf("name: %s\tage: %d\tscore: %v\n",s.Name,s.Age,s.Score)
}

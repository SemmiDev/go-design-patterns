package main

import "fmt"

type studentBuilder struct {
	Name       string
	Identifier string
	Semester   uint32
	PreviousIP float32
}

type StudentBuilder interface {
	SetName(name string) StudentBuilder
	SetIdentifier(identifier string) StudentBuilder
	SetSemester(semester uint32) StudentBuilder
	SetPreviousIP(previousIP float32) StudentBuilder
	Build() Student
}

func (s *studentBuilder) SetIdentifier(identifier string) StudentBuilder {
	s.Identifier = identifier
	return s
}

func (s *studentBuilder) SetSemester(semester uint32) StudentBuilder {
	s.Semester = semester
	return s
}

func (s *studentBuilder) SetPreviousIP(previousIP float32) StudentBuilder {
	s.PreviousIP = previousIP
	return s
}

func (s *studentBuilder) SetName(name string) StudentBuilder {
	s.Name = name
	return s
}

var Fake []*studentBuilder

type Student interface {
	Save() error
	Get() error
	All() error
}

func (s *studentBuilder) Build() Student {
	return &studentBuilder{
		Name:       s.Name,
		Identifier: s.Identifier,
		Semester:   s.Semester,
		PreviousIP: s.PreviousIP,
	}
}

func (s *studentBuilder) Save() error {
	fmt.Println("-------- save student --------")
	Fake = append(Fake, s)
	fmt.Println("------------------------------")
	return nil
}

func (s *studentBuilder) Get() error {
	fmt.Println("--------", s.Name, "--------")
	fmt.Println("Name: ", s.Name)
	fmt.Println("Identifier: ", s.Identifier)
	fmt.Println("Semester: ", s.Semester)
	fmt.Println("Previous IP: ", s.PreviousIP)
	fmt.Println("------------------------------")
	return nil
}

func (s *studentBuilder) All() error {
	for _, v := range Fake {
		fmt.Println("--------", v.Name, "--------")
		fmt.Println("Name: ", v.Name)
		fmt.Println("Identifier: ", v.Identifier)
		fmt.Println("Semester: ", v.Semester)
		fmt.Println("Previous IP: ", v.PreviousIP)
		fmt.Println("------------------------------")
	}
	return nil
}

func main() {
	builder := &studentBuilder{}
	student := builder.
		SetName("sammidev").
		SetIdentifier("202020202").
		SetSemester(2).
		SetPreviousIP(4.0).Build()

	_ = student.Save()
	_ = student.Get()
	_ = student.All()
}

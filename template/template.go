package main

import (
	"fmt"
	"time"
)

type Student struct {
	ID             uint32
	Name           string
	RegistrationAt time.Time
}

func NewStudent(ID uint32, name string, registrationAt time.Time) *Student {
	return &Student{ID: ID, Name: name, RegistrationAt: registrationAt}
}

var JuniorHighSchoolDB []*Student
var SeniorHighSchoolDB []*Student

type Registration interface {
	inputStudent()
}

type AbstractRegistration struct {
	registration Registration
}

func (s *AbstractRegistration) GO() {
	s.registration.inputStudent()
}

type SeniorHighSchool struct {
	*AbstractRegistration
}

var smaCount uint32 = 0
func (s SeniorHighSchool) inputStudent() {
	fmt.Println("\nINPUT SENIOR HIGH SCHOOL REGISTRATION")
	smaCount++
	
	var name string
	fmt.Print("masukkan nama anda : ")
	fmt.Scanln(&name)

	s.prosesStudent(smaCount, name)
}

func (s SeniorHighSchool) prosesStudent(ID uint32, name string) {
	fmt.Println("\nPROCESS SENIOR HIGH SCHOOL REGISTRATION")
	time.Sleep(time.Second)
	fmt.Println("done...")

	std := NewStudent(ID, name, time.Now())
	s.saveStudent(std)
}

func (s SeniorHighSchool) saveStudent(student *Student) {
	fmt.Println("\nSAVE SENIOR HIGH SCHOOL REGISTRATION")
	time.Sleep(time.Second)
	fmt.Println("done...")

	SeniorHighSchoolDB = append(SeniorHighSchoolDB, student)
	s.showStudent()
}

func (s SeniorHighSchool) showStudent() {
	fmt.Println("\nSHOW SENIOR HIGH SCHOOL REGISTRATION")
	time.Sleep(time.Second)
	fmt.Println("done...")

	for _, v := range SeniorHighSchoolDB {
		fmt.Println(v)
	}
}

func main() {
	sam := &AbstractRegistration{&SeniorHighSchool{}}
	sam.GO()

	dev := &AbstractRegistration{&SeniorHighSchool{}}
	dev.GO()
}
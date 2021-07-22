package main

import "fmt"

type AbstractFile interface {
	GetName() string
	PrintList()
}

type File struct {
	name string
}

func (d *File) GetName() string {
	return d.name
}

func (d *File) PrintList() {
	fmt.Println(d.name)
}

type Directory struct {
	includedFiles []AbstractFile
	name          string
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) Add(entry AbstractFile) {
	d.includedFiles = append(d.includedFiles, entry)
}

func (d *Directory) AddAll(entries ...AbstractFile) {
	for _, v := range entries {
		d.includedFiles = append(d.includedFiles, v)
	}
}

func (d *Directory) PrintList() {
	fmt.Println()
	for _, file := range d.includedFiles {
		fmt.Print(d.GetName() + "/")
		file.PrintList()
	}
}

func main() {
	documents := &Directory{name: "DOCUMENTS"}
	pictures := &Directory{name: "PICTURES"}
	videos := &Directory{name: "VIDEOS"}
	rarely := &Directory{name: "RARELY"}

	belajarJava := &File{name: "belajarJava.pdf"}
	belajarGo := &File{name: "belajarGo.pdf"}
	belajarElixir := &File{name: "belajarElixir.pdf"}

	sammi := &File{name: "sammi.jpg"}
	izzah := &File{name: "izzah.jpg"}

	goCrashCourse := &File{"GO-CRASH-COURSE.mp4"}
	KotlinCrashCourse := &File{"KOTLIN-CRASH-COURSE.mp4"}

	documents.Add(belajarJava)
	documents.Add(belajarGo)
	documents.Add(belajarElixir)

	pictures.Add(sammi)
	pictures.Add(izzah)

	rarely.AddAll(belajarJava, belajarElixir, KotlinCrashCourse)

	videos.Add(goCrashCourse)
	videos.Add(KotlinCrashCourse)

	documents.PrintList()
	videos.PrintList()
	pictures.PrintList()
	rarely.PrintList()
}

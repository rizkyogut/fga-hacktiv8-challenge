package main

import (
	"fmt"
	"os"
	"strconv"
)

type studentData struct {
	no         int
	name       string
	address    string
	occupation string
	reason     string
}

func printName(classroom studentData) {
	fmt.Println("Nama: ", classroom.name)
	fmt.Println("Alamat: ", classroom.address)
	fmt.Println("Pekerjaan: ", classroom.occupation)
	fmt.Println("Alasan: ", classroom.reason)
}

func main() {
	var students = []studentData{
		{
			no:         1,
			name:       "Anggela",
			address:    "Jl. Kerja",
			occupation: "Hukum",
			reason:     "Tantangan",
		},
		{
			no:         2,
			name:       "Martis",
			address:    "Jl. Bakti",
			occupation: "Mahasiswa",
			reason:     "Cita - Cita",
		},
	}

	studentId := os.Args[1]
	index, _ := strconv.Atoi(studentId)
	classroom := students[index-1]
	printName(classroom)
}

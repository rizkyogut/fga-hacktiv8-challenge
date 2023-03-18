package main

import (
	"fmt"
	"os"
	"strconv"
)

// buat struct baru isinya field: no, name, address, occupation, reason
type student struct {
	no         int
	name       string
	address    string
	occupation string
	reason     string
}

// buat struct baru: nama field bebas tipe datanya slice studentData
type classroom struct {
	students []student
}

// buat interface isinya:
// getName, getAddress, getOccupation, getReason
type biodata interface {
	getName(no int) string
	getAddress(no int) string
	getOccupation(no int) string
	getReason(no int) string
	// print(i int, b biodata)
}

// implement method2 interface ke struct list studentData
func (c classroom) getName(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.name
		}
	}
	return ""
}

func (c classroom) getAddress(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.address
		}
	}
	return ""
}

func (c classroom) getOccupation(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.occupation
		}
	}
	return ""
}

func (c classroom) getReason(no int) string {
	for _, student := range c.students {
		if no == student.no {
			return student.reason
		}
	}
	return ""
}

func print(i int, b biodata) {
	fmt.Println("Nama\t: ", b.getName(i))
	fmt.Println("Alamat\t: ", b.getAddress(i))
	fmt.Println("Pekerjaan\t: ", b.getOccupation(i))
	fmt.Println("Alasan\t: ", b.getReason(i))
}

// os.Args -> buat ambil argument di command line

func main() {
	studentId := os.Args[1]

	// convert student_id pake strconv.Atoi
	i, _ := strconv.Atoi(studentId)

	// inisialisasi var classroom
	var cr classroom

	// inisialisasi slice studentData
	var murid = []student{
		{
			no:         1,
			name:       "nama1",
			address:    "alamat1",
			occupation: "wirausahawan",
			reason:     "jati diri",
		},
		{
			no:         2,
			name:       "nama2",
			address:    "alamat2",
			occupation: "pegawai",
			reason:     "cari jodoh",
		},
	}

	// murid diassign ke cr.students
	cr.students = murid

	print(i, cr)
}

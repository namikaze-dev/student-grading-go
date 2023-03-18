package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	rd, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	csvRd := csv.NewReader(rd)
	records, err := csvRd.ReadAll()
	if err != nil && err != io.EOF {
		panic(err)
	}

	var students []student
	for _, row := range records[1:] {
		var student student
		student.firstName = row[0]
		student.lastName = row[1]
		student.university = row[2]
		student.test1Score = parseInt(row[3])
		student.test2Score = parseInt(row[4])
		student.test3Score = parseInt(row[5])
		student.test4Score = parseInt(row[6])
		students = append(students, student)
	}

	return students
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func calculateGrade(students []student) []studentStat {
	return nil
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}

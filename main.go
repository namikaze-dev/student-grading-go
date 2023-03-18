package main

import (
	"encoding/csv"
	"io"
	"os"
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

func calculateGrade(students []student) []studentStat {
	var stats []studentStat

	for _, s := range students {
		stat := studentStat{
			student:    s,
			finalScore: avg(s),
		}

		stat.grade = grade(stat.finalScore)
		stats = append(stats, stat)
	}

	return stats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	var topper studentStat
	for _, s := range gradedStudents {
		if s.finalScore > topper.finalScore {
			topper = s
		}
	}

	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	var toppers = map[string]studentStat{}
	for _, s := range gs {
		if _, ok := toppers[s.university]; ok {
			toppers[s.university] = max(s, toppers[s.university])
		} else {
			toppers[s.university] = s
		}
	}
	return toppers
}

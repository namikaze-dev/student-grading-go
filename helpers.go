package main

import "strconv"

func avg(s student) float32 {
	sum := s.test1Score + s.test2Score + s.test3Score + s.test4Score
	return float32(sum) / 4
}

func grade(score float32) Grade {
	if score < 35 {
		return F
	} else if score < 50 {
		return C
	} else if score < 70 {
		return B
	} else {
		return A
	}
}

func parseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

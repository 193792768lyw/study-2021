package main

import (
	"fmt"
	"strings"
)

type IK interface {
	name()
	get()
}

type Ik2 struct {
}

func (receiver Ik2) name() {
	fmt.Println("ik1")
}
func (receiver Ik2) get() {
	fmt.Println("ik1")
}

type Ik1 struct {
	IK
}

func (receiver Ik1) name() {
	fmt.Println("ik777")
}

func main() {
	var dd IK = Ik1{Ik2{}}
	dd.name()

	//context.WithValue()
	//context.Background()
	fmt.Println(maximumTime("0?:3?"))
}
func maximumTime1(time string) string {
	res := strings.Builder{}
	for i := 0; i < 5; i++ {
		if time[i] != '?' {
			res.WriteByte(time[i])
			continue
		}
		if i == 0 && time[i] == '?' {
			if time[i+1] == '?' {
				res.WriteString("23:")
				i = 2
			} else if time[i+1] <= '3' {
				res.WriteString("2")
			} else {
				res.WriteString("1")
			}
			continue
		} else if i == 1 && time[i] == '?' {
			if time[i-1] == '2' {
				res.WriteByte('3')
			} else {
				res.WriteByte('9')
			}
		} else if i == 3 && time[i] == '?' {
			res.WriteByte('5')
		} else if i == 4 && time[i] == '?' {
			res.WriteByte('9')
		}

	}

	return res.String()
}

func maximumTime(time string) string {
	t := []byte(time)
	if t[0] == '?' {
		if '4' <= t[1] && t[1] <= '9' {
			t[0] = '1'
		} else {
			t[0] = '2'
		}
	}
	if t[1] == '?' {
		if t[0] == '2' {
			t[1] = '3'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}

package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name   string
	Age    int
	Height int
}

type ByField struct {
	People  []*Person
	Reverse bool
	Field   string
}

func (p ByField) Len() int {
	return len(p.People)
}

func (p ByField) Less(i, j int) bool {
	if p.Reverse {
		switch p.Field {
		case "age":
			return p.People[i].Age > p.People[j].Age // 倒排按照年龄
		case "height":
			return p.People[i].Height > p.People[j].Height // 倒排按照身高
		}
	} else {
		switch p.Field {
		case "age":
			return p.People[i].Age < p.People[j].Age // 正排按照年龄
		case "height":
			return p.People[i].Height < p.People[j].Height // 正排按照身高
		}
	}
	return false
}

func (p ByField) Swap(i, j int) {
	p.People[i], p.People[j] = p.People[j], p.People[i]
}

func main() {
	people := []*Person{
		{"Alice", 25, 160},
		{"Bob", 30, 170},
		{"Charlie", 20, 150},
	}

	// 正排按照年龄
	sort.Sort(ByField{People: people, Reverse: false, Field: "age"})
	//fmt.Println("正排按照年龄：", people)
	fmt.Println(people[0].Name, people[0].Age, people[1].Name, people[1].Age, people[2].Name, people[2].Age)

	// 倒排按照年龄
	sort.Sort(ByField{People: people, Reverse: true, Field: "age"})
	//fmt.Println("倒排按照年龄：", people)
	fmt.Println(people[0].Name, people[0].Age, people[1].Name, people[1].Age, people[2].Name, people[2].Age)

	// 正排按照身高
	sort.Sort(ByField{People: people, Reverse: false, Field: "height"})
	//fmt.Println("正排按照身高：", people)
	fmt.Println(people[0].Name, people[0].Age, people[1].Name, people[1].Age, people[2].Name, people[2].Age)

	// 倒排按照身高
	sort.Sort(ByField{People: people, Reverse: true, Field: "height"})
	//fmt.Println("倒排按照身高：", people)
	fmt.Println(people[0].Name, people[0].Age, people[1].Name, people[1].Age, people[2].Name, people[2].Age)

}

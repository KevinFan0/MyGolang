package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct {
	Title		string
	Year		int			`json:"released"`
	Color		bool		`json:"Color, omitempty"`
	Actors		[]string
}


var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
	Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
	Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
	Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

// 将一个Go语言中类似movies的结构体slice转为JSON的过程叫编组（marshaling）。编组通过调用json.Marshal函数完成
func jsonMarshal() {
	_, err := json.Marshal(movies)
	data2, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data2)
}

func main()  {
	jsonMarshal()
}
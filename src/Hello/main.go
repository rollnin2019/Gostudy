package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	var testmap map[string]string
	testmap = make(map[string]string)
	testmap["France"] = "paris"
	testmap["Italy"] = "Rome"
	testmap["Japan"] = "Tokyo"

	for i := range testmap {
		fmt.Println(i, "首都是", testmap[i])
	}

	delete(testmap, "Japan")

	for i := range testmap {
		fmt.Println(i, "首都是", testmap[i])
	}
}

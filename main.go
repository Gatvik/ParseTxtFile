package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Struct that describes data from txt file
type Human struct {
	Name string
	Age  int
}

func main() {
	var file = getFile("text.txt")
	var result = getHumanSlice(file)
	for _, v := range result {
		fmt.Printf("Your name is %s and you %d years old\n", v.Name, v.Age)
	}

}

func getFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		panic("Can't read the file")
	}

	return file
}

func getHumanSlice(file []byte) []Human {
	var lines = strings.Fields(string(file))
	var result = []Human{}
	for i := 0; i < len(lines); i++ {
		var line = strings.FieldsFunc(lines[i], splitBySemicolon)
		age, err := strconv.Atoi(line[1])
		if err != nil {
			age = 0
		}
		result = append(result, Human{Name: line[0], Age: age})
	}

	return result
}

func splitBySemicolon(c rune) bool {
	return string(c) == ";"
}

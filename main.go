package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/yarikbratashchuk/servigo_openai_functions/functions"
)

func main(){
	fmt.Println("------------------------------------------")
	fmt.Println("|        Servigo OpenAI Functions        |")
	fmt.Println("------------------------------------------")
	fmt.Println()

	funcs := functions.CreateFunctionList()
	bts, _ := json.Marshal(funcs)

	fmt.Println(string(bts))

	for i := range funcs {
		fmt.Println()
		fmt.Println("------------------------------------------")
		spacesCount := (40 - len(funcs[i].Name)) / 2
		if spacesCount < 1 {
			spacesCount = 0
		}
		spaces := strings.Repeat(" ", spacesCount)
		fmt.Printf("|%s%s%s|\n", spaces, funcs[i].Name, spaces)
		fmt.Println("------------------------------------------")
		fmt.Println()
		bts, _ := json.Marshal(funcs[i])
		fmt.Println(string(bts))
		fmt.Println()
	}
}
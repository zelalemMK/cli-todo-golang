package main

import (
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
)

var (
	data = [][]interface{}{
		{1, "Talk with Asfaw", true},
		{2, "Email R. Edney", true},
		{3, "Go to Jackson's house", false},
		{4, "Ron games", false},
	}
)

func main() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task Name"},
			{Align: simpletable.AlignCenter, Text: "Completed"},
		},
	}

	for _, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
			{Text: row[1].(string)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("%t", row[2])},
		}
		table.Body.Cells = append(table.Body.Cells, r)

	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())

	var input string
	for {
		fmt.Println("Print 1 to exit")
		fmt.Scanln(&input)

		if input == "1" {
			os.Exit(0)
		}
	}
}

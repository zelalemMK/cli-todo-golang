package main

import (
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
)

var (
	data = [][]interface{}{
		{1, "Newton G. Goetz", 543.4},
		{2, "Rebecca R. Edney", 1423.25},
		{3, "John R. Jackson", 7526.12},
		{4, "Ron J. Gomes", 123.84},
		{5, "Penny R. Lewis", 3221.11},
	}
)

func main() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "TAX"},
		},
	}

	subtotal := float64(0)
	for _, row := range data {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", row[0].(int))},
			{Text: row[1].(string)},
			{Align: simpletable.AlignLeft, Text: fmt.Sprintf("$ %.2f", row[2].(float64))},
		}
		table.Body.Cells = append(table.Body.Cells, r)
		subtotal += row[2].(float64)

	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{},
			{Align: simpletable.AlignRight, Text: "Subtotal"},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("$ %.3f", subtotal)},
		},
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

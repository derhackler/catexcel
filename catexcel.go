package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	var worksheet string
	flag.StringVar(&worksheet, "sheet", "", "Name of the worksheet to print. First sheet by default.")

	flag.Parse()
	filename := flag.Arg(0)

	if filename == "" {
		printHelp()
		os.Exit(1)
	} else {
		outputexcel(filename, worksheet)
	}
}

func printHelp() {
	fmt.Println("Usage: catexcel [-sheet SHEET] FILE")
	fmt.Println("Prints the worksheet SHEET in the Excel (xslx) file FILE to STDOUT.")
	fmt.Println()
	fmt.Println("With no SHEET prints the first sheet in the Excel file")
	fmt.Println()
	fmt.Println("Example 1: You have an Excel with 'firstname,lastname,haircolor,birthday'. ")
	fmt.Println("           Search for everyone who is called Smith and has brown hair.")
	fmt.Println("           Print the first name and birthday:")
	fmt.Println("           catexcel example.xslx | awk '/Smith\\tbrown/' | cut -f 1,4")
	fmt.Println()
	fmt.Println("Source: https://github.com/derhackler/catexcel")
}

func outputexcel(file string, worksheet string) {
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		log.Fatal(err)
	}

	if worksheet == "" {
		worksheet = xlsx.GetSheetName(1) // The first sheets index is 1 and not 0
	}

	// Get all the rows in the provided sheet.
	rows := xlsx.GetRows(worksheet)
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

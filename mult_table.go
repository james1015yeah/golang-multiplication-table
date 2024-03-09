package main

import "fmt"

func main() {

	var cols int
	var rows int

	fmt.Println("************************** ")
	fmt.Println("	Multiplication Table	")
	fmt.Println("************************** ")
	fmt.Print("Enter no. of columns: ")
	fmt.Scanln(&cols)
	fmt.Print("Enter no. of rows: ")
	fmt.Scanln(&rows)
	fmt.Println("************************** ")

	if rows > 9 || rows < 2 {
		fmt.Println("rows capacity not met[2-9]")
	} else if cols > 9 || cols < 2 {
		fmt.Println("columns capacity not met[2-9]")
	} else {
		for i := 1; (i <= rows) && (i <= 9); i++ {
			for j := 1; (j <= cols) && (j <= 9); j++ {
				fmt.Print(" ", j*i) //print column values
			}
			fmt.Println(" ") //print new row
		}
	}
}

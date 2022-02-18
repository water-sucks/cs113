package main

import (
	"bufio"
	"fmt"

	// "math"
	"os"
	"strconv"
	"strings"
)

const menu = `MAIN MENU
=========
0 - Show this menu
1 - NOT gate
2 - AND gate
3 - OR gate
4 - NAND gate
5 - NOR gate
6 - Half-adder
7 - Full-adder
8 - Exit`

func main() {
	fmt.Println(menu)

main:
	for {
		var input int
		input = GetInt()
		if input < 0 || input > 8 {
			fmt.Fprintln(os.Stderr, "Number must be between 0 and 8!")
			continue
		}

		var table [][]int
		var title string
		var titles []string

		switch input {
		case 0:
			fmt.Println(menu)
			continue
		case 1:
			table = BuildNotGate()
			title = "NOT Gate"
			titles = []string{"p", "~p"}
		case 2:
			table = BuildAndGate()
			title = "AND Gate"
			titles = []string{"p", "q", "p&q"}
		case 3:
			table = BuildOrGate()
			title = "OR Gate"
			titles = []string{"p", "q", "p|q"}
		case 4:
			table = BuildNandGate()
			title = "NAND Gate"
			titles = []string{"p", "q", "p~&q"}
		case 5:
			table = BuildNorGate()
			title = "NOR Gate"
			titles = []string{"p", "q", "p~|q"}
		case 6:
			table = BuildHalfAdder()
			title = "Half-adder"
			titles = []string{"p", "q", "c", "s"}
		case 7:
			table = BuildFullAdder()
			title = "Full-adder"
			titles = []string{"p", "q", "r", "c", "s"}
		case 8:
			break main
		default:
			panic("Impossible state reached")
		}

		PrintTable(table, title, titles)
	}

	fmt.Println("Goodbye!")
}

// Using a BufReader for better parsing of integers
// than is provided with Go's stdlib fmt.Scanf
func GetInt() int {
	for {
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Unexpected EOF!")
			continue
		}

		input := strings.TrimSpace(raw)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse integer")
			continue
		}

		return num
	}
}

// Skeleton for NOT gate
var skeleton2x2 = [][]int{
	{1, 0},
	{0, 0},
}

// Skeleton for AND, OR, NAND, NOR gates
var skeleton4x3 = [][]int{
	{1, 1, 0},
	{1, 0, 0},
	{0, 1, 0},
	{0, 0, 0},
}

// Skeleton for half-adder
var skeleton4x4 = [][]int{
	{1, 1, 0, 0},
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 0, 0},
}

// Skeleton for full-adder
var skeleton8x5 = [][]int{
	{1, 1, 1, 0, 0},
	{1, 1, 0, 0, 0},
	{1, 0, 1, 0, 0},
	{1, 0, 0, 0, 0},
	{0, 1, 1, 0, 0},
	{0, 1, 0, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
}

//// Original functions used for dynamically generating skeletons, not used
//// because professor prefers hardcoding skeletons here for readability
// // Get binary representation of number with given number of bits
// func BinaryRepresentation(num int, bits int) []int {
// 	representation := make([]int, bits)
//
// 	if num == 0 {
// 		return representation
// 	}
//
// 	for i := range representation {
// 		if bit := num & (1 << (bits - 1 - i)); bit == 0 {
// 			representation[i] = 0
// 		} else {
// 			representation[i] = 1
// 		}
// 	}
//
// 	return representation
// }
//
// // Builds truth table skeleton with inputs filled in
// func MakeTable(combinations int, columns int) [][]int {
// 	bits := int(math.Log2(float64(combinations)))
//
// 	numbers := make([][]int, combinations)
//
// 	for i := range numbers {
// 		numbers[i] = make([]int, columns)
// 		representation := BinaryRepresentation(i, bits)
// 		for j, r := range representation {
// 			numbers[i][j] = r
// 		}
// 	}
//
// 	return numbers
// }

func BuildNotGate() [][]int {
	numbers := make([][]int, 2)
	copy(numbers, skeleton2x2)

	for i := range numbers {
		numbers[i][1] = numbers[i][0] ^ 1
	}

	return numbers
}

func BuildAndGate() [][]int {
	numbers := make([][]int, 4)
	copy(numbers, skeleton4x3)

	for i := range numbers {
		numbers[i][2] = numbers[i][0] & numbers[i][1]
	}

	return numbers
}

func BuildOrGate() [][]int {
	numbers := make([][]int, 4)
	copy(numbers, skeleton4x3)

	for i := range numbers {
		numbers[i][2] = numbers[i][0] | numbers[i][1]
	}

	return numbers
}

func BuildNandGate() [][]int {
	numbers := make([][]int, 4)
	copy(numbers, skeleton4x3)

	for i := range numbers {
		numbers[i][2] = (numbers[i][0] & numbers[i][1]) ^ 1
	}

	return numbers
}

func BuildNorGate() [][]int {
	numbers := make([][]int, 4)
	copy(numbers, skeleton4x3)

	for i := range numbers {
		numbers[i][2] = (numbers[i][0] | numbers[i][1]) ^ 1
	}

	return numbers
}

// Helper for half-adder functions (returns carry, sum respectively)
func HalfAdder(p int, q int) (int, int) {
	return p & q, p ^ q
}

func BuildHalfAdder() [][]int {
	numbers := make([][]int, 4)
	copy(numbers, skeleton4x4)

	for i := range numbers {
		numbers[i][2], numbers[i][3] = HalfAdder(numbers[i][0], numbers[i][1])
	}

	return numbers
}

func BuildFullAdder() [][]int {
	numbers := make([][]int, 8)
	copy(numbers, skeleton8x5)

	for i := range numbers {
		c1, s1 := HalfAdder(numbers[i][0], numbers[i][1]) // First half-adder
		c2, s := HalfAdder(s1, numbers[i][2])             // Second half-adder
		c := c1 | c2                                      // Final carry bit
		numbers[i][3] = c
		numbers[i][4] = s
	}

	return numbers
}

func PrintTitleCell(content string) {
	fmt.Print("| ", content, " ")
}

func PrintCell(content string, length int) {
	spaces := strings.Repeat(" ", (length/2)+1)
	preSpaces := spaces
	if length%2 == 0 {
		preSpaces = strings.Repeat(" ", (length / 2))
	}
	fmt.Print("|", preSpaces, content, spaces)
}

func PrintTable(table [][]int, title string, titles []string) {
	// Divider
	var tableLength int
	for _, t := range titles {
		tableLength += len(t)
	}
	tableLength += (3 * len(titles)) + 1
	divider := strings.Repeat("-", tableLength)

	// Table header
	spaces := strings.Repeat(" ", ((tableLength-len(title))/2)-1)
	fmt.Println(spaces, title, spaces)
	fmt.Println(divider)

	// Title cells
	for _, t := range titles {
		PrintTitleCell(t)
	}

	fmt.Print("|\n")
	fmt.Println(divider)

	// Number cells
	for _, r := range table {
		for i, n := range r {
			PrintCell(fmt.Sprint(n), len(titles[i]))
		}
		fmt.Print("|\n")
		fmt.Println(divider)
	}
}

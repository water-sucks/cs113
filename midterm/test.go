package main

// Varun Narravula: CS-113-02: Midterm Part 2
//
// This program creates a truth table for a 3-bit adder and prints it.
// It takes no inputs, and instead has a set of hardcoded inputs
// (binary representations of the numbers 0-7 as a ThreeBitInput type)
// that are then multipled together with itself as a Cartesian product
// to generate the inputs. There are two helpers (HalfAdder and FullAdder)
// that are used to create the actual four-bit result in BuildRow, which takes
// the two inputs from one resulting array of the Cartesian product
// and returns a TruthTableRow with the two inputs and the FourBitOutput
// computed from the ThreeBitAdder function. This row type has a method to
// pretty-print the result, which is used in a for loop to print out each
// row to the console from the result of the BuildTable function, which returns
// an array of rows.

import "fmt"

func main() {
	fmt.Println("            3-Bit Adder Truth Table         ")
	fmt.Println("--------------------------------------------")
	fmt.Println("| P | Q | R | S | T | U | C1 | C2 | C3 | S |")
	fmt.Println("--------------------------------------------")
	for _, e := range BuildTable() {
		e.print()
	}
}

type ThreeBitInput struct {
	first  int
	second int
	third  int
}

type FourBitOutput struct {
	first  int
	second int
	third  int
	carry  int
}

type TruthTableRow struct {
	pqr  ThreeBitInput
	stu  ThreeBitInput
	wxyz FourBitOutput
}

// Source to create Cartesian products of
var inputs = []ThreeBitInput{
	{
		first:  0,
		second: 0,
		third:  0,
	},
	{
		first:  0,
		second: 0,
		third:  1,
	},
	{
		first:  0,
		second: 1,
		third:  0,
	},
	{
		first:  0,
		second: 1,
		third:  1,
	},
	{
		first:  1,
		second: 0,
		third:  0,
	},
	{
		first:  1,
		second: 0,
		third:  1,
	},
	{
		first:  1,
		second: 1,
		third:  0,
	},
	{
		first:  1,
		second: 1,
		third:  1,
	},
}

// Returns sum, carry in order
func HalfAdder(p int, q int) (int, int) {
	return p ^ q, p & q
}

// Returns sum, carry in order
func FullAdder(p int, q int, r int) (int, int) {
	s1, c1 := HalfAdder(p, q)
	s, c2 := HalfAdder(s1, r)
	c := c1 | c2
	return s, c
}

// Three bit adder that returns a four bit output
func ThreeBitAdder(x ThreeBitInput, y ThreeBitInput) FourBitOutput {
	s1, c1 := HalfAdder(x.third, y.third)
	s2, c2 := FullAdder(c1, x.second, y.second)
	s3, c := FullAdder(c2, x.first, y.first)

	return FourBitOutput{
		first:  s1,
		second: s2,
		third:  s3,
		carry:  c,
	}
}

// Cartesian product function (only works for two inputs of ThreeBitInput type)
func CartesianProduct(one []ThreeBitInput, other []ThreeBitInput) [][]ThreeBitInput {
	output := make([][]ThreeBitInput, len(one)*len(other))
	i := 0
	for _, x := range one {
		for _, y := range other {
			output[i] = []ThreeBitInput{x, y}
			i++
		}
	}
	return output
}

// Simple helper for table creation that creates a truth table row
// given two inputs
func BuildRow(x ThreeBitInput, y ThreeBitInput) TruthTableRow {
	return TruthTableRow{
		pqr:  x,
		stu:  y,
		wxyz: ThreeBitAdder(x, y),
	}
}

func BuildTable() []TruthTableRow {
	inputs := CartesianProduct(inputs, inputs)
	output := make([]TruthTableRow, len(inputs))
	for i := range output {
		output[i] = BuildRow(inputs[i][0], inputs[i][1])
	}
	return output
}

// Obscenely long one-liner, but it prints a row of the truth table
// in order: pqr, then stu, then wxyz's bits in table form for each
// row.
func (r TruthTableRow) print() {
	fmt.Println("|", r.pqr.first, "|", r.pqr.second, "|", r.pqr.third, "|", r.stu.first, "|", r.stu.second, "|", r.stu.third, "|", r.wxyz.first, " |", r.wxyz.second, " |", r.wxyz.third, " |", r.wxyz.carry, "|")
	fmt.Println("--------------------------------------------")
}

package parsers

import (
	"drum-hero/models"
	"fmt"
	"strconv"
)

const (
	duplicateTolerance = 110000
)

// This turns a set of hits from a user's
// input into a Pattern that can then be
// evaluated using the Evaluator from a
// Round struct. It does three things: it
// turns the string output from the Arduino
// into numbers, removes inputs that are below
// a certain threshold (these are duplicates
// that need to be sanitized), and then computes
// intervals on 8th note intervals to then compare
// and set to true based on which intervals are
// are nearest to each number.
func ParseHits(inputs []string, bpm models.BPM, targetLength int) (models.Pattern, error) {
	hits := make([]int, len(inputs))
	for i, r := range inputs {
		time, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println("Unable to convert time from Arduino to integer")
			return nil, err
		}

		hits[i] = time
	}

	hits = Deduplicate(hits, duplicateTolerance)
	hits = hits[:targetLength-1]

	interval := int(bpm.ToMicrosecondDuration(1) / 2)

	// Hypothetical input that contains perfect hits on every
	// eighth note interval
	hypothetical := make([]int, targetLength)
	hypothetical[0] = hits[0]
	for i := 1; i < targetLength; i++ {
		hypothetical[i] = hypothetical[0] + (interval * i)
	}

	pattern := make(models.Pattern, targetLength)

	for _, r := range hits {
		index := FindNearestIndex(hypothetical, r)

		pattern[index] = true
	}

	return pattern, nil
}

// Remove inputs that have a difference smaller
// than the tolerance. The Arduino does detect
// duplicates if hit hard enough to cause vibrations
// in the right way.
func Deduplicate(inputs []int, tolerance int) []int {
	deduplicated := []int{inputs[0]}
	for i := 1; i < len(inputs); i++ {
		if inputs[i]-inputs[i-1] >= tolerance {
			deduplicated = append(deduplicated, inputs[i])
		}
	}
	return deduplicated
}

// Given a number, find the index of the number closest
// to it in the given array. This allows for deciding which
// inputs are set to true in the user-created Pattern.
func FindNearestIndex(inputs []int, num int) int {
	index := 0
	for i, n := range inputs {
		if Abs(num-n) < Abs(num-inputs[index]) {
			index = i
		}
	}
	return index
}

// This is only implemented for floats in the standard library,
// so I'm creating an int version of absolute value here.
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

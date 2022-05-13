package models

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Cleaner alias for pattern inputs
type Pattern []bool

// Dynamic truth table function defined for
// each round's correct input on the fly.
// This is the truth table in the form of
// a boolean expression that is generated
// for each pattern; since this is a simple
// function, it saves space as compared to
// storing a 64x65 slice of slices for an
// 8-measure pattern, for example. There is
// only one critical row in this truth table,
// so it is quite easy to evaluate in this way.
type Evaluator func(Pattern) bool

type Round struct {
	Settings  Settings
	Pattern   Pattern
	Evaluator Evaluator
}

// Alias for speed at which to record hits
type BPM int

// Generate a random Pattern and a corresponding
// boolean Evaluator for it based on specified
// settings; only generates 4/4 bars and
func CreateRound(s Settings) Round {
	length := s.Measures * 8

	// Number of notes to generate should
	// not be zero or too little to make
	// a coherent musical pattern.
	min := length / 2
	max := length / (8 / 7)

	totalHits := rand.Intn(max-min) + min

	hits := rand.Perm(length - 1)
	hits = hits[0:totalHits]

	pattern := make(Pattern, length)

	// Make first note always true to have
	// a reference point when parsing
	pattern[0] = true

	for _, i := range hits {
		pattern[i] = true
	}

	return Round{
		Settings: s,
		Pattern:  pattern,
		Evaluator: func(input Pattern) bool {
			// If any input does not match in `pattern`,
			// then this will turn false. This is how
			// a pattern can be graded.
			acc := true
			for i, h := range pattern {
				if h {
					// Compute and with input
					// (should be 1 to be correct)
					acc = acc && input[i]
				} else {
					// Compute and with inverse of input
					// (should be 0 to be correct)
					acc = acc && !input[i]
				}
			}
			return acc
		},
	}
}

// Convert a BPM and number of repetitions
// to a `time.Duration` type to use for limit
// on time when scanning for input
func (b BPM) ToDuration(times int) time.Duration {
	return time.Duration(float64(times) * float64(time.Second) * (60.0 / float64(b)))
}

// Same as ToDuration, but returns an integer that is
// used to compute the interval needed to parse the
// microseconds output from the Arduino into a Pattern
func (b BPM) ToMicrosecondDuration(times int) float64 {
	return ((60.0 * 1e6) / float64(b)) * float64(times)
}

// Difficulty-BPM mapping; this can be extended.
// if desired.
func (d Difficulty) ToBPM() BPM {
	switch d {
	case Easy:
		return 90
	case Normal:
		return 120
	case Hard:
		return 150
	default:
		return 120
	}
}

// Print a sheet music representation of a Pattern
// to the screen for the user to view and play; this
// is hardcoded in such a way that it can only print
// 4/4 bars that contain at max 8 eighth notes on
// separate lines, but this can be extended.
func (p Pattern) PrettyPrint() {
	for j := 0; j < len(p); j += 8 {
		chunk := p[j:(j + 8)]
		representation := make([][]string, len(chunk))

		for i := 0; i < len(chunk)/2; i++ {
			var section [2]bool
			copy(section[:], chunk[(i*2):(i*2)+2])
			switch section {
			case [2]bool{false, false}:
				for i := range quarterNoteRest {
					representation[i] = append(representation[i], quarterNoteRest[i])
				}
			case [2]bool{true, false}:
				for i := range quarterNote {
					representation[i] = append(representation[i], quarterNote[i])
				}
			case [2]bool{false, true}:
				for i := range eighthNoteAnd {
					representation[i] = append(representation[i], eighthNoteAnd[i])
				}
			default:
				for i := range twoEighthNotes {
					representation[i] = append(representation[i], twoEighthNotes[i])
				}
			}
		}

		for _, r := range representation {
			fmt.Println(strings.Join(r, ""))
		}
		fmt.Println()
	}
}

// String slice representations for each
// possible musical symbol
var (
	quarterNoteRest = []string{
		`     \\         `,
		`      \\        `,
		`       ///      `,
		`      ///       `,
		`       \\       `,
		`      /  \      `,
		`     //   \     `,
	}
	quarterNote = []string{
		`        *|    `,
		`        *     `,
		`        *     `,
		`        *     `,
		`        *     `,
		`     ****     `,
		`     ****     `,
	}
	eighthNoteAnd = []string{
		`                ****|    `,
		`                *        `,
		`      __        *        `,
		`     |__|       *        `,
		`        |       *        `,
		`     __/     ****        `,
		`             ****        `,
	}
	twoEighthNotes = []string{
		`       __________     `,
		`       *        *     `,
		`       *        *     `,
		`       *        *     `,
		`       *        *     `,
		`    ****     ****     `,
		`    ****     ****     `,
	}
)

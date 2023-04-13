package experiment

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type freqMapEntry struct {
	numSamples int
	relFreq    float64
	numStars   int
}

// Experiment holds all the relevant information
type Experiment struct {
	numThrows int
	numDice   int
	numFaces  int
	freqMap   map[int]freqMapEntry
}

// New returns a pointer to an initialized Experiment with the given data
func New(numThrows int, numDice int, numFaces int) *Experiment {
	e := Experiment{
		numThrows: numThrows,
		numDice:   numDice,
		numFaces:  numFaces,
		freqMap:   make(map[int]freqMapEntry, (numFaces-1)*numDice+numFaces),
	}

	return &e
}

// Run executes the simulation
func (e *Experiment) Run() (error) {
	for i := 0; i < e.numThrows; i++ {
		throwResult := e.throw()
		entry := e.freqMap[throwResult]
		entry.numSamples++
		e.freqMap[throwResult] = entry
	}
	err := e.stats()
	if err != nil {
		return fmt.Errorf("cannot run experiment: %w", err)
	}
	return nil
}

// throw simulates the throwing of the given number of dice
func (e *Experiment) throw() (result int) {
	for i := 0; i < e.numDice; i++ {
		result += rand.Intn(e.numFaces) + 1
	}
	return
}

// stats populates the statistics of an experiment after it has been run
func (e *Experiment) stats() (error) {

	// get the total number of samples and max frequency found
	var totalNumSamples int
	var maxFreq int
	for _, v := range e.freqMap {
		totalNumSamples += v.numSamples
		if v.numSamples > maxFreq {
			maxFreq = v.numSamples
		}
	}

	if totalNumSamples != e.numThrows {
		return fmt.Errorf("unexpected number of samples %v didn't match %v", totalNumSamples, e.numThrows)
	}

	// calculate the relative frequency and associated number of '*'
	maxStars := 10
	for i := e.numDice; i <= e.numFaces*e.numDice; i++ {
		entry := e.freqMap[i]
		entry.relFreq = float64(e.freqMap[i].numSamples) / float64(totalNumSamples)
		entry.numStars = int(math.Round(float64(maxStars) / float64(maxFreq) * float64(e.freqMap[i].numSamples)))
		e.freqMap[i] = entry
	}
	return nil
}

// Show will pretty print the results of the experiment
func (e *Experiment) Show() {

	fmt.Println("=== Experiment results ===")
	fmt.Printf("number of throws: %v\n", e.numThrows)
	fmt.Printf("number of faces in dice: %v\n", e.numFaces)
	fmt.Printf("number of dice: %v\n", e.numDice)

	// Length of the number obtained in the throw
	col1Len := int(math.Log10(float64(e.numFaces*e.numDice)) + 1)

	// maxFreq computation
	var maxFreq int
	for _, v := range e.freqMap {
		if v.numSamples > maxFreq {
			maxFreq = v.numSamples
		}
	}

	// Length of the number of times that number happened
	col2Len := math.Floor(math.Log10(float64(maxFreq)) + 1)

	sb := strings.Builder{}

	// number obtained in the throw
	sb.WriteString("%")
	sb.WriteString(fmt.Sprint(col1Len))
	sb.WriteString("d : ")

	// number of times that value happened in the throw
	sb.WriteString("%")
	sb.WriteString(fmt.Sprint(col2Len))
	sb.WriteString("d ")

	// relative frequency
	sb.WriteString("(%.3f) : ")


	// associated number of strings (will be calculated for every line)
	sb.WriteString("%s")

	sb.WriteString("\n")

	// number of times that number happened


	format := sb.String()

	for i := e.numDice; i <= e.numFaces * e.numDice; i++ {
		fmt.Printf(format, i, e.freqMap[i].numSamples, e.freqMap[i].relFreq, strings.Repeat("*", e.freqMap[i].numStars))
	}
}

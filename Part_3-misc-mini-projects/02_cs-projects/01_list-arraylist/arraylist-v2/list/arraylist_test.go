package list_test

import (
	"fmt"
	"math/rand"
	"testing"
	"testing/quick"

	"example.com/arraylist/list"
	"golang.org/x/exp/slices"
)

func TestNew(t *testing.T) {
	l := list.New[string]()
	if !l.Empty() {
		t.Errorf("expected list to be empty, but wasn't")
	}
}

func TestInsertSadPath(t *testing.T) {
	tests := map[string]struct {
		pos         int
		expectedErr error
	}{
		"Insert elem in neg pos": {
			pos:         -1,
			expectedErr: list.ErrPosOutOfRange,
		},
		"Insert elem beyond length": {
			pos:         1,
			expectedErr: list.ErrPosOutOfRange,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			l := list.New[string]()
			err := l.Insert("a", testCase.pos)
			assertError(t, err, testCase.expectedErr)
		})
	}
}

func TestInsertHappyPath(t *testing.T) {
	tests := map[string]struct {
		prevList   []string // state of the list before test (as slice)
		elem       string   // elem to insert
		pos        int    	// pos on which elem will be inserted
		wantedList []string // state of the list after test (as slice)
	}{
		"Insert elem on empty list": {
			prevList:   nil,
			elem:       "a",
			pos:        0,
			wantedList: []string{"a"},
		},
		"Insert elem on list with one elem, pos 0": {
			prevList:   []string{"b"},
			elem:       "a",
			pos:        0,
			wantedList: []string{"a", "b"},
		},
		"Insert elem on list with two elems, pos 0": {
			prevList:   []string{"b", "c"},
			elem:       "a",
			pos:        0,
			wantedList: []string{"a", "b", "c"},
		},
		"Insert elem on list with one elem, pos 1": {
			prevList:   []string{"a"},
			elem:       "b",
			pos:        1,
			wantedList: []string{"a", "b"},
		},
		"Insert elem on list with two elems, pos 1": {
			prevList:   []string{"a", "c"},
			elem:       "b",
			pos:        1,
			wantedList: []string{"a", "b", "c"},
		},
		"Insert elem on list with three elems, pos 0": {
			prevList:   []string{"b", "c", "d"},
			elem:       "a",
			pos:        0,
			wantedList: []string{"a", "b", "c", "d"},
		},
		"Insert elem on list with three elems, pos 1": {
			prevList:   []string{"a", "c", "d"},
			elem:       "b",
			pos:        1,
			wantedList: []string{"a", "b", "c", "d"},
		},
		"Insert elem on list with three elems, pos 2": {
			prevList:   []string{"a", "b", "d"},
			elem:       "c",
			pos:        2,
			wantedList: []string{"a", "b", "c", "d"},
		},
		"Insert elem on list with three elems, pos 3": {
			prevList:   []string{"a", "b", "c"},
			elem:       "d",
			pos:        3,
			wantedList: []string{"a", "b", "c", "d"},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			l := toList(t, testCase.prevList)

			err := l.Insert(testCase.elem, testCase.pos)
			assertNoError(t, err)
			assertListContents(t, l, testCase.wantedList)
			assertListLen(t, l.Len(), len(testCase.prevList)+1)
		})
	}
}

func TestInsertProperties(t *testing.T) {

	assertion := func() (err bool) {

		// Set up the initial slice
		numElems := rand.Intn(list.MaxElems - 1)
		initSlice := make([]string, numElems)
		for i := 0; i < numElems; i++ {
			initSlice[i] = string(fmt.Sprintf("%d", rand.Int()))
		}

		// Get a random pos and elem to insert
		pos := rand.Intn(numElems + 1)
		elem := string(fmt.Sprintf("%d", rand.Int()))

		t.Logf("testing insertion of %v on pos %v on a list with %d element(s)", elem, pos, len(initSlice))

		// Get the resulting slice
		endSlice := make([]string, len(initSlice)+1)
		copy(endSlice, initSlice[0:pos])
		endSlice[pos] = elem
		copy(endSlice[pos+1:], initSlice[pos:])

		// Get the resulting list
		l := toList(t, initSlice)
		l.Insert(elem, int(pos))

		// Assert whether both of them match
		return slices.Equal(toSlice(l), endSlice)
	}

	// quick.Check can provide random input to the assertion function
	// but the randomness made the function to discard most of the inputs
	// so I included the random in the function itself
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("Property based tests for Insert failed")
	}
}

func TestRemoveSadPath(t *testing.T) {
	tests := map[string]struct {
		prevList  []string
		pos       int
		wantError error
	}{
		"Remove elem in negative position": {
			prevList: nil,
			pos: -1,
			wantError: list.ErrPosOutOfRange,
		},
		"Remove elem beyond last on empty list": {
			prevList: nil,
			pos: 0,
			wantError:list.ErrPosOutOfRange,
		},
		"Remove elem beyond last on non-empty list": {
			prevList: []string{"a"},
			pos: 1,
			wantError: list.ErrPosOutOfRange,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			l := toList(t, testCase.prevList)
			err := l.Remove(testCase.pos)
			assertError(t, err, testCase.wantError)
		})
	}
}

func TestRemoveHappyPath(t *testing.T) {

	tests := map[string]struct {
		prevList   []string
		pos        int
		wantedList []string
	}{
		"Remove from list with one elem, pos 0": {
			prevList:   []string{"a"},
			pos:        0,
			wantedList: []string{},
		},
		"Remove from list with two elems, pos 0": {
			prevList:   []string{"a", "b"},
			pos:        0,
			wantedList: []string{"b"},
		},
		"Remove from list with two elems, pos 1": {
			prevList:   []string{"a", "b"},
			pos:        1,
			wantedList: []string{"a"},
		},
		"Remove from list with three elems, pos 0": {
			prevList:   []string{"a", "b", "c"},
			pos:        0,
			wantedList: []string{"b", "c"},
		},
		"Remove from list with three elems, pos 1": {
			prevList:   []string{"a", "b", "c"},
			pos:        1,
			wantedList: []string{"a", "c"},
		},
		"Remove from list with three elems, pos 2": {
			prevList:   []string{"a", "b", "c"},
			pos:        2,
			wantedList: []string{"a", "b"},
		},
		"Remove from list with four elems, pos 0": {
			prevList:   []string{"a", "b", "c", "d"},
			pos:        0,
			wantedList: []string{"b", "c", "d"},
		},
		"Remove from list with four elems, pos 1": {
			prevList:   []string{"a", "b", "c", "d"},
			pos:        1,
			wantedList: []string{"a", "c", "d"},
		},
		"Remove from list with four elems, pos 2": {
			prevList:   []string{"a", "b", "c", "d"},
			pos:        2,
			wantedList: []string{"a", "b", "d"},
		},
		"Remove from list with four elems, pos 3": {
			prevList:   []string{"a", "b", "c", "d"},
			pos:        3,
			wantedList: []string{"a", "b", "c"},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			l := toList(t, testCase.prevList)

			err := l.Remove(testCase.pos)
			assertNoError(t, err)
			assertListContents(t, l, testCase.wantedList)
			assertListLen(t, l.Len(), len(testCase.prevList)-1)
		})
	}
}

func TestRemoveProperties(t *testing.T) {

	assertion := func() (err bool) {

		// Set up the initial slice
		numElems := rand.Intn(list.MaxElems) + 1
		initSlice := make([]string, numElems)
		for i := 0; i < numElems; i++ {
			initSlice[i] = string(fmt.Sprintf("%d", rand.Int()))
		}

		// Get a random pos and elem to remove
		pos := rand.Intn(numElems)

		t.Logf("testing removal of elem in pos %v on a list with %d element(s)", pos, len(initSlice))

		// Get the resulting slice
		endSlice := make([]string, len(initSlice)-1)
		copy(endSlice, initSlice[0:pos])
		copy(endSlice[pos:], initSlice[pos+1:])

		// Get the resulting list
		l := toList(t, initSlice)
		l.Remove(int(pos))

		// Assert whether both of them match
		return slices.Equal(toSlice(l), endSlice)
	}

	// quick.Check can provide random input to the assertion function
	// but the randomness made the function to discard most of the inputs
	// so I included the random in the function itself
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("Property based tests for Remove failed")
	}
}

func TestLenExamples(t *testing.T) {
	tests := map[string]struct{
		prevList []string
		want int
	}{
		"Len on empty list is zero": {
			prevList: nil,
			want: 0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			l := toList(t, testCase.prevList)
			got := l.Len()
			assertListLen(t, got, testCase.want)
		})
	}
}

func TestLenProperties(t *testing.T) {
	assertion := func() (err bool) {

		// Set up the initial slice
		numElems := rand.Intn(list.MaxElems)
		slice := make([]string, numElems)
		for i := 0; i < numElems; i++ {
			slice[i] = string(fmt.Sprintf("%d", rand.Int()))
		}

		t.Logf("testing Len on a list with %d element(s)", len(slice))

		// Get the equivalent list
		l := toList(t, slice)

		// Assert whether both of the lengths match
		return len(slice) == l.Len()
	}

	// quick.Check can provide random input to the assertion function
	// but the randomness made the function to discard most of the inputs
	// so I included the random in the function itself
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("Property based tests for Remove failed")
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no errors, but got %v", err)
	}
}

func assertError(t testing.TB, err error, want error) {
	t.Helper()
	if err != want {
		t.Errorf("Expected error %v, but got %v", want, err)
	}
}

func assertListContents(t testing.TB, got *list.List[string], wanted []string) {
	t.Helper()
	gotElems := toSlice(got)
	if !slices.Equal(gotElems, wanted) {
		t.Errorf("%s: wanted %v, but got %v", t.Name(), wanted, gotElems)
	}
}

func assertListLen(t testing.TB, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("%s: wanted list len to be %d, but was %d", t.Name(), want, got)
	}
}

func toSlice(l *list.List[string]) []string {
	s := []string{}
	l.Traverse(func(e string) {
		s = append(s, e)
	})
	return s
}

func toList(t testing.TB, elems []string) (l *list.List[string]) {
	t.Helper()
	l = list.New[string]()
	for i := len(elems) - 1; i >= 0; i-- {
		err := l.Insert(elems[i], 0)
		assertNoError(t, err)
	}
	return
}

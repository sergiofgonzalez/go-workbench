package stack_test

import (
	"testing"

	"example.com/arraystack/stack"
)

func TestNew(t *testing.T) {
	s := stack.New[string]()

	assertEmpty(t, s)
}

func TestEmpty(t *testing.T) {
	t.Run("empty stack should return true", func(t *testing.T) {
		s := stack.New[string]()

		assertEmpty(t, s)
	})

	t.Run("non-empty stack should return false", func(t *testing.T) {
		s := stack.New[string]()
		s.Push("a")
		assertNotEmpty(t, s)
	})
}

func TestFull(t *testing.T) {
	t.Run("empty stack should return false", func(t *testing.T) {
		s := stack.New[string]()

		assertNotFull(t, s)
	})

	t.Run("stack at its capacity should return true", func(t *testing.T) {
		s := initStack(t, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})

		assertFull(t, s)
	})
}

func TestPush(t *testing.T) {
	tests := map[string]struct {
		initialElems []string
		elem         string
		want         string
		wantErr      error
	}{
		"Push single elem": {
			elem:    "a",
			want:    "a",
			wantErr: nil,
		},
		"Push two elems": {
			initialElems: []string{"a"},
			elem: "b",
			want: "b",
			wantErr: nil,
		},
		"Pushing beyond capacity": {
			initialElems: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
			elem: "11",
			want: "10",
			wantErr: stack.ErrStackFull,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got := initStack(t, testCase.initialElems)

			err := got.Push(testCase.elem)
			assertError(t, err, testCase.wantErr)
			assertNotEmpty(t, got)
			assertTop(t, got, testCase.want)
		})
	}
}

func TestPopHappyPath(t *testing.T) {
	t.Run("Pop from stack with single elem", func(t *testing.T) {
		s := stack.New[string]()
		s.Push("a")
		want := "a"

		top, err := s.Pop()
		assertError(t, err, nil)
		if top != string(want) {
			t.Errorf("expected top to be %q, but was %q", want, top)
		}
		if !s.Empty() {
			t.Errorf("expected the stack to be empty but wasn't")
		}
	})

	t.Run("Pop from stack with two elems", func(t *testing.T) {
		s := stack.New[string]()
		s.Push("a")
		s.Push("b")
		want := "b"

		top, err := s.Pop()
		assertError(t, err, nil)
		if top != string(want) {
			t.Errorf("expected top to be %q, but was %q", want, top)
		}
		if s.Empty() {
			t.Errorf("expected the stack not to be empty but was")
		}
	})

}

func TestPopSadPath(t *testing.T) {
	s := stack.New[string]()

	_, err := s.Pop()
	assertError(t, err, stack.ErrStackEmpty)
}

func TestPeekSadPath(t *testing.T) {
	s := stack.New[string]()

	_, err := s.Peek()
	assertError(t, err, stack.ErrStackEmpty)
}

func assertEmpty(t testing.TB, got *stack.Stack[string]) {
	t.Helper()
	if !got.Empty() {
		t.Errorf("expected stack to be empty, but wasn't")
	}
}

func assertNotEmpty(t testing.TB, got *stack.Stack[string]) {
	t.Helper()
	if got.Empty() {
		t.Errorf("expected stack not to be empty, but was")
	}
}

func assertFull(t testing.TB, got *stack.Stack[string]) {
	t.Helper()
	if !got.Full() {
		t.Errorf("expected stack to be full, but wasn't")
	}
}

func assertNotFull(t testing.TB, got *stack.Stack[string]) {
	t.Helper()
	if got.Full() {
		t.Errorf("expected stack not to be full, but was")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("expected error to be %v, but was %v", want, got)
	}
}

func assertTop(t testing.TB, got *stack.Stack[string], wantTop string) {
	t.Helper()
	top, err := got.Peek()
	assertError(t, err, nil)
	if top != wantTop {
		t.Errorf("expected top of the stack to be %q, but was %q", wantTop, top)
	}
}

func initStack(t testing.TB, elems []string) *stack.Stack[string] {
	s := stack.New[string]()
	for _, elem := range elems {
		err := s.Push(elem)
		if err != nil {
			t.Errorf("error while setting up the stack for testing: %v", err)
		}
	}
	return s
}

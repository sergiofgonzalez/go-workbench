package gordle_test

import (
	"errors"
	"fmt"
	"testing"

	"example.com/gordle/gordle"
)

func TestReadCorpus(t *testing.T) {
	tests := map[string]struct {
		file           string
		expectedLength int
		isSentinelErr  bool
		expectedError  error
	}{
		"Empty corpus": {
			file:          "../corpus/empty.txt",
			isSentinelErr: true,
			expectedError: gordle.ErrCorpusIsEmpty,
		},
		"English corpus small": {
			file:           "../corpus/english_small.txt",
			expectedLength: 3,
			expectedError:  nil,
		},
		"Non-existent file": {
			file:          "../corpus/non_existent_file.txt",
			expectedError: fmt.Errorf("No such file or directory"),
		},
		"English corpus big": {
			file:           "../corpus/english.txt",
			expectedLength: 35,
			expectedError:  nil,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			words, err := gordle.ReadCorpus(testCase.file)

			// If there's an error, it should be expected, otherwise test should fail
			if err != nil && testCase.expectedError != nil {
				// if isSentinelErr, the actual error type should match in the errors tree
				if testCase.isSentinelErr && !errors.Is(err, testCase.expectedError) {
					t.Errorf("%v: unexpected sentinel error: expected %v, but got %v", testName, testCase.expectedError, err)
				}
			} else if err != nil && testCase.expectedError == nil || err == nil && testCase.expectedError != nil {
				t.Errorf("%v: unexpected error: expected %v, but got %v", testName, testCase.expectedError, err)
			}

			if len(words) != testCase.expectedLength {
				t.Errorf("%v: unexpected contents: expected %v, but got %v", testName, testCase.expectedLength, len(words))
			}
		})
	}

}

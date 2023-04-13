package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestAppend(t *testing.T) {
	tests := map[string]struct {
		src     ByteSlice
		data    []byte
		want    []byte
		nWant   int
		errWant error
	}{
		"zero-values": {},
		"appending zero-value": {
			src:  ByteSlice{1},
			want: []byte{1},
			nWant: 0,
			errWant: nil,
		},
		"appending to zero-value": {
			data: ByteSlice{1},
			want: []byte{1},
			nWant: 1,
			errWant: nil,
		},
		"no expansion": {
			src:  ByteSlice(make([]byte, 3, 5)),
			data: ByteSlice{1},
			want: []byte{0, 0, 0, 1},
			nWant: 1,
			errWant: nil,
		},
		"expansion": {
			src:  ByteSlice(make([]byte, 3, 5)),
			data: []byte{1, 2, 3},
			want: []byte{0, 0, 0, 1, 2, 3},
			nWant: 3,
			errWant: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			nGot, errGot := tc.src.Write(tc.data)
			if !slices.Equal(tc.src, tc.want) {
				t.Errorf("%v: got %v, but wanted %v", name, tc.src, tc.want)
			}
			if errGot != tc.errWant {
				t.Errorf("%v: got error %v, but wanted error %v", name, errGot, tc.errWant)
			}
			if nGot != tc.nWant {
				t.Errorf("%v: got %v as n, but wanted %v", name, nGot, tc.nWant)
			}
		})
	}
}

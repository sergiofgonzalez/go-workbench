package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestAppend(t *testing.T) {
	tests := map[string]struct{
		src  ByteSlice
		data []byte
		want []byte
	}{
		"zero-values": {
		},
		"appending zero-value": {
			src: ByteSlice{1},
			want: []byte{1},
		},
		"appending to zero-value": {
			data: ByteSlice{1},
			want: []byte{1},
		},
		"no expansion": {
			src: ByteSlice(make([]byte, 3, 5)),
			data: ByteSlice{1},
			want: []byte{0, 0, 0, 1},
		},
		"expansion": {
			src: ByteSlice(make([]byte, 3, 5)),
			data: []byte{1, 2, 3},
			want: []byte{0, 0, 0, 1, 2, 3},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.src.Append(tc.data)
			if !slices.Equal(got, tc.want) {
				t.Errorf("%v: got %v, but wanted %v", name, got, tc.want)
			}
		})
	}
}


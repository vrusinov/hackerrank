// SPDX-FileCopyrightText: 2023 Vladimir Rusinov
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"testing"
)

func TestCountTriplets(t *testing.T) {
    type test struct {
        arr  []int64
		r    int64
        want int64
    }

	tests := []test{
		{arr: []int64{1, 4, 16, 64}, r: 4, want: 2},
		{arr: []int64{1, 2, 2, 4}, r: 2, want: 2},
		{arr: []int64{1, 3, 9, 9, 27, 81}, r: 3, want: 6},
		{arr: []int64{1, 5, 5, 25, 125}, r: 5, want: 4},
		{arr: []int64{1, 1, 1}, r: 1, want: 1},
		{arr: []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, r: 1, want: 161700},
	}

    for _, tc := range tests {
        got := countTriplets(tc.arr, tc.r)
        if tc.want != got {
            t.Fatalf("expected: %v, got: %v", tc.want, got)
        }
    }
}

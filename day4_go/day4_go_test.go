package main

import "testing"

func CountPointTest(t *testing.T) {
    testcases := []struct{
        in string
        want int
    }{
        {"1 2 3 4 5 6 7 8 9 10 | 11 22 33 44 55 66 77 88 99 10", 1},
        {"1 2 3 4 5 6 7 8 9 10 | 1 2 3 4 5 6 7 8 9 10", 512},
        {"1 2 3 4 12500 6 7 8 9 10 | 12 32 43 44 5 56 27 88 109 12500", 1},
        {"10 60 71 12  7 70 18 63 40 96 |  1 48 83 36 49 21 64 78 91 99 94 56 39 74 45 51 12 32 19 75 15  5 34 79 46", 1},
    }

    for _, tc := range testcases {
        res := CountPoint(tc.in)
        if res != tc.want {
            t.Errorf("res: %v, want: %v", res, tc.want)
        }
    }
}


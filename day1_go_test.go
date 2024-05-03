package main

import "testing"

func TestSum(t *testing.T) {
    testcases := []struct {
        in  string
        want int
    }{
        {"one", 11},
        {"two", 22},
        {"three", 33},
        {"four", 44},
        {"five", 55},
        {"six", 66},
        {"seven", 77},
        {"eight", 88},
        {"nine", 99},
    }

    for _, tc := range testcases {
        s := Sum()
        res := s(tc.in)

        if res != tc.want {
            t.Errorf("Result: %q, want %q", res, tc.want)
        }
    }
}


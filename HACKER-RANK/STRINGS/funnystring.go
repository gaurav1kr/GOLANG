// https://www.hackerrank.com/challenges/funny-string/problem?isFullScreen=true
package main

import (
	"fmt"
	"math"
)

func funnyString(s string) string {
	// Write your code here
	slen := len(s)
	var diff1 float64
	var diff2 float64
	if slen > 2 {
		for i := 0; i < slen/2; i++ {
			if s[i] > s[i+1] {
				diff1 = math.Abs(float64(s[i] - s[i+1]))
			} else {
				diff1 = math.Abs(float64(s[i+1] - s[i]))
			}

			if s[slen-1-i] > s[slen-2-i] {
				diff2 = math.Abs(float64(s[slen-1-i] - s[slen-2-i]))
			} else {
				diff2 = math.Abs(float64(s[slen-2-i] - s[slen-1-i]))
			}
			if diff1 != diff2 {
				return "Not Funny"
			}
		}
	}
	return "Funny"

}

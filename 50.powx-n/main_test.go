package main

import (
	"testing"
	"log"
)

func TestMyPow(t *testing.T) {
	cases := [][]float64{
		{2.0000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
	}
	for _, c := range cases {
		r := myPow(c[0], int(c[1]))
		if r != c[2] {
			log.Fatalf("%v**%v should be %v, get %v", c[0], int(c[1]), c[2], r)
		}
	}
}
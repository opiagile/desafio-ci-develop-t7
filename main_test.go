package main

import (
	"reflect"
	"testing"
)

type Int2IntTestPair struct {
	input1 int
	input2 int
	output int
}

func TestSoma(t *testing.T) {
	var tests = []Int2IntTestPair{
		{1, 1, 2},
		{6, 4, 10},
	}

	for _, pair := range tests {
		result := soma(pair.input1, pair.input2)
		if !reflect.DeepEqual(result, pair.output) {
			t.Errorf("\nSum of %v and %v:\nexpected: %v\nbut got: %v",
				pair.input1, pair.input2, pair.output, result)
		}
	}
}

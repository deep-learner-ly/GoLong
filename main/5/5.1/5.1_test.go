package mypackage

import "testing"

func TestAdd(t *testing.T) {
	//result := Add(2, 3)
	//expected := 5
	//
	//if result != expected {
	//	t.Fatalf("Expected %d, but got %d", expected, result)
	//}

	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{5, -3, 2},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

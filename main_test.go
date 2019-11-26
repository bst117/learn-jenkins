package math

import "testing"

func TestSum(t *testing.T) {
	if Sum(2, 3) != 5 {
		t.Fatal("sum(2, 3) != 5")
	}
}

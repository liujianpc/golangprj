package simplemath

import "testing"

func TestSqrt1(t *testing.T) {
    v := Sqrt(16)
    if v != 4 {
        t.Errorf("Sqrt(16) failed. got %d excepted 4", v)
    }
}

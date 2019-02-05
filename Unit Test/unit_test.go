package main

import "testing"

func TestCheck(t *testing.T) {
    total := CheckErr(nil)
    if total != 0 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

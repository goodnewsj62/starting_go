package gopack

import "testing"


func TestPackNum(t *testing.T){
	x := PackNum(1,2)
	if x == nil {
		t.Error("man this is not working")
	}
}
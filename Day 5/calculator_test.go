package main


import "testing"

func TestAdd (t *testing.T){

	got :=Add(2,3)

	want :=5

	if got !=want {
		t.Errorf("expected %d got %d", want , got)
	}
}

func TestSubtract(t *testing.T){
	got :=Subtract(5,2)

want :=3

if got !=want {
	t.Errorf("expected %d got %d " , want , got)
}
}


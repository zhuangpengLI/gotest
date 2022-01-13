package simplefactory

import "testing"

func TestTpye1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		t.Fatal("Tpye1 test fail")
	}
}

func TestTpye2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		t.Fatal("Tpye2 test fail")
	}
}

package main

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestBasic(t *testing.T) {

	buf, err := ioutil.ReadFile("test_input")
	if err != nil {
		log.Fatal(err)
	}

	if run(buf) != 3 {
		t.Errorf("Test failed. Expected 2, got %d", run(buf))
	}

}

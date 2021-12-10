package test

import (
	"fmt"
	"gin-template/short"
	"testing"
)

func TestTable(t *testing.T) {

	sequence := short.NewSequence()

	nextSequence, _ := sequence.NextSequence()
	fmt.Println(nextSequence)
}

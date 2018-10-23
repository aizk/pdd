package pdd

import (
	"testing"
	"fmt"
)

func TestNewPdd(t *testing.T) {
	p := NewPdd(&Config{})
	fmt.Println(p)
}

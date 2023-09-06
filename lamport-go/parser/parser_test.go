package parser

import (
	"fmt"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
    let x = 10; 
    let y = 20;
  `
	fmt.Printf(input)
}

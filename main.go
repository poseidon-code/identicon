package main

import (
	"flag"
	"fmt"
	"os"

	h "github.com/poseidon-code/go-identicons/pkg/hash"
	m "github.com/poseidon-code/go-identicons/pkg/matrix"
)

func main() {
    size_ptr := flag.Int("size", 5, "sets size of the identicon (range: 4-8)")
    square_ptr := flag.Bool("square", false, "creates a square identicon")
    border_ptr := flag.Bool("border", false, "adds a border to the identicon")
    vertical_ptr := flag.Bool("vertical", false, "creates identicon in portrait dimension (not visible on using --square flag)")
    invert_ptr := flag.Bool("invert", false, "inverts the cell filling of identicon")
    flag.Parse()

    // variable declarations
    var text string
    var hash string
    var matrix [][]int
    var W, H int
    
    // handling text
    if len(flag.Args())>1 {
        fmt.Println("Invalid sequence of flags & arguments passed. \nUse flags first before arguments. e.g.: \ngo-identicons --size 8 lovely")
        os.Exit(1)
    } else if len(flag.Args())==0 {
        fmt.Println("No argument passed for the text. Use like: \ngo-identicons lovely")
        os.Exit(1)
    } else {
        text = flag.Arg(0)
    }

    // handling type (square|wide)
    if *square_ptr {
        hash, W, H = h.GenerateSHA256(text)
    } else {
        hash, W, H = h.GenerateSHA512(text)
    }

    // handling size (4-8)
    if *size_ptr<4 || *size_ptr>8 {
        fmt.Println("Invalid size passed. \nSize must lie between 4 to 8 (inclusive).")
        os.Exit(1)
    }

    // handling vertical dimension (rather than rotating the entire martrix, only the dimensions are switched) (landscape|portrait)
    // handling cell filling (original|invert)
    if *vertical_ptr {
        matrix = m.GenerateSymmetric(hash, *size_ptr, H, W, *invert_ptr)
    } else {
        matrix = m.GenerateSymmetric(hash, *size_ptr, W, H, *invert_ptr)
    }

    // handling border (border|no-border)
    if *border_ptr {
        m.PrintBordered(matrix)
    } else {
        m.Print(matrix)
    }
}
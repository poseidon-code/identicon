package main

import (
	"flag"
	"fmt"
	"os"

	i "github.com/poseidon-code/go-identicons/pkg/identicon"
	m "github.com/poseidon-code/go-identicons/pkg/matrix"
)

func main() {
    // PARSING CMD OPTIONS
    size_ptr        := flag.Int("size", i.Defaults.Size, "sets size of the identicon (range: 4-8)")
    square_ptr      := flag.Bool("square", i.Defaults.Square, "creates a square identicon")
    border_ptr      := flag.Bool("border", i.Defaults.Border, "adds a border to the identicon")
    vertical_ptr    := flag.Bool("vertical", i.Defaults.Vertical, "creates identicon in portrait dimension (not visible on using --square flag)")
    invert_ptr      := flag.Bool("invert", i.Defaults.Invert, "inverts the cell filling of identicon")
    symmetric_ptr   := flag.Bool("symmetric", i.Defaults.Symmetric, "creates symmetric identicon")
    flag.Parse()


    // SETTING OPTIONS
    var options = i.Configuration{
        Size:       *size_ptr,
        Square:     *square_ptr,
        Border:     *border_ptr,
        Vertical:   *vertical_ptr,
        Invert:     *invert_ptr,
        Symmetric:  *symmetric_ptr,
    }

    var identicon i.Identicon


    // PARSING TEXT & SETTING IDENTICON
    // handling text
    if len(flag.Args())>1 {
        fmt.Println("Invalid sequence of flags & arguments passed. \nUse flags first before arguments. e.g.: \ngo-identicons --size 8 lovely")
        os.Exit(1)
    } else if len(flag.Args())==0 {
        fmt.Println("No argument passed for the text. Use like: \ngo-identicons lovely")
        os.Exit(1)
    } else {
        // setting Identicon
        identicon = i.Identicon{
            Options: options,
            Text: flag.Arg(0),
        }
    }


    // GENERATING IDENTICON
    identicon.New()


    // PRINTING
    // handling border (border|no-border)
    if identicon.Options.Border {
        m.PrintBordered(identicon.Matrix)
    } else {
        m.Print(identicon.Matrix)
    }
}
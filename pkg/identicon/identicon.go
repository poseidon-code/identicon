package identicon

import (
	"fmt"
	"os"

	h "github.com/poseidon-code/go-identicons/pkg/hash"
	m "github.com/poseidon-code/go-identicons/pkg/matrix"
)

/* Identicon
 * New() handles all the options & flags passed in the
 * command line regarding the type of the matrix that
 * will be made (like square, portrait dimensions,
 * inversion and size), and sets only Hash, Width, Height,
 * Matrix values of an Identicon (as they need to be generated).
 */

type Identicon struct {
    Options         Configuration
    ImageOptions    ImageConfiguration
    Text            string
    Hash            string
    Width, Height   int
    Matrix          [][]int
}

// creates and sets all the fields of Identicon struct
func (i *Identicon) New() {
    // handling size (4-8)
    if i.Options.Size<4 || i.Options.Size>8 {
        fmt.Println("Invalid size passed. \nSize must lie between 4 to 8 (inclusive).")
        os.Exit(1)
    }


    // handling type (square|wide)
    if i.Options.Square {
        i.Hash, i.Width, i.Height = h.SHA256(i.Text)
    } else {
        i.Hash, i.Width, i.Height = h.SHA512(i.Text)
    }    


    // handling vertical dimension (rather than rotating the entire martrix, only the dimensions are switched) (landscape|portrait)
    if i.Options.Vertical {
        // handling symmetric filling (asymmetric|symmetric)
        if i.Options.Symmetric {
            // handling cell filling (original|invert)
            i.Matrix = m.GenerateSymmetric(i.Hash, i.Options.Size, i.Height, i.Width, i.Options.Invert)
        } else {
            i.Matrix = m.Generate(i.Hash, i.Options.Size, i.Height, i.Width, i.Options.Invert)
        }
    } else {
        if i.Options.Symmetric {
            i.Matrix = m.GenerateSymmetric(i.Hash, i.Options.Size, i.Width, i.Height, i.Options.Invert)
        } else {
            i.Matrix = m.Generate(i.Hash, i.Options.Size, i.Width, i.Height, i.Options.Invert)
        }
    }
}

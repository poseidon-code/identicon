package identicon

import "fmt"

/* TERMINAL PRINTING
 * handles:
 *      --border flag (no-border|border)
 *
 * `Options.Size` is used to determine the size of the identicon, as
 * `Width` & `Height` are just ratios of the dimensions.
 * overall width & height of the identicon to be printed is calculated
 * by multiplying the ratio to the size.
 * (i.e.: i.Width*i.Options.Size, and same for i.Height)
 *
 * IMPORTANT: It is required to have all the fields of struct `Identicon` be filled
 * i.e. Identicon.New() should be called first.
 *
 * NOTE : The size can also be extracted from the `Matrix` directly
 * (i.e width = len(i.Matrix[0]) & height = len(i.Matrix))
 */

func (i *Identicon) Print() {
    w, h := (i.Width*i.Options.Size), (i.Height*i.Options.Size)
    m := i.Matrix
    
    if i.Options.Border {
        w, h = w+4, h+2

        for r:=0; r<h; r++ {
            for c:=0; c<w; c++ {
                if r==0 {
                    if c==0 {
                        fmt.Print("⎡")
                    } else if c==w-1 {
                        fmt.Print("⎤")
                    } else {
                        fmt.Print("⎺⎺")
                    }
                } else if r==h-1 {
                    if c==0 {
                        fmt.Print("⎣")
                    } else if c==w-1 {
                        fmt.Print("⎦")
                    } else {
                        fmt.Print("__")
                    }
                } else {
                    if c==0 {
                        fmt.Print("⎢")
                    } else if c==w-1 {
                        fmt.Print("⎥")
                    } else if c==1 || c==w-2 {
                        fmt.Print("  ")
                    } else {
                        if m[r-1][c-2] == 0 {
                            fmt.Print("  ")
                        } else {
                            fmt.Print("██")
                        }
                    }
                }
            }
            fmt.Println()
        }
    } else {
        for r:=0; r<h; r++ {
            for c:=0; c<w; c++ {
                if m[r][c] == 0 {
                    fmt.Print("  ")
                } else {
                    fmt.Print("██")
                }
            }
            fmt.Println()
        }
    }
}
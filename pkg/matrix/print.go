package matrix

import "fmt"

func Print(m [][]int) {
    w, h := len(m[0]), len(m)

    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            if m[i][j] == 0 {
                fmt.Print("  ")
            } else {
                fmt.Print("██")
            }
        }
        fmt.Println()
    }
}

func PrintBordered(m [][]int) {
    w, h := len(m[0])+4, len(m)+2

    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            if i==0 {
                if j==0 {
                    fmt.Print("⎡")
                } else if j==w-1 {
                    fmt.Print("⎤")
                } else {
                    fmt.Print("⎺⎺")
                }
            } else if i==h-1 {
                if j==0 {
                    fmt.Print("⎣")
                } else if j==w-1 {
                    fmt.Print("⎦")
                } else {
                    fmt.Print("__")
                }
            } else {
                if j==0 {
                    fmt.Print("⎢")
                } else if j==w-1 {
                    fmt.Print("⎥")
                } else if j==1 || j==w-2 {
                    fmt.Print("  ")
                } else {
                    if m[i-1][j-2] == 0 {
                        fmt.Print("  ")
                    } else {
                        fmt.Print("██")
                    }
                }
            }
        }
        fmt.Println()
    }
}

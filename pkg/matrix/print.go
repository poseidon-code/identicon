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

    p := make([][]string, h)
    for i:=0; i<h; i++ {
        p[i] = make([]string, w)
    }

    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            if i==0 {
                if j==0 {
                    p[i][j] = "⎡"
                } else if j==w-1 {
                    p[i][j] = "⎤"
                } else {
                    p[i][j] = "⎺⎺"
                }
            } else if i==h-1 {
                if j==0 {
                    p[i][j] = "⎣"
                } else if j==w-1 {
                    p[i][j] = "⎦"
                } else {
                    p[i][j] = "__"
                }
            } else {
                if j==0 {
                    p[i][j] = "⎢"
                } else if j==w-1 {
                    p[i][j] = "⎥"
                } else if j==1 || j==w-2 {
                    p[i][j] = "  "
                } else {
                    if m[i-1][j-2] == 0 {
                        p[i][j] = "  "
                    } else {
                        p[i][j] = "██"
                    }
                }
            }
        }
    }

    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            fmt.Print(p[i][j])
        }
        fmt.Println()
    }
}

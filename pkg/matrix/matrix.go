package matrix

import "fmt"

func Generate(hash string, s int, w, h *int) [][]int {
    *w, *h = *w*s, *h*s
    m := make([][]int, *h)
    for i:=0; i<*h; i++ {
        m[i] = make([]int, *w)
    }

    k:=1
    for i:=0; i<*h; i++ {
        for j:=0; j<*w; j++ {
            m[i][j] = func()int{if int(hash[k-1]) % 2 != 0 {return 1} else {return 0}}()
            k++
        }
    }

    return m
}

func Print(m [][]int) {
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

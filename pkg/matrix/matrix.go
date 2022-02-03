package matrix

func Generate(hash string, s int, w, h int, invert bool) [][]int {
    w, h = w*s, h*s
    m := make([][]int, h)
    for i:=0; i<h; i++ {
        m[i] = make([]int, w)
    }

    k:=1
    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            if invert {
                m[i][j] = func()int{if int(hash[k-1]) % 2 != 0 {return 0} else {return 1}}()
            } else {
                m[i][j] = func()int{if int(hash[k-1]) % 2 != 0 {return 1} else {return 0}}()
            }
            k++
        }
    }

    return m
}

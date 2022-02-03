package matrix

func put_value(a, x, y int) int {
    if a % 2 != 0 {return x} else {return y}
}

func Generate(hash string, s int, w, h int, invert bool) [][]int {
    w, h = w*s, h*s
    m := make([][]int, h)
    for i:=0; i<h; i++ {
        m[i] = make([]int, w)
    }

    var k int = 1
    var bit int
    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            if invert {
                bit = put_value(int(hash[k-1]), 0, 1)
            } else {
                bit = put_value(int(hash[k-1]), 1, 0)
            }

            m[i][j] = bit
            k++
        }
    }

    return m
}

func GenerateSymmetric(hash string, s int, w, h int, invert bool) [][]int {
    w, h = w*s, h*s
    m := make([][]int, h)
    for i:=0; i<h; i++ {
        m[i] = make([]int, w)
    }

    var k int = 1
    var bit int
    for i:=0; i<h; i++ {
        for j:=0; j<w; j++ {
            if j>=w/2+1 {k++; continue}

            if invert {
                bit = put_value(int(hash[k-1]), 0, 1)
            } else {
                bit = put_value(int(hash[k-1]), 1, 0)
            }

            m[i][j], m[i][w-j-1] = bit, bit
            k++
        }
    }

    return m
}
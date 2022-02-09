package identicon

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

type size struct { w, h int }

var ss = size{1440, 720}
var sm = size{2160, 1080}
var sl = size{2880, 1440}
var sx = size{4320, 2160}

func get_size(s string) (w int, h int) {
    if s=="S" || s=="M" || s=="L" || s=="X" {
        switch s {
        case "S":
            return ss.w, ss.h
        case "M":
            return sm.w, sm.h
        case "L":
            return sl.w, sl.h
        case "X":
            return sx.w, sx.h
        }
    } else {
        fmt.Println("Invalid --image-size value passed.") 
        fmt.Println("--image-size value should be one of S, M, L & X.")
        fmt.Println("i.e.: --image-size=X")
        os.Exit(1)
    }

    return -1, -1
}

func hex_to_rgb(h string) color.Color {
    if len(h)!=6 {
        fmt.Println("Color should be in HEX format of length 6 (range: '000000' to 'ffffff')")
        os.Exit(1)
    }

    b1 := fmt.Sprintf("%s%s", string(h[0]), string(h[1]))
    b2 := fmt.Sprintf("%s%s", string(h[2]), string(h[3]))
    b3 := fmt.Sprintf("%s%s", string(h[4]), string(h[5]))

    r, err := strconv.ParseInt(b1, 16, 64)
    if err!=nil {
        fmt.Printf("Invalid HEX color: '%s' in '%s'\n", b1, h)
        os.Exit(1)
    }
    g, err := strconv.ParseInt(b2, 16, 64)
    if err!=nil {
        fmt.Printf("Invalid HEX color: '%s' in '%s'\n", b2, h)
        os.Exit(1)
    }
    b, err := strconv.ParseInt(b3, 16, 64)
    if err!=nil {
        fmt.Printf("Invalid HEX color: '%s' in '%s'\n", b3, h)
        os.Exit(1)
    }

    return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}


func get_block_size(i Identicon, w, rw int) int {
    var b int
    
    if i.ImageOptions.Portrait && i.Options.Vertical {
        b = w/rw/5
    } else if i.Options.Vertical {
        b = w/rw/10
    } else if i.ImageOptions.Portrait {
        b = w/rw/3
    } else {
        b = w/rw/5
    }

    if i.Options.Square {
        if i.ImageOptions.Portrait && i.Options.Vertical {
            b = w/rw/5
        } else if i.Options.Vertical {
            b = w/rw/10
        } else if i.ImageOptions.Portrait {
            b = w/rw/5
        } else {
            b = w/rw/10
        }
    }

    return b
}


func handle_file_path(path, t string) string {
    if _, err := os.Stat(path); err!=nil {
        if os.IsNotExist(err) {
            fmt.Println("Invalid save directory path. Directory doesn't exists.")
            os.Exit(1)
        }
    }

    sd := fmt.Sprintf("%s/%s.png", path, t)
    return sd
}


func (i *Identicon) Save() {
    w, h := get_size(i.ImageOptions.Size)
    rw, rh := (i.Width*i.Options.Size), (i.Height*i.Options.Size)

    if i.Options.Vertical {rw, rh = rh, rw}
    if i.ImageOptions.Portrait {w, h = h, w}
    
    b := get_block_size(*i, w, rw)

    img := image.NewRGBA(image.Rectangle{image.Point{0,0},image.Point{w,h}})
    fg := hex_to_rgb(i.ImageOptions.FG)
    bg := hex_to_rgb(i.ImageOptions.BG)

    // set background
    for x:=0; x<w; x++ {for y:=0; y<h; y++ {img.Set(x, y, bg)}}

    // set border
    if i.Options.Border {
        offset_w := w/2 - b*(rw/2) - b
        offset_h := h/2 - b*(rh/2) - b
        br := [][]int{
            {0,                 (b*(rw+2)),     0,                  (b/3)      },
            {0,                 (b/3),          0,                  (b*(rh+2)) },
            {0,                 (b*(rw+2)),     (b*(rh+2))-(b/3),   (b*(rh+2)) },
            {(b*(rw+2))-(b/3),  (b*(rw+2)),     0,                  (b*(rh+2)) },
        }
        for _, v := range br {
            for x:=v[0]; x<v[1]; x++ {
                for y:=v[2]; y<v[3]; y++ {
                    img.Set(x+offset_w, y+offset_h, fg)
                }
            }
        }
    }


    // set identicon
    for r:=0; r<len(i.Matrix); r++ {
        for c:=0; c<len(i.Matrix[0]); c++ {
            pos_x := c*b+(w/2)-(b*(rw/2))
            pos_y := r*b+(h/2)-(b*(rh/2))

            if i.Matrix[r][c] == 1 {
                for x:=pos_x; x<pos_x+b; x++ {
                    for y:=pos_y; y<pos_y+b; y++ {
                        img.Set(x, y, fg)
                    }
                }
            }
        }
    }


    file_name := handle_file_path(i.ImageOptions.SaveDir, i.Text)
    f, _ := os.Create(file_name); defer f.Close()
    png.Encode(f, img)
}
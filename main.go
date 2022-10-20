package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	g "github.com/poseidon-code/godenticon"
)

func is_flag_passed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {found = true}
    })
    return found
}

func main() {
    // PARSING COMMANDLINE OPTIONS
    // identicon configurations
    size_ptr        := flag.Int(    "size",         g.IdenticonDefaultOptions.Size,         "sets size of the identicon (range: 4-8)")
    square_ptr      := flag.Bool(   "square",       g.IdenticonDefaultOptions.Square,       "creates a square identicon")
    border_ptr      := flag.Bool(   "border",       g.IdenticonDefaultOptions.Border,       "adds a border to the identicon")
    vertical_ptr    := flag.Bool(   "vertical",     g.IdenticonDefaultOptions.Vertical,     "creates identicon in portrait dimension (not visible on using --square flag)")
    invert_ptr      := flag.Bool(   "invert",       g.IdenticonDefaultOptions.Invert,       "inverts the cell filling of identicon")
    symmetric_ptr   := flag.Bool(   "symmetric",    g.IdenticonDefaultOptions.Symmetric,    "creates symmetric identicon")

    // image configurations
    image_size_ptr      := flag.String( "image-size",       g.ImageDefaultOptions.Size,           "saves image with given resolution preset (S,M,L,X)")
    image_portrait_ptr  := flag.Bool(   "image-portrait",   g.ImageDefaultOptions.Portrait,       "saves image with portrait dimensions")
    fg_ptr              := flag.String( "fg",               g.ImageDefaultOptions.FG,             "sets image's foreground color")
    bg_ptr              := flag.String( "bg",               g.ImageDefaultOptions.BG,             "sets image's background color")
    
    // if --config path is passed, ignore every other flags
    config_ptr          := flag.String( "config",       "",         "path to config.json file")
    save_ptr            := flag.String( "save",         "",         "saves image to the specified directory")
    hash_ptr            := flag.Bool(   "hash",         false,      "allows passing hash directly (instead of text)")

    flag.Parse()

    var identicon g.Identicon
    var identicon_o g.IdenticonConfiguration
    var image_o g.ImageConfiguration


    // SETTING OPTIONS
    if is_flag_passed("config") {
        // handle json configs
        if flag.NFlag()>1 && !is_flag_passed("save") {
            fmt.Println("When --config is passed, all other CLI options will be discarded (except --save).")
        }
        identicon.ReadConfiguration(*config_ptr)
    } else {
        // handle commandline options
        identicon_o = g.IdenticonConfiguration{
            Size:       *size_ptr,
            Square:     *square_ptr,
            Border:     *border_ptr,
            Vertical:   *vertical_ptr,
            Invert:     *invert_ptr,
            Symmetric:  *symmetric_ptr,
        }

        image_o = g.ImageConfiguration{
            Size:       *image_size_ptr,
            Portrait:   *image_portrait_ptr,
            FG:         *fg_ptr,
            BG:         *bg_ptr,
        }

        identicon.IdenticonOptions = identicon_o
        identicon.ImageOptions = image_o
        identicon.CheckConfiguration()
    }


    // SETTING IDENTICON TEXT/HASH
    // handling text/hash
    if len(flag.Args())>1 {
        fmt.Println(
            "Invalid sequence of flags & arguments passed.",
            "\nUse flags before argument. e.g.: ",
            "\nidenticon --size=8 lovely",
        )
        fmt.Println(); flag.Usage()
        os.Exit(1)
    } else if len(flag.Args())==0 {
        fmt.Println(
            "No argument passed for the text. Use like: ",
            "\nidenticon lovely",
        )
        fmt.Println(); flag.Usage()
        os.Exit(1)
    }

    if *hash_ptr {
        identicon.Hash = flag.Arg(0)
        identicon.CheckHash()
    } else {
        identicon.Text = flag.Arg(0)
        identicon.GenerateHash()
    }
    

    // GENERATING IDENTICON
    identicon.GenerateMatrix()
    // variable `identicon` will now have all the required values for further 
    // operation on it, like printing or saving image, etc.


    // PRINTING
    identicon.Print()


    // SAVING IDENTICON IMAGE
    // checking if any other image related flags are passed except `--save`
    // if so, then prompt user to pass `--save` flag also
    // else, when no `--save` or any other image related flags are passed, then do nothing.
    other_image_flags := false
    for _, f := range []string{"image-portrait", "image-size", "fg", "bg"} {
        if is_flag_passed(f) {
            other_image_flags = true
            break
        }
    }

    if is_flag_passed("save") {
        // save image only when `--save` flag is passed
        _, fileName := filepath.Split(*save_ptr)
        parts := strings.Split(fileName, ".")
        extension := strings.ToLower(parts[len(parts)-1])

        if extension == "png" {
            identicon.SaveImage(*save_ptr)
        } else if extension == "svg" {
            identicon.SaveSVG(*save_ptr)
        } else {
            identicon.SaveImage(*save_ptr)
            identicon.SaveSVG(*save_ptr)
        }
    } else if other_image_flags {
        // if any other image related flags are passed without `--save` flag
        fmt.Println("To save image provide --save=<path> flag.")
    }
}

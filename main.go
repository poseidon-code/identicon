package main

import (
	"flag"
	"fmt"
	"os"

	gi "github.com/poseidon-code/godenticon"
)

func main() {
    // PARSING COMMANDLINE OPTIONS
    // identicon configurations
    size_ptr        := flag.Int(    "size",         gi.IdenticonDefaultOptions.Size,         "sets size of the identicon (range: 4-8)")
    square_ptr      := flag.Bool(   "square",       gi.IdenticonDefaultOptions.Square,       "creates a square identicon")
    border_ptr      := flag.Bool(   "border",       gi.IdenticonDefaultOptions.Border,       "adds a border to the identicon")
    vertical_ptr    := flag.Bool(   "vertical",     gi.IdenticonDefaultOptions.Vertical,     "creates identicon in portrait dimension (not visible on using --square flag)")
    invert_ptr      := flag.Bool(   "invert",       gi.IdenticonDefaultOptions.Invert,       "inverts the cell filling of identicon")
    symmetric_ptr   := flag.Bool(   "symmetric",    gi.IdenticonDefaultOptions.Symmetric,    "creates symmetric identicon")

    // image configurations
    image_size_ptr      := flag.String( "image-size",       gi.ImageDefaultOptions.Size,           "saves image with given resolution preset (S,M,L,X)")
    image_portrait_ptr  := flag.Bool(   "image-portrait",   gi.ImageDefaultOptions.Portrait,       "saves image with portrait dimensions")
    fg_ptr              := flag.String( "fg",               gi.ImageDefaultOptions.FG,             "sets image's foreground color")
    bg_ptr              := flag.String( "bg",               gi.ImageDefaultOptions.BG,             "sets image's background color")
    
    // if --config path is passed, ignore every other flags
    config_ptr          := flag.String( "config",       "",         "path to config.json file")
    save_ptr            := flag.Bool(   "save",         false,      "save the identicon as an image with default image options")
    save_dir_ptr        := flag.String( "save-dir",     "",         "saves image to the specified directory")

    flag.Parse()

    var identicon gi.Identicon


    // SETTING OPTIONS
    var identicon_o gi.IdenticonConfiguration
    var image_o gi.ImageConfiguration
    // handle json configs
    if len(*config_ptr)>0 {
        if flag.NFlag()>1 {
            fmt.Println("When --config is passed, all other options will be discarded.")
        }
        identicon.ReadConfiguration(*config_ptr)
    } else {
        // handle commandline options
        identicon_o = gi.IdenticonConfiguration{
            Size:       *size_ptr,
            Square:     *square_ptr,
            Border:     *border_ptr,
            Vertical:   *vertical_ptr,
            Invert:     *invert_ptr,
            Symmetric:  *symmetric_ptr,
        }

        image_o = gi.ImageConfiguration{
            Size:       *image_size_ptr,
            Portrait:   *image_portrait_ptr,
            FG:         *fg_ptr,
            BG:         *bg_ptr,
        }
    }


    // PARSING TEXT & SETTING IDENTICON
    // handling text
    if len(flag.Args())>1 {
        fmt.Println("Invalid sequence of flags & arguments passed. \nUse flags before argument. e.g.: \ngo-identicons --size=8 lovely")
        os.Exit(1)
    } else if len(flag.Args())==0 {
        fmt.Println("No argument passed for the text. Use like: \ngo-identicons lovely")
        os.Exit(1)
    } else {
        // setting Identicon
        identicon = gi.Identicon{
            IdenticonOptions: identicon_o,
            ImageOptions: image_o,
            Text: flag.Arg(0),
        }
    }

    
    // GENERATING IDENTICON
    identicon.GenerateHash()
    identicon.GenerateMatrix()
    // variable `identicon` will now have all the required values for further 
    // operation on it, like printing or saving image, etc.


    // PRINTING
    identicon.Print()


    // SAVING IMAGE
    // checking if any other image related flags are passed except `--save`
    // if so, then prompt user to pass `--save` flag also
    // else, when no `--save` or any other image related flags are passed, then do nothing.
    other_image_flags := false
    handle_image_flags := func(f *flag.Flag){
        if f.Name=="image-portrait" || f.Name=="image-size" || f.Name=="fg" || f.Name=="bg" || f.Name=="save-dir" {
            other_image_flags = true
        }
    }
    flag.Visit(handle_image_flags)

    if *save_ptr {
        // save image only when `--save` flag is passed
        identicon.SaveImage(*save_dir_ptr)
    } else if other_image_flags {
        // if any other image related flags are passed without `--save` flag
        fmt.Println("To save image provide --save flag.")
    }
}
package main

import (
	"flag"
	"fmt"
	"os"

	i "github.com/poseidon-code/go-identicons/pkg/identicon"
)

func main() {
    // PARSING COMMANDLINE OPTIONS
    // identicon configurations
    size_ptr        := flag.Int(    "size",         i.Defaults.Size,         "sets size of the identicon (range: 4-8)")
    square_ptr      := flag.Bool(   "square",       i.Defaults.Square,       "creates a square identicon")
    border_ptr      := flag.Bool(   "border",       i.Defaults.Border,       "adds a border to the identicon")
    vertical_ptr    := flag.Bool(   "vertical",     i.Defaults.Vertical,     "creates identicon in portrait dimension (not visible on using --square flag)")
    invert_ptr      := flag.Bool(   "invert",       i.Defaults.Invert,       "inverts the cell filling of identicon")
    symmetric_ptr   := flag.Bool(   "symmetric",    i.Defaults.Symmetric,    "creates symmetric identicon")

    // image configurations
    save_ptr            := flag.Bool(   "save",             i.ImageDefaults.Save,           "save the identicon as an image with default image options")
    image_portrait_ptr  := flag.Bool(   "image-portrait",   i.ImageDefaults.Portrait,       "saves image with portrait dimensions")
    image_size_ptr      := flag.String( "image-size",       i.ImageDefaults.Size,           "saves image with given resolution preset (S,M,L,X)")
    fg_ptr              := flag.String( "fg",               i.ImageDefaults.FG,             "sets image's foreground color")
    bg_ptr              := flag.String( "bg",               i.ImageDefaults.BG,             "sets image's background color")
    save_dir_ptr        := flag.String( "save-dir",         i.ImageDefaults.SaveDir,        "saves image to the specified directory")

    // if --config path is passed, ignore every other flags
    config_ptr := flag.String( "config", "", "path to config.json file")

    flag.Parse()


    // SETTING OPTIONS
    var options i.Configuration
    var image_options i.ImageConfiguration
    // handle json configs
    if len(*config_ptr)>0 {
        if flag.NFlag()>1 {
            fmt.Println("When --config is passed, all other options will be discarded.")
        }

        if _, err := os.Stat(*config_ptr); err != nil {
            if os.IsNotExist(err) {
                fmt.Println("Invalid file path : ", *config_ptr)
                os.Exit(1)
            }
        }

        options.ReadConfiguration(*config_ptr)
        image_options.ReadConfiguration(*config_ptr)
    } else {
        // handle commandline options
        options = i.Configuration{
            Size:       *size_ptr,
            Square:     *square_ptr,
            Border:     *border_ptr,
            Vertical:   *vertical_ptr,
            Invert:     *invert_ptr,
            Symmetric:  *symmetric_ptr,
        }

        if *save_ptr {
            image_options = i.ImageConfiguration{
                Size: *image_size_ptr,
                Save: *save_ptr,
                SaveDir: *save_dir_ptr,
                Portrait: *image_portrait_ptr,
                FG: *fg_ptr,
                BG: *bg_ptr,
            }
        }
    }


    // PARSING TEXT & SETTING IDENTICON
    var identicon i.Identicon
    // handling text
    if len(flag.Args())>1 {
        fmt.Println("Invalid sequence of flags & arguments passed. \nUse flags before argument. e.g.: \ngo-identicons --size=8 lovely")
        os.Exit(1)
    } else if len(flag.Args())==0 {
        fmt.Println("No argument passed for the text. Use like: \ngo-identicons lovely")
        os.Exit(1)
    } else {
        // setting Identicon
        identicon = i.Identicon{
            Options: options,
            ImageOptions: image_options,
            Text: flag.Arg(0),
        }
    }


    // GENERATING IDENTICON
    identicon.New()
    // variable `identicon` will now have all the required values for further 
    // operation on it, like printing or saving image, etc.


    // PRINTING
    identicon.Print()

    // SAVING IMAGE
    if identicon.ImageOptions.Save {
        identicon.Save()
    } else {
        fmt.Println("To save image provide --save flag.")
    }
}
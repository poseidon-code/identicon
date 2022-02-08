package identicon

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/* Configuration
 * these configuration values are set by passing
 * command line arguments, hence, all the flags
 * are parsed inside `main` package, and the
 * Identicon.Options is also set there after
 * parsing all the flags.
 * If some flags are not passed, then Defaults are used.
 */

type Configuration struct {
    Size        int     `json:"size"`       // sets size of the identicon (range: 4-8)
    Square      bool    `json:"square"`     // creates a square identicon
    Border      bool    `json:"border"`     // adds a border to the identicon
    Vertical    bool    `json:"vertical"`   // creates identicon in portrait dimension (not visible on using --square flag)
    Invert      bool    `json:"invert"`     // inverts the cell filling of identicon
    Symmetric   bool    `json:"symmetric"`  // creates symmetric identicon
}

// default configuration values for Identicon.Options
var Defaults = Configuration{
    Size:       6,
    Square:     false,
    Border:     false,
    Vertical:   false,
    Invert:     false,
    Symmetric:  false,
}

func (o *Configuration) ReadConfiguration(path string) {
    f, _ := os.Open(path); defer f.Close()
    b, _ := ioutil.ReadAll(f)
    *o = Defaults
    json.Unmarshal(b, &o)
}


type ImageConfiguration struct {
    Size        string      `json:"image-size"`         // saves image with given resolution preset (S,M,L,X)
    Save        bool        `json:"save"`               // save the identicon as an image with default image options
    Portrait    bool        `json:"image-portrait"`     // saves image with portrait dimensions
    FG          string      `json:"fg"`                 // sets image's foreground color
    BG          string      `json:"bg"`                 // sets image's background color
}

// default configuration values for Identicon.ImageOptions
var ImageDefaults = ImageConfiguration{
    Size:       "L",
    Save:       false,
    Portrait:   false,
    FG:         "6dff24",
    BG:         "0b2100",
}

func (io *ImageConfiguration) ReadConfiguration(path string) {
    f, _ := os.Open(path); defer f.Close()
    b, _ := ioutil.ReadAll(f)
    *io = ImageDefaults
    json.Unmarshal(b, &io)
}
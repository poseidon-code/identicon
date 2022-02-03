package identicon

type Configuration struct {
    Size        int     // sets size of the identicon (range: 4-8)
    Square      bool    // creates a square identicon
    Border      bool    // adds a border to the identicon
    Vertical    bool    // creates identicon in portrait dimension (not visible on using --square flag)
    Invert      bool    // inverts the cell filling of identicon
    Symmetric   bool    // creates symmetric identicon
}

var Defaults = Configuration{
    Size:       5,
    Square:     false,
    Border:     false,
    Vertical:   false,
    Invert:     false,
    Symmetric:  false,
}
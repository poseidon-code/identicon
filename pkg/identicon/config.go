package identicon

/* Configuration
 * these configuration values are set by passing
 * command line arguments, hence, all the flags
 * are parsed inside `main` package, and the
 * Identicon.Options is also set there after
 * parsing all the flags.
 * If some flags are not passed, then Defaults are used.
 */

type Configuration struct {
    Size        int     // sets size of the identicon (range: 4-8)
    Square      bool    // creates a square identicon
    Border      bool    // adds a border to the identicon
    Vertical    bool    // creates identicon in portrait dimension (not visible on using --square flag)
    Invert      bool    // inverts the cell filling of identicon
    Symmetric   bool    // creates symmetric identicon
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

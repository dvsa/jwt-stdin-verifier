package main

import (
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	flagKey    = flag.String("key", "", "path to key file or '-' to read from stdin")
	flagDebug  = flag.Bool("debug", false, "print out all kinds of debug data")
	flagVerify = flag.String("verify", "", "path to JWT token to verify or '-' to read from stdin")
)

func main() {

	// Parse command line options
	flag.Parse()

	// Do the thing.  If something goes wrong, print error to stderr
	// and exit with a non-zero status code
	if err := verifyToken(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// Helper func:  Read input from specified file or stdin
func loadData(p string) ([]byte, error) {
	if p == "" {
		return nil, fmt.Errorf("no path specified")
	}

	var rdr io.Reader
	if p == "-" {
		rdr = os.Stdin
	} else if p == "+" {
		return []byte("{}"), nil
	} else {
		if f, err := os.Open(p); err == nil {
			rdr = f
			defer f.Close()
		} else {
			return nil, err
		}
	}
	return ioutil.ReadAll(rdr)
}

// Verify a token and output the claims.  This is a great example
// of how to verify and view a token.
func verifyToken() error {
	// get the token
	tokData, err := loadData(*flagVerify)
	if err != nil {
		return fmt.Errorf("couldn't read token: %w", err)
	}

	// trim possible whitespace from token
	tokData = regexp.MustCompile(`\s*$`).ReplaceAll(tokData, []byte{})
	if *flagDebug {
		fmt.Fprintf(os.Stderr, "Token len: %v bytes\n", len(tokData))
	}

	// Parse the token.  Load the key from command line option
	token, err := jwt.Parse(string(tokData), func(t *jwt.Token) (interface{}, error) {
		data, err := loadData(*flagKey)
		if err != nil {
			return nil, err
		}
		return jwt.ParseRSAPublicKeyFromPEM(data)
	})

	// Print some debug data
	if *flagDebug && token != nil {
		fmt.Fprintf(os.Stderr, "Header:\n%v\n", token.Header)
		fmt.Fprintf(os.Stderr, "Claims:\n%v\n", token.Claims)
	}

	// Print an error if we can't parse for some reason
	if err != nil {
		return fmt.Errorf("couldn't parse token: %w", err)
	}

	// Is token invalid?
	if !token.Valid {
		return fmt.Errorf("token is invalid")
	} else {
		return fmt.Errorf("token is valid")
	}
	git
	in
	return nil
}

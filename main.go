package main

import (
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"io/ioutil"
	"os"
)

var (
	flagVerify = flag.String("verify", "", "JWT token to verify")
)

func main() {

	// Parse command line options
	flag.Parse()

	// Do the thing.  If something goes wrong, print error to stderr
	// and exit with a non-zero status code
	if err := verifyToken(); err != nil {
		fmt.Println("INVALID")
		os.Exit(1)
	} else {
		fmt.Println("OK")
	}
}

func loadKey() ([]byte, error) {
	var keyFile = os.Getenv("JWT_VERIFY_KEY_FILE")

	if keyFile == "" {
		return nil, fmt.Errorf("")
	}

	var rdr io.Reader

	if f, err := os.Open(keyFile); err == nil {
		rdr = f
		defer f.Close()
	} else {
		return nil, err
	}

	return ioutil.ReadAll(rdr)
}

// Verify a token and output the claims.  This is a great example
// of how to verify and view a token.
func verifyToken() error {
	// get the token
	if *flagVerify == "" {
		return fmt.Errorf("")
	}

	// Parse the token.  Load the key from command line option
	token, err := jwt.Parse(string(*flagVerify), func(t *jwt.Token) (interface{}, error) {
		data, err := loadKey()
		if err != nil {
			return nil, err
		}
		return jwt.ParseRSAPublicKeyFromPEM(data)
	})

	// Print an error if we can't parse for some reason
	if err != nil {
		return fmt.Errorf("")
	}

	// Is token invalid?
	if !token.Valid {
		return fmt.Errorf("")
	}

	return nil
}

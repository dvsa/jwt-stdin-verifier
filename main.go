package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"io"
	"io/ioutil"
	"os"
)

var (
	flagKeyFile = flag.String("keyfile", "", "path to key file")
)

func main() {
	usage()

	flag.Parse()
	if *flagKeyFile == "" {
		fmt.Fprintln(os.Stderr, "No keyfile provided. Please specify public key with -keyfile <path>")
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := verifyToken(scanner.Text()); err != nil {
			fmt.Fprintln(os.Stdout, "INVALID -", err)
		} else {
			fmt.Println("OK")
		}
	}
}

func loadKey() ([]byte, error) {
	var rdr io.Reader
	if f, err := os.Open(*flagKeyFile); err == nil {
		rdr = f
		defer f.Close()
	} else {
		return nil, err
	}

	return ioutil.ReadAll(rdr)
}

func verifyToken(inputToken string) error {

	if inputToken == "" {
		return fmt.Errorf("no JWT provided")
	}
	
	_, err := jwt.Parse(inputToken, func(t *jwt.Token) (interface{}, error) {
		data, err := loadKey()
		if err != nil {
			return nil, err
		}
		return jwt.ParseRSAPublicKeyFromPEM(data)
	})

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func usage() {
	flag.Usage = func() {
		fmt.Println("    _          _                      _  __       ")
		fmt.Println("   (_)_      _| |_    __   _____ _ __(_)/ _|_   _ ")
		fmt.Println("   | \\ \\ /\\ / / __|___\\ \\ / / _ \\ '__| | |_| | | |")
		fmt.Println("   | |\\ V  V /| ||_____\\ V /  __/ |  | |  _| |_| |")
		fmt.Println("  _/ | \\_/\\_/  \\__|     \\_/ \\___|_|  |_|_|  \\__, |")
		fmt.Println(" |__/                                       |___/ ")
		fmt.Println("Watch stdin and verify JWT it receives")
		fmt.Println("Usage of jwt-verify:")
		flag.PrintDefaults()
	}
}

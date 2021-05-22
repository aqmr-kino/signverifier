package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/clearsign"
)

func main() {
	// Check arguments
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatalln("error: argument error, need to specify verification public key file")
	}

	// Open verification key file
	keyringFile, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Println("error: open key file")
		log.Fatalln(err)
	}

	// Load verification key
	keyring, err := openpgp.ReadArmoredKeyRing(keyringFile)
	if err != nil {
		log.Println("error: load key file")
		log.Fatalln(reflect.TypeOf(err), err)
	}

	// Load clearsigned message from stdin
	dataBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Println("error: read from stdin")
		log.Fatalln(reflect.TypeOf(err), err)
	}

	// Decord message
	msg, _ := clearsign.Decode(dataBytes)
	if msg == nil {
		log.Println("error: decord message")
		log.Fatalln("input message is invalid, it must be clearsign text")
	}

	// Verify signature
	_, err = openpgp.CheckDetachedSignature(keyring, bytes.NewReader(msg.Bytes), msg.ArmoredSignature.Body)
	if err != nil {
		log.Println("error: verify message")
		log.Fatalln(reflect.TypeOf(err), err)
	}

	// Output message to stdout
	os.Stdout.Write(msg.Plaintext)
}

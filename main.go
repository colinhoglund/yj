package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	in, err := read(os.Stdin, os.Args)
	logFatal(err)

	out, err := marshal(in)
	logFatal(err)

	fmt.Println(string(out))
}

func read(stdin io.Reader, args []string) (in []byte, err error) {
	switch len(args) {
	case 1:
		in, err = ioutil.ReadAll(stdin)
	case 2:
		in, err = ioutil.ReadFile(args[1])
	default:
		err = errTooManyArgs
	}

	return
}

func marshal(in []byte) (out []byte, err error) {
	var data interface{}

	switch json.Unmarshal(in, &data) {
	case nil:
		out, err = yaml.JSONToYAML(in)
	default:
		out, err = yaml.YAMLToJSON(in)
	}

	out = bytes.TrimSpace(out)
	return
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

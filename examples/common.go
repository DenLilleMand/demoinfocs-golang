package examples

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// DemoPathFromArgs returns the value of the -demo command line flag.
func DemoPathFromArgs() string {
	fl := new(flag.FlagSet)

	demPathPtr := fl.String("demo", "", "Demo file `path`")

	err := fl.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	demPath := *demPathPtr

	return demPath
}

// RedirectStdout redirects standard output to dev null.
func RedirectStdout(f func()) {
	// Redirect stdout, the resulting image is written to this
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout = w

	// Discard the output in a separate goroutine so writing to stdout can't block indefinitely
	go func() {
		for err := error(nil); err == nil; _, err = io.Copy(ioutil.Discard, r) {
		}
	}()

	f()

	os.Stdout = old
}

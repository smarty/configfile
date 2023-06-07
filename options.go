package configfile

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"
)

func New(pointer any, options ...func(*config)) error {
	if len(options) == 0 {
		options = append(options, Options.JSONFileViaCLI("config", "config.json"))
	}
	var c config
	for _, option := range options {
		option(&c)
	}
	if len(c.flag) > 0 {
		flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
		flags.StringVar(&c.filename, c.flag, c.flagFile, "Path to a JSON configuration file.")
		_ = flags.Parse(os.Args[1:])
	}
	if len(c.filename) > 0 {
		raw, err := os.ReadFile(c.filename)
		if err != nil {
			return err
		}
		c.reader = bytes.NewReader(raw)
	}
	if c.reader == nil {
		return ErrMissingSource
	}
	return json.NewDecoder(c.reader).Decode(pointer)
}

var Options singleton

type config struct {
	flag     string
	flagFile string
	filename string
	reader   io.Reader
}

type singleton struct{}

func (singleton) JSONReader(r io.Reader) func(*config) {
	return func(c *config) { c.reader = r }
}
func (singleton) JSONFile(f string) func(*config) {
	return func(c *config) { c.filename = f }
}
func (singleton) JSONFileViaCLI(flag, defaultFilename string) func(*config) {
	return func(c *config) { c.flag, c.flagFile = flag, defaultFilename }
}

var ErrMissingSource = errors.New("missing config file source")

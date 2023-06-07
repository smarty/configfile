package configfile

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"
)

func New[T any](options ...func(*config[T])) (dto T, err error) {
	if len(options) == 0 {
		options = append(options, OptionJSONFileViaCLI[T]("config", "config.json"))
	}
	var c config[T]
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
			return dto, err
		}
		c.reader = bytes.NewReader(raw)
	}
	if c.reader == nil {
		return dto, ErrMissingSource
	}
	err = json.NewDecoder(c.reader).Decode(&dto)
	return dto, err
}

type config[T any] struct {
	flag     string
	flagFile string
	filename string
	reader   io.Reader
}

func OptionJSONReader[T any](r io.Reader) func(*config[T]) {
	return func(c *config[T]) { c.reader = r }
}
func OptionJSONFile[T any](f string) func(*config[T]) {
	return func(c *config[T]) { c.filename = f }
}
func OptionJSONFileViaCLI[T any](flag, defaultFilename string) func(*config[T]) {
	return func(c *config[T]) { c.flag, c.flagFile = flag, defaultFilename }
}

var ErrMissingSource = errors.New("missing config file source")

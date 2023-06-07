package configfile

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

func LoadJSONFileViaCLI(pointer any, flagName, defaultFilename string) error {
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	filename := flags.String(flagName, defaultFilename, "Path to a JSON configuration file.")
	_ = flags.Parse(os.Args[1:])
	return LoadJSONFile(pointer, *filename)
}
func LoadJSONFile(pointer any, filename string) error {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("configuration error: unable to read JSON configuration file: %w", err)
	}
	return LoadJSON(pointer, bytes.NewReader(raw))
}
func LoadJSON(pointer any, reader io.Reader) error {
	err := json.NewDecoder(reader).Decode(pointer)
	if err != nil {
		return fmt.Errorf("configuration error: unable to parse JSON configuration: %w", err)
	}
	return nil
}

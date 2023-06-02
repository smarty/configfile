package configfile

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
)

func LoadDefault[T any]() (dto T, err error) {
	return LoadJSONFileViaCLI[T]("config", "config.json")
}
func LoadJSONFileViaCLI[T any](flagName, defaultFilename string) (dto T, err error) {
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	filename := flags.String(flagName, defaultFilename, "Path to a JSON configuration file.")
	_ = flags.Parse(os.Args[1:])
	return LoadJSONFile[T](*filename)
}
func LoadJSONFile[T any](filename string) (dto T, err error) {
	raw, err := os.ReadFile(filename)
	if err != nil {
		return dto, fmt.Errorf("configuration error: unable to read JSON configuration file: %w", err)
	}
	return LoadJSON[T](bytes.NewReader(raw))
}
func LoadJSON[T any](reader io.Reader) (dto T, err error) {
	err = json.NewDecoder(reader).Decode(&dto)
	if err != nil {
		err = fmt.Errorf("configuration error: unable to parse JSON configuration: %w", err)
	}
	return dto, err
}

func URLServerName(rawURL string) string {
	endpoint, _ := url.Parse(rawURL)
	if endpoint == nil {
		return ""
	}
	return Coalesce(endpoint.Query().Get("server-name"), endpoint.Hostname())
}
func Coalesce[T comparable](values ...T) (zero T) {
	for _, item := range values {
		if item != zero {
			return item
		}
	}
	return zero
}

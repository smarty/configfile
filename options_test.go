package configfile_test

import (
	"bytes"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/smarty/configfile"
	"github.com/smarty/configfile/should"
)

func TestOptionJSONReader_NilReader(t *testing.T) {
	var c ExampleConfig
	err := configfile.New(&c,
		configfile.Options.JSONReader(nil),
	)
	should.So(t, err, should.NOT.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{})
}
func TestOptionJSONReader_BadJSON(t *testing.T) {
	var c ExampleConfig
	err := configfile.New(&c,
		configfile.Options.JSONReader(bytes.NewBufferString(`BAD JSON`)),
	)
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
}
func TestOptionJSONReader_GoodJSON(t *testing.T) {
	var c ExampleConfig
	err := configfile.New(&c,
		configfile.Options.JSONReader(bytes.NewBufferString(`{"message": "Hello, world!"}`)),
	)
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}

func TestOptionJSONFile_BadJSON(t *testing.T) {
	_, here, _, _ := runtime.Caller(0)
	var c ExampleConfig
	err := configfile.New(&c,
		configfile.Options.JSONFile(filepath.Join(filepath.Dir(here), "config-bad.json")),
	)
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
}
func TestOptionJSONFile_GoodJSON(t *testing.T) {
	_, here, _, _ := runtime.Caller(0)
	var c ExampleConfig
	err := configfile.New(&c,
		configfile.Options.JSONFile(filepath.Join(filepath.Dir(here), "testdata", "config-good.json")),
	)
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}

func TestOptionJSONFileViaCLI_BadJSON(t *testing.T) {
	var c ExampleConfig
	err := configfile.New(&c,
		configfile.Options.JSONFileViaCLI("config", "testdata/config-bad.json"),
	)
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
}
func TestOptionJSONFileViaCLI_GoodJSON(t *testing.T) {
	var c ExampleConfig
	err := configfile.New(&c,
		configfile.Options.JSONFileViaCLI("config", "testdata/config-good.json"),
	)
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}

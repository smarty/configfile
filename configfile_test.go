package configfile_test

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/smarty/configfile/v2"
	"github.com/smarty/configfile/v2/internal/should"
)

type ExampleConfig struct {
	Message string `json:"message"`
}

func TestLoadJSON_BadJSON(t *testing.T) {
	var c ExampleConfig
	err := configfile.LoadJSON(&c, bytes.NewBufferString(`BAD JSON`))
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
	t.Log(err)
}
func TestLoadJSON_GoodJSON(t *testing.T) {
	var c ExampleConfig
	err := configfile.LoadJSON(&c, bytes.NewBufferString(`{"message": "Hello, world!"}`))
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}

func TestLoadJSONFile_BadJSON(t *testing.T) {
	var c ExampleConfig
	_, here, _, _ := runtime.Caller(0)
	err := configfile.LoadJSONFile(&c, filepath.Join(filepath.Dir(here), "testdata", "config-bad.json"))
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
}
func TestLoadJSONFile_GoodJSON(t *testing.T) {
	var c ExampleConfig
	_, here, _, _ := runtime.Caller(0)
	err := configfile.LoadJSONFile(&c, filepath.Join(filepath.Dir(here), "testdata", "config-good.json"))
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}

func TestLoadJSONFileViaCLI_BadJSON(t *testing.T) {
	var c ExampleConfig
	err := configfile.LoadJSONFileViaCLI(&c, "config", "testdata/config-bad.json")
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
}
func TestLoadJSONFileViaCLI_GoodJSON(t *testing.T) {
	var c ExampleConfig
	err := configfile.LoadJSONFileViaCLI(&c, "config", "testdata/config-good.json")
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}
func TestLoadJSONFileViaCLI_MissingFile_Created(t *testing.T) {
	var c ExampleConfig
	err := configfile.LoadJSONFileViaCLI(&c, "config", "testdata/missing.json")
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, errors.Is(err, os.ErrNotExist), should.BeTrue)
	info, err := os.Stat(filepath.Join("testdata", "missing.json"))
	should.So(t, err, should.BeNil)
	should.So(t, info.Size(), should.NOT.Equal, 0)
	_ = os.Remove(filepath.Join("testdata", "missing.json"))
}

package configfile_test

import (
	"bytes"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/smarty/configfile"
	"github.com/smarty/configfile/should"
)

type ExampleConfig struct {
	Message string `json:"message"`
}

func TestLoadJSON_BadJSON(t *testing.T) {
	c, err := configfile.LoadJSON[ExampleConfig](bytes.NewBufferString(`BAD JSON`))
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
}
func TestLoadJSON_GoodJSON(t *testing.T) {
	c, err := configfile.LoadJSON[ExampleConfig](bytes.NewBufferString(`{"message": "Hello, world!"}`))
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}

func TestLoadJSONFile_BadJSON(t *testing.T) {
	_, here, _, _ := runtime.Caller(0)
	c, err := configfile.LoadJSONFile[ExampleConfig](filepath.Join(filepath.Dir(here), "config-bad.json"))
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
}
func TestLoadJSONFile_GoodJSON(t *testing.T) {
	_, here, _, _ := runtime.Caller(0)
	c, err := configfile.LoadJSONFile[ExampleConfig](filepath.Join(filepath.Dir(here), "testdata", "config-good.json"))
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}

func TestLoadJSONFileViaCLI_BadJSON(t *testing.T) {
	c, err := configfile.LoadJSONFileViaCLI[ExampleConfig]("config", "testdata/config-bad.json")
	should.So(t, c, should.Equal, ExampleConfig{})
	should.So(t, err, should.NOT.BeNil)
}
func TestLoadJSONFileViaCLI_GoodJSON(t *testing.T) {
	c, err := configfile.LoadJSONFileViaCLI[ExampleConfig]("config", "testdata/config-good.json")
	should.So(t, err, should.BeNil)
	should.So(t, c, should.Equal, ExampleConfig{Message: "Hello, world!"})
}

func TestCoalesce(t *testing.T) {
	should.So(t, configfile.Coalesce("", "", "a"), should.Equal, "a")
	should.So(t, configfile.Coalesce("", "a", "b"), should.Equal, "a")
	should.So(t, configfile.Coalesce("a", "b", "c"), should.Equal, "a")
}
func TestURLServerName(t *testing.T) {
	should.So(t, configfile.URLServerName("%%%-not-a-url-%%%"), should.Equal, "")
	should.So(t, configfile.URLServerName("amqp://127.0.0.1:5672/"), should.Equal, "127.0.0.1")
	should.So(t, configfile.URLServerName("amqp://127.0.0.1:5672/?server-name=the-server"), should.Equal, "the-server")
}

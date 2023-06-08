package configfile_test

import (
	"encoding/json"
	"testing"

	"github.com/smarty/configfile/v2"
	"github.com/smarty/configfile/v2/internal/should"
)

type URLThing struct {
	Address configfile.URL `json:"address"`
}

func TestURLUnmarshal_NullConvention(t *testing.T) {
	var thing URLThing
	err := json.Unmarshal([]byte(`{"address":null}`), &thing)
	should.So(t, err, should.BeNil)
	should.So(t, thing.Address, should.Equal, configfile.URL{})
}
func TestURLUnmarshal_EmptyString(t *testing.T) {
	var thing URLThing
	err := json.Unmarshal([]byte(`{"address":""}`), &thing)
	should.So(t, err, should.BeNil)
	should.So(t, thing.Address, should.Equal, configfile.URL{})
}
func TestURLUnmarshal_BadValue(t *testing.T) {
	var thing URLThing
	err := json.Unmarshal([]byte(`{"address":42}`), &thing)
	should.So(t, err, should.NOT.BeNil)
	should.So(t, thing.Address, should.Equal, configfile.URL{})
}
func TestURLUnmarshal_BadJSON(t *testing.T) {
	var thing URLThing
	err := json.Unmarshal([]byte(`{"address":"%%%%%% not a url %%%%%"}`), &thing)
	should.So(t, err, should.NOT.BeNil)
	should.So(t, thing.Address, should.Equal, configfile.URL{})
}
func TestURLUnmarshal_GoodJSON(t *testing.T) {
	var thing URLThing
	err := json.Unmarshal([]byte(`{"address":"https://smarty.com"}`), &thing)
	should.So(t, err, should.BeNil)
	should.So(t, thing.Address.URL.String(), should.Equal, "https://smarty.com")
}

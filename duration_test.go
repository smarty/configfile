package configfile_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/smarty/configfile/v2"
	"github.com/smarty/configfile/v2/internal/should"
)

type DurationThing struct {
	Timeout configfile.Duration `json:"timeout"`
}

func TestDurationUnmarshal_NullConvention(t *testing.T) {
	var thing DurationThing
	err := json.Unmarshal([]byte(`{"timeout":null}`), &thing)
	should.So(t, err, should.BeNil)
	should.So(t, thing.Timeout, should.Equal, configfile.Duration{})
}
func TestDurationUnmarshal_BadValue(t *testing.T) {
	var thing DurationThing
	err := json.Unmarshal([]byte(`{"timeout":42}`), &thing)
	should.So(t, err, should.NOT.BeNil)
	should.So(t, thing.Timeout, should.Equal, configfile.Duration{})
}
func TestDurationUnmarshal_BadJSON(t *testing.T) {
	var thing DurationThing
	err := json.Unmarshal([]byte(`{"timeout":"not a duration"}`), &thing)
	should.So(t, err, should.NOT.BeNil)
	should.So(t, thing.Timeout, should.Equal, configfile.Duration{})
}
func TestDurationUnmarshal_GoodJSON(t *testing.T) {
	var thing DurationThing
	err := json.Unmarshal([]byte(`{"timeout":"2s"}`), &thing)
	should.So(t, err, should.BeNil)
	should.So(t, thing.Timeout.Duration, should.Equal, time.Second*2)
}

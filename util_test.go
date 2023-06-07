package configfile_test

import (
	"testing"

	"github.com/smarty/configfile/v2"
	"github.com/smarty/configfile/v2/internal/should"
)

func TestCoalesce(t *testing.T) {
	should.So(t, configfile.Coalesce[string](), should.Equal, "")
	should.So(t, configfile.Coalesce[string](""), should.Equal, "")
	should.So(t, configfile.Coalesce("", "", "a"), should.Equal, "a")
	should.So(t, configfile.Coalesce("", "a", "b"), should.Equal, "a")
	should.So(t, configfile.Coalesce("a", "b", "c"), should.Equal, "a")
}

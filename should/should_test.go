package should_test

import (
	"testing"
	"time"

	"github.com/smarty/configfile/should"
)

func TestPassing(t *testing.T) {
	should.So(t, 1, should.Equal, 1)
	should.So(t, false, should.BeFalse)
	should.So(t, true, should.BeTrue)
	should.So(t, nil, should.BeNil)
	should.So(t, 1, should.NOT.BeNil)
	now1 := time.Now()
	now2 := now1.In(time.UTC)
	should.So(t, now1, should.Equal, now2)
}
func TestFailing(t *testing.T) {
	t.Skip("comment me to see the failures below")
	should.So(nil, 1, should.Equal, 2)
	should.So(t, 1, should.Equal, 2)
	should.So(t, true, should.BeFalse)
	should.So(t, false, should.BeTrue)
	should.So(t, 1, should.BeNil)
	should.So(t, nil, should.NOT.BeNil)
	should.So(t, uint64(1), should.Equal, uint64(2))
	should.So(t, time.Now(), should.Equal, time.Now())
}

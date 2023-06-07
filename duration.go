package configfile

import (
	"encoding/json"
	"fmt"
	"time"
)

type Duration struct{ time.Duration }

func (this *Duration) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return nil
	}
	var v any
	_ = json.Unmarshal(b, &v)
	raw, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid duration: %#v", v)
	}
	this.Duration, err = time.ParseDuration(raw)
	return err
}

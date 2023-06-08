package configfile

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type URL struct{ URL *url.URL }

func (this *URL) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return nil
	}
	var v any
	_ = json.Unmarshal(b, &v)
	raw, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid url: %#v", v)
	}
	this.URL, err = url.Parse(raw)
	return err
}

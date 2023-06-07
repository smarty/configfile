package configfile

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type URL struct{ url.URL }

func (this *URL) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	var v any
	_ = json.Unmarshal(b, &v)
	raw, ok := v.(string)
	if !ok {
		return fmt.Errorf("invalid url: %#v", v)
	}
	u, err := url.Parse(raw)
	if err != nil {
		return err
	}
	this.URL = *u
	return nil
}

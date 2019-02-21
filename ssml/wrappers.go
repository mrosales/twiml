package ssml

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strings"
	"time"
)

type PauseStrength string

func (p PauseStrength) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: string(p)}, nil
}

type Duration time.Duration

func (d Duration) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	durationMillis := int64(time.Duration(d) / time.Millisecond)
	return xml.Attr{Name: name, Value: fmt.Sprintf("%dms", durationMillis)}, nil
}

type URL url.URL

func (u URL) MarshalText() ([]byte, error) {
	raw := url.URL(u)
	return []byte(raw.String()), nil
}

func (u *URL) UnmarshalText(data []byte) error {
	raw, err := url.Parse(string(data))
	if err != nil {
		return err
	}
	*u = URL(*raw)
	return nil
}

var fmtToFormatReplacer = strings.NewReplacer("y", "2006", "m", "01", "d", "02")

func (f DateFormat) Format(t time.Time) string {
	format := f
	if format == "" {
		format = DateFormatDefault
	}
	layout := fmtToFormatReplacer.Replace(string(format))
	return t.Format(layout)
}

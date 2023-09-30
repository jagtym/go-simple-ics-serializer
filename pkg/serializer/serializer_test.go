package serializer

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestSerializeCalendar(t *testing.T) {
	t.Run("test basic example", func(t *testing.T) {
		buff := bytes.Buffer{}

		event := Event{
			UUID:        "example",
			Title:       "This is summary",
			Description: "This is description",
			DateStart:   time.Unix(0, 0),
			DateEnd:     time.Unix(0, 0).Add(time.Minute * 15),
		}
		cal := Calendar{
			Events: []Event{event},
		}
		cal.Serialize(&buff)

		want := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//example//example//EN
BEGIN:VEVENT
UID:example
DTSTAMP:19700101T000000Z
DTSTART:19700101T000000Z
DTEND:19700101T001500Z
SUMMARY:This is summary
DESCRIPTION:This is description
END:VEVENT
END:VCALENDAR`

		got := buff.String()
		if diff := cmp.Diff(got, want); diff != "" {
			t.Error(diff)
		}
	})
}

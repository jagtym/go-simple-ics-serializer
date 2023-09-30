package serializer

import (
	"fmt"
	"io"
	"time"
)

const (
	dateFormat = "20060102T150405Z"
)

type Event struct {
	UUID        string
	Title       string
	Description string
	DateStart   time.Time
	DateEnd     time.Time
}
type Calendar struct {
	Events []Event
}

func (c Calendar) Serialize(w io.Writer) {
	c.serialize(w)
}

func (c Calendar) serialize(w io.Writer) {
	intro := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//example//example//EN`

	fmt.Fprintf(w, "%s\n", intro)

	for _, e := range c.Events {
		e.serialize(w)
	}

	outro := "END:VCALENDAR"
	fmt.Fprintf(w, "%s", outro)
}

func (e Event) serialize(w io.Writer) {
	fmt.Fprintf(w, "BEGIN:VEVENT\n")
	e.addEventData(w)
	fmt.Fprintf(w, "END:VEVENT\n")
}

func (e Event) addEventData(w io.Writer) {
	fmt.Fprintf(w, "UID:%s\n", e.UUID)

	addTimestamps(w, e.DateStart, e.DateEnd)

	fmt.Fprintf(w, "SUMMARY:%s\n", e.Title)
	fmt.Fprintf(w, "DESCRIPTION:%s\n", e.Description)

}

func addTimestamps(w io.Writer, startDate time.Time, endDate time.Time) {
	var (
		start = startDate.UTC().Format(dateFormat)
		end   = endDate.UTC().Format(dateFormat)
	)

	fmt.Fprintf(w, "DTSTAMP:%s\n", start)
	fmt.Fprintf(w, "DTSTART:%s\n", start)
	fmt.Fprintf(w, "DTEND:%s\n", end)
}

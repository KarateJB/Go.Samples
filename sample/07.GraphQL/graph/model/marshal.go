package model

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type Time time.Time
type Date time.Time
type DateTime time.Time
type UUID uuid.UUID

// Creates a marshaller which converts a time.Time (time) to a string
func MashalTime(time time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, e := io.WriteString(w, fmt.Sprintf("%s%s%s", "\"", time.Format("15:04:05-0700"), "\""))
		if e != nil {
			panic(e)
		}
	})
}

// Unmarshalls a string to a time.Time (time)
func UnmarshalTime(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("time must be strings")
	}
	withoutQuotes := strings.ReplaceAll(str, "\"", "")
	i, err := time.Parse("15:04:05-0700", withoutQuotes)
	if err != nil {
		i, err = time.Parse("15:04:05", withoutQuotes)
	}
	return i, err
}

// Creates a marshaller which converts a time.Time (date) to a string
func MashalDate(date time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, e := io.WriteString(w, fmt.Sprintf("%s%s%s", "\"", date.Format("2006-01-02"), "\""))
		if e != nil {
			panic(e)
		}
	})
}

// Unmarshalls a string to a time.Time (date)
func UnmarshalDate(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("date must be strings")
	}
	withoutQuotes := strings.ReplaceAll(str, "\"", "")
	i, err := time.Parse("2006-01-02", withoutQuotes)
	if err != nil {
		i, err = time.Parse("20060102", withoutQuotes)
	}
	return i, err
}

// Creates a marshaller which converts a time.Time (date time) to a string
func MashalDateTime(dateTime time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, e := io.WriteString(w, fmt.Sprintf("%s%s%s", "\"", dateTime.Format("2006-01-02T15:04:05-0700"), "\""))
		if e != nil {
			panic(e)
		}
	})
}

// Unmarshalls a string to a time.Time (date time)
func UnmarshalDateTime(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("datetime must be strings")
	}
	withoutQuotes := strings.ReplaceAll(str, "\"", "")
	i, err := time.Parse("2006-01-02T15:04:05-0700", withoutQuotes)
	if err != nil {
		i, err = time.Parse("2006-01-02T15:04:05", withoutQuotes)
	}
	return i, err
}

// Creates a marshaller which converts a uuid to a string
func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, e := io.WriteString(w, fmt.Sprintf("%s%s%s", "\"", id.String(), "\""))
		if e != nil {
			panic(e)
		}
	})
}

// Unmarshalls a string to a uuid
func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	str, ok := v.(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("ids must be strings")
	}
	withoutQuotes := strings.ReplaceAll(str, "\"", "")
	i, err := uuid.Parse(withoutQuotes)
	return i, err
}

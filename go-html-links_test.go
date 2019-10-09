package links

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

type TestData struct {
	inputs   io.Reader
	result   []string
	hasError bool
}

func TestAll(t *testing.T) {

	data := []TestData{
		{strings.NewReader(``), []string{}, false},
		{strings.NewReader(`<html><body><a href='https://www.google.com'>Google</a></body></html>`), []string{"https://www.google.com"}, false},
		{strings.NewReader(`<html><body><a href='https://www.google.com'>Google</a><a href='https://www.google.com'>Google</a></body></html>`), []string{"https://www.google.com"}, false},
		{strings.NewReader(`<html><body><a href='https://www.google.com'>Google</a><a href='https://www.facebook.com'>Facebook</a></body></html>`), []string{"https://www.google.com", "https://www.facebook.com"}, false},
	}

	for _, dt := range data {
		result := All(dt.inputs)

		if !reflect.DeepEqual(dt.result, result) {
			t.Errorf("All() with args %v : FAILED, expected %v but got value %v", dt.inputs, dt.result, result)
		}
	}
}

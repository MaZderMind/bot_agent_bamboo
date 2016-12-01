package main

import (
	"testing"

	"os"

	. "github.com/bborbe/assert"
	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

func TestCreateRequestConsumer(t *testing.T) {
	createRequestConsumer, err := createRequestConsumer()
	if err := AssertThat(err, NotNilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(createRequestConsumer, NilValue()); err != nil {
		t.Fatal(err)
	}
}

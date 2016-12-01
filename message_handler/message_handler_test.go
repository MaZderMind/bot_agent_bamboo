package message_handler

import (
	"testing"

	"os"

	. "github.com/bborbe/assert"
	"github.com/bborbe/bot_agent/api"
	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

func TestImplementsAgent(t *testing.T) {
	c := New(nil)
	var i *api.MessageHandler
	if err := AssertThat(c, Implements(i)); err != nil {
		t.Fatal(err)
	}
}

func TestMessageWithBamboo(t *testing.T) {
	c := New(nil)
	responses, err := c.HandleMessage(&api.Request{
		Message: "bamboo botname",
		From: &api.User{
			UserName: "username",
		},
		Bot: "botname",
	})
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(responses), Is(1)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(responses[0].Replay, Is(api.ResponseReplay(false))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(responses[0].Message, Is(api.ResponseMessage("bamboo username"))); err != nil {
		t.Fatal(err)
	}
}

func TestMessageWithoutBamboo(t *testing.T) {
	c := New(nil)
	responses, err := c.HandleMessage(&api.Request{Message: "foo"})
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(responses), Is(0)); err != nil {
		t.Fatal(err)
	}
}

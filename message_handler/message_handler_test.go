package message_handler

import (
	"testing"

	"os"

	. "github.com/bborbe/assert"
	"github.com/bborbe/bot_agent/api"
	"github.com/golang/glog"
	"errors"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

type mockDeployer  struct {
	counter int
	result  error
}

func (m *mockDeployer) Deploy() error {
	m.counter++
	return m.result
}

func TestImplementsAgent(t *testing.T) {
	deployer := new(mockDeployer)
	c := New(deployer)
	var i *api.MessageHandler
	if err := AssertThat(c, Implements(i)); err != nil {
		t.Fatal(err)
	}
}

func TestMessageWithBamboo(t *testing.T) {
	deployer := new(mockDeployer)
	c := New(deployer)
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
	if err := AssertThat(responses[0].Message, Is(api.ResponseMessage("deployment triggered succcesful"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(deployer.counter, Is(1)); err != nil {
		t.Fatal(err)
	}
}

func TestMessageWithBambooFailure(t *testing.T) {
	deployer := new(mockDeployer)
	deployer.result = errors.New("fail")
	c := New(deployer)
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
	if err := AssertThat(responses[0].Message, Is(api.ResponseMessage("trigger deployment failed: fail"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(deployer.counter, Is(1)); err != nil {
		t.Fatal(err)
	}
}


func TestMessageWithoutBamboo(t *testing.T) {
	deployer := new(mockDeployer)
	c := New(deployer)
	responses, err := c.HandleMessage(&api.Request{Message: "foo"})
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(len(responses), Is(0)); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(deployer.counter, Is(0)); err != nil {
		t.Fatal(err)
	}
}

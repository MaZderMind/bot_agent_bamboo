package message_handler

import (
	"errors"
	"os"
	"testing"

	. "github.com/bborbe/assert"
	auth_model "github.com/bborbe/auth/model"
	"github.com/bborbe/bot_agent/api"
	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

type mockDeployer struct {
	counter int
	number  int
	result  error
}

func (m *mockDeployer) Deploy(number int) error {
	m.number = number
	m.counter++
	return m.result
}

func TestImplementsAgent(t *testing.T) {
	deployer := new(mockDeployer)
	c := New("/deploy", deployer, func(auth_model.AuthToken) bool {
		return true
	})
	var i *api.MessageHandler
	if err := AssertThat(c, Implements(i)); err != nil {
		t.Fatal(err)
	}
}

func TestMessageWithBamboo(t *testing.T) {
	deployer := new(mockDeployer)
	c := New("/deploy", deployer, func(auth_model.AuthToken) bool {
		return true
	})
	responses, err := c.HandleMessage(&api.Request{
		Message: "/deploy 123",
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
	if err := AssertThat(deployer.number, Is(123)); err != nil {
		t.Fatal(err)
	}
}

func TestMessageWithBambooFailure(t *testing.T) {
	deployer := new(mockDeployer)
	deployer.result = errors.New("fail")
	c := New("/deploy", deployer, func(auth_model.AuthToken) bool {
		return true
	})
	responses, err := c.HandleMessage(&api.Request{
		Message: "/deploy 123",
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
	if err := AssertThat(deployer.number, Is(123)); err != nil {
		t.Fatal(err)
	}
}

package bamboo

import (
	. "github.com/bborbe/assert"
	"github.com/golang/glog"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

func TestImplementsDeployer(t *testing.T) {
	c := NewDeployer("url", "user", "pass")
	var i *Deployer
	if err := AssertThat(c, Implements(i)); err != nil {
		t.Fatal(err)
	}
}

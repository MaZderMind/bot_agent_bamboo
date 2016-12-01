package bamboo

import (
	"os"
	"testing"

	. "github.com/bborbe/assert"
	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

func TestImplementsDeployer(t *testing.T) {
	c := NewDeployer(nil, "url", "user", "pass")
	var i *Deployer
	if err := AssertThat(c, Implements(i)); err != nil {
		t.Fatal(err)
	}
}

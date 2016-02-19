package deploy

import (
	"log"
	"os"
	"testing"

	tools "rsprd.com/spread/pkg/deploy/test"

	"github.com/stretchr/testify/assert"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"k8s.io/kubernetes/test/integration/framework"
)

var KubeClient *client.Client
var Etcd *tools.EtcdInstance

func TestMain(m *testing.M) {
	logger := nil
	// setup etcd
	Etcd = tools.NewEtcdInstance(logger)
	Etcd.Start()
	etcdClient := Etcd.Client()

	// setup kube api server
	config := tools.NewMasterConfig(etcdClient)
	master := framework.NewMasterComponents(&framework.Config{MasterConfig: config})
	KubeClient = master.RestClient

	// run tests
	status := m.Run()

	// cleanup
	master.Stop(true, true)
	master.ApiServer.Close() // fix until stop works on api server
	Etcd.Stop()

	os.Exit(status)
}

func TestClusterSetup(t *testing.T) {
	defer Etcd.Reset()

	_, err := NewKubeClusterFromContext("")
	assert.NoError(t, err)
}

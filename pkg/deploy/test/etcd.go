package test

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"time"

	etcd "github.com/coreos/etcd/client"
	"github.com/coreos/etcd/etcdserver"
	"github.com/coreos/etcd/etcdserver/etcdhttp"
	"github.com/coreos/etcd/pkg/transport"
	"github.com/coreos/etcd/pkg/types"
	"github.com/coreos/etcd/rafthttp"
)

// nextListenPort is the next port to be tried
var nextListenPort int64 = 20000

// EtcdInstance manages an etcd server for testing.
// Much of this code has been modified from etcds integration suite which has been copied for stability.
type EtcdInstance struct {
	etcdserver.ServerConfig
	ClientListener net.Listener
	etcd           *etcdserver.EtcdServer
	http           *httptest.Server
	*log.Logger
}

func NewEtcdInstance(l *log.Logger) *EtcdInstance {
	listen := newLocalListener(l)
	clients, err := types.NewURLs([]string{"http://" + listen.Addr().String()})
	if err != nil {
		l.Fatal(err)
	}

	return &EtcdInstance{
		ServerConfig: etcdserver.ServerConfig{
			ClientURLs:          clients,
			Name:                "erin",
			InitialClusterToken: "FREE50",
			NewCluster:          true,
		},
		ClientListener: listen,
		Logger:         l,
	}
}

func (i *EtcdInstance) Start() {
	var err error
	i.DataDir, err = ioutil.TempDir(os.TempDir(), "spread-etcd")
	if err != nil {
		i.Fatalf("unable to create data directory: %v", err)
	}

	i.etcd, err = etcdserver.NewServer(&i.ServerConfig)
	if err != nil {
		i.Fatalf("could not start etcd: %v", err)
	}

	i.http = &httptest.Server{
		Listener: i.ClientListener,
		Config:   &http.Server{Handler: etcdhttp.NewClientHandler(i.etcd, i.ServerConfig.ReqTimeout())},
	}

	i.etcd.Start()
	i.http.Start()
}

func (i *EtcdInstance) Stop() {
	if i.etcd == nil {
		i.Println("tried to stop server that was already stopped")
	}

	if i.etcd != nil {
		i.etcd.Stop()
		i.etcd = nil
	}

	if i.http != nil {
		i.http.CloseClientConnections()
		i.http.Close()
		i.http = nil
	}

	os.RemoveAll(i.DataDir)
}

func (i *EtcdInstance) Reset() {
	i.Stop()
	i.Start()
}

func (i *EtcdInstance) Client() etcd.Client {
	cfg := etcd.Config{Transport: mustNewTransport(i.Logger, transport.TLSInfo{}), Endpoints: i.ClientURLs.StringSlice()}
	c, err := etcd.New(cfg)
	if err != nil {
		i.Fatal(err)
	}
	return c
}

func newLocalListener(t *log.Logger) net.Listener {
	nextListenPort++
	l, err := net.Listen("tcp", "127.0.0.1:"+strconv.FormatInt(nextListenPort, 10))
	if err != nil {
		t.Fatal(err)
	}
	return l
}

func mustNewTransport(t *log.Logger, tlsInfo transport.TLSInfo) *http.Transport {
	// tick in integration test is short, so 1s dial timeout could play well.
	tr, err := transport.NewTimeoutTransport(tlsInfo, time.Second, rafthttp.ConnReadTimeout, rafthttp.ConnWriteTimeout)
	if err != nil {
		t.Fatal(err)
	}
	return tr
}

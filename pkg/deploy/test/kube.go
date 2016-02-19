package test

import (
	etcd "github.com/coreos/etcd/client"
	kube "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/testapi"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apiserver"
	"k8s.io/kubernetes/pkg/genericapiserver"
	kubeletclient "k8s.io/kubernetes/pkg/kubelet/client"
	"k8s.io/kubernetes/pkg/master"
	etcdstorage "k8s.io/kubernetes/pkg/storage/etcd"
	"k8s.io/kubernetes/pkg/storage/etcd/etcdtest"
	"k8s.io/kubernetes/plugin/pkg/admission/admit"
	"k8s.io/kubernetes/test/integration/framework"
)

// Returns a basic master config.
func NewMasterConfig(etcd etcd.Client) *master.Config {
	storageVersions := make(map[string]string)

	etcdStorage := etcdstorage.NewEtcdStorage(etcd, testapi.Default.Codec(), etcdtest.PathPrefix(), false)
	storageVersions[kube.GroupName] = testapi.Default.GroupVersion().String()
	expEtcdStorage := framework.NewExtensionsEtcdStorage(etcd)
	storageVersions[extensions.GroupName] = testapi.Extensions.GroupVersion().String()

	storageDestinations := genericapiserver.NewStorageDestinations()
	storageDestinations.AddAPIGroup(kube.GroupName, etcdStorage)
	storageDestinations.AddAPIGroup(extensions.GroupName, expEtcdStorage)

	return &master.Config{
		Config: &genericapiserver.Config{
			StorageDestinations: storageDestinations,
			StorageVersions:     storageVersions,
			APIPrefix:           "/api",
			APIGroupPrefix:      "/apis",
			Authorizer:          apiserver.NewAlwaysAllowAuthorizer(),
			AdmissionControl:    admit.NewAlwaysAdmit(),
			Serializer:          kube.Codecs,
		},
		KubeletClient: kubeletclient.FakeKubeletClient{},
	}
}

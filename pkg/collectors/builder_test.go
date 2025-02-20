// Copyright (c) 2020 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package collectors

import (
	"bytes"
	"reflect"
	"testing"

	ocinfrav1 "github.com/openshift/api/config/v1"
	ocpclient "github.com/openshift/client-go/config/clientset/versioned"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsstore "k8s.io/kube-state-metrics/pkg/metrics_store"
	koptions "k8s.io/kube-state-metrics/pkg/options"
	"k8s.io/kube-state-metrics/pkg/whiteblacklist"
	clusterclient "open-cluster-management.io/api/client/cluster/clientset/versioned"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

var (
	ctx context.Context = context.TODO()
)

func TestBuilder_WithApiserver(t *testing.T) {
	type fields struct {
		apiserver         string
		kubeconfig        string
		namespaces        koptions.NamespaceList
		ctx               context.Context
		enabledCollectors []string
		whiteBlackList    whiteBlackLister
	}
	type args struct {
		apiserver string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "apiserver",
			fields: fields{
				apiserver:         "",
				kubeconfig:        "",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
				// whiteBlackList:    whiteBlackLister{func(s) { return false }, func(s) { return false }},
			},
			args: args{
				apiserver: "apiserver",
			},
			want: &Builder{
				apiserver:         "apiserver",
				kubeconfig:        "",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				apiserver:         tt.fields.apiserver,
				kubeconfig:        tt.fields.kubeconfig,
				namespaces:        tt.fields.namespaces,
				ctx:               tt.fields.ctx,
				enabledCollectors: tt.fields.enabledCollectors,
				whiteBlackList:    tt.fields.whiteBlackList,
			}
			if got := b.WithApiserver(tt.args.apiserver); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.WithApiserver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_WithKubeConfig(t *testing.T) {
	type fields struct {
		apiserver         string
		kubeconfig        string
		namespaces        koptions.NamespaceList
		ctx               context.Context
		enabledCollectors []string
		whiteBlackList    whiteBlackLister
	}
	type args struct {
		kubeconfig string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "kubeconfig",
			fields: fields{
				apiserver:         "",
				kubeconfig:        "",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
				// whiteBlackList:    whiteBlackLister{func(s) { return false }, func(s) { return false }},
			},
			args: args{
				kubeconfig: "kubeconfig",
			},
			want: &Builder{
				apiserver:         "",
				kubeconfig:        "kubeconfig",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				apiserver:         tt.fields.apiserver,
				kubeconfig:        tt.fields.kubeconfig,
				namespaces:        tt.fields.namespaces,
				ctx:               tt.fields.ctx,
				enabledCollectors: tt.fields.enabledCollectors,
				whiteBlackList:    tt.fields.whiteBlackList,
			}
			if got := b.WithKubeConfig(tt.args.kubeconfig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.WithKubeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_WithEnabledCollectors(t *testing.T) {
	type fields struct {
		apiserver         string
		kubeconfig        string
		namespaces        koptions.NamespaceList
		ctx               context.Context
		enabledCollectors []string
		whiteBlackList    whiteBlackLister
	}
	type args struct {
		c []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "collectors",
			fields: fields{
				apiserver:         "",
				kubeconfig:        "",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{},
				// whiteBlackList:    whiteBlackLister{func(s) { return false }, func(s) { return false }},
			},
			args: args{
				c: []string{"col1", "col2"},
			},
			want: &Builder{
				apiserver:         "",
				kubeconfig:        "",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				apiserver:         tt.fields.apiserver,
				kubeconfig:        tt.fields.kubeconfig,
				namespaces:        tt.fields.namespaces,
				ctx:               tt.fields.ctx,
				enabledCollectors: tt.fields.enabledCollectors,
				whiteBlackList:    tt.fields.whiteBlackList,
			}
			if got := b.WithEnabledCollectors(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.WithEnabledCollectors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_WithNamespaces(t *testing.T) {
	type fields struct {
		apiserver         string
		kubeconfig        string
		namespaces        koptions.NamespaceList
		ctx               context.Context
		enabledCollectors []string
		whiteBlackList    whiteBlackLister
	}
	type args struct {
		n koptions.NamespaceList
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "namespace",
			fields: fields{
				apiserver:         "",
				kubeconfig:        "",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
				// whiteBlackList:    whiteBlackLister{func(s) { return false }, func(s) { return false }},
			},
			args: args{
				n: []string{"ns1"},
			},
			want: &Builder{
				apiserver:         "",
				kubeconfig:        "",
				namespaces:        []string{"ns1"},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				apiserver:         tt.fields.apiserver,
				kubeconfig:        tt.fields.kubeconfig,
				namespaces:        tt.fields.namespaces,
				ctx:               tt.fields.ctx,
				enabledCollectors: tt.fields.enabledCollectors,
				whiteBlackList:    tt.fields.whiteBlackList,
			}
			if got := b.WithNamespaces(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.WithNamespaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_WithWhiteBlackList(t *testing.T) {
	w, _ := whiteblacklist.New(map[string]struct{}{}, map[string]struct{}{})
	type fields struct {
		apiserver         string
		kubeconfig        string
		namespaces        koptions.NamespaceList
		ctx               context.Context
		enabledCollectors []string
		whiteBlackList    whiteBlackLister
	}
	type args struct {
		l whiteBlackLister
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Builder
	}{
		{
			name: "whiteBlackList",
			fields: fields{
				apiserver:         "",
				kubeconfig:        "",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
				whiteBlackList:    nil,
			},
			args: args{
				l: w,
			},
			want: &Builder{
				apiserver:         "",
				kubeconfig:        "",
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"col1", "col2"},
				whiteBlackList:    w,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				apiserver:         tt.fields.apiserver,
				kubeconfig:        tt.fields.kubeconfig,
				namespaces:        tt.fields.namespaces,
				ctx:               tt.fields.ctx,
				enabledCollectors: tt.fields.enabledCollectors,
				whiteBlackList:    tt.fields.whiteBlackList,
			}
			if got := b.WithWhiteBlackList(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Builder.WithWhiteBlackList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuilder_buildManagedClusterCollectorWithClient(t *testing.T) {
	const headers = `# HELP acm_managed_cluster_info Managed cluster information
# TYPE acm_managed_cluster_info gauge
`
	envTest, kubeconfig, _, _ := setupEnvTest(t)
	_, err := envtest.InstallCRDs(envTest.Config, envtest.CRDInstallOptions{
		Paths: []string{"../../test/unit/resources/crds"},
	})
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = envTest.ControlPlane.KubeCtl().Run("create", "ns", "imported-cluster")
	if err != nil {
		t.Fatal(err)
	}
	ocpClient, _ := ocpclient.NewForConfig(envTest.Config)
	clusterClient, _ := clusterclient.NewForConfig(envTest.Config)

	version := &ocinfrav1.ClusterVersion{
		ObjectMeta: metav1.ObjectMeta{
			Name: "version",
		},
		Spec: ocinfrav1.ClusterVersionSpec{
			ClusterID: "mycluster_id",
		},
	}

	_, err = ocpClient.ConfigV1().ClusterVersions().Create(context.TODO(), version, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}

	w, _ := whiteblacklist.New(map[string]struct{}{}, map[string]struct{}{})
	type fields struct {
		apiserver         string
		kubeconfig        string
		namespaces        koptions.NamespaceList
		ctx               context.Context
		enabledCollectors []string
		whiteBlackList    whiteBlackLister
	}
	type args struct {
		ocpClient     *ocpclient.Clientset
		clusterClient *clusterclient.Clientset
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *metricsstore.MetricsStore
	}{
		{
			name: "buildManagedClusterCollectorWithClient",
			fields: fields{
				apiserver:         envTest.ControlPlane.APIURL().String(),
				kubeconfig:        kubeconfig,
				namespaces:        koptions.NamespaceList{},
				ctx:               ctx,
				enabledCollectors: []string{"managedclusterinfos"},
				whiteBlackList:    w,
			},
			args: args{
				ocpClient:     ocpClient,
				clusterClient: clusterClient,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Builder{
				apiserver:         tt.fields.apiserver,
				kubeconfig:        tt.fields.kubeconfig,
				namespaces:        tt.fields.namespaces,
				ctx:               tt.fields.ctx,
				enabledCollectors: tt.fields.enabledCollectors,
				whiteBlackList:    tt.fields.whiteBlackList,
			}
			got := b.buildManagedClusterInfoCollectorWithClient(tt.args.ocpClient, tt.args.clusterClient)
			if got == nil {
				t.Errorf("Builder.buildManagedClusterCollectorWithClient() = %v, want %v", got, tt.want)
			}
			buf := new(bytes.Buffer)
			got.WriteAll(buf)
			if buf.String() != headers {
				t.Errorf("Expected headers \n%s\ngot\n%s", headers, buf.String())
			}
		})
	}
}

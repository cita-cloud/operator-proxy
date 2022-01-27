package kubeapi

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	citacloudv1 "github.com/cita-cloud/cita-cloud-operator/api/v1"
)

var scheme = runtime.NewScheme()

func init() {
	// Register all types of our clientset into the standard scheme.
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(citacloudv1.AddToScheme(scheme))
}

func loadClientConfig() (*rest.Config, error) {
	// The default loading rules try to read from the files specified in the
	// environment or from the home directory.
	loader := clientcmd.NewDefaultClientConfigLoadingRules()

	// The deferred loader tries an in-cluster config if the default loading
	// rules produce no results.
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loader, &clientcmd.ConfigOverrides{},
	).ClientConfig()
}

var K8sClient client.Client

func InitK8sClient() error {
	config, err := loadClientConfig()
	if err != nil {
		return err
	}
	// Match the settings applied by sigs.k8s.io/controller-runtime@v0.6.0;
	// see https://github.com/kubernetes-sigs/controller-runtime/issues/365.
	if config.QPS == 0.0 {
		config.QPS = 20.0
		config.Burst = 30.0
	}

	k8sManager, err := ctrl.NewManager(config, ctrl.Options{
		Scheme:             scheme,
		LeaderElection:     false,
		MetricsBindAddress: "0",
	})
	if err != nil {
		return err
	}

	err = k8sManager.GetCache().IndexField(context.Background(), &citacloudv1.Account{}, "spec.chain", func(o client.Object) []string {
		var res []string
		res = append(res, o.(*citacloudv1.Account).Spec.Chain)
		return res
	})
	if err != nil {
		return err
	}

	err = k8sManager.GetCache().IndexField(context.Background(), &citacloudv1.ChainNode{}, "spec.chainName", func(o client.Object) []string {
		var res []string
		res = append(res, o.(*citacloudv1.ChainNode).Spec.ChainName)
		return res
	})
	if err != nil {
		return err
	}

	go func() {
		_ = k8sManager.Start(ctrl.SetupSignalHandler())
	}()

	K8sClient = k8sManager.GetClient()
	return nil
}

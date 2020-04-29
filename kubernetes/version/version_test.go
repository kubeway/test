package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func TestVersion(t *testing.T) {
	kubecfg, err := getConfig("", "/Users/qiulin.nql/Downloads/long_text_2020-04-29-14-50-17.txt")
	if err != nil {
		t.Fatalf("get config error:%s", err)
	}
	kc := kubernetes.NewForConfigOrDie(kubecfg)
	versioner := kc.Discovery()
	v, err := versioner.ServerVersion()
	if err != nil {
		t.Fatalf("get serverversion error:%s", err)
	}

	vbts, _ := json.Marshal(v)
	t.Logf("version info:%s", vbts)
}

// GetConfig returns a rest.Config to be used for kubernetes client creation.
// It does so in the following order:
//   1. Use the passed kubeconfig/masterURL.
//   2. Fallback to the KUBECONFIG environment variable.
//   3. Fallback to in-cluster config.
//   4. Fallback to the ~/.kube/config.
func getConfig(masterURL, kubeconfig string) (*rest.Config, error) {
	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
	}
	// If we have an explicit indication of where the kubernetes config lives, read that.
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	}
	// If not, try the in-cluster config.
	if c, err := rest.InClusterConfig(); err == nil {
		return c, nil
	}
	// If no in-cluster config, try the default location in the user's home directory.
	if usr, err := user.Current(); err == nil {
		if c, err := clientcmd.BuildConfigFromFlags("", filepath.Join(usr.HomeDir, ".kube", "config")); err == nil {
			return c, nil
		}
	}

	return nil, fmt.Errorf("could not create a valid kubeconfig")
}

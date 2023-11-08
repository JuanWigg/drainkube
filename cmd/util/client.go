/*
Copyright © 2023 Juan Wiggenhauser <jgwiggenahuser@gmail.com>
*/
package util

import (
	"flag"
	"fmt"
	"path/filepath"
	"sync"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var lock = &sync.Mutex{}

var clientInstance *kubernetes.Clientset

func init() {
	if flag.Lookup("kubeconfig") == nil {
		if home := homedir.HomeDir(); home != "" {
			flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
	}
}

func GetInstance() *kubernetes.Clientset {
	if clientInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if clientInstance == nil {
			var kubeconfig string
			flag.Parse()
			kubeconfig = flag.Lookup("kubeconfig").Value.(flag.Getter).Get().(string)

			// use the current context in kubeconfig
			config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				panic(err.Error())
			}

			// create the clientset
			clientInstance, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}
			return clientInstance
		} else {
			fmt.Println("Client instance already created.")
		}
	} else {
		fmt.Println("Client instance already created.")
	}
	return clientInstance
}

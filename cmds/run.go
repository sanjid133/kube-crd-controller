package cmds

import (
	"github.com/spf13/cobra"

	"fmt"
	clientset "github.com/sanjid133/crd-controller/client/clientset/versioned"
	informers "github.com/sanjid133/crd-controller/client/informers/externalversions"
	"github.com/sanjid133/crd-controller/pkg/controller"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func NewCmdRun() *cobra.Command {
	var kubeconfig string
	var master string
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "run the server",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			err := runController(kubeconfig, master)
			if err != nil {
				fmt.Println(err)
			}

		},
	}
	cmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "kubeconfig file")
	cmd.Flags().StringVar(&master, "master", "", "master url")
	return cmd
}

func runController(kubeconfig, master string) error {
	cfg, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		return err
	}

	kubeclient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return err
	}

	demoClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		return err
	}
	kif := kubeinformers.NewSharedInformerFactory(kubeclient, time.Second*30)
	dif := informers.NewSharedInformerFactory(demoClient, time.Second*30)
	ctrl := controller.NewController(kubeclient, demoClient, kif, dif)

	stop := make(chan struct{})
	defer close(stop)
	go kif.Start(stop)
	go dif.Start(stop)

	ctrl.Run(2, stop)

	return nil
}

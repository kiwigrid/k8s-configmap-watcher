package main

import (
	"bytes"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var VERSION = "latest"

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Usage: "Kube config path for outside of cluster access",
		},
	}

	app.Action = func(c *cli.Context) error {
		var err error
		if err != nil {
			logrus.Error(err)
			return err
		}
		configFile := c.String("config")
		label := os.Getenv("LABEL")
		folder := os.Getenv("FOLDER")
		namespace := os.Getenv("NAMESPACE")
		requestUrl := os.Getenv("REQ_URL")
		requestMethod := os.Getenv("REQ_METHOD")
		requestPayload := os.Getenv("REQ_PAYLOAD")

		if requestMethod == "" {
			requestMethod = "GET"
		}

		if namespace == "" {
			file, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
			if err != nil {
				logrus.Fatalf("Err: %v", err)
			}
			namespace = string(file)
		} else {
			if namespace == "ALL" {
				namespace = v1.NamespaceAll
			}
		}

		go startWatchConfigMaps(configFile, namespace, label, folder, requestUrl, requestMethod, requestPayload)
		for {
			time.Sleep(5 * time.Second)
		}
	}
	app.Run(os.Args)
}

func startWatchConfigMaps(pathToCfg, namespace, label, folder, requestUrl, requestMethod, requestPayload string) {
	clientSet, _ := getClient(pathToCfg)
	listOptions := metav1.ListOptions{
	}

	watcher, err := clientSet.CoreV1().ConfigMaps(namespace).Watch(listOptions)
	if err != nil {
		logrus.Errorf("%v", err)
	}
	for {
		select {

		case event, ok := <-watcher.ResultChan():
			if !ok {
				logrus.Error("watcher stopped")
				watcher, err = clientSet.CoreV1().ConfigMaps(namespace).Watch(listOptions)
				logrus.Info("start new watcher")
				if err != nil {
					logrus.Errorf("%v", err)
				}
				break
			}
			configMap, ok := event.Object.(*v1.ConfigMap)
			if !ok {
				logrus.Fatal("unexpected type")
				continue
			}
			if _, ok := configMap.Labels[label]; ok {
				for k, v := range configMap.Data {
					logrus.Infof("create or update file %s", k)

					err := ioutil.WriteFile(fmt.Sprintf("%s/%s", folder, k), []byte(v), 0644)
					if err != nil {
						logrus.Errorf("error writing file %v", err)
					}
				}
				if requestUrl != "" && requestMethod != "" {

					var body io.Reader = nil
					if requestPayload != "" {
						body = bytes.NewBuffer([]byte(requestPayload))
					}
					req, err := http.NewRequest(requestMethod, requestUrl, body)
					if err != nil {
						logrus.Errorf("error:", err)
						continue
					}
					resp, err := http.DefaultClient.Do(req)
					if err != nil {
						logrus.Errorf("error:", err)
						continue
					}
					resp.Body.Close()
					if resp.StatusCode != 200 {
						logrus.Infof("error: Received response code %s , expected %s", resp.StatusCode, 200)
						continue
					}
				}
			}
		}
	}
}

func getClient(pathToCfg string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	if pathToCfg == "" {
		// in cluster access
		logrus.Info("Using in cluster config")
		config, err = rest.InClusterConfig()
	} else {
		logrus.Info("Using out of cluster config")
		config, err = clientcmd.BuildConfigFromFlags("", pathToCfg)
	}
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)

}

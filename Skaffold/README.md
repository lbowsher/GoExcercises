## Kubernetes Pod Manager

This is a simple example in golang to manage kubernetes pods.

Every 15 seconds: 

- a pod with a random name is created and added to the cluster
- Each pod is listed with their name and creation timestamp

The general plan for the app is:

* **building** a single Go file app and with a multistage `Dockerfile` using local docker to build
* **tagging** using the default tagPolicy (`gitCommit`)
* **deploying** starts as a single container pod using `kubectl`

The mainly intended to be run using Skaffold and Kubernetes kind for local development.

### Install

In order to run the app yourself, you need to install a few dependencies: 

`go get -u k8s.io/client-go@master`

`go get -u k8s.io/apimachinery/pkg/apis/meta/v1`

`go get -u k8s.io/api/core/v1`

See [here](https://github.com/kubernetes/client-go) for more details about installing go dependencies.

Then create a new kubectl service account, this tutorial is setup for one named `go-tutorials-service-account`.

Finally, give yourself the needed permissions to add a pod: 

`kubectl apply -f kubeclient-rbac.yaml`

### Setup

To run this example yourself, first clone the repo, then simply run

`kind create cluster`

Then `skaffold dev`

And you should be in business
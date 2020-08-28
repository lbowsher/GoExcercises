// ensures that the versions of go, k8s api, apimachinery, and client-go are all using compatible versions
module workspace

go 1.14

require (
	cloud.google.com/go v0.51.0 // indirect
	github.com/Azure/go-autorest/autorest v0.9.6 // indirect
	github.com/evanphx/json-patch v4.2.0+incompatible // indirect
	github.com/gophercloud/gophercloud v0.1.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20180305231024-9cad4c3443a7 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	//k8s.io/client-go v0.0.0-20200802132507-00dbcca6ee44 // indirect
	k8s.io/klog v0.4.0 // indirect
	k8s.io/klog/v2 v2.3.0 // indirect
	sigs.k8s.io/structured-merge-diff v0.0.0-20190525122527-15d366b2352e // indirect
	k8s.io/api v0.18.6
  	k8s.io/apimachinery v0.18.6
  	k8s.io/client-go v0.18.6
)

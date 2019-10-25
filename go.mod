module github.com/wwitzel3/octant-certificates.certmanager.k8s.io

go 1.13

require (
	github.com/jetstack/cert-manager v0.10.0
	github.com/pkg/errors v0.8.1
	github.com/vmware/octant v0.7.0
	k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
)

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab

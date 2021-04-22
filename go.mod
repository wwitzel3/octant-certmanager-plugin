module github.com/wwitzel3/octant-certificates.certmanager.k8s.io

go 1.16

require (
	github.com/jetstack/cert-manager v1.3.1
	github.com/pkg/errors v0.9.1
	github.com/vmware-tanzu/octant v0.19.0
	k8s.io/apimachinery v0.19.4
)

replace k8s.io/client-go => k8s.io/client-go v0.19.3

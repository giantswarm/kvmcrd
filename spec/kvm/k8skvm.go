package kvm

import "github.com/giantswarm/kvmcrd/spec/kvm/k8skvm"

type K8sKVM struct {
	Docker      k8skvm.Docker `json:"docker" yaml:"docker"`
	StorageType string        `json:"storageType" yaml:"storageType"`
}

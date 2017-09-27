package kvm

import "github.com/giantswarm/kvmcrd/spec/kvm/endpointupdater"

type EndpointUpdater struct {
	Docker endpointupdater.Docker `json:"docker" yaml:"docker"`
}

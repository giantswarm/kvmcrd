package kvmcrd

import (
	"github.com/giantswarm/microerror"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomObject represents the KVM CRD's custom object, the CRO. It holds the
// specifications of the resource the KVM operator acts on behalf of.
type CustomObject struct {
	apismetav1.TypeMeta   `json:",inline"`
	apismetav1.ObjectMeta `json:"metadata,omitempty"`

	Spec Spec `json:"spec" yaml:"spec"`
}

func ToCustomObject(v interface{}) (CustomObject, error) {
	customObjectPointer, ok := v.(*CustomObject)
	if !ok {
		return CustomObject{}, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", &CustomObject{}, v)
	}
	customObject := *customObjectPointer

	return customObject, nil
}

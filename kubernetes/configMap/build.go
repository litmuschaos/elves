package configMap

import (
	"errors"
	"fmt"
	//templatespec "github.com/litmuschaos/kube-helper/kubernetes/podtemplatespec"
	//batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
)

type Builder struct {
	configMap *ConfigMap
	errs      []error
}

func NewBuilder() *Builder {
	return &Builder{
		configMap: &ConfigMap{
			object: &corev1.ConfigMap{},
		},
	}
}

func (b *Builder) WithName(name string) *Builder {
	if len(name) == 0 {
		b.errs = append(
			b.errs,
			errors.New("Failed to build ConfigMap object: missing ConfigMap Name"),
		)
		return b
	}
	b.configMap.object.Name = name
	return b
}

func (b *Builder) WithData(data map[string]string) *Builder {
	if len(data) == 0 {
		b.errs = append(
			b.errs,
			errors.New("Failed to build ConfigMap object: missing Data"),
		)
		return b
	}
	b.configMap.object.Data = data
	return b
}

func (b *Builder) Build() (*corev1.ConfigMap, error) {
	if len(b.errs) > 0 {
		return nil, fmt.Errorf("%+v", b.errs)
	}
	return b.configMap.object, nil
}

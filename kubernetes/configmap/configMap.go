package configmap

import (
	corev1 "k8s.io/api/core/v1"
)

// ConfigMap is a struct with contains th*corev1.ConfigMap
type ConfigMap struct {
	object *corev1.ConfigMap
}

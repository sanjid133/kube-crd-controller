package controller

import (
	demov1 "github.com/sanjid133/crd-controller/apis/democontroller/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newSecret(secretType, namespace string, secret demov1.InfoSpec) *core.Secret {
	return &core.Secret{
		Type: core.SecretType(secretType),
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret.SecretName,
			Namespace: namespace,
		},
		Data: convertData(secret.Data),
	}
}

func convertData(data map[string]string) map[string][]byte {
	retdata := make(map[string][]byte)
	for k, v := range data {
		retdata[k] = []byte(v)
	}
	return retdata
}

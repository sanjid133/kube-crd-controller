package controller

import (
	"fmt"
	"github.com/appscode/go/runtime"
	demov1 "github.com/sanjid133/crd-controller/apis/democontroller/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
)

func (c *Controller) demoSyncHandler(key string) error {
	fmt.Println("handling demo controller from here...")
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(err)
		return nil
	}

	secdb, err := c.secdbLister.SecDbs(namespace).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			err = c.deleteSecretToSync(namespace, name)
			runtime.HandleError(fmt.Errorf("secdb %s with name %s is not found on queue", key, name))
			return nil
		}
		return err
	}

	secretType := secdb.Spec.SecretType
	if secretType == "" {
		runtime.HandleError(fmt.Errorf("secret type must be specified"))
		return nil
	}

	currentSecrets := make(map[string]string, 0)
	for _, secret := range secdb.Spec.Info {
		s, err := c.secretsLister.Secrets(secdb.Namespace).Get(secret.SecretName)
		if errors.IsNotFound(err) {
			newSec := newSecret(secdb.Spec.SecretType, secdb.Namespace, secret)
			newSec.OwnerReferences = []metav1.OwnerReference{
				*metav1.NewControllerRef(secdb, schema.GroupVersionKind{
					Group:   demov1.SchemeGroupVersion.Group,
					Version: demov1.SchemeGroupVersion.Version,
					Kind:    demov1.KindSecDb,
				}),
			}
			newSec.Labels = map[string]string{
				demov1.SecDbLabel: secdb.Name,
			}
			s, err = c.kubeClient.CoreV1().Secrets(secdb.Namespace).Create(newSec)
		}
		if err != nil {
			fmt.Println("error in getting/creatting secret ", err)
			return err
		}
		if !metav1.IsControlledBy(s, secdb) {
			msg := fmt.Sprintf("Resource %q already exists and is not managed by Something", s.Name)
			return fmt.Errorf(msg)
		}
		fmt.Println(s.Name, " is controlled by ", secdb.Name)
		currentSecrets[secret.SecretName] = s.Labels[demov1.SecDbLabel]
		if isDataupdated(s.Data, secret.Data) {
			fmt.Println("data is updated...")
			s.Data = convertData(secret.Data)
			s, err = c.kubeClient.CoreV1().Secrets(secdb.Namespace).Update(s)
			if err != nil {
				return err
			}
		}
	}

	realSecrets, err := c.kubeClient.CoreV1().Secrets(secdb.Namespace).List(metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(map[string]string{
			demov1.SecDbLabel: secdb.Name,
		}).String(),
	})

	for _, rs := range realSecrets.Items {
		if _, found := currentSecrets[rs.Name]; !found && rs.Labels[demov1.SecDbLabel] == secdb.Name {
			c.kubeClient.CoreV1().Secrets(secdb.Namespace).Delete(rs.Name, &metav1.DeleteOptions{})
		}
	}

	return nil
}

func isDataupdated(oldData map[string][]byte, newData map[string]string) bool {
	for k, v := range oldData {
		if string(v) != newData[k] {
			return true
		}
	}
	return false
}

func (c *Controller) deleteSecretToSync(namespace, label string) error {
	realSecrets, err := c.kubeClient.CoreV1().Secrets(namespace).List(metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(map[string]string{
			demov1.SecDbLabel: label,
		}).String(),
	})
	if err != nil {
		return err
	}
	for _, rs := range realSecrets.Items {
		c.kubeClient.CoreV1().Secrets(namespace).Delete(rs.Name, &metav1.DeleteOptions{})
	}
	return nil
}

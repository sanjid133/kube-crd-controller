/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	v1 "github.com/sanjid133/crd-controller/apis/democontroller/v1"
	scheme "github.com/sanjid133/crd-controller/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecDbsGetter has a method to return a SecDbInterface.
// A group's client should implement this interface.
type SecDbsGetter interface {
	SecDbs(namespace string) SecDbInterface
}

// SecDbInterface has methods to work with SecDb resources.
type SecDbInterface interface {
	Create(*v1.SecDb) (*v1.SecDb, error)
	Update(*v1.SecDb) (*v1.SecDb, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.SecDb, error)
	List(opts meta_v1.ListOptions) (*v1.SecDbList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SecDb, err error)
	SecDbExpansion
}

// secDbs implements SecDbInterface
type secDbs struct {
	client rest.Interface
	ns     string
}

// newSecDbs returns a SecDbs
func newSecDbs(c *DemocontrollerV1Client, namespace string) *secDbs {
	return &secDbs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secDb, and returns the corresponding secDb object, and an error if there is any.
func (c *secDbs) Get(name string, options meta_v1.GetOptions) (result *v1.SecDb, err error) {
	result = &v1.SecDb{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secdbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecDbs that match those selectors.
func (c *secDbs) List(opts meta_v1.ListOptions) (result *v1.SecDbList, err error) {
	result = &v1.SecDbList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secdbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secDbs.
func (c *secDbs) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("secdbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a secDb and creates it.  Returns the server's representation of the secDb, and an error, if there is any.
func (c *secDbs) Create(secDb *v1.SecDb) (result *v1.SecDb, err error) {
	result = &v1.SecDb{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("secdbs").
		Body(secDb).
		Do().
		Into(result)
	return
}

// Update takes the representation of a secDb and updates it. Returns the server's representation of the secDb, and an error, if there is any.
func (c *secDbs) Update(secDb *v1.SecDb) (result *v1.SecDb, err error) {
	result = &v1.SecDb{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secdbs").
		Name(secDb.Name).
		Body(secDb).
		Do().
		Into(result)
	return
}

// Delete takes name of the secDb and deletes it. Returns an error if one occurs.
func (c *secDbs) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secdbs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secDbs) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secdbs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched secDb.
func (c *secDbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.SecDb, err error) {
	result = &v1.SecDb{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("secdbs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

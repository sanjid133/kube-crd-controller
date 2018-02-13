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

package fake

import (
	democontroller_v1 "github.com/sanjid133/crd-controller/apis/democontroller/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSecDbs implements SecDbInterface
type FakeSecDbs struct {
	Fake *FakeDemocontrollerV1
	ns   string
}

var secdbsResource = schema.GroupVersionResource{Group: "democontroller.k8s.io", Version: "v1", Resource: "secdbs"}

var secdbsKind = schema.GroupVersionKind{Group: "democontroller.k8s.io", Version: "v1", Kind: "SecDb"}

// Get takes name of the secDb, and returns the corresponding secDb object, and an error if there is any.
func (c *FakeSecDbs) Get(name string, options v1.GetOptions) (result *democontroller_v1.SecDb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secdbsResource, c.ns, name), &democontroller_v1.SecDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*democontroller_v1.SecDb), err
}

// List takes label and field selectors, and returns the list of SecDbs that match those selectors.
func (c *FakeSecDbs) List(opts v1.ListOptions) (result *democontroller_v1.SecDbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secdbsResource, secdbsKind, c.ns, opts), &democontroller_v1.SecDbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &democontroller_v1.SecDbList{}
	for _, item := range obj.(*democontroller_v1.SecDbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secDbs.
func (c *FakeSecDbs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secdbsResource, c.ns, opts))

}

// Create takes the representation of a secDb and creates it.  Returns the server's representation of the secDb, and an error, if there is any.
func (c *FakeSecDbs) Create(secDb *democontroller_v1.SecDb) (result *democontroller_v1.SecDb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secdbsResource, c.ns, secDb), &democontroller_v1.SecDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*democontroller_v1.SecDb), err
}

// Update takes the representation of a secDb and updates it. Returns the server's representation of the secDb, and an error, if there is any.
func (c *FakeSecDbs) Update(secDb *democontroller_v1.SecDb) (result *democontroller_v1.SecDb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secdbsResource, c.ns, secDb), &democontroller_v1.SecDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*democontroller_v1.SecDb), err
}

// Delete takes name of the secDb and deletes it. Returns an error if one occurs.
func (c *FakeSecDbs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secdbsResource, c.ns, name), &democontroller_v1.SecDb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecDbs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secdbsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &democontroller_v1.SecDbList{})
	return err
}

// Patch applies the patch and returns the patched secDb.
func (c *FakeSecDbs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *democontroller_v1.SecDb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secdbsResource, c.ns, name, data, subresources...), &democontroller_v1.SecDb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*democontroller_v1.SecDb), err
}

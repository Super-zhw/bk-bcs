/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	meshv1 "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/mesh/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAppSvcs implements AppSvcInterface
type FakeAppSvcs struct {
	Fake *FakeMeshV1
	ns   string
}

var appsvcsResource = schema.GroupVersionResource{Group: "mesh", Version: "v1", Resource: "appsvcs"}

var appsvcsKind = schema.GroupVersionKind{Group: "mesh", Version: "v1", Kind: "AppSvc"}

// Get takes name of the appSvc, and returns the corresponding appSvc object, and an error if there is any.
func (c *FakeAppSvcs) Get(ctx context.Context, name string, options v1.GetOptions) (result *meshv1.AppSvc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(appsvcsResource, c.ns, name), &meshv1.AppSvc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppSvc), err
}

// List takes label and field selectors, and returns the list of AppSvcs that match those selectors.
func (c *FakeAppSvcs) List(ctx context.Context, opts v1.ListOptions) (result *meshv1.AppSvcList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(appsvcsResource, appsvcsKind, c.ns, opts), &meshv1.AppSvcList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &meshv1.AppSvcList{ListMeta: obj.(*meshv1.AppSvcList).ListMeta}
	for _, item := range obj.(*meshv1.AppSvcList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested appSvcs.
func (c *FakeAppSvcs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(appsvcsResource, c.ns, opts))

}

// Create takes the representation of a appSvc and creates it.  Returns the server's representation of the appSvc, and an error, if there is any.
func (c *FakeAppSvcs) Create(ctx context.Context, appSvc *meshv1.AppSvc, opts v1.CreateOptions) (result *meshv1.AppSvc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(appsvcsResource, c.ns, appSvc), &meshv1.AppSvc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppSvc), err
}

// Update takes the representation of a appSvc and updates it. Returns the server's representation of the appSvc, and an error, if there is any.
func (c *FakeAppSvcs) Update(ctx context.Context, appSvc *meshv1.AppSvc, opts v1.UpdateOptions) (result *meshv1.AppSvc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(appsvcsResource, c.ns, appSvc), &meshv1.AppSvc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppSvc), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAppSvcs) UpdateStatus(ctx context.Context, appSvc *meshv1.AppSvc, opts v1.UpdateOptions) (*meshv1.AppSvc, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(appsvcsResource, "status", c.ns, appSvc), &meshv1.AppSvc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppSvc), err
}

// Delete takes name of the appSvc and deletes it. Returns an error if one occurs.
func (c *FakeAppSvcs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(appsvcsResource, c.ns, name), &meshv1.AppSvc{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAppSvcs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(appsvcsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &meshv1.AppSvcList{})
	return err
}

// Patch applies the patch and returns the patched appSvc.
func (c *FakeAppSvcs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *meshv1.AppSvc, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(appsvcsResource, c.ns, name, pt, data, subresources...), &meshv1.AppSvc{})

	if obj == nil {
		return nil, err
	}
	return obj.(*meshv1.AppSvc), err
}

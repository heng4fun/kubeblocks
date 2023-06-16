/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	scheme "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComponentClassDefinitionsGetter has a method to return a ComponentClassDefinitionInterface.
// A group's client should implement this interface.
type ComponentClassDefinitionsGetter interface {
	ComponentClassDefinitions() ComponentClassDefinitionInterface
}

// ComponentClassDefinitionInterface has methods to work with ComponentClassDefinition resources.
type ComponentClassDefinitionInterface interface {
	Create(ctx context.Context, componentClassDefinition *v1alpha1.ComponentClassDefinition, opts v1.CreateOptions) (*v1alpha1.ComponentClassDefinition, error)
	Update(ctx context.Context, componentClassDefinition *v1alpha1.ComponentClassDefinition, opts v1.UpdateOptions) (*v1alpha1.ComponentClassDefinition, error)
	UpdateStatus(ctx context.Context, componentClassDefinition *v1alpha1.ComponentClassDefinition, opts v1.UpdateOptions) (*v1alpha1.ComponentClassDefinition, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ComponentClassDefinition, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ComponentClassDefinitionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComponentClassDefinition, err error)
	ComponentClassDefinitionExpansion
}

// componentClassDefinitions implements ComponentClassDefinitionInterface
type componentClassDefinitions struct {
	client rest.Interface
}

// newComponentClassDefinitions returns a ComponentClassDefinitions
func newComponentClassDefinitions(c *AppsV1alpha1Client) *componentClassDefinitions {
	return &componentClassDefinitions{
		client: c.RESTClient(),
	}
}

// Get takes name of the componentClassDefinition, and returns the corresponding componentClassDefinition object, and an error if there is any.
func (c *componentClassDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComponentClassDefinition, err error) {
	result = &v1alpha1.ComponentClassDefinition{}
	err = c.client.Get().
		Resource("componentclassdefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComponentClassDefinitions that match those selectors.
func (c *componentClassDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComponentClassDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ComponentClassDefinitionList{}
	err = c.client.Get().
		Resource("componentclassdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested componentClassDefinitions.
func (c *componentClassDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("componentclassdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a componentClassDefinition and creates it.  Returns the server's representation of the componentClassDefinition, and an error, if there is any.
func (c *componentClassDefinitions) Create(ctx context.Context, componentClassDefinition *v1alpha1.ComponentClassDefinition, opts v1.CreateOptions) (result *v1alpha1.ComponentClassDefinition, err error) {
	result = &v1alpha1.ComponentClassDefinition{}
	err = c.client.Post().
		Resource("componentclassdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(componentClassDefinition).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a componentClassDefinition and updates it. Returns the server's representation of the componentClassDefinition, and an error, if there is any.
func (c *componentClassDefinitions) Update(ctx context.Context, componentClassDefinition *v1alpha1.ComponentClassDefinition, opts v1.UpdateOptions) (result *v1alpha1.ComponentClassDefinition, err error) {
	result = &v1alpha1.ComponentClassDefinition{}
	err = c.client.Put().
		Resource("componentclassdefinitions").
		Name(componentClassDefinition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(componentClassDefinition).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *componentClassDefinitions) UpdateStatus(ctx context.Context, componentClassDefinition *v1alpha1.ComponentClassDefinition, opts v1.UpdateOptions) (result *v1alpha1.ComponentClassDefinition, err error) {
	result = &v1alpha1.ComponentClassDefinition{}
	err = c.client.Put().
		Resource("componentclassdefinitions").
		Name(componentClassDefinition.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(componentClassDefinition).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the componentClassDefinition and deletes it. Returns an error if one occurs.
func (c *componentClassDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("componentclassdefinitions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *componentClassDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("componentclassdefinitions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched componentClassDefinition.
func (c *componentClassDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComponentClassDefinition, err error) {
	result = &v1alpha1.ComponentClassDefinition{}
	err = c.client.Patch(pt).
		Resource("componentclassdefinitions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

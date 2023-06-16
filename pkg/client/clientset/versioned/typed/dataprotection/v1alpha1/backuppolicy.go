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

	v1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	scheme "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupPoliciesGetter has a method to return a BackupPolicyInterface.
// A group's client should implement this interface.
type BackupPoliciesGetter interface {
	BackupPolicies(namespace string) BackupPolicyInterface
}

// BackupPolicyInterface has methods to work with BackupPolicy resources.
type BackupPolicyInterface interface {
	Create(ctx context.Context, backupPolicy *v1alpha1.BackupPolicy, opts v1.CreateOptions) (*v1alpha1.BackupPolicy, error)
	Update(ctx context.Context, backupPolicy *v1alpha1.BackupPolicy, opts v1.UpdateOptions) (*v1alpha1.BackupPolicy, error)
	UpdateStatus(ctx context.Context, backupPolicy *v1alpha1.BackupPolicy, opts v1.UpdateOptions) (*v1alpha1.BackupPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BackupPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BackupPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupPolicy, err error)
	BackupPolicyExpansion
}

// backupPolicies implements BackupPolicyInterface
type backupPolicies struct {
	client rest.Interface
	ns     string
}

// newBackupPolicies returns a BackupPolicies
func newBackupPolicies(c *DataprotectionV1alpha1Client, namespace string) *backupPolicies {
	return &backupPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupPolicy, and returns the corresponding backupPolicy object, and an error if there is any.
func (c *backupPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupPolicy, err error) {
	result = &v1alpha1.BackupPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backuppolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupPolicies that match those selectors.
func (c *backupPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BackupPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backuppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupPolicies.
func (c *backupPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backuppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backupPolicy and creates it.  Returns the server's representation of the backupPolicy, and an error, if there is any.
func (c *backupPolicies) Create(ctx context.Context, backupPolicy *v1alpha1.BackupPolicy, opts v1.CreateOptions) (result *v1alpha1.BackupPolicy, err error) {
	result = &v1alpha1.BackupPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backuppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backupPolicy and updates it. Returns the server's representation of the backupPolicy, and an error, if there is any.
func (c *backupPolicies) Update(ctx context.Context, backupPolicy *v1alpha1.BackupPolicy, opts v1.UpdateOptions) (result *v1alpha1.BackupPolicy, err error) {
	result = &v1alpha1.BackupPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backuppolicies").
		Name(backupPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backupPolicies) UpdateStatus(ctx context.Context, backupPolicy *v1alpha1.BackupPolicy, opts v1.UpdateOptions) (result *v1alpha1.BackupPolicy, err error) {
	result = &v1alpha1.BackupPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backuppolicies").
		Name(backupPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backupPolicy and deletes it. Returns an error if one occurs.
func (c *backupPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backuppolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backuppolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backupPolicy.
func (c *backupPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupPolicy, err error) {
	result = &v1alpha1.BackupPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backuppolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/apecloud/kubeblocks/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Backups returns a BackupInformer.
	Backups() BackupInformer
	// BackupPolicies returns a BackupPolicyInformer.
	BackupPolicies() BackupPolicyInformer
	// BackupTools returns a BackupToolInformer.
	BackupTools() BackupToolInformer
	// RestoreJobs returns a RestoreJobInformer.
	RestoreJobs() RestoreJobInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Backups returns a BackupInformer.
func (v *version) Backups() BackupInformer {
	return &backupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupPolicies returns a BackupPolicyInformer.
func (v *version) BackupPolicies() BackupPolicyInformer {
	return &backupPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupTools returns a BackupToolInformer.
func (v *version) BackupTools() BackupToolInformer {
	return &backupToolInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// RestoreJobs returns a RestoreJobInformer.
func (v *version) RestoreJobs() RestoreJobInformer {
	return &restoreJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

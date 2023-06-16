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

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	dataprotectionv1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	extensionsv1alpha1 "github.com/apecloud/kubeblocks/apis/extensions/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apps.kubeblocks.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("backuppolicytemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().BackupPolicyTemplates().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().Clusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ClusterDefinitions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ClusterVersions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("componentclassdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ComponentClassDefinitions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("componentresourceconstraints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ComponentResourceConstraints().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("configconstraints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().ConfigConstraints().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("opsrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apps().V1alpha1().OpsRequests().Informer()}, nil

		// Group=dataprotection.kubeblocks.io, Version=v1alpha1
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("backups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().Backups().Informer()}, nil
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("backuppolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().BackupPolicies().Informer()}, nil
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("backuptools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().BackupTools().Informer()}, nil
	case dataprotectionv1alpha1.SchemeGroupVersion.WithResource("restorejobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataprotection().V1alpha1().RestoreJobs().Informer()}, nil

		// Group=extensions.kubeblocks.io, Version=v1alpha1
	case extensionsv1alpha1.SchemeGroupVersion.WithResource("addons"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().Addons().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

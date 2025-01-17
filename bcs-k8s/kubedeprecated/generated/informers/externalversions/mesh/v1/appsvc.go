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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	meshv1 "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/apis/mesh/v1"
	versioned "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/generated/clientset/versioned"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/Tencent/bk-bcs/bcs-k8s/kubedeprecated/generated/listers/mesh/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AppSvcInformer provides access to a shared informer and lister for
// AppSvcs.
type AppSvcInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AppSvcLister
}

type appSvcInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAppSvcInformer constructs a new informer for AppSvc type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAppSvcInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAppSvcInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAppSvcInformer constructs a new informer for AppSvc type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAppSvcInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeshV1().AppSvcs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MeshV1().AppSvcs(namespace).Watch(context.TODO(), options)
			},
		},
		&meshv1.AppSvc{},
		resyncPeriod,
		indexers,
	)
}

func (f *appSvcInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAppSvcInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *appSvcInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&meshv1.AppSvc{}, f.defaultInformer)
}

func (f *appSvcInformer) Lister() v1.AppSvcLister {
	return v1.NewAppSvcLister(f.Informer().GetIndexer())
}

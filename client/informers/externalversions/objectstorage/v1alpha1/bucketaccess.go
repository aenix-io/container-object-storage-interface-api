/*
Copyright 2022 The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	objectstoragev1alpha1 "sigs.k8s.io/container-object-storage-interface-api/apis/objectstorage/v1alpha1"
	versioned "sigs.k8s.io/container-object-storage-interface-api/client/clientset/versioned"
	internalinterfaces "sigs.k8s.io/container-object-storage-interface-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "sigs.k8s.io/container-object-storage-interface-api/client/listers/objectstorage/v1alpha1"
)

// BucketAccessInformer provides access to a shared informer and lister for
// BucketAccesses.
type BucketAccessInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BucketAccessLister
}

type bucketAccessInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewBucketAccessInformer constructs a new informer for BucketAccess type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBucketAccessInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBucketAccessInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredBucketAccessInformer constructs a new informer for BucketAccess type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBucketAccessInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ObjectstorageV1alpha1().BucketAccesses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ObjectstorageV1alpha1().BucketAccesses().Watch(context.TODO(), options)
			},
		},
		&objectstoragev1alpha1.BucketAccess{},
		resyncPeriod,
		indexers,
	)
}

func (f *bucketAccessInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBucketAccessInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bucketAccessInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&objectstoragev1alpha1.BucketAccess{}, f.defaultInformer)
}

func (f *bucketAccessInformer) Lister() v1alpha1.BucketAccessLister {
	return v1alpha1.NewBucketAccessLister(f.Informer().GetIndexer())
}
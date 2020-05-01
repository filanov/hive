// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	versioned "github.com/openshift/hive/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openshift/hive/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/hive/pkg/client/listers/hive/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DNSZoneInformer provides access to a shared informer and lister for
// DNSZones.
type DNSZoneInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DNSZoneLister
}

type dNSZoneInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDNSZoneInformer constructs a new informer for DNSZone type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDNSZoneInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDNSZoneInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDNSZoneInformer constructs a new informer for DNSZone type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDNSZoneInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HiveV1().DNSZones(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HiveV1().DNSZones(namespace).Watch(context.TODO(), options)
			},
		},
		&hivev1.DNSZone{},
		resyncPeriod,
		indexers,
	)
}

func (f *dNSZoneInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDNSZoneInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dNSZoneInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hivev1.DNSZone{}, f.defaultInformer)
}

func (f *dNSZoneInformer) Lister() v1.DNSZoneLister {
	return v1.NewDNSZoneLister(f.Informer().GetIndexer())
}

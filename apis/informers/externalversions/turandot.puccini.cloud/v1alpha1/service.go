// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/tliron/turandot/apis/clientset/versioned"
	internalinterfaces "github.com/tliron/turandot/apis/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/tliron/turandot/apis/listers/turandot.puccini.cloud/v1alpha1"
	turandotpuccinicloudv1alpha1 "github.com/tliron/turandot/resources/turandot.puccini.cloud/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceInformer provides access to a shared informer and lister for
// Services.
type ServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServiceLister
}

type serviceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceInformer constructs a new informer for Service type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceInformer constructs a new informer for Service type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TurandotV1alpha1().Services(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TurandotV1alpha1().Services(namespace).Watch(context.TODO(), options)
			},
		},
		&turandotpuccinicloudv1alpha1.Service{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&turandotpuccinicloudv1alpha1.Service{}, f.defaultInformer)
}

func (f *serviceInformer) Lister() v1alpha1.ServiceLister {
	return v1alpha1.NewServiceLister(f.Informer().GetIndexer())
}

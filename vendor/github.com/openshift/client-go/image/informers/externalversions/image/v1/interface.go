// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/openshift/client-go/image/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Images returns a ImageInformer.
	Images() ImageInformer
	// ImageStreams returns a ImageStreamInformer.
	ImageStreams() ImageStreamInformer
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

// Images returns a ImageInformer.
func (v *version) Images() ImageInformer {
	return &imageInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ImageStreams returns a ImageStreamInformer.
func (v *version) ImageStreams() ImageStreamInformer {
	return &imageStreamInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
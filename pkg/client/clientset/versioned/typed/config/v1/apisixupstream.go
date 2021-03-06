/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/gxthrj/apisix-ingress-types/pkg/apis/config/v1"
	scheme "github.com/gxthrj/apisix-ingress-types/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApisixUpstreamsGetter has a method to return a ApisixUpstreamInterface.
// A group's client should implement this interface.
type ApisixUpstreamsGetter interface {
	ApisixUpstreams(namespace string) ApisixUpstreamInterface
}

// ApisixUpstreamInterface has methods to work with ApisixUpstream resources.
type ApisixUpstreamInterface interface {
	Create(*v1.ApisixUpstream) (*v1.ApisixUpstream, error)
	Update(*v1.ApisixUpstream) (*v1.ApisixUpstream, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ApisixUpstream, error)
	List(opts metav1.ListOptions) (*v1.ApisixUpstreamList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ApisixUpstream, err error)
	ApisixUpstreamExpansion
}

// apisixUpstreams implements ApisixUpstreamInterface
type apisixUpstreams struct {
	client rest.Interface
	ns     string
}

// newApisixUpstreams returns a ApisixUpstreams
func newApisixUpstreams(c *ApisixV1Client, namespace string) *apisixUpstreams {
	return &apisixUpstreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apisixUpstream, and returns the corresponding apisixUpstream object, and an error if there is any.
func (c *apisixUpstreams) Get(name string, options metav1.GetOptions) (result *v1.ApisixUpstream, err error) {
	result = &v1.ApisixUpstream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apisixupstreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApisixUpstreams that match those selectors.
func (c *apisixUpstreams) List(opts metav1.ListOptions) (result *v1.ApisixUpstreamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ApisixUpstreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apisixupstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apisixUpstreams.
func (c *apisixUpstreams) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apisixupstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apisixUpstream and creates it.  Returns the server's representation of the apisixUpstream, and an error, if there is any.
func (c *apisixUpstreams) Create(apisixUpstream *v1.ApisixUpstream) (result *v1.ApisixUpstream, err error) {
	result = &v1.ApisixUpstream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apisixupstreams").
		Body(apisixUpstream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apisixUpstream and updates it. Returns the server's representation of the apisixUpstream, and an error, if there is any.
func (c *apisixUpstreams) Update(apisixUpstream *v1.ApisixUpstream) (result *v1.ApisixUpstream, err error) {
	result = &v1.ApisixUpstream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apisixupstreams").
		Name(apisixUpstream.Name).
		Body(apisixUpstream).
		Do().
		Into(result)
	return
}

// Delete takes name of the apisixUpstream and deletes it. Returns an error if one occurs.
func (c *apisixUpstreams) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apisixupstreams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apisixUpstreams) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apisixupstreams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apisixUpstream.
func (c *apisixUpstreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ApisixUpstream, err error) {
	result = &v1.ApisixUpstream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apisixupstreams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

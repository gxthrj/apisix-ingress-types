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

// ApisixServicesGetter has a method to return a ApisixServiceInterface.
// A group's client should implement this interface.
type ApisixServicesGetter interface {
	ApisixServices(namespace string) ApisixServiceInterface
}

// ApisixServiceInterface has methods to work with ApisixService resources.
type ApisixServiceInterface interface {
	Create(*v1.ApisixService) (*v1.ApisixService, error)
	Update(*v1.ApisixService) (*v1.ApisixService, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ApisixService, error)
	List(opts metav1.ListOptions) (*v1.ApisixServiceList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ApisixService, err error)
	ApisixServiceExpansion
}

// apisixServices implements ApisixServiceInterface
type apisixServices struct {
	client rest.Interface
	ns     string
}

// newApisixServices returns a ApisixServices
func newApisixServices(c *ApisixV1Client, namespace string) *apisixServices {
	return &apisixServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apisixService, and returns the corresponding apisixService object, and an error if there is any.
func (c *apisixServices) Get(name string, options metav1.GetOptions) (result *v1.ApisixService, err error) {
	result = &v1.ApisixService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apisixservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApisixServices that match those selectors.
func (c *apisixServices) List(opts metav1.ListOptions) (result *v1.ApisixServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ApisixServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apisixservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apisixServices.
func (c *apisixServices) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apisixservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apisixService and creates it.  Returns the server's representation of the apisixService, and an error, if there is any.
func (c *apisixServices) Create(apisixService *v1.ApisixService) (result *v1.ApisixService, err error) {
	result = &v1.ApisixService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apisixservices").
		Body(apisixService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apisixService and updates it. Returns the server's representation of the apisixService, and an error, if there is any.
func (c *apisixServices) Update(apisixService *v1.ApisixService) (result *v1.ApisixService, err error) {
	result = &v1.ApisixService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apisixservices").
		Name(apisixService.Name).
		Body(apisixService).
		Do().
		Into(result)
	return
}

// Delete takes name of the apisixService and deletes it. Returns an error if one occurs.
func (c *apisixServices) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apisixservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apisixServices) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apisixservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apisixService.
func (c *apisixServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ApisixService, err error) {
	result = &v1.ApisixService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apisixservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

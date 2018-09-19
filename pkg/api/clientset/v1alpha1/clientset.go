// Copyright 2018 Nimrod Shneor <nimrodshn@gmail.com>
// and other contributors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"github.com/nimrodshn/openshift-provisioner/pkg/api/types/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// ClusterInterface details the methods on the Cluster types.
type ClusterInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.ClusterList, error)
	Get(name string, options metav1.GetOptions) (*v1alpha1.Cluster, error)
	Create(*v1alpha1.Cluster) (*v1alpha1.Cluster, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

// ClusterV1Alpha1Client implements the ClusterInterface
type ClusterV1Alpha1Client struct {
	restClient rest.Interface
	namespace  string
}

// Cluster Pass namespace to cluster object.
func (c *ClusterV1Alpha1Client) Cluster(namespace string) *ClusterV1Alpha1Client {
	return &ClusterV1Alpha1Client{
		restClient: c.restClient,
		namespace:  namespace,
	}
}

// NewForConfig constructs a new Cluster Client or returns an error.
func NewForConfig(c *rest.Config) (*ClusterV1Alpha1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &schema.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}
	config.UserAgent = rest.DefaultKubernetesUserAgent()
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ClusterV1Alpha1Client{
		restClient: client,
	}, nil
}

// List lists the current list of clusters
func (c *ClusterV1Alpha1Client) List(namespace string, opts metav1.ListOptions) (*v1alpha1.ClusterList, error) {
	result := v1alpha1.ClusterList{}
	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

// Get retrieves a client with a given name.
func (c *ClusterV1Alpha1Client) Get(name string, namespace string, options metav1.GetOptions) (*v1alpha1.Cluster, error) {
	result := v1alpha1.Cluster{}
	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("clusters").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

// Create creates a  new Cluster custom resource.
func (c *ClusterV1Alpha1Client) Create(cluster *v1alpha1.Cluster, namespace string) (*v1alpha1.Cluster, error) {
	result := v1alpha1.Cluster{}
	err := c.restClient.
		Post().
		Namespace(c.namespace).
		Resource("clusters").
		Body(cluster).
		Do().
		Into(&result)

	return &result, err
}

// Watch returns a watch (channel) for changes on a resource defined by opts.
func (c *ClusterV1Alpha1Client) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namesapce(c.namespace).
		Resource("clusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

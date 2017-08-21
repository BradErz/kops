/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha2 "k8s.io/kops/pkg/client/clientset_generated/internalclientset/typed/kops/v1alpha2"
)

type FakeKopsV1alpha2 struct {
	*testing.Fake
}

func (c *FakeKopsV1alpha2) Clusters(namespace string) v1alpha2.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeKopsV1alpha2) Federations(namespace string) v1alpha2.FederationInterface {
	return &FakeFederations{c, namespace}
}

func (c *FakeKopsV1alpha2) InstanceGroups(namespace string) v1alpha2.InstanceGroupInterface {
	return &FakeInstanceGroups{c, namespace}
}

func (c *FakeKopsV1alpha2) SSHSecrets(namespace string) v1alpha2.SSHSecretInterface {
	return &FakeSSHSecrets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKopsV1alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

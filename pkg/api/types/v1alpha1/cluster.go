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
	"gitlab.cee.redhat.com/service/uhc-clusters-service/pkg/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Cluster is the internal k8s struct representing an
// Openshift Cluster.
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata, omitempty"`
	Spec              api.Cluster `json:"spec"`
}

// ClusterList representes a list of clusters.
type ClusterList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata, omitempty"`
	Items             []Cluster `json:"items"`
}

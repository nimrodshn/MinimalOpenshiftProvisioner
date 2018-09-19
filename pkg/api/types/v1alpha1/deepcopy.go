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
	"runtime"
)

// DeepCopyInto copies the information from one cluster to another.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = in.Spec
}

// DeepCopyObject implements runtim Object DeepCopyObject for Cluster.
func (in *Cluster) DeepCopyObject() runtime.Object {
	out := Cluster{}
	in.DeepCopyInto(&out)
	return &out
}

// DeepCopyObject implements a runtime object DeepCopyObject method for ClusterList.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	out := ClusterList{}
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta

	if in.Items != nil {
		out.List.Items = make([]Cluster, len(in.Items))
		for idx := range in.Items {
			in.Items[idx].DeepCopyInto(&out.Items[idx])
		}
	}
	return &out
}

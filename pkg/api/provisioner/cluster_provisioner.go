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

package provisioner

import (
	"gitlab.cee.redhat.com/service/uhc-clusters-service/pkg/api"
	"os/exec"
)

// ClustersProvisioner is a minial openshift provisioner using Openshift Ansible.
type ClustersProvisioner struct{}

// Provision provisions an openshift cluster.
func (p *ClustersProvisioner) Provision(api.Cluster) error {}

// SetCredentials sets the appropriate credentials
func (p *ClustersProvisioner) SetCredentials() error {}

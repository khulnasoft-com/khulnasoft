// Copyright (c) 2020 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package khulnasoftlib

type KhulnasoftImage struct {
	File    string
	Context string `yaml:"context,omitempty"`
}

type khulnasoftPort struct {
	Number int32 `yaml:"port"`
}

type khulnasoftTask struct {
	Command string `yaml:"command,omitempty"`
	Init    string `yaml:"init,omitempty"`
}

type KhulnasoftFile struct {
	Image             interface{}  `yaml:"image,omitempty"`
	Ports             []khulnasoftPort `yaml:"ports,omitempty"`
	Tasks             []khulnasoftTask `yaml:"tasks,omitempty"`
	CheckoutLocation  string       `yaml:"checkoutLocation,omitempty"`
	WorkspaceLocation string       `yaml:"workspaceLocation,omitempty"`
}

// SetImageName configures a pre-built docker image by name
func (cfg *KhulnasoftFile) SetImageName(name string) {
	if name == "" {
		return
	}
	cfg.Image = name
}

// SetImage configures a Dockerfile as workspace image
func (cfg *KhulnasoftFile) SetImage(img KhulnasoftImage) {
	cfg.Image = img
}

// AddPort adds a port to the list of exposed ports
func (cfg *KhulnasoftFile) AddPort(port int32) {
	cfg.Ports = append(cfg.Ports, khulnasoftPort{
		Number: port,
	})
}

// AddTask adds a workspace startup task
func (cfg *KhulnasoftFile) AddTask(task ...string) {
	if len(task) > 1 {
		cfg.Tasks = append(cfg.Tasks, khulnasoftTask{
			Command: task[0],
			Init:    task[1],
		})
	} else {
		cfg.Tasks = append(cfg.Tasks, khulnasoftTask{
			Command: task[0],
		})
	}
}

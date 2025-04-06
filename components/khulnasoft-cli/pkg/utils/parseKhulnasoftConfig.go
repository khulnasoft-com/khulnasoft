// Copyright (c) 2023 Khulnasoft GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package utils

import (
	"errors"
	"os"
	"path/filepath"

	khulnasoft "github.com/khulnasoft-com/khulnasoft/khulnasoft-protocol"
	yaml "gopkg.in/yaml.v2"
)

func ParseKhulnasoftConfig(repoRoot string) (*khulnasoft.KhulnasoftConfig, error) {
	if repoRoot == "" {
		return nil, errors.New("repoRoot is empty")
	}
	data, err := os.ReadFile(filepath.Join(repoRoot, ".khulnasoft.yml"))
	if err != nil {
		// .khulnasoft.yml not exist is ok
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, errors.New("read .khulnasoft.yml file failed: " + err.Error())
	}
	var config *khulnasoft.KhulnasoftConfig
	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, errors.New("unmarshal .khulnasoft.yml file failed" + err.Error())
	}
	return config, nil
}

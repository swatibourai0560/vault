// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !enterprise

package plugincatalog

import (
	"fmt"

	"github.com/hashicorp/vault/sdk/helper/pluginutil"
)

func (c *PluginCatalog) unpackPluginArtifact(plugin pluginutil.SetPluginInput) (bool, string, []byte, error) {
	return false, plugin.Command, plugin.Sha256, fmt.Errorf("enterprise-only feature: plugin artifact unpacking")
}

func (c *PluginCatalog) sanitize() error {
	return nil
}

/*

Copyright (c) 2022 - Present. Blend Labs, Inc. All rights reserved
Use of this source code is governed by a MIT license that can be found in the LICENSE file.

*/

package sourceutil

import (
	"fmt"
	"os"
	"time"
)

// CreateTempDir creates a temporary directory with a given prefix.
func CreateTempDir(prefix string) (*TempDir, error) {
	fullPath := fmt.Sprintf("%s_%d", prefix, time.Now().UTC().UnixNano())
	if err := os.MkdirAll(fullPath, 0755); err != nil {
		return nil, err
	}
	return &TempDir{fullPath}, nil
}

// TempDir is a directory that can be cleaned up with Close.
type TempDir struct {
	Path string
}

// Close removes the directory.
func (td TempDir) Close() error {
	return os.RemoveAll(td.Path)
}

/*

Copyright (c) 2022 - Present. Blend Labs, Inc. All rights reserved
Use of this source code is governed by a MIT license that can be found in the LICENSE file.

*/

package fileutil

import (
	"bufio"
	"os"

	"github.com/blend/go-sdk/ex"
)

// ReadLines reads a file and calls the handler for each line.
func ReadLines(filePath string, handler func(string) error) error {
	f, err := os.Open(filePath)
	if err != nil {
		return ex.New(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		err = handler(line)
		if err != nil {
			return ex.New(err)
		}
	}
	return nil
}

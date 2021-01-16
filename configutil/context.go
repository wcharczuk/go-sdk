/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package configutil

import (
	"context"
)

type configFilePathsKey struct{}

// WithConfigPaths adds config file paths to the context.
func WithConfigPaths(ctx context.Context, paths []string) context.Context {
	return context.WithValue(ctx, configFilePathsKey{}, paths)
}

// GetConfigPaths gets the config file paths from a context..
func GetConfigPaths(ctx context.Context) []string {
	if raw := ctx.Value(configFilePathsKey{}); raw != nil {
		if typed, ok := raw.([]string); ok {
			return typed
		}
	}
	return nil
}

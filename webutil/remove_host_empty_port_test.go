/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package webutil

import (
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestRemoveEmptyHostPort(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("foobar.com", RemoveHostEmptyPort("foobar.com"))
	assert.Equal("foobar.com", RemoveHostEmptyPort("foobar.com:"))
	assert.Equal("[00::00]", RemoveHostEmptyPort("[00::00]:"))
}

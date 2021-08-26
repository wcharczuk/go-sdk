/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package r2

import (
	"net/http"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestOptTransport(t *testing.T) {
	assert := assert.New(t)

	r := New(TestURL, OptTransport(&http.Transport{}))
	assert.NotNil(r.Client.Timeout)
}

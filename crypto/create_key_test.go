/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package crypto

import (
	"crypto/hmac"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func Test_CreateKey(t *testing.T) {
	t.Parallel()

	its := assert.New(t)

	key, err := CreateKey(32)
	its.Nil(err)
	its.Len(key, 32)

	key2, err := CreateKey(32)
	its.Nil(err)
	its.Len(key2, 32)

	its.False(hmac.Equal(key, key2))
}

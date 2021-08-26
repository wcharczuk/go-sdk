/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package configutil

import (
	"context"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestLazyInt32(t *testing.T) {
	its := assert.New(t)

	isNil := LazyInt32(nil)
	var value int32 = 0
	hasValue := LazyInt32(&value)
	var value2 int32 = 2
	hasValue2 := LazyInt32(&value2)

	var setValue int32
	its.Nil(SetInt32(&setValue, isNil, hasValue, hasValue2)(context.TODO()))
	its.Equal(2, setValue)
}

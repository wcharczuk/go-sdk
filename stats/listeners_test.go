/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package stats

import (
	"testing"

	"github.com/blend/go-sdk/assert"
	"github.com/blend/go-sdk/logger"
)

func TestAddErrorListeners(t *testing.T) {
	assert := assert.New(t)

	log := logger.None()
	AddErrorListeners(nil, nil)
	assert.False(log.HasListener(logger.Warning, ListenerNameStats))
	assert.False(log.HasListener(logger.Error, ListenerNameStats))
	assert.False(log.HasListener(logger.Fatal, ListenerNameStats))
	AddErrorListeners(log, NewMockCollector(32))
	assert.True(log.HasListener(logger.Warning, ListenerNameStats))
	assert.True(log.HasListener(logger.Error, ListenerNameStats))
	assert.True(log.HasListener(logger.Fatal, ListenerNameStats))
}

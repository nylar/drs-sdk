package drs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeviceStatus(t *testing.T) {
	ConfigSetup()
	result, err := UpdateDeviceStatus("", "")
	assert.NotNil(t, err)
	assert.False(t, result)
	assert.Equal(t, 400, err.Code)
	result, err = UpdateDeviceStatus("TEST", "2005-11-06 01:54:00")
	assert.NotNil(t, err)
	assert.False(t, result)
	assert.Equal(t, 400, err.Code)

	result, err = UpdateDeviceStatus("TEST", "2008-01-01T17:08:00Z")
	assert.Nil(t, err)
	assert.True(t, result)
}
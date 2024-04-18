package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidTimestampt(t *testing.T) {
	validTime := "2024-04-16T22:58:50+07:00"
	noTz := "2024-04-16T22:58:50"
	noSeparator := "2024-04-16 22:58:50+07:00"
	noSeparatorAndTz := "2024-04-16 22:58:50"
	random := "26/06/1900"

	assert.Equal(t, true, IsValidTimestampt(validTime))
	assert.Equal(t, false, IsValidTimestampt(noTz))
	assert.Equal(t, false, IsValidTimestampt(noSeparator))
	assert.Equal(t, false, IsValidTimestampt(noSeparatorAndTz))
	assert.Equal(t, false, IsValidTimestampt(random))
}

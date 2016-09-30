package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgramme(t *testing.T) {
	assert := assert.New(t)
	index := 1
	name := "programme"
	pid := "123"
	episode := 1
	series := 1
	thumbnail := "http://thumbnail.url"
	url := "http://programme.url"
	np := NewProgramme(index, series, episode, name, pid, thumbnail, url)
	assert.Equal(np.index, 1)
	assert.Equal(np.pid, "123")
	assert.Equal(np.episode, 1)
	assert.Equal(np.thumbnail, "http://thumbnail.url")
	assert.Equal(np.url, "http://programme.url")
	assert.Equal(np.name, "programme")
	assert.Equal(np.series, 1)
}

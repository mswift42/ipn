package main

import (
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

const mostpopular = "iplayermostpopular.html"

func TestFindTitle(t *testing.T) {
	assert := assert.New(t)
	html := loadTestHtml(mostpopular)
	html.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		title := findTitle(s)
		assert.NotEqual(title, "")
	})
	assert.Equal(len(html.Find(".list-item").Nodes), 40)
	html.Find(".list-item").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title := findTitle(s)
		assert.Equal(title, "Strictly Come Dancing")
		return false
	})
}
func TestFindSubtitle(t *testing.T) {
	assert := assert.New(t)
	html := loadTestHtml(mostpopular)
	assert.Equal(len(html.Find(".list-item").Nodes), 40)
	html.Find(".list-item").EachWithBreak(func(i int, s *goquery.Selection) bool {
		subtitle := findSubtitle(s)
		assert.Equal(subtitle, "Series 14: Week 3")
		return false
	})
}

func TestUrl(t *testing.T) {
	assert := assert.New(t)
	html := loadTestHtml(mostpopular)
	html.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		url := findUrl(s)
		assert.NotEqual(url, "")
	})
	html.Find(".list-item").EachWithBreak(func(i int, s *goquery.Selection) bool {
		url := findUrl(s)
		assert.Equal(url, "www.bbbc.co.uk/iplayer/episode/b07zhnf6/strictly-come-dancing-series-14-week-3")
		return false
	})
}

func TestNewProgramme(t *testing.T) {
	assert := assert.New(t)
	index := 1
	name := "programme"
	subtitle := "subtitle"
	pid := "123"
	episode := 1
	series := 1
	thumbnail := "http://thumbnail.url"
	url := "http://programme.url"
	np := NewProgramme(index, series, episode, name, subtitle, pid, thumbnail, url)
	assert.Equal(np.index, 1)
	assert.Equal(np.pid, "123")
	assert.Equal(np.episode, 1)
	assert.Equal(np.thumbnail, "http://thumbnail.url")
	assert.Equal(np.url, "http://programme.url")
	assert.Equal(np.name, "programme")
	assert.Equal(np.series, 1)
	assert.Equal(np.subtitle, "subtitle")
	testhtml := loadTestHtml("iplayermostpopular.html")
	assert.NotNil(testhtml)
}

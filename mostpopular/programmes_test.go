package mostpopular

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

const mostpopular = "iplayermostpopular.html"

func loadTestHtml(filename string) *goquery.Document {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}
	return doc
}
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
func TestFindThumbnail(t *testing.T) {
	assert := assert.New(t)
	html := loadTestHtml(mostpopular)
	html.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		thumbnail := findThumbnail(s)
		assert.NotEqual(thumbnail, "")
		assert.Equal(strings.HasSuffix(thumbnail, "jpg"), true)
		assert.Equal(strings.HasPrefix(thumbnail, "http://ichef.bbci"), true)
	})

}

func TestFindPid(t *testing.T) {
	assert := assert.New(t)
	html := loadTestHtml(mostpopular)
	html.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		pid := findPid(s)
		assert.NotEqual(pid, "")
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
		assert.Equal(strings.HasPrefix(url, "www.bbc.co.uk"), true)
	})
	html.Find(".list-item").EachWithBreak(func(i int, s *goquery.Selection) bool {
		url := findUrl(s)
		assert.Equal(url, "www.bbc.co.uk/iplayer/episode/b07zhnf6/strictly-come-dancing-series-14-week-3")
		return false
	})
}

func TestFindSynopsis(t *testing.T) {
	assert := assert.New(t)
	html := loadTestHtml(mostpopular)
	html.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		synopsis := findSynopsis(s)
		assert.NotEqual(synopsis, "")
	})
	html.Find(".list-item").EachWithBreak(func(i int, s *goquery.Selection) bool {
		synopsis := findSynopsis(s)
		assert.Equal(synopsis, "The couples must deliver an amazing routine to a classic movie track.")
		return false
	})
}

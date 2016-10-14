package main

import (
	"reflect"
	"strings"
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

func TestNewProgramme(t *testing.T) {
	assert := assert.New(t)
	index := 1
	name := "programme"
	subtitle := "subtitle"
	pid := "123"
	episode := 1
	series := 1
	synopsis := "this is programme 1"
	thumbnail := "http://thumbnail.url"
	url := "http://programme.url"
	np := NewProgramme(index, series, episode, name, subtitle, synopsis, pid, thumbnail, url)
	assert.Equal(np.index, 1)
	assert.Equal(np.pid, "123")
	assert.Equal(np.episode, 1)
	assert.Equal(np.thumbnail, "http://thumbnail.url")
	assert.Equal(np.url, "http://programme.url")
	assert.Equal(np.title, "programme")
	assert.Equal(np.series, 1)
	assert.Equal(np.subtitle, "subtitle")
	assert.Equal(np.synopsis, "this is programme 1")
	testhtml := loadTestHtml("iplayermostpopular.html")
	assert.NotNil(testhtml)
}

func Test_mostPopular(t *testing.T) {
	tests := []struct {
		name string
		want []*Programme
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := mostPopular(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. mostPopular() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_loadTestHtml(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want *goquery.Document
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := loadTestHtml(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. loadTestHtml() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_findTitle(t *testing.T) {
	type args struct {
		s *goquery.Selection
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findTitle(tt.args.s); got != tt.want {
			t.Errorf("%q. findTitle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_findSubtitle(t *testing.T) {
	type args struct {
		s *goquery.Selection
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findSubtitle(tt.args.s); got != tt.want {
			t.Errorf("%q. findSubtitle() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_findUrl(t *testing.T) {
	type args struct {
		s *goquery.Selection
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findUrl(tt.args.s); got != tt.want {
			t.Errorf("%q. findUrl() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_findThumbnail(t *testing.T) {
	type args struct {
		s *goquery.Selection
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findThumbnail(tt.args.s); got != tt.want {
			t.Errorf("%q. findThumbnail() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_findPid(t *testing.T) {
	type args struct {
		s *goquery.Selection
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := findPid(tt.args.s); got != tt.want {
			t.Errorf("%q. findPid() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}

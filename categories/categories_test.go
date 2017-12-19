package tv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const filmspage1 = "../tv/filmspage1.html"

func TestHasNextPage(t *testing.T) {
	assert := assert.New(t)
	th := tv.TestHtmlURL(filmspage1)
	doc, err := th.UrlDoc()
	assert.Nil(err)
	assert.NotNil(doc)
	s := doc.Find(".page > a").AttrOr("href", "")
	assert.Equal(s, "hallo")
}

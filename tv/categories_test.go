package tv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const filmspage1 = "../tv/filmspage1.html"

func TestHasNextPage(t *testing.T) {
	assert := assert.New(t)
	th := TestHtmlURL{filmspage1}
	c := make(chan *IplayerDocumentResult)
	go th.loadDocument(c)
	idr := <-c
	assert.Nil(idr.Error)
	assert.NotNil(idr.idoc)
	s := idr.idoc.Find(".page > a").AttrOr("href", "")
	assert.Equal(s, "hallo")
}

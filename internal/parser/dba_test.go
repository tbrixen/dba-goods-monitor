package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoldItemsThatRedirects(t *testing.T) {
	dba := NewDba()

	result := dba.IsActive("https://www.dba.dk/annonce/id-1110121538")

	assert.False(t, result)
}

func TestInactiveListingPage(t *testing.T) {
	dba := NewDba()

	result := dba.IsActive("https://www.dba.dk/prime-fuji-xf-56mm-f12/id-1110117857/")

	assert.False(t, result)
}

func TestStillActive(t *testing.T) {
	dba := NewDba()

	result := dba.IsActive("https://www.dba.dk/airpods-max-med-smart-case/id-1000455752/")

	assert.True(t, result)
}

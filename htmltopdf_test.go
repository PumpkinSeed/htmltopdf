package htmltopdf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample1(t *testing.T) {
	err := Example("examples/invoice-template-1/invoice.html", "./invoice1.pdf")

	assert.NoError(t, err)
}

func TestExample2(t *testing.T) {
	err := Example("examples/invoice-template-2/invoice.html", "./invoice2.pdf")

	assert.NoError(t, err)
}

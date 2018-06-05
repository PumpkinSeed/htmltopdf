package htmltopdf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	err := Example("examples/invoice-template-1/invoice.html")

	assert.NoError(t, err)
}

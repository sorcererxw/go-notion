package notion

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage_MarshalJSON(t *testing.T) {
	b, _ := json.Marshal(&Page{})
	var p Page
	json.Unmarshal(b, &p)
	assert.Equal(t, ObjectPage, p.Object)
}

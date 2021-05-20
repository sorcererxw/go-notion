package notion

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlock_MarshalJSON(t *testing.T) {
	b, _ := json.Marshal(&Block{})
	var block Block
	json.Unmarshal(b, &block)
	assert.Equal(t, ObjectBlock, block.Object)
}

package filecache

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMeta_MetadataField_RoundTrip(t *testing.T) {
	// Sample metadata as arbitrary JSON
	sampleMetadata := json.RawMessage(`{"foo":"bar","num":42,"arr":[1,2,3]}`)

	opts := &ItemOptions{
		Name:     "test-item",
		TTL:      2 * time.Hour,
		Fields:   Values{"a": "b"},
		Metadata: sampleMetadata,
	}

	// Convert ItemOptions to meta
	m := newMeta("test-key", opts, 0)
	assert.Equal(t, sampleMetadata, m.Metadata, "Metadata should be copied to meta")

	// Convert meta back to ItemOptions
	opts2 := metaToOptions(m)
	assert.Equal(t, sampleMetadata, opts2.Metadata, "Metadata should be preserved in round-trip")

	// Check other fields are also preserved
	assert.Equal(t, opts.Name, opts2.Name)
	assert.Equal(t, opts.TTL, opts2.TTL)
	assert.Equal(t, opts.Fields, opts2.Fields)
}

func TestMeta_MetadataField_Empty(t *testing.T) {
	opts := &ItemOptions{
		Name:     "empty-meta",
		TTL:      0,
		Fields:   nil,
		Metadata: nil,
	}

	m := newMeta("empty-key", opts, 0)
	assert.Nil(t, m.Metadata, "Metadata should be nil if not set")

	opts2 := metaToOptions(m)
	assert.Nil(t, opts2.Metadata, "Metadata should remain nil after round-trip")
}
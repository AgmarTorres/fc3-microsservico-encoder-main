package utils_test

import (
	"testing"

	"github.com/encoder/framework/utils"
	"github.com/stretchr/testify/require"
)

func TestIsJson(t *testing.T) {
	json := `{
		"id": "teste id",
		"file_path": "convite.mp4", 
		"status": "pending"
	}`

	err := utils.IsJson(json)
	require.Nil(t, err)

	json = `bru`

	err = utils.IsJson(json)
	require.Error(t, err)
}

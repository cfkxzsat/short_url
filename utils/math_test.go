package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//在init执行后，这些变量保留
func TestInit(t *testing.T) {
	assert.Equal(t, 62, Length, "Token length error")
	// fmt.Println("Token=>", Token)
	// fmt.Println("len=>", Length)
}

func TestIdToString(t *testing.T) {
	id := 125
	expectValue := "21"
	assert.Equal(t, expectValue, IdToString(id))
}

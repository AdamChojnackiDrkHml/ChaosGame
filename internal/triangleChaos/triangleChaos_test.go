package trianglechaos

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateChaosTriangle(t *testing.T) {
	iterations := 100000
	size := 1000
	stPt := [2]int{400, 350}

	bitmap := CreateChaosTriangle(size, stPt, iterations)
	fmt.Print(bitmap[699*1000+175], bitmap[699*1000+674], bitmap[274*1000+425])
	assert.True(t, bitmap[699*1000+175] == 0 || bitmap[699*1000+674] == 0 || bitmap[274*1000+425] == 0, "Something went wrong in bitmap setting")
}

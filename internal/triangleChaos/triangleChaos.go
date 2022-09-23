package trianglechaos

import (
	"crypto/rand"
	"fmt"
)

const (
	X = 1
	Y = 0
)

func CreateChaosTriangle(size int, startingPoint [2]int, iterations int) []uint8 {
	bitmap := make([][]uint8, size)

	for i := range bitmap {
		bitmap[i] = make([]uint8, size)
		for j := range bitmap[i] {
			bitmap[i][j] = 255
		}
	}

	A := [2]int{size - 1, 0}
	B := [2]int{size - 1, size - 1}
	C := [2]int{size - 1 - int(float64(size/2)*1.7), size / 2}
	fmt.Print(A, B, C)
	fmt.Println(C)
	refPts := [3][2]int{A, B, C}
	var refPt [2]int

	bitmap[A[Y]][A[X]] = 0
	bitmap[B[Y]][B[X]] = 0
	bitmap[C[Y]][C[X]] = 0
	bitmap[startingPoint[Y]][startingPoint[X]] = 0

	for i := 0; i < iterations; i++ {
		ra := make([]byte, 1)
		_, _ = rand.Read(ra)
		fmt.Println(ra)
		refPt = refPts[ra[0]%3]
		startingPoint[Y], startingPoint[X] = startingPoint[Y]/2+refPt[Y]/2, startingPoint[X]/2+refPt[X]/2
		bitmap[startingPoint[Y]][startingPoint[X]] = 0
		fmt.Println(refPt, startingPoint)
		break
	}

	flattenBitmap := make([]uint8, 0)

	for _, n := range bitmap {
		flattenBitmap = append(flattenBitmap, n...)
	}
	return flattenBitmap
}

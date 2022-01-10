package core

var dataset = []int{1, 2, 8, 17, 22, 37, 78, 113} // Mock dataset that refer to dataset in Question 1 ([1, X, 8, 17, Y, Z, 78, 113])
const (
	xPosition = 1 // Position of X
	yPosition = 4 // Position of Y
	zPosition = 5 // Position of Z
)

type XYZ struct {
	X int
	Y int
	Z int
}

// GetXYZByPosition is a function to get X, Y, Z by position from the dataset.
func GetXYZByPosition() XYZ {
	x := dataset[xPosition]
	y := dataset[yPosition]
	z := dataset[zPosition]
	return XYZ{x, y, z}
}

// GetXYZByRemoveKnowData is a function to get X, Y, Z by remove data you know.
func GetXYZByRemoveKnowData() XYZ {
	knowData := []int{1, 8, 17, 78, 113}
	datasetTemp := dataset
	for i := 0; i < len(knowData); i++ {
	tempLoop:
		for j := 0; j < len(datasetTemp); j++ {
			if knowData[i] == datasetTemp[j] {
				datasetTemp = append(datasetTemp[:j], datasetTemp[j+1:]...)
				break tempLoop
			}
		}
	}
	return XYZ{datasetTemp[0], datasetTemp[1], datasetTemp[2]}
}

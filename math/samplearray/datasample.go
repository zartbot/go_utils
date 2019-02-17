package samplearray

import "errors"

//DataSample is a function to proccess array sample
func Sample(arrLength int, samplenum int, totalnum int) (int, int, error) {
	var sampleInterval int = 1
	var samplestart int = 0
	//if samplenum <0 , will set auto sample interval  by total num
	if samplenum <= 0 {
		if totalnum < 1 {
			return 0, 0, errors.New("Invalid samplenum and totalnum, if sample number < 0, it will request totalnum > 1")
		}
		sampleInterval = arrLength / totalnum
		if sampleInterval == 0 {
			sampleInterval = 1
		}
		if arrLength < totalnum {
			sampleInterval = 1
		}
	} else {
		sampleInterval = samplenum
	}

	samplestart = arrLength%sampleInterval - 1

	//make sure the last value[max index] to be fetchout
	if samplestart < 0 {
		samplestart = samplestart + sampleInterval
	}

	//data is less than sample number, send out all
	if samplestart > arrLength {
		samplestart = 0
		sampleInterval = 1
	}
	return samplestart, sampleInterval, nil
}

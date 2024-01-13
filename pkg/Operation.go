package pkg

import ("math")

func Average(Values []float64) (float64,int) {
	Sum := float64(0)
	for i := 0; i < len(Values); i++ {
		Sum = Sum + Values[i]
	}
	AverageFloat := Sum / float64(len(Values))
	AverageInt := int(math.Round(float64(AverageFloat)))

	return AverageFloat, AverageInt
}

func Median(Values []float64) int {
	var Minimum float64
	var index int
	var AscendingOrder []float64
	for len(Values) > 0{
		a := 0
		if len(Values) == 1 {
			AscendingOrder = append(AscendingOrder, Values[0])
			break
		}
		Minimum = Values[a]
		for b := 0; b < len(Values); b++ {
			if Values[b] <= Minimum {
				Minimum = Values[b]
				index = b
			}
		}
		Values = append(Values[:index], Values[(index+1):]...)
		AscendingOrder = append(AscendingOrder, Minimum)
	}

	var MedianFloat float64
	var MedianInt int
	if len(AscendingOrder)%2 == 0 {
		Middleindex1 := (len(AscendingOrder)/2) - 1
		Middleindex2 := Middleindex1 + 1
		MedianFloat = float64((AscendingOrder[Middleindex1]+AscendingOrder[Middleindex2])/2)
		MedianInt = int(math.Round(MedianFloat))

	} else {
		Middleindex:= (len(AscendingOrder)/2)
		MedianFloat = AscendingOrder[Middleindex]
		MedianInt = int(math.Round(MedianFloat))
	}

	return MedianInt
}


func Variance(Average float64, Values []float64) (float64, int) {
	Numerator := float64(0)
	var Difference float64
	var SquareDifference float64
	for i:=0; i<len(Values); i++ {
		Difference = (Values[i] - Average) 
		SquareDifference = Difference * Difference
		Numerator += SquareDifference
	}
	VarianceFloat := Numerator/ float64(len(Values))
	VarianceInt := int(math.Round(VarianceFloat))

	return VarianceFloat, VarianceInt
}

func StandardDeviation(Variance float64) int {
	StandardDeviationFloat := math.Sqrt(Variance)
	StandardDeviationInt := int(math.Round(StandardDeviationFloat))

	return StandardDeviationInt
}
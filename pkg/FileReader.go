package pkg

import ("strconv"
        "os"
	    "bufio")

func FileReader(File string) ([]float64,error) {
	var err error
	file, err := os.Open(File)
	defer file.Close()
	// Create a scanner for the file
	scanner := bufio.NewScanner(file)
	// Read each line of the file
	var Value float64
	var Values []float64
	for scanner.Scan() {
		Value, err = strconv.ParseFloat(scanner.Text(), 64)
		Values = append(Values,float64(Value) )
	}
	return Values,err
}
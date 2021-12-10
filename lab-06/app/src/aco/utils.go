package aco

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getConverted(s string) (float64, error) {
	var (
		buf float64
		err error
	)
	if buf, err = strconv.ParseFloat(s, 64); err != nil {
		return 0, err
	}
	return buf, nil
}

func ParceConfigFile(filename string) (*Config) {
	path := "../meta/" + filename
	data, err := ioutil.ReadFile(path)
	if (err != nil) {
		fmt.Println("Error while reading: no such file or directory.")
		return nil
	}
	constants := strings.Fields(string(data))
	conf := new(Config)

	if conf.TRACE_WEIGHT, err = getConverted(constants[0]); err != nil {
		fmt.Println("Error while reading: invalid data in conf file.")
		return nil
	}

	if conf.TOUR_VISIB, err = getConverted(constants[1]); err != nil {
		fmt.Println("Error while reading: invalid data in conf file.")
		return nil
	}

	if conf.EVAP_RATE, err = getConverted(constants[2]); err != nil {
		fmt.Println("Error while reading: invalid data in conf file.")
		return nil
	}

	if conf.Q, err = getConverted(constants[3]); err != nil {
		fmt.Println("Error while reading: invalid data in conf file.")
		return nil
	}

	if conf.PHERO_INIT, err = getConverted(constants[4]); err != nil {
		fmt.Println("Error while reading: invalid data in conf file.")
		return nil
	}

	if conf.PHERO_EPS, err = getConverted(constants[5]); err != nil {
		fmt.Println("Error while reading: invalid data in conf file.")
		return nil
	}


	return conf  
}


func ParceGraphFile(filename string) *[][]int {
	path := "../meta/" + filename
	data, erropen := ioutil.ReadFile(path)
	if (erropen != nil) {
		fmt.Println("Error while reading: no such file or directory.")
		return nil
	}

	graph := make([][]int, 0)
	space := 0
	string_data := string(data)
	string_matrix := strings.Split(string_data[:len(string_data) - 1], "\n")

	for _, string_row := range string_matrix {
		if(string(string_row[len(string_row) - 1]) == " ") {
			space = 1
		}
		string_items := strings.Split(string_row[:len(string_row) - space], " ")

		row := make([]int, 0)

		for _, string_item := range string_items {
			item, errconv := strconv.Atoi(string_item)
			if errconv != nil {
				fmt.Printf("Error while reading: cannot convert %v to int.", string_item)
				return nil
			} else {
				row = append(row, item)
			}
		}
		graph = append(graph, row)
	}



	return &graph
}
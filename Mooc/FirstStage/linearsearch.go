package main

// LinearSearchV1 V1
func LinearSearchV1(tars [8]int, target int) int {
	for i, v := range tars {
		if v == target {
			return i
		}
	}
	return -1
}

// LinearSearchV2 V2
// Use the interface{}
// Use Age field to judge if or not
func LinearSearchV2(target interface{}, tars interface{}) int {
	for i, v := range tars.([]Student) {
		if v.Age == target.(Student).Age {
			return i
		}
	}
	return -1
}

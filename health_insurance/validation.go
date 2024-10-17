package main

// this function is used to check the age is valid or not
func isvalidage(age int) (result bool) {
	if age > 1 {

		result = true //if valid
	}
	return // if not valid
}

//validate the indicator count
func check_indicator(indicator int) (result bool) {
	//validate the indicator count
	if indicator > 0 || indicator < 8 {
		result = true //if indicator is not valid
	}
	return //if indicator is valid

}

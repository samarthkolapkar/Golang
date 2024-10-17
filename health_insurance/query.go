package main

import "fmt"

//This function is used to count persons have done the surgery
func number_of_surgery(modified_data []modify_struct) (count_surgery int) {
	for _, data := range modified_data {
		if data.numberof_surgery == 1 {
			count_surgery++
		}
	}
	return
}

// This function is used to count the number of person having the history of cancer
func number_of_historycancer(modified_data []modify_struct) (count_cancer int) {
	for _, data := range modified_data {
		if data.historyof_cancer == 1 {
			count_cancer++
		}
	}
	return
}

//This function is used to count the number of persons having the any allergy
func number_of_anyallergy(modified_data []modify_struct) (count_allergry int) {
	for _, data := range modified_data {
		if data.known_allergies == 1 {
			count_allergry++
		}

	}
	return
}

//This function is used to count the number of persons having the chronic_disease
func number_of_chronicdisease(modified_data []modify_struct) (count_chronic int) {
	for _, data := range modified_data {
		if data.chronic_disease == 1 {
			count_chronic++
		}
	}
	return
}

//This function is used to count the number of persons having the bloodpressure problem
func number_of_bloodpressure(modified_data []modify_struct) (count_pressure int) {
	for _, data := range modified_data {
		if data.blood_pressure == 1 {
			count_pressure++
		}
	}
	return
}

//This function is used to count the number of persons having the diabetes
func number_of_diabetes(modified_data []modify_struct) (count int) {

	for _, data := range modified_data {
		if data.diabetes == 1 {
			count++
		}
	}
	return
}

//This function is used to count the number of persons done any transplant
func number_of_transplant(modified_data []modify_struct) (count_transplant int) {
	for _, data := range modified_data {
		if data.any_transplant == 1 {
			count_transplant++
		}
	}
	return
}

//quries to check
func queries_toCheck(query_no int, modified_data []modify_struct) {

	switch query_no {
	case 1:
		result := number_of_diabetes(modified_data)
		fmt.Println("Count of persons having the diabetics is", result)
	case 2:
		result := number_of_transplant(modified_data)
		fmt.Println("Count of persons having the transplants is", result)

	case 3:
		result := number_of_bloodpressure(modified_data)
		fmt.Println("Count of persons having the problem of blood pressure is", result)
	case 4:
		result := number_of_chronicdisease(modified_data)
		fmt.Println("Count of persons having chronic disease is:", result)

	case 5:
		result := number_of_anyallergy(modified_data)
		fmt.Println("Count of persons having any allergy is:", result)

	case 6:
		result := number_of_historycancer(modified_data)
		fmt.Println("Count of persons having history of cancer is:", result)

	case 7:
		result := number_of_surgery(modified_data)
		fmt.Println("Count of persons having surgery done is:", result)
	}

}

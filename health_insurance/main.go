package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type modify_struct struct {
	age              int
	diabetes         int
	blood_pressure   int
	any_transplant   int
	chronic_disease  int
	BMI              int
	known_allergies  int
	historyof_cancer int
	numberof_surgery int
	premium_price    int
}
type health struct {
	age              int
	diabetes         int
	blood_pressure   int
	any_transplant   int
	chronic_disease  int
	height           int
	weight           int
	known_allergies  int
	historyof_cancer int
	numberof_surgery int
	premium_price    int
}

//checking the fields and count the matching indicator
func countMatchingIndicator(indicator int, fields []int) (count int) {
	for _, field := range fields {
		//count := 0
		if field == 1 {
			count++
			if count == indicator {
				// match_count++
				//total += data.premium_price
				break
			}

		}
	}
	return
}

//Calculate the total for matching indicator count
func Calculatetotalpremium(modified_data []modify_struct, match_count int) (total int) {
	// total := 0
	for _, data := range modified_data {

		total += data.premium_price
	}

	return
}

//calculating the average for matching indicator
func calulate_average(total, match_count, indicator int) float64 {
	if match_count == 0 {
		return 0
	}
	return float64(total) / (float64(match_count) / float64(indicator))
}

//function to calculate the premium for an indicator
func indicator_premium(indicator int, modified_data []modify_struct) (average float64) {
	//var count int
	var total int
	var match_count int
	// var average float64

	for _, data := range modified_data {
		//declare the fields to check
		fields := []int{
			data.diabetes,
			data.blood_pressure,
			data.any_transplant,
			data.chronic_disease,
			int(data.BMI),
			data.known_allergies,
			data.historyof_cancer,
			data.numberof_surgery,
		}
		count_match := countMatchingIndicator(indicator, fields)
		match_count += count_match
		if count_match > 0 {
			total = Calculatetotalpremium(modified_data, match_count)

			// total += data.premium_price
			// fmt.Println("Total", total)

		}
	}

	// fmt.Println(average)
	return calulate_average(total, match_count, indicator)
}

//This function is used to calculate the average premium
func average_premium(total, count int) (average float64) {
	if count > 0 {
		average = float64(total) / float64(count) //avrage
	}
	return average
}

//function to calculate the average premium for specified age
func average_query(age int, modified_data []modify_struct) float64 {
	var count int
	var total int
	for _, data := range modified_data {
		if data.age == age {
			total += data.premium_price // total for the particular age
			count++
		}
	}

	return average_premium(total, count)

}

//convert the integer values to string to store it into the csv file
func converttoString(row modify_struct) []string {
	return []string{
		strconv.Itoa(row.age),
		strconv.Itoa(row.diabetes), //convert the values int to string
		strconv.Itoa(row.blood_pressure),
		strconv.Itoa(row.any_transplant),
		strconv.Itoa(row.chronic_disease),
		strconv.Itoa(int(row.BMI)),
		strconv.Itoa(row.known_allergies),
		strconv.Itoa(row.historyof_cancer),
		strconv.Itoa(row.numberof_surgery),
		strconv.Itoa(row.premium_price),
	}

}

//this function is used to write the header to the csv file using writer interface
func writeheader(data *csv.Writer) {
	header := []string{
		"age",
		"diabetes",
		"blood_pressure",
		"any_transplant",
		"chronic_disease",
		"BMI",
		"known_allergies",
		"historyof_cancer",
		"numberof_surgery",
		"premium_price",
	}
	data.Write(header)
}

//This function is used to create the file and schedule a writer for csv file
func createfile() *csv.Writer {
	file, err := os.Create("modified_medicalpremium.csv") // create the file to store the data
	if err != nil {
		fmt.Println("Error while creating the file!!")
	}
	//defer file.Close() //schdule a function call to excuted after surrounding functions returns
	data := csv.NewWriter(file)
	// defer data.Flush()
	return data
}

//This function is used to write the data into an another file
func writetofile(modified_data []modify_struct) {

	data := createfile()
	//write the buffered data
	data.Flush()
	writeheader(data)

	//convert the data into string and stored it into the csv file
	for _, row := range modified_data {

		csv_data := converttoString(row)

		//write the data to the csv file
		data.Write(csv_data)
	}
	data.Flush()
	fmt.Println("Successfully created the file!!")
	fmt.Println("You can view the updated file!!")

}

//This function is used to update the surgery values in 0 and 1
func update_surgery(modified_data modify_struct) modify_struct {
	// var modified_data modify_struct
	//update the value of number of surgery to either 0 or 1
	if modified_data.numberof_surgery >= 1 {
		modified_data.numberof_surgery = 1 //update the number of surgery to 1
	} else {
		modified_data.numberof_surgery = 0 //update the number of surgery to 0
	}
	// writetofile(modified_data)
	// fmt.Println(modified_data)
	return modified_data
}

//This function is used to convert the BMI to 0 or 1 (1-normal and 0-otherwise)
func convert_BMI(modified_data modify_struct) modify_struct {
	if float64(modified_data.BMI) > 24.9 {
		modified_data.BMI = 1 //update the value of BMI to 1
	} else {
		modified_data.BMI = 0 //update the value of BMI to 0
	}

	return modified_data

	// fmt.Println(modified_data)
}

//This function is used to calculate the BMI
func calculate_BMI(data1 health) modify_struct {
	var modified_data modify_struct
	//create the BMI column ignoring the height and weights
	BMI := float64(data1.weight) / (float64(data1.height) * float64(data1.height)) * 10000 // ((data1.height) * (data1.height))
	modified_data.BMI = int(BMI)
	//update the value of every field to new struct
	modified_data.age = data1.age
	modified_data.diabetes = data1.diabetes
	modified_data.blood_pressure = data1.blood_pressure
	modified_data.any_transplant = data1.any_transplant
	modified_data.chronic_disease = data1.chronic_disease
	modified_data.known_allergies = data1.known_allergies
	modified_data.historyof_cancer = data1.historyof_cancer
	modified_data.numberof_surgery = data1.numberof_surgery
	modified_data.premium_price = data1.premium_price
	modified_data = convert_BMI(modified_data)
	// update_surgery(modified_data)
	// fmt.Println(modified_data)
	return modified_data
}

//This function is used to convert the string values to integer
func convert_toInt(data []string) health {
	var data1 health
	//convert the string data to integer
	data1.age, _ = strconv.Atoi(data[0])

	data1.diabetes, _ = strconv.Atoi(data[1])
	data1.blood_pressure, _ = strconv.Atoi(data[2])
	data1.any_transplant, _ = strconv.Atoi(data[3])
	data1.chronic_disease, _ = strconv.Atoi(data[4])
	data1.height, _ = strconv.Atoi(data[5])
	data1.weight, _ = strconv.Atoi(data[6])
	data1.known_allergies, _ = strconv.Atoi(data[7])
	data1.historyof_cancer, _ = strconv.Atoi(data[8])
	data1.numberof_surgery, _ = strconv.Atoi(data[9])
	data1.premium_price, _ = strconv.Atoi(data[10])
	//calculate_BMI(data1)
	return data1
}

//This function is used to process the data
func processData(data []string) modify_struct {

	healthdata := convert_toInt(data)     //convert the string values to interger
	modified := calculate_BMI(healthdata) //calculate the BMI calculation and update the data
	modified = update_surgery(modified)   //update the surgery field

	return modified
}

//This function is used to read csv and append the data to the modified slice
func readCSV() []modify_struct {
	file, err := os.Open("Medicalpremium.csv")
	if err != nil {
		fmt.Println("Error while opening the file:", err)
		return nil
	}
	defer file.Close()
	filereader := csv.NewReader(file) //csv reader to read the file
	var modified_data []modify_struct
	for {
		data, err := filereader.Read() //read the data and store it into the data variable
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading CSV:", err) //prompt the error
			return nil
		}
		//appending_to_struct(data)
		modified := processData(data)                   //process data and stored it into modified data
		modified_data = append(modified_data, modified) //update the modification
	}
	return modified_data
}

//This function is used to take the age as a input
func getInputAge() int {
	var age int
	fmt.Printf("Enter the age for which you want to calculate the premium:")
	fmt.Scan(&age)
	return age
}

//This function is used to take the indicator as a input
func get_indicator() int {
	var indicator int
	fmt.Printf("Enter the health indicator:")
	fmt.Scan(&indicator)
	return indicator
}

func main() {
	var age int
	var indicator int
	modified_data := readCSV()
	if modified_data == nil {
		log.Fatal("Failed to read CSV data.")
	}
	writetofile(modified_data)
	//filename := "modified_medicalpremium.csv"
	age = getInputAge()
	// final_quries(filename, age)
	if isvalidage(age) {
		average := average_query(age, modified_data)
		fmt.Printf("Average primium for age %d is %.02f \n", age, average)
	} else {
		fmt.Println("Age is not valid!!")
	}

	indicator = get_indicator()

	if check_indicator(indicator) {
		result := indicator_premium(indicator, modified_data)
		fmt.Printf("Average premium for indicator %d is %.02f\n", indicator, result)
	} else {
		log.Fatal("You have enter the indicator either zero or greater than 8!!")
	}

	for {
		var choice, query_no int
		fmt.Println("\t1:Number of person have Diabetics\n\t2:Number of person having any transplant\n\t3:Number of have blood pressur problem")
		fmt.Println("\t4:Number of person have chronic disease\n\t5:Number of persons have any allergey")
		fmt.Println("\t6:Number of person have history of cancerfamily\n\t7:Number of persons done the surgery")
		fmt.Printf("Choose the query you want to peform:")
		fmt.Scan(&query_no)
		queries_toCheck(query_no, modified_data)
		fmt.Print("Do you want to continue!! Press 1 to continue 0 to stop!!")
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}
	}
}

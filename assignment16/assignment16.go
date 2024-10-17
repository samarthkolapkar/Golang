package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

//this function will get the timezones
func GetZones(zonedir string) []string {
	var time_zones []string
	files, err := ioutil.ReadDir(zonedir)
	if err != nil {
		fmt.Println("Error while loading the file!!")
		return time_zones
	}
	for _, f := range files {
		if f.IsDir() {
			subtime_zone := GetZones(filepath.Join(zonedir, f.Name()))
			time_zones = append(time_zones, subtime_zone...)
		} else {
			time_zones = append(time_zones, strings.TrimPrefix(filepath.Join(zonedir, f.Name()), "/"))
		}
	}
	return time_zones
}

//This function will check the the timezone is valid or not
func isTimezone_valid(timezone string) bool {
	zonedir := "/usr/share/zoneinfo"
	timezone_list := GetZones(zonedir)
	for _, tz := range timezone_list {
		if timezone == tz {
			return true
		}
	}
	return false

	// timezone_list := []string{
	// 	"Asia/Kolkata",
	// 	"Africa/Abidjan",
	// 	"Africa/Accra",
	// 	"Africa/Addis_Ababa",
	// 	"Africa/Algiers",
	// 	"Africa/Asmara",
	// 	"Africa/Bamako",
	// 	"Africa/Bangui",
	// 	"Africa/Banjul",
	// 	"Africa/Bissau",
	// 	"Africa/Blantyre",
	// 	"Africa/Brazzaville",
	// 	"Africa/Bujumbura",
	// 	"Africa/Cairo",
	// 	"Africa/Casablanca",
	// 	"Africa/Ceuta",
	// 	"Africa/Conakry",
	// 	"Africa/Dakar",
	// 	"Africa/Dar_es_Salaam",
	// 	"Africa/Djibouti",
	// 	"Africa/Douala",
	// 	"Africa/El_Aaiun",
	// 	"Africa/Freetown",
	// 	"Africa/Gaborone",
	// 	"Africa/Harare",
	// 	"Africa/Johannesburg",
	// 	"Africa/Juba",
	// 	"Africa/Kampala",
	// 	"Africa/Khartoum",
	// 	"Africa/Kigali",
	// 	"Africa/Kinshasa",
	// 	"Africa/Lagos",
	// 	"Africa/Libreville",
	// 	"Africa/Lome",
	// 	"Africa/Luanda",
	// 	"Africa/Lubumbashi",
	// 	"Africa/Lusaka",
	// 	"Africa/Malabo",
	// 	"Africa/Maputo",
	// 	"Africa/Maseru",
	// 	"Africa/Mbabane",
	// 	"Africa/Mogadishu",
	// 	"Africa/Monrovia",
	// 	"Africa/Nairobi",
	// 	"Africa/Ndjamena",
	// 	"Africa/Niamey",
	// 	"Africa/Nouakchott",
	// 	"Africa/Ouagadougou",
	// 	"Africa/Porto-Novo",
	// 	"Africa/Sao_Tome",
	// 	"Africa/Tripoli",
	// 	"Africa/Tunis",
	// 	"Africa/Windhoek",
	// 	"America/Adak",
	// 	"America/Anchorage",
	// 	"America/Anguilla",
	// 	"America/Antigua",
	// 	"America/Araguaina",
	// 	"America/Argentina/Buenos_Aires",
	// 	"America/Argentina/Catamarca",
	// 	"America/Argentina/Cordoba",
	// 	"America/Argentina/Jujuy",
	// 	"America/Argentina/La_Rioja",
	// 	"America/Argentina/Mendoza",
	// 	"America/Argentina/Rio_Gallegos",
	// 	"America/Argentina/Salta",
	// 	"America/Argentina/San_Juan",
	// 	"America/Argentina/San_Luis",
	// 	"America/Argentina/Tucuman",
	// 	"America/Argentina/Ushuaia",
	// 	"America/Aruba",
	// 	"America/Asuncion",
	// 	"America/Atikokan",
	// 	"America/Bahia",
	// 	"America/Bahia_Banderas",
	// 	"America/Barbados",
	// 	"America/Belem",
	// 	"America/Belize",
	// 	"America/Blanc-Sablon",
	// 	"America/Boa_Vista",
	// 	"America/Bogota",
	// 	"America/Boise",
	// 	"America/Cambridge_Bay",
	// 	"America/Campo_Grande",
	// 	"America/Cancun",
	// 	"America/Caracas",
	// 	"America/Cayenne",
	// 	"America/Cayman",
	// 	"America/Chicago",
	// 	"America/Chihuahua",
	// 	"America/Costa_Rica",
	// 	"America/Creston",
	// 	"America/Cuiaba",
	// 	"America/Curacao",
	// 	"America/Danmarkshavn",
	// 	"America/Dawson",
	// 	"America/Dawson_Creek",
	// 	"America/Denver",
	// 	"America/Detroit",
	// 	"America/Dominica",
	// 	"America/Edmonton",
	// 	"America/Eirunepe",
	// 	"America/El_Salvador",
	// 	"America/Fort_Nelson",
	// 	"America/Fortaleza",
	// 	"America/Glace_Bay",
	// 	"America/Godthab",
	// 	"America/Goose_Bay",
	// 	"America/Grand_Turk",
	// 	"America/Grenada",
	// 	"America/Guadeloupe",
	// 	"America/Guatemala",
	// 	"America/Guayaquil",
	// 	"America/Guyana",
	// 	"America/Halifax",
	// 	"America/Havana",
	// 	"America/Hermosillo",
	// 	"America/Indiana/Indianapolis",
	// 	"America/Indiana/Knox",
	// 	"America/Indiana/Marengo",
	// }
	// for i := range timezone_list {
	// 	if timezone == timezone_list[i] {
	// 		return true
	// 	}
	// }
	// return false
}

//This function will load the time of the timezone that user will enter
func specified_timezone(t time.Time, timezone string) (string, error) {
	// currentTime := time.Now().UTC()
	//change the enter format as speciic format as available globally
	timezone = strings.Title(timezone)
	location_time, err := time.LoadLocation(timezone)
	if isTimezone_valid(timezone) {
		// return "", fmt.Errorf("Enter time zone is not a valid!!")
		fmt.Printf("%s is valid time zone\n", timezone)
	}

	location_timezone := t.In(location_time)
	supported_timezone := location_timezone.Format("2 Jan 2006 T 15:04:05Z")
	return supported_timezone, err

}
func main() {
	currentTime := time.Now().UTC()
	var timezone string
	fmt.Scan(&timezone)
	result, _ := specified_timezone(currentTime, timezone)
	// if err != nil {
	// 	fmt.Println("Error while loading location!!")
	// }
	fmt.Println(result)

}

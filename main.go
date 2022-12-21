package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// ### CONSTANTS ###
const CONFIG_PATH string = "config.json"

// struct contains information that should remain secret and not be included in the online repository
type config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Host         string `json:"api_URL"`
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// loads config into config struct from given file name
func (c *config) loadConfig() {
	jsonFile, err := os.Open(CONFIG_PATH)
	chkError(err)
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(c)
	chkError(err)
}

// saves config to file from struct
func (c config) saveConfig() {
	jsonFile, err := json.MarshalIndent(c, "", " ")
	chkError(err)
	err = os.WriteFile(CONFIG_PATH, jsonFile, 0644)
	chkError(err)
}

// region FlightCircle Structs
// ### FLIGHT CIRCLE STRUCTS ###
type Meta struct {
	Error  bool `json:"error"`
	Status int  `json:"status"`
}

// /access
type FCAccess struct {
	Data struct {
		Msg string `json:"msg"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /aircraft/{ FboID }(/{ AircraftID })
type FCAircraft struct {
	Data []struct {
		ID            string `json:"ID"`
		FboID         string `json:"FboID"`
		TailNumber    string `json:"tail_number"`
		Serial        string `json:"serial"`
		Year          string `json:"year"`
		Make          string `json:"make"`
		Model         string `json:"model"`
		Category      string `json:"category"`
		PreferredName string `json:"preferred_name"`
		HourlyRate    string `json:"hourly_rate"`
		Equipment     string `json:"equipment"`
		Requirements  string `json:"requirements"`
		Seats         string `json:"seats"`
		Range         string `json:"range"`
		Engines       string `json:"engines"`
		HobbsTotal    string `json:"hobbs_total"`
		TachTotal     string `json:"tach_total"`
		Description   string `json:"description"`
		Photo         string `json:"photo"`
		Status        string `json:"status"`
		Enabled       string `json:"enabled"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /squawks/{ FboID }(/{ ID })
type FCSquawks struct {
	Data []struct {
		AircraftID   string      `json:"AircraftID"`
		FboID        string      `json:"FboID"`
		ID           string      `json:"ID"`
		UserID       string      `json:"UserID"`
		TailNumber   string      `json:"tail_number"`
		ActionsTaken interface{} `json:"actions_taken"`
		Created      time.Time   `json:"created"`
		Description  string      `json:"description"`
		Status       string      `json:"status"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /maintenancereminders/{ FboID }(/{ ID })
type FCMaintenanceReminders struct {
	Data []struct {
		AircraftID    string `json:"AircraftID"`
		FboID         string `json:"FboID"`
		ID            string `json:"ID"`
		Aircraft      string `json:"aircraft"`
		Label         string `json:"label"`
		Notes         string `json:"notes"`
		PreferredName string `json:"preferred_name"`
		TailNumber    string `json:"tail_number"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /instructors/{ FboID }(/{ InstructorID })
type FCInstructors struct {
	Data []struct {
		FboID         string `json:"FboID"`
		InstructorID  string `json:"InstructorID"`
		UserID        string `json:"UserID"`
		ClassroomRate string `json:"classroom_rate"`
		Email         string `json:"email"`
		Enabled       string `json:"enabled"`
		FirstName     string `json:"first_name"`
		FlightRate    string `json:"flight_rate"`
		LastName      string `json:"last_name"`
		LastRental    string `json:"last_rental"`
		MiddleName    string `json:"middle_name"`
		Phone         string `json:"phone"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /schedules/{ FboID }/{ InstructorID }(/{ year }(/{ month }(/{ day }(/{ eyear }(/{ emonth }(/{ eday }))))))
type FCSchedules struct {
	Data []struct {
		AircraftID      string `json:"AircraftID"`
		FboID           string `json:"FboID"`
		ID              string `json:"ID"`
		InstructorID    string `json:"InstructorID"`
		UserID          string `json:"UserID"`
		TimezoneString  string `json:"timezone_string"`
		ArrivalDate     string `json:"arrival_date"`
		DepartDate      string `json:"depart_date"`
		ReservationType string `json:"reservation_type"`
		PilotName       string `json:"pilot_name"`
		InstructorName  string `json:"instructor_name"`
		Aircraft        string `json:"aircraft"`
		PreferredName   string `json:"preferred_name"`
		TailNumber      string `json:"tail_number"`
		AircraftStatus  string `json:"aircraft_status"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /flights/{ FboID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCFlights struct {
	Data []struct {
		ScheduleID                 string    `json:"ScheduleID"`
		ArrivalDate                time.Time `json:"arrival_date"`
		DepartDate                 time.Time `json:"depart_date"`
		PublicNotes                string    `json:"public_notes"`
		ReservationType            string    `json:"reservation_type"`
		InstructorID               string    `json:"InstructorID"`
		AircraftID                 string    `json:"AircraftID"`
		UserID                     string    `json:"UserID"`
		CheckinID                  string    `json:"CheckinID"`
		CheckinDate                time.Time `json:"checkin_date"`
		HobbsIn                    string    `json:"hobbs_in"`
		HobbsOut                   string    `json:"hobbs_out"`
		TachIn                     string    `json:"tach_in"`
		TachOut                    string    `json:"tach_out"`
		AircraftCharge             string    `json:"aircraft_charge"`
		AircraftRate               string    `json:"aircraft_rate"`
		InstructorFlightTime       string    `json:"instructor_flight_time"`
		InstructorFlightTimeCharge string    `json:"instructor_flight_time_charge"`
		InstructorFlightRate       string    `json:"instructor_flight_rate"`
		InstructorGroundTime       string    `json:"instructor_ground_time"`
		InstructorGroundTimeCharge string    `json:"instructor_ground_time_charge"`
		InstructorGroundRate       string    `json:"instructor_ground_rate"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /cancellations/{ FboID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCCancellations struct {
	Data []struct {
		CancelledBy        string    `json:"cancelled_by"`
		CancelledByID      int       `json:"CancelledByID"`
		CancelledDate      time.Time `json:"cancelled_date"`
		DepartDate         time.Time `json:"depart_date"`
		ReturnDate         time.Time `json:"return_date"`
		UserID             int       `json:"UserID"`
		ScheduleID         string    `json:"ScheduleID"`
		Notice             string    `json:"notice"`
		PublicNotes        string    `json:"public_notes"`
		CancellationReason string    `json:"cancellation_reason"`
		CancellationNotes  string    `json:"cancellation_notes"`
		Resource           string    `json:"resource"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /ledger/{ FboID }/{ UserID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCLedger struct {
	Data []struct {
		EntryType   string      `json:"entry_type"`
		EntryDate   time.Time   `json:"entry_date"`
		Description string      `json:"description"`
		Notes       string      `json:"notes"`
		AircraftID  interface{} `json:"AircraftID"`
		ReceiptLink string      `json:"receipt_link"`
		UserID      int         `json:"UserID"`
		User        string      `json:"user"`
		Amount      string      `json:"amount"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /user/describe
type FCSelf struct {
	Data []struct {
		UserID           string `json:"UserID"`
		AopaID           string `json:"aopa_id"`
		FirstName        string `json:"first_name"`
		MiddleName       string `json:"middle_name"`
		LastName         string `json:"last_name"`
		FboID            string `json:"FboID"`
		OrganizationName string `json:"OrganizationName"`
		TimezoneString   string `json:"timezone_string"`
		Email            string `json:"email"`
		CustomFields     struct {
			OrganizationDate     string `json:"organization_date"`
			OrganizationApproval string `json:"organization_approval"`
		} `json:"custom_fields"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /users/{ FboID }
type FCUsers struct {
	Data []struct {
		CustomerID              string `json:"CustomerID"`
		FirstName               string `json:"first_name"`
		LastName                string `json:"last_name"`
		Email                   string `json:"email"`
		DateOfBirth             string `json:"Date of Birth"`
		Address                 string `json:"address"`
		Address2                string `json:"address2"`
		City                    string `json:"city"`
		State                   string `json:"state"`
		ZipCode                 string `json:"zip_code"`
		Phone                   string `json:"phone"`
		EmergencyContactName    string `json:"emergency_contact_name"`
		EmergencyContactPhone   string `json:"emergency_contact_phone"`
		Balance                 string `json:"balance"`
		Created                 string `json:"created"`
		LastMedical             string `json:"Last Medical"`
		MedicalExpiration       string `json:"Medical Expiration"`
		LastFAAFlightReview     string `json:"Last FAA Flight Review"`
		RenterSInsuranceExpires string `json:"Renters Insurance Expires"`
		LastLocalFlightReview   string `json:"Last Local Flight Review"`
		CompanyName             string `json:"company_name"`
		Status                  string `json:"Status"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /user/schedule/next/{ FboID }(/{ UserID })
type FCNext struct {
	Data []struct {
		AircraftID      string    `json:"AircraftID"`
		FboID           string    `json:"FboID"`
		ID              string    `json:"ID"`
		InstructorID    string    `json:"InstructorID"`
		UserID          string    `json:"UserID"`
		TimezoneString  string    `json:"timezone_string"`
		ArrivalDate     time.Time `json:"arrival_date"`
		DepartDate      time.Time `json:"depart_date"`
		ReservationType string    `json:"reservation_type"`
		PilotName       string    `json:"pilot_name"`
		InstructorName  string    `json:"instructor_name"`
		Aircraft        string    `json:"aircraft"`
		PreferredName   string    `json:"preferred_name"`
		TailNumber      string    `json:"tail_number"`
		AircraftStatus  string    `json:"aircraft_status"`
	} `json:"data"`
	Meta `json:"meta"`
}

// /user/schedules/{ FboID }/{ UserID }(/{ year }(/{ month }(/{ day }(/{ eyear }(/{ emonth }(/{ eday }))))))
type FCUserSchedule struct {
	Data []struct {
		AircraftID      string `json:"AircraftID"`
		FboID           string `json:"FboID"`
		ID              string `json:"ID"`
		InstructorID    string `json:"InstructorID"`
		UserID          string `json:"UserID"`
		TimezoneString  string `json:"timezone_string"`
		ArrivalDate     string `json:"arrival_date"`
		DepartDate      string `json:"depart_date"`
		ReservationType string `json:"reservation_type"`
		PilotName       string `json:"pilot_name"`
		InstructorName  string `json:"instructor_name"`
		Aircraft        string `json:"aircraft"`
		PreferredName   string `json:"preferred_name"`
		TailNumber      string `json:"tail_number"`
		AircraftStatus  string `json:"aircraft_status"`
	} `json:"data"`
	Meta `json:"meta"`
}

//endregion

// ### MAIN ###
func main() {
	//load config
	var c config
	c.loadConfig()

	//getToken(cfg)
	var fcSelf FCSelf
	apiCall("/user/describe", &fcSelf, &c)
	var fcAircraft FCAircraft
	apiCall("/aircraft/"+fcSelf.Data[0].FboID, &fcAircraft, &c)
	var fcUsers FCUsers
	apiCall("/users/"+fcSelf.Data[0].FboID, &fcUsers, &c)

}
func apiCall(uri string, v any, c *config) {
	chkError(json.Unmarshal(tryGet(uri, c), v))
}

func httpGet(host, uri string, bearer string) (respBody []byte, status int) {
	req, err := http.NewRequest("GET", host+uri, nil)
	chkError(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Bearer "+bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	chkError(err)
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	chkError(err)
	status = resp.StatusCode

	return
}

func httpPost(host, uri string, requestBody []byte) (respBody []byte, status int) {
	req, err := http.NewRequest("POST", host+uri, bytes.NewReader(requestBody))
	chkError(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	chkError(err)
	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	chkError(err)
	status = resp.StatusCode

	return
}

func getTokenFromRefresh(c *config) (statusCode int) {
	//use a refresh token to get an access token
	requestBody, err := json.Marshal(map[string]string{
		"refresh_token": c.RefreshToken,
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
	})
	chkError(err)

	//make request
	respBody, statusCode := httpPost(c.Host, "/token", requestBody)

	//unmarshal response
	err = json.Unmarshal(respBody, &c)
	chkError(err)

	return
}

func tryGet(uri string, c *config) (respBody []byte) {
	respBody, status := httpGet(c.Host, uri, c.AccessToken)
	if status != 200 {
		getToken(c)
		respBody, status = httpGet(c.Host, uri, c.AccessToken)
		if status != 200 {
			log.Fatalln(respBody, status)
		}
	}
	return
}

func getTokensFromCode(c *config) (statusCode int) {

	//create authorization code request body
	requestBody, err := json.Marshal(map[string]string{
		"code":          c.Code,
		"client_id":     c.ClientID,
		"client_secret": c.ClientSecret,
	})
	chkError(err)

	//make request
	respBody, statusCode := httpPost(c.Host, "/token", requestBody)

	//unmarshal response
	json.Unmarshal(respBody, &c)

	return
}

func getToken(c *config) {
	if getTokenFromRefresh(c) != 200 {
		if getTokensFromCode(c) != 200 {
			log.Fatalln("Failed to get token")
		}
	}
	c.saveConfig()
}

func chkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

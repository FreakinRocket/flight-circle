package main

import (
	"fmt"
	"time"

	"github.com/FreakinRocket/zapi"
)

// region FlightCircle Data Structs
// ### FLIGHT CIRCLE STRUCTS ###

// access
type Access struct {
	Msg string `json:"msg"`
}

// aircraft
type Aircraft struct {
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
}

// squawk
type Squawk struct {
	AircraftID   string      `json:"AircraftID"`
	FboID        string      `json:"FboID"`
	ID           string      `json:"ID"`
	UserID       string      `json:"UserID"`
	TailNumber   string      `json:"tail_number"`
	ActionsTaken interface{} `json:"actions_taken"`
	Created      time.Time   `json:"created"`
	Description  string      `json:"description"`
	Status       string      `json:"status"`
}

// maintenance reminder
type MXReminder struct {
	AircraftID    string `json:"AircraftID"`
	FboID         string `json:"FboID"`
	ID            string `json:"ID"`
	Aircraft      string `json:"aircraft"`
	Label         string `json:"label"`
	Notes         string `json:"notes"`
	PreferredName string `json:"preferred_name"`
	TailNumber    string `json:"tail_number"`
}

// instructor
type Instructor struct {
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
}

// schedule
type Schedule struct {
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
}

// flight
type Flight struct {
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
}

// cancellation
type Cancellation struct {
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
}

// ledger
type Ledger struct {
	EntryType   string      `json:"entry_type"`
	EntryDate   time.Time   `json:"entry_date"`
	Description string      `json:"description"`
	Notes       string      `json:"notes"`
	AircraftID  interface{} `json:"AircraftID"`
	ReceiptLink string      `json:"receipt_link"`
	UserID      int         `json:"UserID"`
	User        string      `json:"user"`
	Amount      string      `json:"amount"`
}

// self
type Self struct {
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
}

// user
type User struct {
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
}

// next
type Next struct {
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
}

// user schedule
type UserSchedule struct {
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
}

//endregion

// region FlightCircle End Point Structs
// meta
type Meta struct {
	Error  bool `json:"error"`
	Status int  `json:"status"`
}

// /access
type FCAccess struct {
	Acesses []Access `json:"data"`
	Meta    `json:"meta"`
}

// /aircraft/{ FboID }(/{ AircraftID })
type FCAircraft struct {
	Aircraft []Aircraft `json:"data"`
	Meta     `json:"meta"`
}

// /squawks/{ FboID }(/{ ID })
type FCSquawks struct {
	Squawks []Squawk `json:"data"`
	Meta    `json:"meta"`
}

// /maintenancereminders/{ FboID }(/{ ID })
type FCMaintenanceReminders struct {
	MXReminders []MXReminder `json:"data"`
	Meta        `json:"meta"`
}

// /instructors/{ FboID }(/{ InstructorID })
type FCInstructors struct {
	Instructors []Instructor `json:"data"`
	Meta        `json:"meta"`
}

// /schedules/{ FboID }/{ InstructorID }(/{ year }(/{ month }(/{ day }(/{ eyear }(/{ emonth }(/{ eday }))))))
type FCSchedules struct {
	Schedules []Schedule `json:"data"`
	Meta      `json:"meta"`
}

// /flights/{ FboID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCFlights struct {
	Flights []Flight `json:"data"`
	Meta    `json:"meta"`
}

// /cancellations/{ FboID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCCancellations struct {
	Cancellations []Cancellation `json:"data"`
	Meta          `json:"meta"`
}

// /ledger/{ FboID }/{ UserID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCLedger struct {
	Ledgers []Ledger `json:"data"`
	Meta    `json:"meta"`
}

// /user/describe
type FCSelf struct {
	Selfs []Self `json:"data"`
	Meta  `json:"meta"`
}

// /users/{ FboID }
type FCUsers struct {
	Users []User `json:"data"`
	Meta  `json:"meta"`
}

// /user/schedule/next/{ FboID }(/{ UserID })
type FCNext struct {
	Nexts []Next `json:"data"`
	Meta  `json:"meta"`
}

// /user/schedules/{ FboID }/{ UserID }(/{ year }(/{ month }(/{ day }(/{ eyear }(/{ emonth }(/{ eday }))))))
type FCUserSchedule struct {
	UserSchedules []UserSchedule `json:"data"`
	Meta          `json:"meta"`
}

//endregion

// ### MAIN ###
func main() {
	//load config
	var c zapi.Config
	c.LoadConfig()

	//get current date and time
	currentTime := time.Now()
	//year, month, day := currentTime.Date()
	weekday := int(currentTime.Weekday())
	_, weeknum := currentTime.ISOWeek()

	//calculate days left in 2 week block
	daysLeft := 6 - weekday
	if weeknum%2 == 1 {
		daysLeft += 7
	}

	//calculate end date of 2 week block
	endTime := currentTime.AddDate(0, 0, daysLeft)
	//endYearInt, endMonthMonth, endDayInt := endTime.Date()

	//calculate start date of 2 week block
	startTime := currentTime.AddDate(0, 0, -1*(13-daysLeft))
	//startYearInt, startMonthMonth, startDayInt := startTime.Date()

	fmt.Println(startTime.Format("01/02/06"))
	fmt.Println(currentTime.Format("01/02/06"))
	fmt.Println(endTime.Format("01/02/06"))

	endYear := endTime.Format("2006")
	endMonth := endTime.Format("01")
	endDay := endTime.Format("02")
	startYear := startTime.Format("2006")
	startMonth := startTime.Format("01")
	startDay := startTime.Format("02")

	//get information about currently logged in user
	var fcSelf FCSelf
	zapi.ApiCall("/user/describe", &fcSelf, &c)

	//list all aircraft from a given FboID
	var fcAircraft FCAircraft
	zapi.ApiCall("/aircraft/"+fcSelf.Selfs[0].FboID, &fcAircraft, &c)

	//create list of active aircraft
	var aircraft []Aircraft
	for _, a := range fcAircraft.Aircraft {
		if a.Status != "0" {
			aircraft = append(aircraft, a)
		}
	}

	//get schedule entries for current two week block
	var fcSchedules FCSchedules
	callString := fmt.Sprint("/schedules/", fcSelf.Selfs[0].FboID, "/all/", startYear, "/", startMonth, "/", startDay, "/", endYear, "/", endMonth, "/", endDay)
	zapi.ApiCall(callString, &fcSchedules, &c)

	//filter to entries only with a plane attached
	var schedules []Schedule
	scheduleCounts := make(map[string]int)

	for _, s := range fcSchedules.Schedules {
		if s.Aircraft != "" {
			schedules = append(schedules, s)
			//add a count per each tail number
			scheduleCounts[s.TailNumber] += 1
		}
	}
	fmt.Println(scheduleCounts)
}

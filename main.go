package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FreakinRocket/zapi"
	"github.com/FreakinRocket/zjson"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// region MARS Structs
type Fleet struct {
	Meta struct {
		LastUpdated time.Time `json:"last_updated"`
	} `json:"meta"`
	Planes         map[string]Aircraft `json:"aircraft"`
	FilePath       string              `json:"-"`
	ScheduleCounts map[string]int      `json:"schedule_counts"`
	FboID          string              `json:"fbo_id"`
}

type Aircraft struct {
	TailNumber     string    `json:"tail_number"`
	DaysSinceFlown int       `json:"days_since_flown"`
	LastFlown      time.Time `json:"last_flown"`
}

func NewFleet() *Fleet {
	var f Fleet
	f.Planes = make(map[string]Aircraft)
	f.ScheduleCounts = make(map[string]int)
	return &f
}

//endregion

// region FlightCircle Data Structs
// ### FLIGHT CIRCLE STRUCTS ###

// access
type FC_Access struct {
	Msg string `json:"msg"`
}

// aircraft
type FC_Aircraft struct {
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
type FC_Squawk struct {
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
type FC_MXReminder struct {
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
type FC_Instructor struct {
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
type FC_Schedule struct {
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
type FC_Flight struct {
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
type FC_Cancellation struct {
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
type FC_Ledger struct {
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
type FC_Self struct {
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
type FC_User struct {
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
type FC_Next struct {
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
type FC_UserSchedule struct {
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
type FCMeta struct {
	Error  bool `json:"error"`
	Status int  `json:"status"`
}

// /access
type FCAccess struct {
	Acesses []FC_Access `json:"data"`
	FCMeta  `json:"meta"`
}

// /aircraft/{ FboID }(/{ AircraftID })
type FCAircraft struct {
	Aircraft []FC_Aircraft `json:"data"`
	FCMeta   `json:"meta"`
}

// /squawks/{ FboID }(/{ ID })
type FCSquawks struct {
	Squawks []FC_Squawk `json:"data"`
	FCMeta  `json:"meta"`
}

// /maintenancereminders/{ FboID }(/{ ID })
type FCMaintenanceReminders struct {
	MXReminders []FC_MXReminder `json:"data"`
	FCMeta      `json:"meta"`
}

// /instructors/{ FboID }(/{ InstructorID })
type FCInstructors struct {
	Instructors []FC_Instructor `json:"data"`
	FCMeta      `json:"meta"`
}

// /schedules/{ FboID }/{ InstructorID }(/{ year }(/{ month }(/{ day }(/{ eyear }(/{ emonth }(/{ eday }))))))
type FCSchedules struct {
	Schedules []FC_Schedule `json:"data"`
	FCMeta    `json:"meta"`
}

// /flights/{ FboID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCFlights struct {
	Flights []FC_Flight `json:"data"`
	FCMeta  `json:"meta"`
}

// /cancellations/{ FboID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCCancellations struct {
	Cancellations []FC_Cancellation `json:"data"`
	FCMeta        `json:"meta"`
}

// /ledger/{ FboID }/{ UserID }?year={ year }&month={ month }&day={ day }&eyear={ eyear }&emonth={ emonth }&eday={ eday }
type FCLedger struct {
	Ledgers []FC_Ledger `json:"data"`
	FCMeta  `json:"meta"`
}

// /user/describe
type FCSelf struct {
	Selfs  []FC_Self `json:"data"`
	FCMeta `json:"meta"`
}

// /users/{ FboID }
type FCUsers struct {
	Users  []FC_User `json:"data"`
	FCMeta `json:"meta"`
}

// /user/schedule/next/{ FboID }(/{ UserID })
type FCNext struct {
	Nexts  []FC_Next `json:"data"`
	FCMeta `json:"meta"`
}

// /user/schedules/{ FboID }/{ UserID }(/{ year }(/{ month }(/{ day }(/{ eyear }(/{ emonth }(/{ eday }))))))
type FCUserSchedule struct {
	UserSchedules []FC_UserSchedule `json:"data"`
	FCMeta        `json:"meta"`
}

//endregion

// region SendGrid End Point Structs

//region end

// region Connectors Struct
type Connectors struct {
	EmailFrom     string `json:"email_from"`
	EmailTo       string `json:"email_to"`
	EmailFromName string `json:"email_from_name"`
	EmailToName   string `json:"email_to_name"`
	SendGridKey   string `json:"sendgrid_key"`
	FilePath      string `json:"-"`
}

//endregion

// ### CONSTANTS ###
const FLEET_PATH string = "fleet.json"
const FLIGHTCIRCLE_PATH string = "flightcircle.json"
const GMAIL_PATH string = "gmail.json"
const CONNECTORS_PATH string = "connectors.json"

// ### MAIN ###
func main() {

	// region Initialization
	//load flight circle config
	var fc zapi.Config
	fc.FilePath = FLIGHTCIRCLE_PATH
	zjson.LoadJSON(&fc, fc.FilePath)

	//load Fleet
	mars := NewFleet()
	mars.FilePath = FLEET_PATH
	zjson.LoadJSON(&mars, mars.FilePath)
	mars.Meta.LastUpdated = time.Now()

	//load gmail config
	var gm zapi.Config
	gm.FilePath = GMAIL_PATH
	zjson.LoadJSON(&gm, gm.FilePath)

	//load connectors config
	var con Connectors
	con.FilePath = CONNECTORS_PATH
	zjson.LoadJSON(&con, con.FilePath)

	//endregion

	// region Main Code
	//get information about currently logged in user
	var fcSelf FCSelf
	zapi.Call("/user/describe", &fcSelf, &fc)

	//set FBO ID for easier access
	mars.FboID = fcSelf.Selfs[0].FboID

	//list all aircraft from a given FboID
	var fcAircraft FCAircraft
	zapi.Call("/aircraft/"+mars.FboID, &fcAircraft, &fc)

	//create list of active aircraft
	aircraft := make(map[string]FC_Aircraft)
	for _, a := range fcAircraft.Aircraft {
		if a.Status != "0" {
			aircraft[a.ID] = a
		}
	}

	//get schedule entries for current two week block
	var fcSchedules FCSchedules
	startDate, endDate := calcDateBlock()
	callString := fmt.Sprint("/schedules/", mars.FboID, "/all", makeFCDateString(startDate, endDate))
	zapi.Call(callString, &fcSchedules, &fc)

	//filter to entries only with a plane attached
	var schedules []FC_Schedule
	scheduleCounts := make(map[string]int)

	for _, s := range fcSchedules.Schedules {
		if s.Aircraft != "" {
			schedules = append(schedules, s)
			//add a count per each tail number
			scheduleCounts[s.AircraftID] += 1
		}
	}

	//get flights over past 30 days
	var fcFlights FCFlights
	startDate, endDate = calc30DaysAgo()
	callString = fmt.Sprint("/flights/", mars.FboID, makeFCDateStringReq(startDate, endDate))
	zapi.Call(callString, &fcFlights, &fc)

	flights := make(map[string][]FC_Flight)
	for _, f := range fcFlights.Flights {
		if f.AircraftID != "" {
			flights[f.AircraftID] = append(flights[f.AircraftID], f)
		}
	}

	//loop through the flights and get the most recent date an aircraft was flown
	for k, v := range flights {
		//create copy of current plane
		entry := mars.Planes[k]
		entry.TailNumber = mars.Planes[k].TailNumber

		//for each flight of this aircraft, check if it is more recent than the saved LastFlown, if it is, replace last flown
		for _, fs := range v {
			//if the stored last flown date is older than this flight, set last flown to this flight
			if fs.DepartDate.After(entry.LastFlown) {
				entry.LastFlown = fs.DepartDate
			}
		}

		//update days since flown
		entry.DaysSinceFlown = int(time.Since(entry.LastFlown.Local()).Hours() / 24)
		mars.Planes[k] = entry
		//fs.DepartDate
	}
	for k, v := range mars.Planes {
		fmt.Println(aircraft[k].TailNumber + " " + v.LastFlown.Local().Format("01/02/06") + "  " + strconv.Itoa(v.DaysSinceFlown) + " days since last flown.")
	}

	zjson.SaveJSON(&mars, FLEET_PATH)
	//endregion

	// region SendGrid test
	from := mail.NewEmail(con.EmailFromName, con.EmailFrom)
	subject := "SP034"
	to := mail.NewEmail(con.EmailToName, con.EmailTo)
	plainTextContent := "Awesome"
	htmlContent := "<strong>Awesome</strong>"
	marsJson, _ := json.MarshalIndent(mars, "", " ")
	marsb64 := base64.URLEncoding.EncodeToString(marsJson)

	attachment := mail.NewAttachment()

	attachment.SetFilename(FLEET_PATH).SetContent(marsb64).SetType("application/json")

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent).AddAttachment(attachment)
	client := sendgrid.NewSendClient(con.SendGridKey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Good")
	}

	//endregion
}

// region My Functions
// returns a string with a leading "/"
func makeFCDateString(startDate, endDate time.Time) (dateString string) {

	startYear := startDate.Format("2006")
	startMonth := startDate.Format("01")
	startDay := startDate.Format("02")
	endYear := endDate.Format("2006")
	endMonth := endDate.Format("01")
	endDay := endDate.Format("02")

	return fmt.Sprint("/", startYear, "/", startMonth, "/", startDay, "/", endYear, "/", endMonth, "/", endDay)
}

func calcDateBlock() (startDate, endDate time.Time) {
	//get current date and time
	currentTime := time.Now()

	//get day of week (0-6; 0=Sun), and get week number (1-53)
	weekday := int(currentTime.Weekday())
	_, weeknum := currentTime.ISOWeek()

	//calculate days left in 2 week block
	daysLeft := 6 - weekday
	if weeknum%2 == 1 {
		daysLeft += 7
	}

	//calculate end date of 2 week block
	endDate = currentTime.AddDate(0, 0, daysLeft)

	//calculate start date of 2 week block
	startDate = currentTime.AddDate(0, 0, -1*(13-daysLeft))

	return
}

// returns a string with a leading "?"
func makeFCDateStringReq(startDate, endDate time.Time) (dateString string) {
	startYear := startDate.Format("2006")
	startMonth := startDate.Format("01")
	startDay := startDate.Format("02")
	endYear := endDate.Format("2006")
	endMonth := endDate.Format("01")
	endDay := endDate.Format("02")

	return fmt.Sprint("?year=", startYear, "&month=", startMonth, "&day=", startDay, "&eyear=", endYear, "&emonth=", endMonth, "&eday=", endDay)
}

// returns the date 30 days ago
func calc30DaysAgo() (oldDate, today time.Time) {
	currentTime := time.Now()
	return currentTime.AddDate(0, 0, -1*30), currentTime
}

//endregion

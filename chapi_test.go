package chapi_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jimsmart/chapi"
)

// TODO(js) Documentation - need example code.

var apiKey string

const facebookCompanyNumber = "06331310"

const icelandCompanyNumber = "02800588"
const jcbCompanyNumber = "00561597"
const richardBransonOfficerID = "fPsul1-gLgzfRlgRvGBL14iNV3c"
const lordBamfordOfficerID = "KwkjxuswE9qwWKLU0ndEaau9cq0"

func init() {
	apiKey = os.Getenv("COMPANIES_HOUSE_API_KEY")
	if len(apiKey) == 0 {
		panic("COMPANIES_HOUSE_API_KEY environment variable not set")
	}
}

func TestDefaultHTTPClient(t *testing.T) {
	if chapi.DefaultHTTPClient == nil {
		t.Error("Expected DefaultHTTPClient to not be nil")
	}
}

// TODO(js) Look into what exactly Search is returning?

// func TestClient_searchCompanies1(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.SearchCompanies("Iceland Limited", 5, -1)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	l := len(res.Items)
// 	if l != 5 {
// 		t.Error("Expected 5 results, got", l)
// 	}

// 	fmt.Printf("%+v\n", res.Items[0])
// }

// func TestClient_searchCompanies2(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.SearchCompanies("Bamford Excavators", 5, -1)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	l := len(res.Items)
// 	if l != 5 {
// 		t.Error("Expected 5 results, got", l)
// 	}

// 	fmt.Printf("%+v\n", res.Items[0])
// }

// func TestClient_searchOfficers1(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.SearchOfficers("Sir Richard Branson", 5, -1)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	l := len(res.Items)
// 	if l != 5 {
// 		t.Error("Expected 5 results, got", l)
// 	}

// 	fmt.Printf("%+v\n", res.Items[0])
// }

// func TestClient_companyProfile1(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.CompanyProfile(icelandCompanyNumber)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// l := len(res.Items)
// 	// if l != 5 {
// 	// 	t.Error("Expected 5 results, got", l)
// 	// }

// 	fmt.Printf("%+v\n", res)
// }

// func TestClient_companyProfile2(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.CompanyProfile(jcbCompanyNumber)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	// l := len(res.Items)
// 	// if l != 5 {
// 	// 	t.Error("Expected 5 results, got", l)
// 	// }

// 	fmt.Printf("%+v\n", res)
// }

// func TestClient_companyRegisteredOfficeAddress(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.CompanyRegisteredOfficeAddress(icelandCompanyNumber)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Printf("%+v\n", res)
// }

// func TestClient_companyOfficers1(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.CompanyOfficers(icelandCompanyNumber, "", "appointed_on", -1, -1)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Printf("%+v\n", res)
// }

// // func TestClient_companyOfficers2(t *testing.T) {

// // 	// 404

// // 	c := chapi.NewClient(apiKey)
// // 	res, err := c.CompanyOfficers(icelandCompanyNumber, "directors", "", -1, -1)
// // 	if err != nil {
// // 		t.Error(err)
// // 	}

// // 	fmt.Printf("%+v\n", res)
// // }

// // func TestClient_companyOfficers3(t *testing.T) {

// // 	// 404

// // 	c := chapi.NewClient(apiKey)
// // 	res, err := c.CompanyOfficers(icelandCompanyNumber, "secretaries", "", -1, -1)
// // 	if err != nil {
// // 		t.Error(err)
// // 	}

// // 	fmt.Printf("%+v\n", res)
// // }

// // func TestClient_companyOfficers4(t *testing.T) {

// // 	// 404

// // 	c := chapi.NewClient(apiKey)
// // 	res, err := c.CompanyOfficers(icelandCompanyNumber, "llp_members", "", -1, -1)
// // 	if err != nil {
// // 		t.Error(err)
// // 	}

// // 	fmt.Printf("%+v\n", res)
// // }

// func TestClient_companyOfficers5(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.CompanyOfficers(jcbCompanyNumber, "", "appointed_on", -1, -1)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Printf("%+v\n", res)
// }

// func TestClient_companyRegisters(t *testing.T) {

// 	// 404

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.CompanyRegisters(icelandCompanyNumber)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Printf("%+v\n", res)
// }

// func TestClient_officerAppointments1(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.OfficerAppointments(richardBransonOfficerID, 5, -1)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Printf("%+v\n", res)
// }

func TestClient_officerAppointments2(t *testing.T) {

	c := chapi.NewClient(apiKey)
	res, err := c.OfficerAppointments(lordBamfordOfficerID, -1, -1)
	if err != nil {
		t.Error(err)
	}

	// fmt.Printf("%+v\n", res)

	for _, appt := range res.Items {
		fmt.Printf("(%s) %s - %s\n", appt.AppointedTo.CompanyNumber, appt.AppointedTo.CompanyName, appt.Occupation)
		// fmt.Printf("(%s) %s - %s\n", appt.AppointedTo.CompanyNumber, appt.AppointedTo.CompanyName, appt.OfficerRole)
		// fmt.Printf("(%s) %s - %s %s\n", appt.AppointedTo.CompanyNumber, appt.AppointedTo.CompanyName, appt.Occupation, appt.OfficerRole)
	}

}

// func TestClient_officerNaturalDisqualifications(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.OfficerNaturalDisqualifications(richardBransonOfficerID)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Printf("%+v\n", res)
// }

// func TestClient_officerCorporateDisqualifications(t *testing.T) {

// 	c := chapi.NewClient(apiKey)
// 	res, err := c.OfficerCorporateDisqualifications(richardBransonOfficerID)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	fmt.Printf("%+v\n", res)
// }

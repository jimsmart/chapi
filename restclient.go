package chapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/jimsmart/chapi/ch"
)

var APIKey = ""

// BaseURL of Companies House REST API.
var BaseURL = "https://api.companieshouse.gov.uk"

// DefaultHTTPClient for requests.
//
// To use an http.Client with different settings, either override this global variable,
// or instantiate a Client (or RESTClient) instance with the appropriate local configuration.
var DefaultHTTPClient = &http.Client{}

// RESTClient provides low-level access to the Companies House API, with all methods returning
// the body of the HTTP response in []byte format (raw JSON).
//
// To work with this same API but with unmarshalled JSON structs instead, see Client.
//
// In all methods: should an otherwise successful HTTP request return a response containing
// a status that is not 2xx or 3xx, an error value of type RESTStatusError will be returned.
type RESTClient struct {
	APIKey     string
	HTTPClient *http.Client
}

// RESTStatusError is returned as an error value if the http.Response
// has a StatusCode that is not 2xx or 3xx.
//
// If possible, any content found in the response body is unmarshalled into ErrorResource.
type RESTStatusError struct {
	Status        string
	StatusCode    int
	Body          []byte
	ErrorResource *ch.ErrorResource
}

func (e *RESTStatusError) Error() string {
	msg := e.Status
	if len(e.Body) > 2 {
		msg += fmt.Sprintf(" %s %+v", string(e.Body), e.ErrorResource)
	}
	return msg
}

// newRestStatusError creates a new instance of RESTStatusError,
// and attempts to unmarshall the body into an ErrorResource.
func newRESTStatusError(status string, statusCode int, body []byte) *RESTStatusError {
	rerr := &RESTStatusError{
		Status:     status,
		StatusCode: statusCode,
		Body:       body,
	}
	if len(body) == 0 {
		return rerr
	}
	var r ch.ErrorResource
	err := json.Unmarshal(body, &r)
	if err != nil {
		log.Printf("json.Unmarshal error %s for %s", err, body)
		// panic(err)
		return rerr
	}
	rerr.ErrorResource = &r
	return rerr
}

//

func (c *RESTClient) Search(q string, limit, offset int) ([]byte, error) {
	params := searchParams(q, limit, offset)
	return c.request("/search", params)
}

func (c *RESTClient) SearchCompanies(q string, limit, offset int) ([]byte, error) {
	params := searchParams(q, limit, offset)
	return c.request("/search/companies", params)
}

func (c *RESTClient) SearchOfficers(q string, limit, offset int) ([]byte, error) {
	params := searchParams(q, limit, offset)
	return c.request("/search/officers", params)
}

func (c *RESTClient) SearchDisqualifiedOfficers(q string, limit, offset int) ([]byte, error) {
	params := searchParams(q, limit, offset)
	return c.request("/search/disqualified-officers", params)
}

func (c *RESTClient) CompanyProfile(companyNumber string) ([]byte, error) {
	return c.request("/company/"+companyNumber, nil)
}

func (c *RESTClient) CompanyRegisteredOfficeAddress(companyNumber string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/registered-office-address", nil)
}

func (c *RESTClient) CompanyOfficers(companyNumber, registerType, sortOrder string, limit, offset int) ([]byte, error) {
	params := listParams(limit, offset)
	if len(registerType) != 0 {
		params.Add("register_view", "true")
		params.Add("register_type", registerType)
	}
	if len(sortOrder) != 0 {
		params.Add("order_by", sortOrder)
	}
	return c.request("/company/"+companyNumber+"/officers", params)
}

func (c *RESTClient) CompanyFilingHistory(companyNumber, categoryFilter string, limit, offset int) ([]byte, error) {
	params := listParams(limit, offset)
	if len(categoryFilter) != 0 {
		params.Add("category", categoryFilter)
	}
	return c.request("/company/"+companyNumber+"/filing-history", params)
}

func (c *RESTClient) CompanyFilingHistoryTransaction(companyNumber, transactionID string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/filing-history/"+transactionID, nil)
}

func (c *RESTClient) CompanyInsolvency(companyNumber string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/insolvency", nil)
}

func (c *RESTClient) CompanyCharges(companyNumber string, limit, offset int) ([]byte, error) {
	params := listParams(limit, offset)
	return c.request("/company/"+companyNumber+"/charges", params)
}

func (c *RESTClient) CompanyCharge(companyNumber, chargeID string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/charges/"+chargeID, nil)
}

func (c *RESTClient) OfficerAppointments(officerID string, limit, offset int) ([]byte, error) {
	params := listParams(limit, offset)
	return c.request("/officers/"+officerID+"/appointments", params)
}

func (c *RESTClient) OfficerNaturalDisqualifications(officerID string) ([]byte, error) {
	return c.request("/disqualified-officers/natural/"+officerID, nil)
}

func (c *RESTClient) OfficerCorporateDisqualifications(officerID string) ([]byte, error) {
	return c.request("/disqualified-officers/corporate/"+officerID, nil)
}

func (c *RESTClient) CompanyUKEstablishments(companyNumber string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/uk-establishments", nil)
}

func (c *RESTClient) PSCs(companyNumber string, registerView bool, limit, offset int) ([]byte, error) {
	params := listParams(limit, offset)
	if registerView {
		params.Add("register_view", "true")
	}
	return c.request("/company/"+companyNumber+"/persons-with-significant-control", params)
}

func (c *RESTClient) PSCIndividual(companyNumber, pscID string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/persons-with-significant-control/individual/"+pscID, nil)
}

func (c *RESTClient) PSCCorporateEntity(companyNumber, pscID string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/persons-with-significant-control/corporate-entity/"+pscID, nil)
}

func (c *RESTClient) PSCLegalPerson(companyNumber, pscID string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/persons-with-significant-control/legal-person/"+pscID, nil)
}

func (c *RESTClient) PSCStatements(companyNumber string, registerView bool, limit, offset int) ([]byte, error) {
	params := listParams(limit, offset)
	if registerView {
		params.Add("register_view", "true")
	}
	return c.request("/company/"+companyNumber+"/persons-with-significant-control-statements", params)
}

func (c *RESTClient) PSCStatement(companyNumber, statementID string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/persons-with-significant-control-statements/"+statementID, nil)
}

func (c *RESTClient) PSCSuperSecure(companyNumber, superSecureID string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/persons-with-significant-control/super-secure/"+superSecureID, nil)
}

func (c *RESTClient) CompanyRegisters(companyNumber string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/registers", nil)
}

func (c *RESTClient) CompanyExemptions(companyNumber string) ([]byte, error) {
	return c.request("/company/"+companyNumber+"/exemptions", nil)
}

func (c *RESTClient) DocumentMetadata(documentID string) ([]byte, error) {
	return c.request("/document/"+documentID, nil)
}

func (c *RESTClient) DocumentContent(documentID string) ([]byte, error) {
	// TODO(js) Does DocumentContent need a param for setting Accept content-type header?
	return c.request("/document/"+documentID+"/content", nil)
}

//

func (c *RESTClient) request(urlSlug string, params url.Values) ([]byte, error) {

	apiKey := c.APIKey
	if len(apiKey) == 0 {
		apiKey = APIKey
		if len(apiKey) == 0 {
			panic("missing APIKey")
		}
	}

	req, err := http.NewRequest("GET", BaseURL+urlSlug, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(apiKey, "")

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = DefaultHTTPClient
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if !success(res) {
		return nil, newRESTStatusError(res.Status, res.StatusCode, body)
	}

	return body, nil
}

// success returns true if Response.Status is 2xx or 3xx.
func success(r *http.Response) bool {
	s := r.Status[:1]
	return s == "2" || s == "3"
}

//

func listParams(limit, offset int) url.Values {
	params := make(url.Values)
	if limit != -1 {
		params.Add("items_per_page", strconv.Itoa(limit))
	}
	if offset != -1 {
		params.Add("start_index", strconv.Itoa(offset))
	}
	return params
}

func searchParams(q string, limit, offset int) url.Values {
	params := listParams(limit, offset)
	params.Add("q", q)
	return params
}

//

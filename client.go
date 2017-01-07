package chapi

import (
	"encoding/json"

	"github.com/jimsmart/chapi/ch"
)

// TODO(js) Look into decoding/unmarshalling streaming JSON ?

// Client provides higher-level API to access to the Companies House API,
// with all methods returning unmarshalled JSON structs.
//
// To work with this same API but with raw JSON bytes, see RESTClient.
type Client struct {
	*RESTClient
}

// NewClient creats a new instance of Client.
//
// By default, it will use the API key provided by the APIKey package variable
// and the HTTPClient provided by the DefaultHTTPClient package variable.
func NewClient() *Client {
	return &Client{RESTClient: &RESTClient{}}
}

// NewClientWithKey creats a new instance of Client, configured to use the given API key
// and the HTTPClient provided by the DefaultHTTPClient package variable.
func NewClientWithKey(apiKey string) *Client {
	return &Client{
		RESTClient: &RESTClient{
			APIKey: apiKey,
		},
	}
}

// Search performs a query across all Companies House indexed information.
// To search against specific resource types,
// see SearchCompanies, SearchOfficers or SearchDisqualifiedOfficers.
//
// Where q is the term being searched for,
// limit is the number of search results to return 'per page',
// and offset is the index of the first item to return.
//
// Passing a value of -1 for limit or offset will cause the
// parameter to be ignored and not sent to the server.
//
// See https://developer.companieshouse.gov.uk/api/docs/search/search.html
func (c *Client) Search(q string, limit, offset int) (*ch.SearchResource, error) {
	data, err := c.RESTClient.Search(q, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.SearchResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// SearchCompanies performs a search of company information.
//
// Where q is the term being searched for,
// limit is the number of search results to return 'per page',
// and offset is the index of the first item to return.
//
// Passing a value of -1 for limit or offset will cause the
// parameter to be ignored and not sent to the server.
//
// See https://developer.companieshouse.gov.uk/api/docs/search/companies/companysearch.html
func (c *Client) SearchCompanies(q string, limit, offset int) (*ch.CompanySearchResource, error) {
	data, err := c.RESTClient.SearchCompanies(q, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.CompanySearchResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// SearchOfficers performs a search of officer information.
//
// Where q is the term being searched for,
// limit is the number of search results to return 'per page',
// and offset is the index of the first item to return.
//
// Passing a value of -1 for limit or offset will cause the
// parameter to be ignored and not sent to the server.
//
// See https://developer.companieshouse.gov.uk/api/docs/search/officers/officersearch.html
func (c *Client) SearchOfficers(q string, limit, offset int) (*ch.OfficerSearchResource, error) {
	data, err := c.RESTClient.SearchOfficers(q, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.OfficerSearchResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// SearchDisqualifiedOfficers performs a search of disqualified officer information.
//
// Where q is the term being searched for,
// limit is the number of search results to return 'per page',
// and offset is the index of the first item to return.
//
// Passing a value of -1 for limit or offset will cause the
// parameter to be ignored and not sent to the server.
//
// See https://developer.companieshouse.gov.uk/api/docs/search/disqualified-officers/disqualifiedofficersearch.html
func (c *Client) SearchDisqualifiedOfficers(q string, limit, offset int) (*ch.DisqualifiedOfficerSearchResource, error) {
	data, err := c.RESTClient.SearchDisqualifiedOfficers(q, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.DisqualifiedOfficerSearchResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// CompanyProfile gets the basic company information.
//
// See https://developer.companieshouse.gov.uk/api/docs/company/company_number/readCompanyProfile.html
func (c *Client) CompanyProfile(companyNumber string) (*ch.CompanyProfileResource, error) {
	data, err := c.RESTClient.CompanyProfile(companyNumber)
	if err != nil {
		return nil, err
	}
	var r ch.CompanyProfileResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// CompanyRegisteredOfficeAddress gets the current address of a company.
//
// See https://developer.companieshouse.gov.uk/api/docs/company/company_number/registered-office-address/readRegisteredOfficeAddress.html
func (c *Client) CompanyRegisteredOfficeAddress(companyNumber string) (*ch.RegisteredOfficeAddressResource, error) {
	data, err := c.RESTClient.CompanyRegisteredOfficeAddress(companyNumber)
	if err != nil {
		return nil, err
	}
	var r ch.RegisteredOfficeAddressResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// type CORegisterType string
// type COSortOrder string

// const (
// 	Directors   CORegisterType = "directors"
// 	Secretaries CORegisterType = "secretaries"
// 	LLPMembers  CORegisterType = "llp-members"

// 	AppointedOn     COSortOrder = "appointed_on"
// 	AppointedOnDesc COSortOrder = "-appointed_on"
// 	ResignedOn      COSortOrder = "resigned_on"
// 	ResignedOnDesc  COSortOrder = "-resigned_on"
// 	Surname         COSortOrder = "surname"
// 	SurnameDesc     COSortOrder = "-surname"
// )

// register_type string (optional) "", "directors", "secretaries", "llp-members"
// register_view bool (optional) set to true if register_type != ""
// order_by string (optional) "", "appointed_on", "resigned_on", "surname" - can also have "-" prefix.

// TODO(js) Documentation.

// TODO actual string for LLP Members register_type is "llp_members" (2016-10-02)

// CompanyOfficers lists the company officers.
//
// See https://developer.companieshouse.gov.uk/api/docs/company/company_number/officers/officerList.html
func (c *Client) CompanyOfficers(companyNumber, registerType, sortOrder string, limit, offset int) (*ch.OfficerListResource, error) {
	data, err := c.RESTClient.CompanyOfficers(companyNumber, registerType, sortOrder, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.OfficerListResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// CompanyFilingHistory gets the filing history of the company.
//
// See https://developer.companieshouse.gov.uk/api/docs/company/company_number/filing-history/getFilingHistoryList.html
func (c *Client) CompanyFilingHistory(companyNumber, categoryFilter string, limit, offset int) (*ch.FilingHistoryListResource, error) {
	data, err := c.RESTClient.CompanyFilingHistory(companyNumber, categoryFilter, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.FilingHistoryListResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// CompanyFilingHistoryTransaction gets a single item from the filing history of the company.
//
// See https://developer.companieshouse.gov.uk/api/docs/company/company_number/filing-history/transaction_id/getFilingHistoryItem.html
func (c *Client) CompanyFilingHistoryTransaction(companyNumber, transactionID string) (*ch.FilingHistoryItemResource, error) {
	data, err := c.RESTClient.CompanyFilingHistoryTransaction(companyNumber, transactionID)
	if err != nil {
		return nil, err
	}
	var r ch.FilingHistoryItemResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

// CompanyInsolvency gets company insolvency information.
//
// See https://developer.companieshouse.gov.uk/api/docs/company/company_number/insolvency/readCompanyInsolvency.html
func (c *Client) CompanyInsolvency(companyNumber string) (*ch.CompanyInsolvencyResource, error) {
	data, err := c.RESTClient.CompanyInsolvency(companyNumber)
	if err != nil {
		return nil, err
	}
	var r ch.CompanyInsolvencyResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) CompanyCharges(companyNumber string, limit, offset int) (*ch.ChargeListResource, error) {
	data, err := c.RESTClient.CompanyCharges(companyNumber, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.ChargeListResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) CompanyCharge(companyNumber, chargeID string) (*ch.ChargeDetailsResource, error) {
	data, err := c.RESTClient.CompanyCharge(companyNumber, chargeID)
	if err != nil {
		return nil, err
	}
	var r ch.ChargeDetailsResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) OfficerAppointments(officerID string, limit, offset int) (*ch.AppointmentListResource, error) {
	data, err := c.RESTClient.OfficerAppointments(officerID, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.AppointmentListResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) OfficerNaturalDisqualifications(officerID string) (*ch.NaturalDisqualificationResource, error) {
	data, err := c.RESTClient.OfficerNaturalDisqualifications(officerID)
	if err != nil {
		return nil, err
	}
	var r ch.NaturalDisqualificationResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) OfficerCorporateDisqualifications(officerID string) (*ch.CorporateDisqualificationResource, error) {
	data, err := c.RESTClient.OfficerCorporateDisqualifications(officerID)
	if err != nil {
		return nil, err
	}
	var r ch.CorporateDisqualificationResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) CompanyUKEstablishments(companyNumber string) (*ch.CompanyUKEstablishmentsResource, error) {
	data, err := c.RESTClient.CompanyUKEstablishments(companyNumber)
	if err != nil {
		return nil, err
	}
	var r ch.CompanyUKEstablishmentsResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) PSCs(companyNumber string, registerView bool, limit, offset int) (*ch.PSCListResource, error) {
	data, err := c.RESTClient.PSCs(companyNumber, registerView, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.PSCListResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) PSCIndividual(companyNumber, pscID string) (*ch.PSCIndividualResource, error) {
	data, err := c.RESTClient.PSCIndividual(companyNumber, pscID)
	if err != nil {
		return nil, err
	}
	var r ch.PSCIndividualResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) PSCCorporateEntity(companyNumber, pscID string) (*ch.PSCCorporateEntityResource, error) {
	data, err := c.RESTClient.PSCCorporateEntity(companyNumber, pscID)
	if err != nil {
		return nil, err
	}
	var r ch.PSCCorporateEntityResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) PSCLegalPerson(companyNumber, pscID string) (*ch.PSCLegalPersonResource, error) {
	data, err := c.RESTClient.PSCLegalPerson(companyNumber, pscID)
	if err != nil {
		return nil, err
	}
	var r ch.PSCLegalPersonResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) PSCStatements(companyNumber string, registerView bool, limit, offset int) (*ch.PSCStatementListResource, error) {
	data, err := c.RESTClient.PSCStatements(companyNumber, registerView, limit, offset)
	if err != nil {
		return nil, err
	}
	var r ch.PSCStatementListResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) PSCStatement(companyNumber, statementID string) (*ch.PSCStatementResource, error) {
	data, err := c.RESTClient.PSCStatement(companyNumber, statementID)
	if err != nil {
		return nil, err
	}
	var r ch.PSCStatementResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) PSCSuperSecure(companyNumber, superSecureID string) (*ch.PSCSuperSecureResource, error) {
	data, err := c.RESTClient.PSCSuperSecure(companyNumber, superSecureID)
	if err != nil {
		return nil, err
	}
	var r ch.PSCSuperSecureResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) CompanyRegisters(companyNumber string) (*ch.CompanyRegisterResource, error) {
	data, err := c.RESTClient.CompanyRegisters(companyNumber)
	if err != nil {
		return nil, err
	}
	var r ch.CompanyRegisterResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) CompanyExemptions(companyNumber string) (*ch.CompanyExemptionsResource, error) {
	data, err := c.RESTClient.CompanyExemptions(companyNumber)
	if err != nil {
		return nil, err
	}
	var r ch.CompanyExemptionsResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) DocumentMetadata(documentID string) (*ch.DocumentMetadataResource, error) {
	data, err := c.RESTClient.DocumentMetadata(documentID)
	if err != nil {
		return nil, err
	}
	var r ch.DocumentMetadataResource
	err = json.Unmarshal(data, &r)
	return &r, err
}

func (c *Client) DocumentContent(documentID string) ([]byte, error) {
	return c.RESTClient.DocumentContent(documentID)
}

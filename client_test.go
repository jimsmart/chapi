package chapi_test

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jimsmart/chapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("Client", func() {

	apiKey := os.Getenv("COMPANIES_HOUSE_API_KEY")
	It("needs an API key", func() {
		Expect(apiKey).ToNot(Equal(""))
	})

	Context("a new client with a specified API key", func() {

		ch := chapi.NewClientWithKey(apiKey)

		It("should have a non-nil RESTClient", func() {
			Expect(ch.RESTClient).ToNot(BeNil())
		})
		It("should have an API key in its RESTClient", func() {
			Expect(ch.RESTClient.APIKey).ToNot(Equal(""))
		})
		It("should be able to Search() without error", func() {
			_, err := ch.Search("Richard Branson", 1, -1)
			Expect(err).To(BeNil())
		})
	})

	Context("a new client using the default package-level APIKey", func() {

		chapi.APIKey = apiKey
		ch := chapi.NewClient()

		It("should have a non-nil RESTClient", func() {
			Expect(ch.RESTClient).ToNot(BeNil())
		})
		It("should be able to Search() without error", func() {
			_, err := ch.Search("Richard Branson", 1, -1)
			Expect(err).To(BeNil())
		})
	})

	ch := chapi.NewClientWithKey(apiKey)

	Context("when calling Search(), asking for 10 results", func() {

		res, err := ch.Search("Richard Branson", 10, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return 10 results", func() {
			Expect(res.ItemsPerPage).To(Equal(10))
			Expect(res.StartIndex).To(Equal(0))
		})
		It("should return an expected result", func() {
			Expect(res.Items).To(ContainElement(MatchFields(IgnoreExtras, Fields{
				"Kind":  Equal("searchresults#officer"),
				"Title": Equal("Sir Richard Charles Nicholas BRANSON"),
				"Links": MatchAllFields(Fields{
					"Self": Not(Equal("")),
				}),
			})))
		})
		// TODO(js) We seem to be missing the ID extractors ...?
	})

	Context("when calling SearchCompanies(), asking for 5 results", func() {

		res, err := ch.SearchCompanies("Facebook UK Ltd", 5, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return 5 results", func() {
			Expect(res.ItemsPerPage).To(Equal(5))
			Expect(res.StartIndex).To(Equal(0))
		})
		It("should return an expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"Kind": Equal("search#companies"),
				"Items": ContainElement(MatchFields(IgnoreExtras, Fields{
					"Title":         Equal("FACEBOOK UK LTD"),
					"CompanyNumber": Equal("06331310"),
					"Links": MatchAllFields(Fields{
						"Self": Not(Equal("")),
					}),
				})),
			}))
		})
		// TODO(js) We seem to be missing the ID extractors ...?
	})

	Context("when calling SearchOfficers(), asking for 10 results", func() {

		res, err := ch.SearchOfficers("Richard Branson", 10, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return 10 results", func() {
			Expect(res.ItemsPerPage).To(Equal(10))
			Expect(res.StartIndex).To(Equal(0))
		})
		It("should return an expected result", func() {
			Expect(res.Items).To(ContainElement(MatchFields(IgnoreExtras, Fields{
				"Kind":  Equal("searchresults#officer"),
				"Title": Equal("Sir Richard Charles Nicholas BRANSON"),
				"DateOfBirth": MatchAllFields(Fields{
					"Month": Equal(7),
					"Year":  Equal(1950),
				}),
				"Links": MatchAllFields(Fields{
					"Self": Not(Equal("")),
				}),
			})))
		})

		// TODO(js) We seem to be missing the ID extractors ...?
	})

	Context("when calling SearchDisqualifiedOfficers(), asking for 10 results", func() {

		res, err := ch.SearchDisqualifiedOfficers("John Smith", 10, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return 10 results", func() {
			Expect(res.ItemsPerPage).To(Equal(10))
			Expect(res.StartIndex).To(Equal(0))
		})
		It("should return some results", func() {
			Expect(res.Items).To(ContainElement(MatchFields(IgnoreExtras, Fields{
				"Kind": Equal("searchresults#disqualified-officer"),
			})))
		})

		// TODO(js) We seem to be missing the ID extractors ...?
	})

	Context("when calling CompanyProfile()", func() {

		res, err := ch.CompanyProfile("06331310")

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return the expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"CompanyName":   Equal("FACEBOOK UK LTD"),
				"CompanyNumber": Equal("06331310"),
				"Links": MatchFields(IgnoreExtras, Fields{
					"Self":          Not(Equal("")),
					"FilingHistory": Not(Equal("")),
					"Officers":      Not(Equal("")),
				}),
			}))
		})
	})

	Context("when calling CompanyRegisteredOfficeAddress()", func() {

		res, err := ch.CompanyRegisteredOfficeAddress("06331310")

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return the expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"AddressLine1": Equal("10 Brock Street"),
				"AddressLine2": Equal("Regent's Place"),
				"Locality":     Equal("London"),
				"PostalCode":   Equal("NW1 3FG"),
				"Country":      Equal("England"),
			}))
		})
	})

	Context("when calling CompanyOfficers()", func() {

		res, err := ch.CompanyOfficers("02627406", "", "", -1, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return the expected result", func() {

			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"Links": MatchAllFields(Fields{
					"Self": Not(Equal("")),
				}),
				"Items": ContainElement(MatchFields(IgnoreExtras, Fields{
					"Name": Equal("DYSON, James, Sir"),
					"DateOfBirth": MatchFields(IgnoreExtras, Fields{
						"Month": Equal(5),
						"Year":  Equal(1947),
					}),
					"Nationality": Equal("British"),
					"Occupation":  Equal("Designer"),
					"OfficerRole": Equal("director"),
					"ResignedOn":  Equal("2010-06-18"),
					"Links": MatchAllFields(Fields{
						"Officer": MatchAllFields(Fields{
							"Appointments": Not(Equal("")),
						}),
					}),
				})),
			}))
		})
	})

	Context("when calling CompanyFilingHistory(), asking for 10 results", func() {

		res, err := ch.CompanyFilingHistory("02627406", "", 10, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return 10 results", func() {
			Expect(res.ItemsPerPage).To(Equal(10))
			Expect(res.StartIndex).To(Equal(0))
		})
		It("should return the expected result", func() {
			Expect(res.Items).To(ContainElement(MatchFields(IgnoreExtras, Fields{
				"Type":          Equal("AA"),
				"Category":      Equal("accounts"),
				"Description":   Equal("accounts-with-accounts-type-full"),
				"PaperFiled":    Equal(true),
				"Date":          Not(Equal("")),
				"TransactionID": Not(Equal("")),
				"Links": MatchAllFields(Fields{
					"Self":             Not(Equal("")),
					"DocumentMetadata": Not(Equal("")),
				}),
			})))
		})
	})

	Context("when calling CompanyFilingHistoryTransaction()", func() {

		res, err := ch.CompanyFilingHistoryTransaction("02627406", "MzA4MTM5MTMwMWFkaXF6a2N4")

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return the expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"Type":        Equal("AR01"),
				"Category":    Equal("annual-return"),
				"Description": Equal("annual-return-company-with-made-up-date-full-list-shareholders"),
				"Date":        Not(Equal("")),
				"AssociatedFilings": ContainElement(MatchFields(IgnoreExtras, Fields{
					"Date":        Not(Equal("")),
					"Description": Not(Equal("")),
					"Type":        Not(Equal("")),
				})),
				"Links": MatchAllFields(Fields{
					"Self":             Not(Equal("")),
					"DocumentMetadata": Not(Equal("")),
				}),
			}))
		})
	})

	Context("when calling CompanyInsolvency()", func() {

		// TODO(js) I am unable to find a Company Number for a company that has insolvency cases.

		_, err := ch.CompanyInsolvency("NF001705")

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		// It("should return the expected result", func() {
		// 	Expect(*res).To(MatchFields(IgnoreExtras, Fields{
		// 		//
		// 	}))
		// })
	})

	Context("when calling CompanyCharges()", func() {

		res, err := ch.CompanyCharges("NF001705", -1, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return an expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"TotalCount":      Equal(2),
				"UnfilteredCount": Equal(2),
				"Items": ContainElement(MatchFields(IgnoreExtras, Fields{
					"CreatedOn":   Not(Equal("")),
					"DeliveredOn": Not(Equal("")),
					"PersonsEntitled": ContainElement(MatchAllFields(Fields{
						"Name": Equal("Barclays Bank PLC"),
					})),
					"Classification": MatchAllFields(Fields{
						"Description": Not(Equal("")),
						"Type":        Not(Equal("")),
					}),
					"Links": MatchAllFields(Fields{
						"Self": Not(Equal("")),
					}),
				})),
			}))
		})
	})

	Context("when calling CompanyCharge()", func() {

		res, err := ch.CompanyCharge("NF001705", "eZSEl_fk_3LqlhAHRHEFH69egpE")

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return an expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"CreatedOn":   Not(Equal("")),
				"DeliveredOn": Not(Equal("")),
				"PersonsEntitled": ContainElement(MatchAllFields(Fields{
					"Name": Equal("Barclays Bank PLC"),
				})),
				"Classification": MatchAllFields(Fields{
					"Description": Not(Equal("")),
					"Type":        Not(Equal("")),
				}),
				"Links": MatchAllFields(Fields{
					"Self": Not(Equal("")),
				}),
			}))
		})
	})

	Context("when calling OfficerAppointments(), asking for 10 results", func() {

		res, err := ch.OfficerAppointments("fPsul1-gLgzfRlgRvGBL14iNV3c", -1, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return an expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"Kind": Equal("personal-appointment"),
				"Name": Equal("Richard Charles Nicholas BRANSON"),
				"DateOfBirth": MatchAllFields(Fields{
					"Month": Equal(7),
					"Year":  Equal(1950),
				}),
				"Items": ContainElement(MatchFields(IgnoreExtras, Fields{
					"Occupation": Equal("Company Director"),
					"AppointedTo": MatchFields(IgnoreExtras, Fields{
						"CompanyName": ContainSubstring("VIRGIN"),
					}),
				})),
			}))
		})

		// TODO(js) We seem to be missing the ID extractors ...?
	})

	//

	// TODO(js) This test returns 404 - Is that simply because Branson has no disqualifications? Look into this - maybe ask CH for some test data?

	// Context("when calling OfficerNaturalDisqualifications()", func() {

	// 	res, err := ch.OfficerNaturalDisqualifications("fPsul1-gLgzfRlgRvGBL14iNV3c")

	// 	It("should not return an error", func() {
	// 		Expect(err).To(BeNil())
	// 	})

	// 	printJSON(res)

	// 	// It("should return an expected result", func() {
	// 	// 	Expect(*res).To(MatchFields(IgnoreExtras, Fields{
	// 	// 	// ...
	// 	// 	}))
	// 	// })
	// })

	// -- Same problem as above :(

	// Context("when calling OfficerCorporateDisqualifications()", func() {

	// 	res, err := ch.OfficerCorporateDisqualifications("fPsul1-gLgzfRlgRvGBL14iNV3c")

	// 	It("should not return an error", func() {
	// 		Expect(err).To(BeNil())
	// 	})

	// 	printJSON(res)

	// 	// It("should return an expected result", func() {
	// 	// 	Expect(*res).To(MatchFields(IgnoreExtras, Fields{
	// 	// 	// ...
	// 	// 	}))
	// 	// })
	// })

	Context("when calling CompanyUKEstablishments()", func() {

		res, err := ch.CompanyUKEstablishments("02627406")

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return an expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				// "Ttems":
				"Kind": Equal("related-companies"),
				"Links": MatchAllFields(Fields{
					"Self": Not(Equal("")),
				}),
			}))
		})
	})

	Context("when calling PSCs()", func() {

		res, err := ch.PSCs("02627406", false, -1, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return an expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"TotalResults": Not(Equal(0)),
				// "Kind": Not(Equal("")),
				"Links": MatchFields(IgnoreExtras, Fields{
					"Self": Not(Equal("")),
				}),
				"Items": ContainElement(MatchFields(IgnoreExtras, Fields{
					"Name": Equal("Dyson James Limited"),
					"Address": MatchFields(IgnoreExtras, Fields{
						"AddressLine1": Not(Equal("")),
						"Locality":     Not(Equal("")),
						"PostalCode":   Not(Equal("")),
						"Country":      Equal("England"),
					}),
					"NaturesOfControl": ContainElement("ownership-of-shares-75-to-100-percent"),
					"Links": MatchFields(IgnoreExtras, Fields{
						"Self": Not(Equal("")),
					}),
				})),
			}))
		})
	})

	// TODO(js) Need to find a PSC individual.

	// Context("when calling PSCIndividual()", func() {

	// 	res, err := ch.PSCIndividual("02627406", "")

	// 	It("should not return an error", func() {
	// 		Expect(err).To(BeNil())
	// 	})
	// 	It("should return an expected result", func() {
	// 		Expect(*res).To(MatchFields(IgnoreExtras, Fields{
	// 		}))
	// 	})
	// })

	Context("when calling PSCCorporateEntity()", func() {

		res, err := ch.PSCCorporateEntity("02627406", "fxqbqgRvv3y8SuMWArNKkfezAAw")

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})
		It("should return an expected result", func() {
			Expect(*res).To(MatchFields(IgnoreExtras, Fields{
				"Address": MatchFields(IgnoreExtras, Fields{
					"AddressLine1": Not(Equal("")),
					"Locality":     Not(Equal("")),
					"PostalCode":   Not(Equal("")),
					"Country":      Equal("England"),
				}),
				"Identification": MatchFields(IgnoreExtras, Fields{
					"CountryRegistered":  Not(Equal("")),
					"LegalAuthority":     Not(Equal("")),
					"LegalForm":          Not(Equal("")),
					"PlaceRegistered":    Not(Equal("")),
					"RegistrationNumber": Not(Equal("")),
				}),
				"Kind": Equal("corporate-entity-person-with-significant-control"),
				"Links": MatchFields(IgnoreExtras, Fields{
					"Self": Not(Equal("")),
				}),
				"Name":             Equal("Dyson James Limited"),
				"NaturesOfControl": Not(BeEmpty()),
			}))
		})
	})

	// TODO(js) Need to find a PSC legal person.

	// Context("when calling PSCLegalPerson()", func() {

	// 	res, err := ch.PSCLegalPerson("02627406", "")

	// 	It("should not return an error", func() {
	// 		Expect(err).To(BeNil())
	// 	})
	// 	It("should return an expected result", func() {
	// 		Expect(*res).To(MatchFields(IgnoreExtras, Fields{
	// 		}))
	// 	})
	// })

	// TODO(js) This also gives 404 :(

	// Context("when calling PSCStatements()", func() {

	// 	res, err := ch.PSCStatements("02627406", false, -1, -1)

	// 	It("should not return an error", func() {
	// 		Expect(err).To(BeNil())
	// 	})
	// 	printJSON(res)
	// 	It("should return an expected result", func() {
	// 		Expect(*res).To(MatchFields(IgnoreExtras, Fields{
	// 		//
	// 		}))
	// 	})
	// })

})

func printJSON(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "   ")
	log.Println(string(b))
}

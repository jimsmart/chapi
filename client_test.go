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
	It("needs an API keyto run these tests", func() {
		Expect(apiKey).ToNot(Equal(""))
	})

	Context("with a specified API key", func() {

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

	Context("using the default package-level APIKey", func() {

		ch := chapi.NewClient()

		It("should have a non-nil RESTClient", func() {
			Expect(ch.RESTClient).ToNot(BeNil())
		})
		It("should be able to Search() without error", func() {
			chapi.APIKey = apiKey
			Expect(chapi.APIKey).ToNot(Equal(""))
			_, err := ch.Search("Richard Branson", 1, -1)
			Expect(err).To(BeNil())
			chapi.APIKey = ""
		})
	})

	Context("without an API key, and with no global key", func() {

		ch := chapi.NewClient()

		It("should panic when calling any method", func() {
			shouldPanic := func() {
				chapi.APIKey = ""
				ch.Search("Richard Branson", 1, -1)
			}
			Expect(shouldPanic).To(Panic())
		})
	})

	ch := chapi.NewClientWithKey(apiKey)

	XDescribe("Search()", func() {
		Context("asking for 10 results", func() {

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
	})

	Describe("SearchCompanies()", func() {
		Context("asking for 5 results", func() {

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
	})

	Describe("SearchOfficers()", func() {
		Context("asking for 10 results", func() {

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
	})

	Describe("SearchDisqualifiedOfficers()", func() {
		Context("asking for 10 results", func() {

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
	})

	Describe("CompanyProfile()", func() {
		Context("with a valid company number", func() {

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

		Context("with an invalid company number", func() {
			_, err := ch.CompanyProfile("foo")
			shouldError404(err, "company-profile-not-found")
		})
	})

	Describe("CompanyRegisteredOfficeAddress()", func() {
		Context("with a valid company number", func() {

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
		Context("with an invalid company number", func() {
			_, err := ch.CompanyRegisteredOfficeAddress("foo")
			// TODO(js) This used to return an error message.
			// shouldError404(err, "registered-office-address-not-found")
			shouldError404(err, "")
		})
	})

	Describe("CompanyOfficers()", func() {
		Context("with a valid company number", func() {

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
		Context("with an invalid company number", func() {
			_, err := ch.CompanyOfficers("foo", "", "", -1, -1)
			shouldError404(err, "")
		})
	})

	Describe("CompanyFilingHistory()", func() {
		Context("with a valid company number, asking for 10 results", func() {

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
		// TODO(js) Inconsistent behaviour - this doesn't seem to produce an error.
		// Context("with an invalid company number", func() {
		// 	_, err := ch.CompanyFilingHistory("foo", "", -1, -1)
		// 	shouldError(err)
		// 	shouldError404(err, "foo")
		// })
	})

	Describe("CompanyFilingHistoryTransaction()", func() {
		Context("with a valid company number and transaction id", func() {

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
		Context("with an invalid company number", func() {
			_, err := ch.CompanyFilingHistoryTransaction("foo", "bar")
			shouldError404(err, "filing-history-item-not-found")
		})
	})

	Describe("CompanyInsolvency()", func() {
		Context("with a valid company number - that has no insolvencies", func() {

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
	})

	Describe("CompanyCharges()", func() {
		Context("with a valid company number", func() {

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
		// TODO(js) Inconsistent - doesn't return populated ErrorResource.
		Context("with an invalid company number", func() {
			_, err := ch.CompanyCharges("foo", -1, -1)
			shouldError(err)
			// shouldError404(err, "foo")
		})
	})

	Describe("CompanyCharge()", func() {
		Context("with a valid company number and charge id", func() {

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
		Context("with an invalid company number and charge id", func() {
			_, err := ch.CompanyCharge("foo", "bar")
			shouldError404(err, "company-mortgages-not-found")
		})
	})

	Describe("OfficerAppointments()", func() {
		Context("with a valid officer id, asking for 10 results", func() {

			res, err := ch.OfficerAppointments("fPsul1-gLgzfRlgRvGBL14iNV3c", 10, -1)

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
		// TODO(js) Inconsistent - does not return populated ErrorResource.
		Context("with an invalid officer id", func() {
			_, err := ch.OfficerAppointments("foo", -1, -1)
			shouldError(err)
			// shouldError404(err, "foo")
		})
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

	Describe("CompanyUKEstablishments()", func() {
		Context("with a valid company number", func() {

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
		Context("with an invalid company number", func() {
			_, err := ch.CompanyUKEstablishments("foo")
			shouldError404(err, "company-profile-not-found")
		})
	})

	Describe("PSCs()", func() {
		Context("with a valid company number", func() {

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
							"Country":      Equal("United Kingdom"),
						}),
						"NaturesOfControl": ContainElement("ownership-of-shares-75-to-100-percent"),
						"Links": MatchFields(IgnoreExtras, Fields{
							"Self": Not(Equal("")),
						}),
					})),
				}))
			})
		})
		Context("with an invalid company number", func() {
			_, err := ch.PSCs("foo", false, -1, -1)
			shouldError404(err, "company-psc-not-found")
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

	Describe("PSCCorporateEntity()", func() {
		Context("with a valid company number and corporate entity id", func() {

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
						"Country":      Equal("United Kingdom"),
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
		Context("with an invalid company number and corporate entity id", func() {
			_, err := ch.PSCCorporateEntity("foo", "bar")
			shouldError404(err, "company-psc-details-not-found")
		})
	})

	// TODO(js) Need to find a PSC legal person.

	// Context("when calling PSCLegalPerson()", func() {

	// 	res, err := ch.PSCLegalPerson("02627406", "")

	// 	It("should not return an error", func() {
	// 		Expect(err).To(BeNil())
	// 	})
	// 	printJSON(res)
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

	// TODO(js) This one relies on an id from the previous query :/

	// Context("when calling PSCStatement()", func() {

	// 	res, err := ch.PSCStatement("02627406", "foo")

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

	// TODO(js) This one relies on a superSecureID - I suspect they're hard to get?

	// Context("when calling PSCStatement()", func() {

	// 	res, err := ch.PSCSuperSecure("02627406", "foo")

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

	// TODO(js) 404

	// Context("when calling CompanyRegisters()", func() {

	// 	res, err := ch.CompanyRegisters("NF001705")

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

	// TODO(js) 404

	// Context("when calling CompanyExemptions()", func() {

	// 	res, err := ch.CompanyExemptions("02627406")

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

	// TODO(js) 404 - although it simply times-out from the webpage :(

	// Context("when calling DocumentMetadata()", func() {

	// 	res, err := ch.DocumentMetadata("ged7Kn26fOPGwTr0MACEDypCDiWP8wq2-eNSO1vblk0")

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

func shouldError(err error) {
	It("should return an error", func() {
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).ToNot(Equal(""))
	})
}

func shouldError404(err error, msg string) {
	shouldError(err)
	It("should return a 404 and a valid ErrorResource", func() {
		Expect(err).ToNot(BeNil())
		e, ok := err.(*chapi.RESTStatusError)
		Expect(ok).To(BeTrue())
		Expect(e.StatusCode).To(Equal(404))
		Expect(e.Status).To(ContainSubstring("404"))
		if msg != "" {
			// log.Printf("%s %v", msg, e.ErrorResource)
			Expect(e.ErrorResource).ToNot(BeNil())
			Expect(*e.ErrorResource).To(MatchFields(IgnoreExtras, Fields{
				"Errors": ContainElement(MatchFields(IgnoreExtras, Fields{
					"Error": Equal(msg),
				})),
			}))
		}
	})
}

func printJSON(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "   ")
	log.Println(string(b))
}

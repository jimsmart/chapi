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

const richardBransonID = "fPsul1-gLgzfRlgRvGBL14iNV3c"

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

	Context("when calling OfficerAppointments(), asking for 10 results", func() {

		res, err := ch.OfficerAppointments(richardBransonID, -1, -1)

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

})

func logPrintJSON(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "   ")
	log.Println(string(b))
}

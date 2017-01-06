package chapi_test

import (
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

	ch := chapi.NewClient(apiKey)

	Context("a new client", func() {

		It("should have a non-nil RESTClient", func() {
			Expect(ch.RESTClient).ToNot(BeNil())
		})
		It("should have an API key in its RESTClient", func() {
			Expect(ch.RESTClient.APIKey).ToNot(Equal(""))
		})
	})

	Context("when calling Search()", func() {

		res, err := ch.Search("Richard Branson", -1, -1)

		It("should not return an error", func() {
			Expect(err).To(BeNil())
		})

		It("should return an expected result", func() {
			Expect(res.Items).To(ContainElement(MatchFields(IgnoreExtras, Fields{
				"Kind":  Equal("searchresults#officer"),
				"Title": Equal("Sir Richard Charles Nicholas BRANSON"),
			})))
		})

		// TODO(js) We seem to be missing the ID extractors ...?

		Context("asking for only 10 results", func() {

			res, err := ch.Search("Richard Branson", 10, -1)

			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return a page of 10 results", func() {
				Expect(res.ItemsPerPage).To(Equal(10))
				Expect(res.StartIndex).To(Equal(0))
			})
		})
	})

	Context("when calling OfficerAppointments() with Richard Branson's ID", func() {

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
						"CompanyName": Equal("VIRGIN ATLANTIC LIMITED"),
					}),
				})),
			}))
		})

		// TODO(js) We seem to be missing the ID extractors ...?

		Context("asking for only 10 results", func() {

			res, err := ch.OfficerAppointments(richardBransonID, 10, -1)

			It("should not return an error", func() {
				Expect(err).To(BeNil())
			})

			It("should return a page of 10 results", func() {
				Expect(res.ItemsPerPage).To(Equal(10))
				Expect(res.StartIndex).To(Equal(0))
			})
		})
	})

})

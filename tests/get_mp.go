package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"github.com/burythehammer/golang-api-test-quickstart/api/mpapi"
	"github.com/burythehammer/golang-api-test-quickstart/utils"
	"github.com/burythehammer/golang-api-test-quickstart/model"
)

var _ = Describe("Employee API - Employees", func() {

	apiKey := utils.GetApiKey()

	BeforeEach(func() {
		utils.RegisterErrors()
	})

	AfterEach(func() {
	})

	Context("No MP", func() {
		It("Can't find incorrect MP id", func() {
			statusCode, _, err := mpapi.GetMp(utils.RandomString(10), apiKey)

			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusNotFound), "Status code") // should fail - turns out the api returns 200 if not found... ¯\_(ツ)_/¯
		})

		It("Can't find incorrect MP by postcode", func() {
			statusCode, _, err := mpapi.GetMpByPostcode(utils.RandomString(10), apiKey)

			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusNotFound), "Status code") // turns out the api returns 200 if not found... ¯\_(ツ)_/¯
		})

		It("Can't find incorrect MP by constituency", func() {
			statusCode, _, err := mpapi.GetMpByConstituency(utils.RandomString(10), apiKey)

			Expect(err).To(Not(BeNil()), "Error")
			Expect(statusCode).To(Equal(http.StatusNotFound), "Status code")
		})

		It("Gets a list of MPs", func() {
			statusCode, mpList, err := mpapi.ListMps(apiKey)
			Expect(err).To(BeNil(), "Error")
			Expect(statusCode).To(Equal(http.StatusOK), "Status code")
			Expect(mpList).To(HaveLen(650), "Response body length")
		})

	})

	Context("Single MP", func() {

		mp := model.Mp{PersonId: "25310", FullName: "Helen Hayes", Constituency: "Dulwich and West Norwood", Party: "Labour"}
		mpPostcode := "SE218HY"

		It("Gets MP", func() {
			statusCode, mpResponse, err := mpapi.GetMp(mp.PersonId, apiKey)

			Expect(err).To(BeNil(), "Error")
			Expect(statusCode).To(Equal(http.StatusOK), "Status code")
			Expect(mpResponse.FullName).To(Equal(mp.FullName), "Response body")
		})

		It("Gets an MP by postcode", func() {
			statusCode, mpResponse, err := mpapi.GetMpByPostcode(mpPostcode, apiKey)

			Expect(err).To(BeNil(), "Error")
			Expect(statusCode).To(Equal(http.StatusOK), "Status code")
			Expect(mpResponse.FullName).To(Equal(mp.FullName), "Response body")
		})

		It("Gets an MP by constituency", func() {
			statusCode, mpResponse, err := mpapi.GetMpByConstituency(mp.Constituency, apiKey)

			Expect(err).To(BeNil(), "Error")
			Expect(statusCode).To(Equal(http.StatusOK), "Status code")
			Expect(mpResponse.FullName).To(Equal(mp.FullName), "Response body")
		})

		It("Gets a list of MPs containing the MP", func() {
			statusCode, mpList, err := mpapi.ListMps(apiKey)

			Expect(err).To(BeNil(), "Error")
			Expect(statusCode).To(Equal(http.StatusOK), "Status code")

			mpNames := []string{}

			for i := 0; i < len(mpList); i++ {
				mpNames = append(mpNames, mpList[i].Name)
			}

			Expect(mpNames).To(ContainElement(mp.FullName), "Response body")
		})

	})
})
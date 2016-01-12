package golexa_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golexa"
)

var ()

func init() {
	golexa.AlexaResponseFile = "./output/alexa_marker_responses.json"
}

var _ = Describe("Alexa Framework", func() {
	It("can find the Alexa Responses JSON seed file", func() {
		Expect(golexa.AlexaResponseFile).Should(BeAnExistingFile())
	})

	It("can load the Alexa Responses JSON seed file", func() {
		golexa.SeedAlexaResponses, _ = golexa.LoadAlexaResponseMarkers(golexa.AlexaResponseFile)
		Expect(len(golexa.SeedAlexaResponses)).Should(BeNumerically(">", 0))
	})

	It("can lookup a simple output message", func() {
		tmp := golexa.GetAlexaMarkers([]string{"help"})
		Expect(tmp).Should(Equal("You can check system status, get metrics, or run a job.  Which would you like to do?"))
	})

	It("can parrot back a variable message", func() {
		tmp_err := "This is a sample error message."
		tmp := golexa.GetAlexaMarkers([]string{"error", tmp_err})
		Expect(tmp).Should(Equal(tmp_err))
	})

	It("can substitute multiple variables into the response template", func() {
		tmp := golexa.GetAlexaMarkers([]string{"test_2_variables", "Alexa", "George"})
		Expect(tmp).Should(ContainSubstring("Alexa"))
		Expect(tmp).Should(ContainSubstring("George"))
	})

	// Test the error conditions.
	It("will reply with an error message if it expect parms, but none are passed", func() {
		tmp := golexa.GetAlexaMarkers([]string{})
		Expect(tmp).Should(Equal("You did not specify any items.  You might want to check your code."))
	})
	It("will reply with an error message if a bad key is passed", func() {
		tmp := golexa.GetAlexaMarkers([]string{"idontknow"})
		Expect(tmp).Should(Equal("I could not find the appropriate response in the structure map."))
	})
	It("will reply with an error message the number of arguments passed does not match what was expected", func() {
		tmp := golexa.GetAlexaMarkers([]string{"test_2_variables", "Alexa"})
		Expect(tmp).Should(Equal("There was a mistmatch in requested and assigned parameters for test_2_variables"))
	})
})

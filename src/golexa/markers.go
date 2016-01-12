package golexa

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
)

// Alexa template responses structure.
type AlexaTemplateResponse struct {
	Name      string   `json:"name"`
	Purpose   string   `json:"purpose"`
	Responses []string `json:"responses"`
	IsActive  bool     `json:"is_active"`
}

//
var (
	alexa_text         *regexp.Regexp // Regular expression holding the Alexa template pattern.
	SeedAlexaResponses map[string]AlexaTemplateResponse
	randAlexaResponse  int
	AlexaResponseFile  string
)

func init() {
	alexa_text, _ = regexp.Compile("[%]+([a-z0-9_]+)[%]+")
}

//
func LoadAlexaResponseMarkers(seedFile string) (map[string]AlexaTemplateResponse, error) {
	file, err := os.Open(seedFile)
	seeders := make(map[string]AlexaTemplateResponse)

	// If there is a problem with the file, err on the side of caution and
	// reject the request.
	if err != nil {
		log.Printf("error: Unable to open file/%s", err.Error())
		return seeders, err
	}
	defer file.Close()

	// Decode the json into something we can process.  The JSON is set up to load
	// into a map.  We could also do an array and move it to a map, but why?
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&seeders)
	if err != nil {
		log.Printf("error: Could not decode Config JSON/%s", err.Error())
		return seeders, err
	}

	return seeders, nil
} // func

// GetAlexaMarkers will pull the Alexa tags out of a string of text and returns them in a string.
func GetAlexaMarkers(values []string) string {
	var ret_value string
	var alexa_key AlexaTemplateResponse

	// Did the user pass anything in?
	if len(values) == 0 {
		//log.Printf("error: No items specified")
		return "You did not specify any items.  You might want to check your code."
	}

	// Is the key legitimate?
	if alexa_key = SeedAlexaResponses[values[0]]; alexa_key.Name == "" {
		//log.Printf("error: %s is not a recognized Alexa key", values[0])
		return "I could not find the appropriate response in the structure map."
	}

	// If there is more than one responses, randomize the response from the slice to give the user a more conversational feel.
	// Otherwise, just grab the only response and use it.
	if len(alexa_key.Responses) > 1 {
		randAlexaResponse = rand.Intn(len(alexa_key.Responses) - 1)
	} else {
		randAlexaResponse = 0
	}

	tmp := alexa_text.FindAllStringSubmatch(alexa_key.Responses[randAlexaResponse], -1)
	if len(values)-1 != len(tmp) {
		//log.Printf("error: Mismatched function parameters (%d) and Alexa markers (%d)", len(values)-1, len(tmp))
		return "There was a mistmatch in requested and assigned parameters for " + alexa_key.Name
	}

	// OK, everything looks good, let's match it up.
	i := 1
	ret_value = alexa_key.Responses[randAlexaResponse]
	for _, tag := range tmp {
		ret_value = strings.Replace(ret_value, tag[0], values[i], -1)
		i++
	} // for

	return ret_value
} // func

package alexa

// The primary wrapper of an Alexa request.
type AlexaRequestPacket struct {
	Version string              `json:"version" xml:"version"`
	Session AlexaRequestSession `json:"session" xml:"session"`
	Request json.RawMessage     `json:"request" xml:"request"`
}

// Session attributes for an Alexa request.
type AlexaRequestSession struct {
	New         bool              `json:"new" xml:"new"`
	SessionId   string            `json:"sessionId" xml:"sessionId"`
	Application AlexaApplication  `json:"application" xml:"application"`
	Attributes  map[string]string `json:"attributes" xml:"attributes"`
	User        AlexaUser         `json:"user" xml:"user"`
}

// Alexa launch request that is used to identify a request is a proper Alexa request.
type AlexaLaunchRequest struct {
	Type      string `json:"type" xml:"type"`
	Timestamp string `json:"timestamp" xml:"timestamp"`
	RequestId string `json:"requestId" xml:"requestId"`
}

type AlexaApplication struct {
	ApplicationId string `json:"applicationId" xml:"applicationId"`
}

// Filled out from the Intents defined as part of your Alexa Skill.
type AlexaIntentRequest struct {
	Type      string      `json:"type" xml:"type"`
	Timestamp string      `json:"timestamp" xml:"timestamp"`
	RequestId string      `json:"requestId" xml:"requestId"`
	Intent    AlexaIntent `json:"intent" xml:"intent"`
}

// Used to peek under the hood to see what kind of data is in the RawMessage.
type AlexaRequestStub struct {
	Type string `json:"type" xml:"type"`
}

// Inbound variable data tied to the slot definitions.
type AlexaIntent struct {
	Name  string                       `json:"name" xml:"name"`
	Slots map[string]AlexaIntentString `json:"slots" xml:"slots"`
}

// What the user intends to do.
type AlexaIntentString struct {
	Name  string `json:"name" xml:"name"`
	Value string `json:"value" xml:"value"`
}

type AlexaUser struct {
	UserId string `json:"userId" xml:"userId"`
}

// A SessionEndedRequest is an object that represents a request made to an Alexa skill to notify that a session was ended.
type AlexaSessionEndedRequest struct {
	Type      string `json:"type" xml:"type"`
	Timestamp string `json:"timestamp" xml:"timestamp"`
	RequestId string `json:"requestId" xml:"requestId"`
	Reason    string `json:"reason" xml:"reason"`
}

// Alexa Response Body Format
type AlexaResponseBody struct {
	Version           string            `json:"version" xml:"version"`
	SessionAttributes map[string]string `json:"sessionAttributes,omitempty" xml:"sessionAttributes"`
	Response          AlexaResponse     `json:"response" xml:"response"`
}

// Alexa Response Body Format
type AlexaResponse struct {
	OutputSpeech     AlexaResponseOutputSpeech `json:"outputSpeech" xml:"outputSpeech"`
	Card             AlexaResponseCard         `json:"card,omitempty" xml:"card"`
	Reprompt         AlexaResponseReprompt     `json:"reprompt,omitempty" xml:"reprompt"`
	ShouldEndSession bool                      `json:"shouldEndSession" xml:"shouldEndSession"`
}

// Alexa Response Body Format
type AlexaResponseCard struct {
	Type    string `json:"type" xml:"type"`
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
} //

// Alexa Response Body Format
type AlexaResponseReprompt struct {
	OutputSpeech AlexaResponseOutputSpeech `json:"outputSpeech" xml:"outputSpeech"`
} //

// Alexa Response Body Format
type AlexaResponseOutputSpeech struct {
	Type string `json:"type" xml:"type"`
	Text string `json:"text" xml:"text"`
	SSML string `json:"ssml" xml:"ssml"`
} //

// Alexa template responses structure.
type AlexaTemplateResponse struct {
	Name      string   `json:"name"`
	Purpose   string   `json:"purpose"`
	Responses []string `json:"responses"`
	IsActive  bool     `json:"is_active"`
}

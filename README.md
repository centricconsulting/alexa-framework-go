# Golexa
A simple set of tools for interfacing your Go application with Alexa on an Amazon Echo.  This is still a work in progress and will change as I refine it into the projects I'm working on now.  Feel free to make any changes you want and send me a pull request.

## Getting Started

## Markers
By using markers, you can avoid having to hard-code your responses to whatever Alexa is requesting.  

A marker entry is a simple JSON map that has a name on which it is matched in your code, and an array of potential responses that Golexa will randomly select from in order to make your voice responses sound more conversational and natural.  In this particular example, there is only one response.

    "help":{
        "name":"help",
        "purpose": "General help message.",
        "responses": ["You can check system status, get metrics, or run a job.  Which would you like to do?"],
        "is_active": true
    }

The call to the marker function returns a simple string so the output can be sent straight to Alexa.  The function signature looks like this:

An example of the marker call looks like this:

`func GetAlexaMarkers(values []string) string {}`

Where `values[0]` is the key to the marker you want to match, and `values[1..n]` are the substitution parameters you want to use.  You must pass the same number of parameters to the marker as is being expected otherwise you will generate an error message.  In addition, they should be in the same order as they're expected in the template.  Currently, the names of the template placeholders can be whatever you want (e.g. `%%1%%`, `%%customer_id%%`), but eventually, they may respect the name in order to remove the sequence requirement just mentioned.

`resp.Response.OutputSpeech.Text = golexa.GetAlexaMarkers([]string{"error", tmp_err})`

If you have variables you want Alexa to use, simply add them to the `[]string` slice in the order you want Alexa to say them.  For example, if you have a marker definition that looks like this:

    "test_2_variables":{
        "name":"test_2_variables",
        "purpose": "To make sure two variables can be replaced.",
        "responses": ["My name is %%my_name%% and your name is %%your_name%%.","Nice to meet you %%your_name%%. My name is %%my_name%%."],
        "is_active": true
    }

You would call it like this.

`resp.Response.OutputSpeech.Text = golexa.GetAlexaMarkers([]string{"test_2_variables", "Alexa", "George"})`

### Getting Started with Markers
To use Golexa's Markers, you'll first need to create a marker file as shown in the repo. Then define and load the file like this:

    golexa.AlexaResponseFile = "./output/alexa_marker_responses.json"
    golexa.SeedAlexaResponses, _ = golexa.LoadAlexaResponseMarkers(golexa.AlexaResponseFile)

## Rule Engine
Like the Markers are used for creating Alexa's outbound speech, the Golexa Rules Engine is used for mapping Alexa's inbound requests.

## Structs
A model of the structures needed to interact with Alexa and an Amazon Echo.

## Tests
You can run the tests from the command-line: `go test`

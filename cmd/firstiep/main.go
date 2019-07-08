package main

import (
	"fmt"

	"github.com/AlecAivazis/survey"
	"github.com/jakewarren/firstiep"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name: "tlp",
		Prompt: &survey.Select{
			Message: "Choose the TLP level:",
			Options: []string{"(skip)", "RED", "AMBER", "GREEN", "WHITE"},
			Default: "(skip)",
			Help:    "Recipients are permitted to redistribute the information received within the redistribution scope as defined by the enumerations. The enumerations “RED”, “AMBER”, “GREEN”, “WHITE” are to be interpreted as described in the FIRST Traffic Light Protocol Policy ",
		},
	},
	{
		Name: "EncryptInTransit",
		Prompt: &survey.Select{
			Message: "Encrypt in transit?",
			Options: []string{"(skip)", "MUST", "MAY"},
			Default: "(skip)",
			Help:    "States whether the received information has to be encrypted when it is retransmitted by the recipient.",
		},
	},
	{
		Name: "EncryptAtREST",
		Prompt: &survey.Select{
			Message: "Encrypt at rest?",
			Options: []string{"(skip)", "MUST", "MAY"},
			Default: "(skip)",
			Help:    "States whether the received information has to be encrypted by the Recipient when it is stored at rest.",
		},
	},
	{
		Name: "PermittedActions",
		Prompt: &survey.Select{
			Message: "Permitted actions?",
			Options: []string{"(skip)", "NONE", "CONTACT FOR INSTRUCTION", "INTERNALLY VISIBLE ACTIONS", "EXTERNALLY VISIBLE INDIRECT ACTIONS", "EXTERNALLY VISIBLE DIRECT ACTIONS"},
			Default: "(skip)",
			Help:    "States the permitted actions that Recipients can take upon information received.",
		},
	},
	{
		Name: "AffectedPartyNotifications",
		Prompt: &survey.Select{
			Message: "Notify affected parties?",
			Options: []string{"(skip)", "MAY", "MUST NOT"},
			Default: "(skip)",
			Help:    "Recipients are permitted notify affected third parties of a potential compromise or threat. Examples include permitting National CSIRTs to send notifications to affected constituents, or a service provider contacting affected customers.",
		},
	},
	{
		Name: "Attribution",
		Prompt: &survey.Select{
			Message: "Allowed to attribute the provider?",
			Options: []string{"(skip)", "MAY", "MUST", "MUST NOT"},
			Default: "(skip)",
			Help:    "Recipients could be required to attribute or anonymize the Provider when redistributing the information received.",
		},
	},
	{
		Name: "ObfuscateAffectedParties",
		Prompt: &survey.Select{
			Message: "Obfuscate affected parties?",
			Options: []string{"(skip)", "MAY", "MUST", "MUST NOT"},
			Default: "(skip)",
			Help: `Recipients could be required to obfuscate or anonymize information that could be used to identify the affected parties before redistributing the information received.
Examples include removing affected parties IP addresses, or removing the affected parties names but leaving the affected parties industry vertical prior to sending a notification.`,
		},
	},
	{
		Name: "UnmodifiedResale",
		Prompt: &survey.Select{
			Message: "Allow unmodified resale?",
			Options: []string{"(skip)", "MAY", "MUST NOT"},
			Default: "(skip)",
			Help:    "States whether the recipient MAY or MUST NOT resell the information received unmodified or in a semantically equivalent format. e.g. transposing the information from a .csv file format to a .json file format would be considered semantically equivalent.",
		},
	},
}

func main() {
	i := firstiep.New()

	answers := struct {
		TLP                        string
		EncryptInTransit           string
		EncryptAtREST              string
		PermittedActions           string
		AffectedPartyNotifications string
		Attribution                string
		ObfuscateAffectedParties   string
		UnmodifiedResale           string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if answers.TLP != "(skip)" {
		i.TLP = answers.TLP
	}
	if answers.EncryptInTransit != "(skip)" {
		i.EncryptInTransit = answers.EncryptInTransit
	}
	if answers.EncryptAtREST != "(skip)" {
		i.EncryptAtREST = answers.EncryptAtREST
	}
	if answers.PermittedActions != "(skip)" {
		i.PermittedActions = answers.PermittedActions
	}
	if answers.AffectedPartyNotifications != "(skip)" {
		i.AffectedPartyNotifications = answers.AffectedPartyNotifications
	}
	if answers.Attribution != "(skip)" {
		i.Attribution = answers.Attribution
	}
	if answers.ObfuscateAffectedParties != "(skip)" {
		i.ObfuscateAffectedParties = answers.ObfuscateAffectedParties
	}
	if answers.UnmodifiedResale != "(skip)" {
		i.UnmodifiedResale = answers.UnmodifiedResale
	}

	fmt.Println("\n", i)
}

package firstiep

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/go-validator/validator"
	"github.com/oklog/ulid"
)

var entropy *ulid.MonotonicEntropy

func init() {
	entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
}

func (r *IEP) String() string {
	json, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(json)

}

func New() *IEP {
	var i IEP

	i.ID = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	i.Version = 1.0

	return &i
}

func (r *IEP) Validate() error {

	if errs := validator.Validate(r); errs != nil {
		return errs
	}

	return nil
}

// IEP defines the information exchange policy
// https://www.first.org/iep/
type IEP struct {
	// Provides a unique ID to identify a specific IEP implementation.
	ID string `json:"id" validate:"nonzero"`
	// This statement can be used to provide a name for an IEP implementation.
	Name string `json:"name,omitempty"`
	// States the version of the IEP framework that has been used. e.g. 1.0
	Version int64 `json:"version,omitempty"`
	// This statement can be used to provide a URL reference to the specific IEP implementation.
	Reference string `json:"reference,omitempty"`
	// States the UTC date that the IEP is effective from.
	StartDate string `json:"start-date,omitempty"`
	// States the UTC date that the IEP is effective until.
	EndDate string `json:"end-date,omitempty"`
	// States whether the received information has to be encrypted when it is retransmitted by the recipient.
	EncryptInTransit string `json:"encrypt-in-transit,omitempty" validate:"regexp=(^$|MUST|MAY)"`
	// States whether the received information has to be encrypted by the Recipient when it is stored at rest.
	EncryptAtREST string `json:"encrypt-at-rest,omitempty" validate:"regexp=(^$|MUST|MAY)"`
	// States the permitted actions that Recipients can take upon information received.
	PermittedActions string `json:"permitted-actions,omitempty" validate:"regexp=(^$|NONE|CONTACT FOR INSTRUCTION|INTERNALLY VISIBLE ACTIONS|EXTERNALLY VISIBLE INDIRECT ACTIONS|EXTERNALLY VISIBLE DIRECT ACTIONS)"`
	// Recipients are permitted notify affected third parties of a potential compromise or threat. Examples include permitting National CSIRTs to send notifications to affected constituents, or a service provider contacting affected customers.
	AffectedPartyNotifications string `json:"affected-party-notifications,omitempty" validate:"regexp=(^$|MAY|MUST NOT)"`
	// Recipients are permitted to redistribute the information received within the redistribution scope as defined by the enumerations. The enumerations “RED”, “AMBER”, “GREEN”, “WHITE” are to be interpreted as described in the FIRST Traffic Light Protocol Policy
	TLP string `json:"tlp,omitempty" validate:"regexp=(^$|RED|AMBER|GREEN|WHITE)"`
	// Recipients could be required to attribute or anonymize the Provider when redistributing the information received.
	Attribution string `json:"attribution,omitempty" validate:"regexp=(^$|MUST|MAY|MUST NOT)"`
	// Recipients could be required to obfuscate or anonymize information that could be used to identify the affected parties before redistributing the information received. 	Examples include removing affected parties IP addresses, or removing the affected parties names but leaving the affected parties industry vertical prior to sending a notification.
	ObfuscateAffectedParties string `json:"obfuscate-affected-parties,omitempty" validate:"regexp=(^$|MAY|MUST|MUST NOT)"`
	// States whether the recipient MAY or MUST NOT resell the information received unmodified or in a semantically equivalent format. e.g. transposing the information from a .csv file format to a .json file format would be considered semantically equivalent.
	UnmodifiedResale string `json:"unmodified-resale,omitempty" validate:"regexp=(^$|MUST NOT|MAY)"`
	// This statement can be used to convey a description or reference to any applicable licenses, agreements, or conditions between the producer and receiver. e.g. specific terms of use , contractual language, agreement name, or a URL.
	ExternalReference string `json:"external-reference,omitempty"`
}

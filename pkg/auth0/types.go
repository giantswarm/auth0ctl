package auth0

type ResouceServer struct {
	Name                                      string `json:"name"`
	Identifier                                string `json:"identifier"`
	AllowOfflineAccess                        bool   `json:"allow_offline_access"`
	TokenLifetime                             int    `json:"token_lifetime"`
	TokenLifetimeWeb                          int    `json:"token_lifetime_for_web"`
	SigningAlgorithm                          string `json:"signing_alg"`
	SkipConsentForVerifiableFirstPartyClients bool   `json:"skip_consent_for_verifiable_first_party_clients"`
	EnforcePolicy                             bool   `json:"enforce_policies"`
}

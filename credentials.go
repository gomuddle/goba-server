package goba_server

// Credentials are credentials for interacting with the Server.
type Credentials struct {
	// Username is the Credentials' username.
	Username string `json:"username,omitempty"`

	// Password is the Credentials' password.
	Password string `json:"password,omitempty"`
}

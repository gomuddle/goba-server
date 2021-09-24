package goba_server

// Credentials are credentials for interacting with the Server.
type Credentials struct {
	// Username is the Credential's username.
	Username string `json:"username,omitempty"`

	// Password is the Credential's password.
	Password string `json:"password,omitempty"`
}

package model

// User model
type User struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	FullName    string `json:"fullName,omitempty"`
	SaIdentity  string `json:"saIdentity,omitempty"`
	AccountType string `json:"accountType,omitempty"`
	CreatedOn   int64  `json:"createdOn,omitempty"`
	ModifiedOn  int64  `json:"modifiedOn,omitempty"`
}

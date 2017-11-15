package model

// OutputRule model
type OutputRule struct {
	ID              string                 `json:"id,omitempty"`
	UserUD          string                 `json:"uid,omitempty"`
	AppID           string                 `json:"aid,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Enabled         bool                   `json:"enabled,omitempty"`
	Error           error                  `json:"error,omitempty"`
	Index           int                    `json:"index,omitempty"`
	Name            string                 `json:"name,omitempty"`
	CreatedOn       int                    `json:"createdOn,omitempty"`
	InvalidatedOn   int                    `json:"invalidatedOn,omitempty"`
	LanguageVersion int                    `json:"languageVersion,omitempty"`
	ModifiedOn      int                    `json:"modifiedOn,omitempty"`
	Rule            map[string]interface{} `json:"rule,omitempty"`
	Warning         RuleWarningOutput      `json:"warning,omitempty"`
	IsTestable      bool                   `json:"isTestable,omitempty"`
	Scope           string                 `json:"scope,omitempty"`
}

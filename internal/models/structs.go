package models

type Passport struct {
	Uid            string      `json:"uid"`
	SourceUid      interface{} `json:"sourceUid"`
	ReceptionDate  string      `json:"receptionDate"`
	PassportStatus struct {
		Id           int         `json:"id"`
		Name         string      `json:"name"`
		Description  interface{} `json:"description"`
		Color        string      `json:"color"`
		Subscription bool        `json:"subscription"`
	} `json:"passportStatus"`
	InternalStatus struct {
		Name    string `json:"name"`
		Percent int    `json:"percent"`
	} `json:"internalStatus"`
	Clones []interface{} `json:"clones"`
}

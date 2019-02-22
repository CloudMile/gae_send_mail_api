package model

// Form is creatd by http post x-www-form-urlencoded
type Form struct {
	To      string `json:"to"`
	CC      string `json:"cc"`
	BCC     string `json:"bcc"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

package types

import "time"

type WebsiteSSLDTO struct {
	PageResult
	Items []WebsiteSSL `json:"items"`
}

type ListWebsiteSSLRes struct {
	Response
	Data WebsiteSSLDTO `json:"data"`
}

type WebsiteSSL struct {
	ID            uint      `json:"id"`
	PrimaryDomain string    `json:"primaryDomain"`
	Domains       string    `json:"domains"`
	Provider      string    `json:"provider"`
	Organization  string    `json:"organization"`
	AutoRenew     bool      `json:"autoRenew"`
	ExpireDate    time.Time `json:"expireDate"`
	StartDate     time.Time `json:"startDate"`
	Status        string    `json:"status"`
}
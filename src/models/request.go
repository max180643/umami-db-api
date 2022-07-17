package models

type TrackRequest struct {
	WebsiteName *string `json:"websiteName" binding:"required"`
	Hostname    *string `json:"hostname" binding:"required"`
	Ip          string  `json:"ip"`
	UserAgent   string  `json:"userAgent"`
	Url         string  `json:"url"`
}

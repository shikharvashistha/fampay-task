package types

type Response struct {
	VideoData []VideoDataResponse `json:"video_data,omitempty"`
	PageNo    int                 `json:"page_no,omitempty"`
}

type VideoDataResponse struct {
	VideoTitle    string `json:"video_title,omitempty"`
	Description   string `json:"description,omitempty"`
	Publishing    string `json:"publishing,omitempty"`
	ThumnailsURLs string `json:"thumnails_urls,omitempty"`
}
type CronRequest struct {
	Name     string   `json:"name,omitempty"`
	Query    string   `json:"query,omitempty"`
	Interval string   `json:"interval,omitempty"`
	APIKeys  []string `json:"api_keys,omitempty"`
}
type CronResponse struct {
	ID string `json:"id,omitempty"`
}
type CronDeleteRequest struct {
	ID string `json:"id,omitempty"`
}

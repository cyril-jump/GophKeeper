package raw

// BlobData struct
type BlobData struct {
	ID       int    `json:"id,omitempty"`
	Data     string `json:"data"`
	Metadata string `json:"metadata"`
}

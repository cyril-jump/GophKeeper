package domain

type TextData struct {
	ID       int    `json:"id,omitempty"`
	Text     string `json:"text"`
	Metadata string `json:"metadata"`
}

type CredData struct {
	ID       int    `json:"id,omitempty"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Metadata string `json:"metadata"`
}

type CardData struct {
	ID         int    `json:"id,omitempty"`
	CardNumber string `json:"card_number"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	CVC        string `json:"cvc"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Metadata   string `json:"metadata"`
}

type BlobData struct {
	ID       int    `json:"id,omitempty"`
	Data     []byte `json:"data"`
	Metadata string `json:"metadata"`
}

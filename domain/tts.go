package domain

type TtsRequest struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}

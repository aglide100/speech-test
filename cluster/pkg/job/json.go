package job

type ReturningJob struct {
	Id	string `json:"Id,omitempty"`
	Content string `json:"Content,omitempty"`
	Speaker string `json:"Speaker,omitempty"`
	Title string `json:"Title,omitempty"`
	PlayingTime float32 `json:"PlayingTime,omitempty"`
}
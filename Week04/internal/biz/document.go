package biz

type Document struct {
	Content   string `json:"content"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Topic     string `json:"topic"`
	Watermark string `json:"watermark,omitempty"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

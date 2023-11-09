package model

type ImageCard struct {
	Type string `json:"type"`
	Src  string `json:"src"`
}

type ImageGroupCard struct {
	Type     string      `json:"type"`
	Elements []ImageCard `json:"elements"`
}

type TextCard struct {
	Type string   `json:"type"`
	Text TextData `json:"text"`
}

type TextData struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

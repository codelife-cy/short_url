package short

import "time"

type Req struct {
	ShortURL string `json:"shortURL"`
	LongURL  string `json:"longURL"`
	Note     string `json:"note"`
}

type Resp struct {
	ShortURL   string `json:"shortURL"`
	LongURL    string `json:"longURL"`
	NewLongURL string `json:"newLongUrl"`
}

type SequenceTable struct {
	Id        uint64 `gorm:"primary_key"`
	Stub      string
	Timestamp time.Time
}

func (SequenceTable) TableName() string {
	return "sequence"
}

type ShorterTable struct {
	Id         uint64 `gorm:"primary_key"`
	LongUrl    string
	ShortUrl   string
	Note       string
	CreateTime time.Time
}

func (ShorterTable) TableName() string {
	return "short"
}

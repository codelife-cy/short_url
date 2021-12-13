package short

import (
	"time"
)

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
	ID        uint `gorm:"primarykey"`
	Stub      string
	Timestamp time.Time
}

func (SequenceTable) TableName() string {
	return "sequence"
}

type ShorterTable struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	LongUrl   string
	ShortUrl  string
	Note      string
}

func (ShorterTable) TableName() string {
	return "short"
}

package short

import (
	"gin-template/base"
	"gin-template/config"
	"gin-template/util/db"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"time"
)

type Short interface {
	i()
	ShortURL(longURL, note string) (string, error)
	Expand(shortURL string) (string, error)
	UpdateURL(shortURL, longURL string) (bool, error)
}

type Shorter struct {
	db db.Repo
}

func NewShorter() Shorter {
	repo := db.Db
	return Shorter{
		db: repo,
	}
}

func (s Shorter) i() {
	panic("implement me")
}

func (s Shorter) ShortURL(longURL, note string) (string, error) {
	//_, err := url.Parse(longURL)
	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		return "", errors.Wrap(err, "requested url is malformed")
	}
	sequence := NewSequence()
	conf := config.Get()
	var shortStr string
	for {
		nextSequence, err := sequence.NextSequence()
		if err != nil {
			log.Printf("get next sequence error. %v", err)
			return "", errors.Wrap(err, "get next sequence error")
		}
		shortStr = base.CreateShortStr(nextSequence)
		// 需要过滤的短域名
		if conf.Common.BlackShortUrlMap[shortStr] {
			continue
		} else {
			break
		}
	}
	writer := s.db.GetDbW()
	shorterTable := &ShorterTable{
		ShortUrl:   shortStr,
		LongUrl:    longURL,
		Note:       note,
		CreateTime: time.Now(),
	}
	create := writer.Create(shorterTable)
	if create.Error != nil {
		log.Printf("short write db insert error. %v", create.Error)
		return "", errors.Wrap(create.Error, "short write db insert error")
	}
	shortURL := (&url.URL{Scheme: conf.Common.Schema, Host: conf.Common.DomainName, Path: shortStr}).String()
	return shortURL, nil
}

func (s Shorter) Expand(shortURL string) (string, error) {
	reader := s.db.GetDbR()
	shorterTable := new(ShorterTable)
	find := reader.Where(&ShorterTable{ShortUrl: shortURL}).Find(shorterTable)
	longUrl := shorterTable.LongUrl
	return longUrl, errors.Wrap(find.Error, "find shortURL err")
}

func (s Shorter) UpdateURL(shortURL, longURL string) (bool, error) {
	panic("implement me")
}

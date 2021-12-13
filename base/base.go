package base

import (
	"gin-template/config"
	"math"
)

func CreateShortStr(seq uint) (shortURL string) {
	conf := config.Get()
	baseStrLen := uint(len(conf.Common.BaseString))
	domainLen := conf.Common.DomainLength
	charSeq := make([]rune, 0, domainLen)
	num := math.Pow(float64(baseStrLen), float64(domainLen-1))
	seq = uint(num) + seq
	if seq != 0 {
		for seq != 0 {
			mod := seq % baseStrLen
			div := seq / baseStrLen
			charSeq = append(charSeq, rune(conf.Common.BaseString[mod]))
			seq = div
		}
	} else {
		charSeq = append(charSeq, rune(conf.Common.BaseString[seq]))
	}
	tmpShortURL := string(charSeq)
	shortURL = reverse(tmpShortURL)
	return
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

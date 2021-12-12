package short

import (
	"gin-template/util/db"
)

type SequenceDB struct {
	db db.Repo
}

func NewSequence() SequenceDB {
	repo := db.Db
	return SequenceDB{db: repo}
}

func (seq SequenceDB) NextSequence() (uint64, error) {
	w := seq.db.GetDbW()
	var sequenceTb SequenceTable
	last := w.Exec(`REPLACE INTO sequence(stub) VALUES ("sequence")`).Select("id").Last(&sequenceTb)
	if last.Error != nil {
		return 0, last.Error
	}
	sequence := sequenceTb.Id
	return sequence - 1, nil
}

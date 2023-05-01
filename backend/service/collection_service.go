package service

import (
	"encoding/json"
)

type Collection struct {
	ID             int             `json:"id" db:"id"`
	CollectionList json.RawMessage `json:"collection" db:"collection" binding:"required"`
}

func (c *Collection) CudCollection() (Collection, error) {
	r, e := db.Exec("update blog_collection set collection=?", c.CollectionList)
	if e != nil {
		return Collection{}, e
	}
	id, _ := r.LastInsertId()
	return Collection{int(id), c.CollectionList}, nil
}

func (c Collection) GetCollection() (Collection, error) {
	var collection Collection
	e := db.Get(&collection, "select * from blog_collection limit 1")
	return collection, e
}

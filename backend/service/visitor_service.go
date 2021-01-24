package service

import (
	"backend/utils"
	"time"
)

type Visitor struct {
	ID        uint64 `json:"id" db:"id"`
	Uri       string `json:"uri" db:"uri" binding:"required"`
	Referer   string `json:"referer" db:"referer"`
	Ua        string `json:"ua" db:"ua" binding:"required"`
	IP        string `json:"ip" db:"ip"`
	VisitTime int    `json:"visit_time" db:"visit_time"`
}

func (v Visitor) NewVisitor() error {
	visitTime := time.Now().Format(utils.AppInfo.TimeFormat)
	_, e := db.Exec("insert into blog_visitor (uri,referer,ua,ip,visit_time) values (?,?,?,?,?)", v.Uri, v.Referer, v.Ua, v.IP, visitTime)
	return e
}

func (v Visitor) GetVisitorCount() (count int8, e error) {
	e = db.Get(&count, "select count('id') from blog_visitor")
	return
}

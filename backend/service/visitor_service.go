package service

import (
	"backend/utils"
	"strings"
	"time"
)

type Visitor struct {
	ID        uint64 `json:"id" db:"id"`
	Uri       string `json:"uri" db:"uri" binding:"required"`
	Referer   string `json:"referer" db:"referer"`
	Ua        string `json:"ua" db:"ua" binding:"required"`
	IP        string `json:"ip" db:"ip"`
	VisitDate string `json:"visit_date" db:"visit_date"`
	VisitTime int    `json:"visit_time" db:"visit_time"`
}

type CountByDate struct {
	Count string
	Date  string
}

type CountByUA struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (v Visitor) NewVisitor() error {
	visitTime := time.Now().Format(utils.AppInfo.TimeFormat)
	visitDate := strings.Split(visitTime, " ")[0]
	_, e := db.Exec("insert into blog_visitor (uri,referer,ua,ip,visit_date,visit_time) values (?,?,?,?,?,?)", v.Uri, v.Referer, v.Ua, v.IP, visitDate, visitTime)
	return e
}

func (v Visitor) GetVisitorCount() (count int8, e error) {
	e = db.Get(&count, "select count('id') from blog_visitor")
	return
}

func (v Visitor) GetCountByDate() (map[string][]string, error) {
	rows, e := db.Queryx("SELECT COUNT(id) AS count, visit_date AS date FROM blog_visitor where DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= visit_date GROUP BY visit_date")

	if e != nil {
		return nil, e
	}
	res := make([]CountByDate, 0)
	for rows.Next() {
		var r CountByDate
		e = rows.StructScan(&r)
		res = append(res, r)
	}

	counts := make([]string, 0)
	dates := make([]string, 0)
	for _, value := range res {
		counts = append(counts, value.Count)
		dates = append(dates, value.Date)
	}
	data := make(map[string][]string)
	data["counts"] = counts
	data["dates"] = dates
	return data, rows.Err()
}

func (v Visitor) GetCountByUA() ([]CountByUA, error) {
	rows, e := db.Queryx("SELECT ua AS name,COUNT(id) AS value FROM blog_visitor GROUP BY ua")

	if e != nil {
		return nil, e
	}
	res := make([]CountByUA, 0)
	for rows.Next() {
		var r CountByUA
		e = rows.StructScan(&r)
		res = append(res, r)
	}

	return res, e
}

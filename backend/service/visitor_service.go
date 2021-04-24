package service

import (
	"backend/utils"
	"strconv"
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
	VisitTime string `json:"visit_time" db:"visit_time"`
}

type CountByDate struct {
	Count string
	Date  string
}

func (v Visitor) NewVisitor() error {
	visitTime := time.Now().Format(utils.AppInfo.TimeFormat)
	visitDate := strings.Split(visitTime, " ")[0]
	_, e := db.Exec("insert into blog_visitor (uri,referer,ua,ip,visit_date,visit_time) values (?,?,?,?,?,?)", v.Uri, v.Referer, v.Ua, v.IP, visitDate, visitTime)
	return e
}

func (v Visitor) GetVisitorCount() (count uint64, e error) {
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

func (v Visitor) GetCountByUA() (data []map[string]string, e error) {
	visitors := make([]Visitor, 0)

	if e = db.Select(&visitors, "SELECT * FROM blog_visitor"); e != nil {
		return
	}
	uas := map[string]int{"IE": 0, "Opera": 0, "Chrome": 0, "Firefox": 0, "Android": 0, "iPhone": 0, "WeChat": 0}

	for _, visitor := range visitors {
		switch {
		case strings.Contains(visitor.Ua, "Trident"):
			uas["IE"] += 1
		case strings.Contains(visitor.Ua, "Presto"):
			uas["Opera"] += 1
		case strings.Contains(visitor.Ua, "Chrome"):
			uas["Chrome"] += 1
		case strings.Contains(visitor.Ua, "Firefox"):
			uas["Firefox"] += 1
		case strings.Contains(visitor.Ua, "Android"):
			uas["Android"] += 1
		case strings.Contains(visitor.Ua, "iPhone"):
			uas["iPhone"] += 1
		case strings.Contains(visitor.Ua, "MicroMessenger"):
			uas["WeChat"] += 1
		default:
			uas[visitor.Ua] += 1
		}
	}

	for key, value := range uas {
		if value != 0 {
			data = append(data, map[string]string{"name": key, "value": strconv.Itoa(value)})
		}
	}
	return
}

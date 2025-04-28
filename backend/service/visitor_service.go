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

var (
	lowerBotKeywords    = []string{"bot", "spider", "crawler", "ahrefsbot", "yandexbot", "bingbot", "duckduckbot", "slurp", "baiduspider", "sogou"}
	lowerMobileKeywords = []string{"mobile", "android", "iphone", "ipad", "ipod", "blackberry", "iemobile", "opera mini"}
	lowerDesktopSysKeys = []string{"windows nt", "macintosh", "x11", "ubuntu"} // 系统关键词
)

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
		rows.StructScan(&r)
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
	uas := map[string]int{"Bot": 0, "Mobile": 0, "Desktop": 0, "Other": 0}

	rows, err := db.Query("SELECT ua FROM blog_visitor")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ua string
		if err = rows.Scan(&ua); err != nil {
			return nil, err
		}

		lowerUA := strings.ToLower(ua)

		category := classifyUA(lowerUA)
		uas[category]++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	for key, value := range uas {
		if value != 0 {
			data = append(data, map[string]string{"name": key, "value": strconv.Itoa(value)})
		}
	}
	return
}

func classifyUA(lowerUA string) string {
	if utils.IfContainStr(lowerUA, lowerBotKeywords) {
		return "Bot"
	}
	if utils.IfContainStr(lowerUA, lowerMobileKeywords) {
		return "Mobile"
	}
	if utils.IfContainStr(lowerUA, lowerDesktopSysKeys) {
		return "Desktop"
	}
	return "Other"
}

package service

type Tag struct {
	ID      int    `json:"id" db:"id"`
	TagName string `json:"tag_name" db:"tag_name" binding:"required,max=16"`
}

func (t *Tag) Create() (Tag, error) {
	r, e := db.Exec("insert into blog_tag (tag_name) values (?)", t.TagName)
	if e != nil {
		return Tag{}, e
	}
	id, _ := r.LastInsertId()
	return Tag{int(id), t.TagName}, nil
}

func (t *Tag) Delete() error {
	_, e := db.Exec("delete from blog_tag where id=?", t.ID)
	return e
}

func (t *Tag) Edit() error {
	_, e := db.Exec("update blog_tag set tag_name=? where id=?", t.TagName, t.ID)
	return e
}

func (t Tag) GetOne() (Tag, error) {
	var tag Tag
	e := db.Get(&tag, "select * from blog_tag where id=?", t.ID)
	return tag, e
}

func (t Tag) GetAll() ([]Tag, error) {
	tags := make([]Tag, 0)
	err := db.Select(&tags, "select * from blog_tag")
	return tags, err
}

func (t Tag) GetTagCount() (count int8, e error) {
	e = db.Get(&count, "select count('id') from blog_tag")
	return
}

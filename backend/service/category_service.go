package service

type Category struct {
	ID           int    `json:"id" db:"id"`
	CategoryName string `json:"category_name" db:"category_name" binding:"required,max=16"`
}

func (c *Category) Create() (Category, error) {
	r, e := db.Exec("insert into blog_category (category_name) values (?)", c.CategoryName)
	if e != nil {
		return Category{}, e
	}
	id, _ := r.LastInsertId()
	return Category{int(id), c.CategoryName}, nil
}

func (c *Category) Delete() error {
	_, e := db.Exec("delete from blog_category where id=?", c.ID)
	return e
}

func (c *Category) Edit() error {
	_, e := db.Exec("update blog_category set category_name=? where id=?", c.CategoryName, c.ID)
	return e
}

func (c *Category) GetOne() (Category, error) {
	var category Category
	e := db.Get(&category, "select * from blog_category where id=?", c.ID)
	return category, e
}

func (c Category) GetAll() ([]Category, error) {
	categories := make([]Category, 0)
	err := db.Select(&categories, "select * from blog_category")
	return categories, err
}

func (c Category) GetCategoryCount() (count int8, e error) {
	e = db.Get(&count, "select count('id') from blog_category")
	return
}

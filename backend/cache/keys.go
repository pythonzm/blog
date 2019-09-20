package cache

import (
	"fmt"
)

func ArticleKey(limit, page int, admin bool, category, tag string) string {
	if admin {
		return fmt.Sprintf("article_%d_%d_%s_%s_%s", limit, page, "admin", category, tag)
	} else {
		return fmt.Sprintf("article_%d_%d_%s_%s", limit, page, category, tag)
	}
}

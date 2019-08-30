package cache

import (
	"fmt"
)

func ArticleKey(limit, page int) string {
	return fmt.Sprintf("article_%d_%d", limit, page)
}

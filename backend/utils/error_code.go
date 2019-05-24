package utils

const (
	Success       = 200
	InvalidParams = 400
	AuthError     = 401

	TagExistError    = 10001
	TagGetError      = 10002
	TagNotExistError = 10003
	TagGetAllError   = 10004
	TagAddError      = 10006
	TagEditError     = 10007
	TagDeleteError   = 10008

	CategoryExistError    = 20001
	CategoryGetError      = 20002
	CategoryNotExistError = 20003
	CategoryGetAllError   = 20004
	CategoryAddError      = 20006
	CategoryEditError     = 20007
	CategoryDeleteError   = 20008

	ArticleExistError    = 30001
	ArticleGetError      = 30002
	ArticleNotExistError = 30003
	ArticleGetAllError   = 30004
	ArticleCountError    = 30005
	ArticleAddError      = 30006
	ArticleEditError     = 30007
	ArticleDeleteError   = 30008

	TokenCheckError   = 40001
	TokenTimeoutError = 40002
	TokenGenError     = 40003
)

var MsgFlags = map[int]string{
	Success:       "success",
	InvalidParams: "请求参数错误",
	AuthError:     "用户名或密码错误",

	TagExistError:    "标签已存在",
	TagGetError:      "标签获取失败",
	TagNotExistError: "标签不存在",
	TagGetAllError:   "获取所有标签失败",
	TagAddError:      "添加标签失败",
	TagEditError:     "编辑标签失败",
	TagDeleteError:   "删除标签失败",

	CategoryExistError:    "分类已存在",
	CategoryGetError:      "分类获取失败",
	CategoryNotExistError: "分类不存在",
	CategoryGetAllError:   "获取所有分类失败",
	CategoryAddError:      "添加分类失败",
	CategoryEditError:     "编辑分类失败",
	CategoryDeleteError:   "删除分类失败",

	ArticleExistError:    "文章已存在",
	ArticleGetError:      "文章获取失败",
	ArticleNotExistError: "文章不存在",
	ArticleGetAllError:   "获取所有文章失败",
	ArticleCountError:    "统计文章失败",
	ArticleAddError:      "添加文章失败",
	ArticleEditError:     "编辑文章失败",
	ArticleDeleteError:   "删除文章失败",

	TokenCheckError:   "Token鉴权失败",
	TokenTimeoutError: "Token已超时",
	TokenGenError:     "Token生成失败",
}

func ErrorText(code int) string {
	return MsgFlags[code]
}

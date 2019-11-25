package utils

const (
	Success = 20000

	InvalidParams = 40000
	AuthError     = 40001

	TokenCheckError   = 40002
	TokenTimeoutError = 40003
	TokenGenError     = 40004

	TagExistError    = 40005
	TagGetError      = 40006
	TagNotExistError = 40007
	TagGetAllError   = 40008
	TagCreateError   = 40009
	TagEditError     = 40010
	TagDeleteError   = 40011

	CategoryExistError    = 40012
	CategoryGetError      = 40013
	CategoryNotExistError = 40014
	CategoryGetAllError   = 40015
	CategoryCreateError   = 40016
	CategoryEditError     = 40017
	CategoryDeleteError   = 40018

	ArticleExistError    = 40019
	ArticleGetError      = 40020
	ArticleNotExistError = 40021
	ArticleGetAllError   = 40022
	ArticleCountError    = 40023
	ArticleCreateError   = 40024
	ArticleEditError     = 40025
	ArticleDeleteError   = 40026

	UserGetError      = 40027
	UserEditError     = 40028
	RestPasswordError = 40031

	GetUploadImageError  = 40029
	SaveUploadImageError = 40030

	CommentGetAllError = 40032
	SearchArticleError = 40033
)

var MsgFlags = map[int]string{
	Success:       "success",
	InvalidParams: "请求参数错误",
	AuthError:     "用户名或密码错误",

	TagExistError:    "标签已存在",
	TagGetError:      "标签获取失败",
	TagNotExistError: "标签不存在",
	TagGetAllError:   "获取所有标签失败",
	TagCreateError:   "添加标签失败",
	TagEditError:     "编辑标签失败",
	TagDeleteError:   "删除标签失败",

	CategoryExistError:    "分类已存在",
	CategoryGetError:      "分类获取失败",
	CategoryNotExistError: "分类不存在",
	CategoryGetAllError:   "获取所有分类失败",
	CategoryCreateError:   "添加分类失败",
	CategoryEditError:     "编辑分类失败",
	CategoryDeleteError:   "删除分类失败",

	ArticleExistError:    "文章已存在",
	ArticleGetError:      "文章获取失败",
	ArticleNotExistError: "文章不存在",
	ArticleGetAllError:   "获取所有文章失败",
	ArticleCountError:    "统计文章失败",
	ArticleCreateError:   "添加文章失败",
	ArticleEditError:     "编辑文章失败",
	ArticleDeleteError:   "删除文章失败",

	TokenCheckError:   "Token鉴权失败",
	TokenTimeoutError: "Token已超时",
	TokenGenError:     "Token生成失败",

	UserGetError:      "用户信息获取失败",
	UserEditError:     "用户信息更新失败",
	RestPasswordError: "修改密码失败",

	GetUploadImageError:  "获取上传文件信息失败",
	SaveUploadImageError: "保存上传文件失败",

	CommentGetAllError: "获取所有评论失败",
	SearchArticleError: "搜索文章失败",
}

func ErrorText(code int) string {
	return MsgFlags[code]
}


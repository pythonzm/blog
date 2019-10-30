package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ImageAvatarInfo struct {
	ImageFullUrl  string
	ImageRelPath  string
	AvatarFullUrl string
	AvatarRelPath string
}

func MD5ImageName(name string) string {
	ext := filepath.Ext(name)
	fileName := EncodeMD5(name + fmt.Sprintf("%v", time.Now().Unix()))

	return fileName + ext
}

func GenImageAvatarInfo(t, n string) (i ImageAvatarInfo) {
	if t == "avatar" {
		i.ImageRelPath = ""
		i.ImageFullUrl = ""
		i.AvatarRelPath = filepath.Join(AppInfo.UploadBasePath, AppInfo.AvatarRelPath, n)
		i.AvatarFullUrl = fmt.Sprintf(`%s/%s/%s/%s`, AppInfo.ApiBaseUrl, AppInfo.StaticBasePath, AppInfo.AvatarRelPath, n)
	} else if t == "image" {
		i.ImageRelPath = filepath.Join(AppInfo.UploadBasePath, AppInfo.ImageRelPath, today(), n)
		i.ImageFullUrl = fmt.Sprintf(`%s/%s/%s/%s/%s`, AppInfo.ApiBaseUrl, AppInfo.StaticBasePath, AppInfo.ImageRelPath, today(), n)
		i.AvatarRelPath = ""
		i.AvatarFullUrl = ""
	}
	return
}

func today() string {
	t := time.Now().Format("20060102")
	return t
}

// 如果不存在返回true，存在则返回false
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func IsNotExistMkDir(src string) error {
	if CheckExist(src) {
		e := MkDir(src)
		return e
	}
	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModeDir)
	return err
}


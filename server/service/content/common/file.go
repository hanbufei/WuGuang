package common

import (
	"strings"
)

// 判断文件类型
func GetExtName(extName string) (result string) {
	extName = strings.ToLower(extName)
	if extName == "txt" {
		result = "txt"
	}
	if extName == "md" {
		result = "md"
	}
	if strings.HasPrefix(extName, "htm") {
		result = "html"
	}
	return
}

//获取文本中的图片外链列表
//func GetImgSrcList(text string) (imgSrcList []string,err error)  {
//	if !strings.Contains(text,"<img src=\"http"){
//		err = gerror.New("不存在图片外链")
//		return
//	}
//	re := regexp.MustCompile(`<img[^>]+src=["http']([^"']+)["'][^>]*>`)
//	matches := re.FindAllStringSubmatch(text, -1)
//	for _, match := range matches {
//		imgSrcList = append(imgSrcList, match[1])
//	}
//	return
//}

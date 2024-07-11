package tools

import (
	"backend/models"
	"fmt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
)

type AlgoliaArticle struct {
	ObjectID string `json:"objectID"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

// Algolia免费套餐每个record的大小限制是10K
// 如果record过大需要将文本切割
func (aa *AlgoliaArticle) splitNum() int {
	if 9000 < len(aa.Content) && len(aa.Content) < 16000 {
		return 2
	} else if len(aa.Content) >= 16000 {
		return 4
	} else {
		return 0
	}
}

func (aa *AlgoliaArticle) AddRecord() error {
	num := aa.splitNum()
	var objects []AlgoliaArticle
	if num != 0 {
		parts, _ := splitStringIntoParts(aa.Content, num)

		for i := 0; i < num; i++ {
			object := AlgoliaArticle{
				ObjectID: fmt.Sprintf("%s-%d", aa.ObjectID, i),
				ID:       aa.ID,
				Title:    aa.Title,
				Content:  parts[i],
			}
			objects = append(objects, object)

		}
	} else {
		objects = append(objects, *aa)
	}
	_, err := models.AlgoliaIndex.SaveObjects(objects, opt.ExposeIntermediateNetworkErrors(true))
	fmt.Println("add record success")
	return err
}

func (aa *AlgoliaArticle) DeleteRecord() error {
	_, err := models.AlgoliaIndex.DeleteBy(opt.Filters(fmt.Sprintf("id:%d", aa.ID)))
	fmt.Println("delete record success")
	return err
}

func (aa *AlgoliaArticle) UpdateRecord() error {
	num := aa.splitNum()
	if num == 0 {
		_, err := models.AlgoliaIndex.PartialUpdateObject(aa)
		fmt.Println("update record success")
		return err
	} else {
		err := aa.DeleteRecord()
		if err != nil {
			return err
		}
		fmt.Println("update record success")
		return aa.AddRecord()
	}
}

// 将字符串分成指定数量的部分
func splitStringIntoParts(s string, numParts int) ([]string, error) {
	if numParts <= 0 {
		return nil, fmt.Errorf("number of parts must be greater than 0")
	}

	n := len(s)
	partSize := n / numParts
	remainder := n % numParts

	parts := make([]string, 0, numParts)
	start := 0

	for i := 0; i < numParts; i++ {
		end := start + partSize
		if remainder > 0 {
			end++
			remainder--
		}
		parts = append(parts, s[start:end])
		start = end
	}

	return parts, nil
}

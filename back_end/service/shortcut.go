package service

import (
	"back_end/model"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Shortcut struct {
}

func (s *Shortcut) Link_add(link string, title string) error {

	url := link
	url_icon := ""
	baseurls := strings.Split(url, "/")
	baseurl := baseurls[0] + "//" + baseurls[2]

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	doc.Find("head").Each(func(index int, head *goquery.Selection) {
		head.Find("link[rel='icon'], link[rel='shortcut icon'], meta[name='msapplication-TileImage']").Each(func(i int, link *goquery.Selection) {
			href, exists := link.Attr("href")
			if exists {
				url_icon = href
			}
		})
	})

	if url_icon == "" {
		url_icon = baseurl + "/favicon.ico"
	} else {
		if (strings.Split(url_icon, "/")[0] != "https:") && (strings.Split(url_icon, "/")[0] != "http:") {
			url_icon = baseurl + url_icon
		}
	}

	if title == "" {
		if doc.Find("title").Text() != "" {
			title = doc.Find("title").Text()
		} else {
			title = baseurls[2]
		}
	}

	shortcut := &model.Shortcut{}
	shortcut.LinkOri = link
	shortcut.LinkIcon = url_icon
	shortcut.LinkTitle = title

	if err := model.DB.Model(&model.Shortcut{}).Where("link_ori = ?", link).FirstOrCreate(&shortcut).Error; err != nil {
		return err
	}
	return nil
}

func (s *Shortcut) Links_get() (error, []model.Shortcut) {
	shortcuts := []model.Shortcut{}
	if err := model.DB.Model(&model.Shortcut{}).Find(&shortcuts).Error; err != nil {
		return err, nil
	}
	return nil, shortcuts
}

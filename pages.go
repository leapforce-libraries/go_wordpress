package wordpress

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	"strings"
)

type Page struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Link  string `json:"link"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered string `json:"rendered"`
	} `json:"content"`
}

func (wp *Service) GetPages(pageID *int, excludeIDs *[]string) (*[]Page, *errortools.Error) {
	if wp == nil {
		return nil, errortools.ErrorMessage("Service pointer is nil")
	}

	perPage := 10
	page := 1

	pages := []Page{}

	urlString := fmt.Sprintf("%s/pages", wp.BaseURL())
	if pageID != nil {
		urlString = fmt.Sprintf("%s/%v", urlString, *pageID)
	} else {
		urlString = fmt.Sprintf("%s?per_page=%v", urlString, perPage)
		if excludeIDs != nil {
			urlString = fmt.Sprintf("%s&exclude=%s", urlString, strings.Join(*excludeIDs, ","))
		}
	}

	for {
		pages_ := []Page{}

		e := wp.Get(fmt.Sprintf("%s&page=%v", urlString, page), &pages_)
		//fmt.Println(fmt.Sprintf("%s&page=%v", urlString, page))
		if e != nil {
			return nil, e
		}

		pages = append(pages, pages_...)

		if len(pages_) < perPage {
			break
		}

		//fmt.Println(len(pages_))

		if len(pages_) == 0 {
			break
		}

		page++
	}

	return &pages, nil
}

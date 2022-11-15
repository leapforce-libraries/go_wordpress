package wordpress

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
)

type Type struct {
	ID   int `json:"id"`
	GUID struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	Slug  string `json:"slug"`
	Link  string `json:"link"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered string `json:"rendered"`
	} `json:"content"`
}

func (wp *Service) GetTypes(typeName string, typeID *int) (*[]Type, *errortools.Error) {
	if wp == nil {
		return nil, errortools.ErrorMessage("Service pointer is nil")
	}

	perPage := 100
	page := 1

	types := []Type{}

	urlString := fmt.Sprintf("%s/%s", wp.BaseURL(), typeName)
	if typeID != nil {
		urlString = fmt.Sprintf("%s/%v", urlString, *typeID)
	}

	for {
		types_ := []Type{}

		e := wp.Get(fmt.Sprintf("%s?per_page=%v&page=%v", urlString, perPage, page), &types_)
		if e != nil {
			return nil, e
		}

		types = append(types, types_...)

		//if len(types_) < perPage {
		//	break
		//}

		if len(types_) == 0 {
			break
		}

		page++
	}

	return &types, nil
}

package wordpress

import (
	"fmt"

	types "github.com/leapforce-libraries/go_types"
)

type Type struct {
	ID    int    `json:"id"`
	Slug  string `json:"slug"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
}

func (wp *WordPress) GetTypes(typeName string, typeID *int) (*[]Type, error) {
	if wp == nil {
		return nil, &types.ErrorString{"WordPress pointer is nil"}
	}

	perPage := 100
	page := 1

	types := []Type{}

	urlString := fmt.Sprintf("%s/%s", wp.BaseURL(), typeName)
	if typeID != nil {
		urlString = fmt.Sprintf("%s/%v", urlString, *typeID)
	}

	for true {
		types_ := []Type{}

		err := wp.Get(fmt.Sprintf("%s?per_page=%v&page=%v", urlString, perPage, page), &types_)
		if err != nil {
			return nil, err
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

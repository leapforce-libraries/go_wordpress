package wordpress

import (
	"fmt"

	types "github.com/leapforce-libraries/go_types"
)

type Type struct {
	ID int `json:"id"`
}

func (wp *WordPress) GetTypes(typeName string, typeID *int) (*[]Type, error) {
	if wp == nil {
		return nil, &types.ErrorString{"WordPress pointer is nil"}
	}

	urlString := fmt.Sprintf("%s/%s", wp.BaseURL(), typeName)
	if typeID != nil {
		urlString = fmt.Sprintf("%s/%v", urlString, *typeID)
	}
	fmt.Println(urlString)

	types := []Type{}

	err := wp.Get(urlString, &types)
	if err != nil {
		return nil, err
	}

	return &types, nil
}

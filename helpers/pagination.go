package helpers

import (
	"pagination/initialisers"
)

type PaginationData struct {
	NextPage        int
	PreviousPage    int
	CurrentPage     int
	TotalPages      int
	TwoPreviousPage int
	TwoNextPage     int
	ThreeNextPage   int
	Offset          int
	BaseUrl         string
}

func GetPagintionData(pageNo int, perPage int, model interface{}, baseUrl string) PaginationData {
	var totalRows int64
	initialisers.DB.Model(model).Count(&totalRows)
	totalPages := float64(totalRows / int64(perPage))

	offset := (pageNo - 1) * perPage

	return PaginationData{NextPage: pageNo + 1, PreviousPage: pageNo - 1, CurrentPage: pageNo, TotalPages: int(totalPages), TwoNextPage: pageNo + 2, TwoPreviousPage: pageNo - 2, ThreeNextPage: pageNo + 3, Offset: offset, BaseUrl: baseUrl}
}

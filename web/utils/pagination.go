package utils

import (
	"access-point/web/model"
	"math"
	"net/http"
	"strconv"
)

const (
	maxLimit = 100.0
	pageKey = "pageNumber"
	limitKey = "itemsPerPage"
	searchKey = "search"
	sortByKey = "sortBy"
	sortOrderKey = "sortOrder"
)

func parsePage(r *http.Request) int {
	pageStr := r.URL.Query().Get("pageKey")
	page, _ := strconv.ParseInt(pageStr, 10, 32)
	page = int64(math.Max(1.0, float64(page)))
	return int(page)
}

func parseLimit(r *http.Request) int {
	limitStr := r.URL.Query().Get("limitKey")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	limit = int64(math.Max(0.0, math.Min(maxLimit, float64(limit))))
	return int(limit)
}

func CountTotalPages(limit, totalItems int) int {
	return int(math.Ceil(float64(totalItems) / math.Max(1.0, float64(limit))))
}

func GetPaginationParams(r *http.Request, defaultSortBy, defaultSortOrder string) model.PaginationParams {
	params := model.PaginationParams {
		Page:  1,
		Limit:   10,
		Search:   "",
		SortBy:    defaultSortBy,
		SortOrder:   defaultSortOrder,
	}

	for k := range r.URL.Query() {
		switch k {
		case "pageKey":
			params.Page = parsePage(r)

		case "limitKey":
			params.Limit = parseLimit(r)

		case "searchKey":
			params.Search = r.URL.Query().Get("searchKey")

		case "sortByKey":
			params.SortBy = r.URL.Query().Get("sortByKey")
			
		case "sortOrderKey":
			params.SortOrder = r.URL.Query().Get("sortOrderKey")

		default:

		}
	}

	return params
}

func GetSortingData(r *http.Request, defaultSortBy, defaultSortOrder string) (sortBy, sortOrder string) {
	params := GetPaginationParams(r, defaultSortBy, defaultSortOrder)
	return params.SortBy, params.SortOrder
}

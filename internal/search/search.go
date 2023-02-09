package search

import "finder/internal/api"

func Search(payload api.Payload) api.Result {
	res := api.Result{}
	ids := make([]int, 0, 64)
	ids = append(ids, 1)
	ids = append(ids, 2, 3, 4)

	res.ProductIds = ids

	return res
}

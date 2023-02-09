package api

type Result struct {
	Error      string `json:"error,omitempty"`
	ProductIds []int  `json:"productIds,omitempty"`
}

type Payload struct {
	Query    string `json:"query"`      // поисковый запрос
	Category int    `json:"categoryId"` // ID категории, может быть пустым
}

package paginate

/*
Pagination creates a Paginator with query params
*/
type Pagination struct {
	// Limit is the limit of interfaces we want
	Limit int `json:"limit,omitempty;query:limit"`
	// Page is the number of pages we have
	Page int `json:"page,omitempty;query:page"`
	// Sort gets the sorting order
	Sort string `json:"sort,omitempty;query:sort"`
	// IfNext is a boolean if a next page is present
	IfNext *bool `json:"if_next"`
	// Rows is the body to be sent in response
	Rows interface{} `json:"rows"`

	Query string `json:"query"`
}

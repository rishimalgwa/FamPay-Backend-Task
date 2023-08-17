package paginate

// GetOffset gets the offset for the gorm query
func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

// GetLimit gets the limit and checks if limit does not exist
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

// GetPage gets the page
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

// GetSort gets the sorting order
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	} else if p.Sort == "asc" {
		p.Sort = "created_at asc"
	}
	return p.Sort
}

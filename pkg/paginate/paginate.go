package paginate

import (
	"gorm.io/gorm"
)

// Paginate implements the paginator
func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {

	// Get Limit
	pagination.Limit = pagination.GetLimit()

	// We return a gorm.DB query
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit() + 1).Order(pagination.GetSort())
	}
}

/*
WithWhere : allows pagination with a where clause in the query.

Accepts: Value, query, args and Pagination struct along with gorm.DB with a PreloadIdentifier

Returns: Paginated Value
*/
func WithWhere(value interface{}, query interface{}, args interface{}, pagination *Pagination, db *gorm.DB, PreloadIdentifier string) func(db *gorm.DB) *gorm.DB {

	// Get Limit
	pagination.Limit = pagination.GetLimit()

	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset(pagination.GetOffset()).
			Limit(pagination.GetLimit()+1).
			Order(pagination.GetSort()).
			Where(query, args)
	}

}

package repository

import (
	"app/internal"
)

// NewRepositoryTicketMap creates a new repository for tickets in a map
func NewRepositoryTicketMap(lastId int, db map[int]internal.TicketAttributes) *RepositoryTicketMap {
	defaultDb := make(map[int]internal.TicketAttributes)
	if db != nil {
		defaultDb = db
	}
	return &RepositoryTicketMap{
		lastId: lastId,
		db:     defaultDb,
	}
}

// RepositoryTicketMap implements the repository interface for tickets in a map
type RepositoryTicketMap struct {
	// db represents the database in a map
	// - key: id of the ticket
	// - value: ticket
	db map[int]internal.TicketAttributes
	// lastId represents the last id of the ticket
	lastId int
}

// GetAll returns all the tickets
func (r *RepositoryTicketMap) Get() (t map[int]internal.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

// GetTicketsByDestinationCountry returns the tickets filtered by destination country
func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(country string) (t map[int]internal.TicketAttributes, err error) {
	// create a copy of the map
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return
}

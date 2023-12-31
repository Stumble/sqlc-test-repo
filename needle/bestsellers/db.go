// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v2.1.0-1-g1f618f69-wicked-fork

package bestsellers

import (
	"github.com/stumble/dcache"
	"github.com/stumble/wpgx"
)

// BeforeDump allows you to edit result before dump.
type BeforeDump func(m *Bestseller)

func New(db wpgx.WGConn, cache *dcache.DCache) *Queries {
	return &Queries{db: db, cache: cache}
}

type Queries struct {
	db    wpgx.WGConn
	cache *dcache.DCache
}

func (q *Queries) WithTx(tx *wpgx.WTx) *Queries {
	return &Queries{
		db:    tx,
		cache: q.cache,
	}
}

func (q *Queries) WithCache(cache *dcache.DCache) *Queries {
	return &Queries{
		db:    q.db,
		cache: cache,
	}
}

var Schema = `
CREATE TABLE IF NOT EXISTS BestSellers (
   ID          INT GENERATED ALWAYS AS IDENTITY,
   Name        VARCHAR(255) NOT NULL,
   FirstBSTime TIMESTAMP DEFAULT NOW(),
   PRIMARY KEY(ID)
);
`

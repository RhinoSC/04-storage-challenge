package migrator

import (
	"app/internal"
)

type MigratorProduct struct {
	ld internal.LoaderProduct
	rp internal.RepositoryProduct
}

func NewMigratorProduct(ld internal.LoaderProduct, rp internal.RepositoryProduct) *MigratorProduct {
	return &MigratorProduct{
		ld: ld,
		rp: rp,
	}
}

func (m *MigratorProduct) Migrate() (err error) {
	// load the data
	c, err := m.ld.Load()
	if err != nil {
		return
	}
	// save the data
	for _, v := range c {
		err = m.rp.Save(&v)
		if err != nil {
			return
		}
	}
	return
}

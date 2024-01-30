package migrator

import (
	"app/internal"
)

type MigratorSale struct {
	ld internal.LoaderSale
	rp internal.RepositorySale
}

func NewMigratorSale(ld internal.LoaderSale, rp internal.RepositorySale) *MigratorSale {
	return &MigratorSale{
		ld: ld,
		rp: rp,
	}
}

func (m *MigratorSale) Migrate() (err error) {
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

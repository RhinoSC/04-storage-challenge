package migrator

import (
	"app/internal"
)

type MigratorCustomer struct {
	ld internal.LoaderCustomer
	rp internal.RepositoryCustomer
}

func NewMigratorCustomer(ld internal.LoaderCustomer, rp internal.RepositoryCustomer) *MigratorCustomer {
	return &MigratorCustomer{
		ld: ld,
		rp: rp,
	}
}

func (m *MigratorCustomer) Migrate() (err error) {
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

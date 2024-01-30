package migrator

import (
	"app/internal"
)

type MigratorInvoice struct {
	ld internal.LoaderInvoice
	rp internal.RepositoryInvoice
}

func NewMigratorInvoice(ld internal.LoaderInvoice, rp internal.RepositoryInvoice) *MigratorInvoice {
	return &MigratorInvoice{
		ld: ld,
		rp: rp,
	}
}

func (m *MigratorInvoice) Migrate() (err error) {
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

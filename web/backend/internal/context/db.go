package context

import "fmt"

type DbIf interface {
	Release() error
}

func newDb(dbType, dbPath string) (DbIf, error) {
	switch dbType {
	case "bbolt":
		return newBboltDb(dbPath)
	}
	return nil, fmt.Errorf("unsupported db type: %s", dbType)
}

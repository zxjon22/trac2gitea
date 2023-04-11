package gitea

import (
	"gorm.io/gorm/clause"
)

// DB-agnostic GREATEST function
func (accessor *DefaultAccessor) Greatest(expr string, args ...interface{}) clause.Expr {
	greatest := "GREATEST("
	if accessor.dbType == "sqlite3" {
		greatest = "MAX("
	}

	return clause.Expr{SQL: greatest + expr + ")", Vars: args}
}

package gitea

import (
	"errors"
	"fmt"
	"strings"

	"github.com/stevejefferson/trac2gitea/log"
	"gorm.io/gorm/clause"
)

// DB-agnostic GREATEST function
func (accessor *DefaultAccessor) Greatest(expr string, args ...interface{}) clause.Expr {
	name := accessor.db.Dialector.Name()
	log.Debug(name)
	greatest := "GREATEST("
	if accessor.dbType == "sqlite3" {
		greatest = "MAX("
	} else if accessor.dbType == "mssql" {
		// GREATEST is not available until SQL Server 2022
		// This is a bit hacky but works enough for our use case
		params := strings.Split(expr, ",")
		for i := range params {
			params[i] = strings.TrimSpace(params[i])
			if params[i] == "?" {
				params[i] = fmt.Sprintf("@p%d", i)
			}
		}

		if len(params) > 2 {
			panic(errors.New("Cannot create GREATEST with more the 2 parms for mssql"))
		}

		expr = fmt.Sprintf("IIF(%s > %s, %s, %s)", params[0], params[1], params[0], params[1])
		argv := map[string]interface{}{}
		for i := range args {
			argv[fmt.Sprintf("@p%d", i)] = args[i]
		}

		return clause.Expr{SQL: expr, Vars: args}
	}

	return clause.Expr{SQL: greatest + expr + ")", Vars: args}
}

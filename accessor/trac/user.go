// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package trac

import "github.com/pkg/errors"

// GetUserNames retrieves the names of all users mentioned in Trac tickets, wiki pages etc., passing each one to the provided "handler" function.
func (accessor *DefaultAccessor) GetUserNames(handlerFn func(userName string) error) error {
	// find every conceivable place where a user name may be hiding
	// some of these may be redundant but it does no harm
	rows, err := accessor.db.Query(`
		SELECT owner FROM ticket WHERE owner IS NOT NULL and owner != ''
		UNION SELECT author FROM attachment
		UNION SELECT author FROM ticket_change
		UNION SELECT oldvalue FROM ticket_change WHERE field='owner' AND oldvalue != ''
		UNION SELECT newvalue FROM ticket_change WHERE field='owner' AND newvalue != ''
		UNION SELECT author FROM wiki`)
	if err != nil {
		err = errors.Wrapf(err, "retrieving Trac users")
		return err
	}

	for rows.Next() {
		var userName string
		if err = rows.Scan(&userName); err != nil {
			err = errors.Wrapf(err, "retrieving Trac user")
			return err
		}

		if err = handlerFn(userName); err != nil {
			return err
		}

	}

	return nil
}

// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

// CommitTransaction commits a Gitea transaction.
func (accessor *DefaultAccessor) CommitTransaction() error {
	err := accessor.db.Commit().Error
	if err != nil {
		return err
	}

	if !accessor.dbOnly {
		return accessor.commitWikiRepo()
	}

	return nil
}

// RollbackTransaction rolls back a Gitea transaction.
func (accessor *DefaultAccessor) RollbackTransaction() error {
	err := accessor.db.Rollback().Error
	if err != nil {
		return err
	}

	if !accessor.dbOnly {
		return accessor.rollbackWikiRepo()
	}

	return nil
}

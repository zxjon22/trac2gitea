// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (accessor *DefaultAccessor) getRepoID(userName string, repoName string) (int64, error) {
	var id int64 = NullID

	err := accessor.db.Model(&Repository{}).
		Where("owner_name=? AND name=?", userName, repoName).
		Limit(1).
		Pluck("id", &id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrapf(err, "retrieving id of repository %s for user %s", repoName, userName)
		return NullID, err
	}

	return id, nil
}

// UpdateRepoIssueCounts updates issue counts for our chosen Gitea repository.
func (accessor *DefaultAccessor) UpdateRepoIssueCounts() error {
	// TODO: All these bulk updates are wrong? Don't filter by repo_id in the subselect?
	err := accessor.db.Model(&Repository{}).
		Where("id=?", accessor.repoID).
		Update("num_issues", accessor.db.Model(&Issue{}).
			Where("repo_id=?", accessor.repoID).
			Select("count(id)")).
		Error

	if err == nil {
		err = accessor.db.Model(&Repository{}).
			Where("id=?", accessor.repoID).
			Update("num_closed_issues", accessor.db.Model(&Issue{}).
				Select("count(id)").
				Where("is_closed=? AND repo_id=?", 1, accessor.repoID)).
			Error
	}

	if err != nil {
		err = errors.Wrapf(err, "updating number of milestones for repository %d", accessor.repoID)
		return err
	}

	return nil
}

// UpdateRepoMilestoneCounts updates milestone counts for our chosen Gitea repository.
func (accessor *DefaultAccessor) UpdateRepoMilestoneCounts() error {
	err := accessor.db.Model(&Repository{}).
		Where("id=?", accessor.repoID).
		Update("num_milestones", accessor.db.Model(&Milestone{}).
			Where("repo_id=?", accessor.repoID).
			Select("count(id)")).
		Error

	if err == nil {
		err = accessor.db.Model(&Repository{}).
			Where("id=?", accessor.repoID).
			Update("num_closed_milestones", accessor.db.Model(&Milestone{}).
				Select("count(id)").
				Where("is_closed=? AND repo_id=?", 1, accessor.repoID)).
			Error
	}

	if err != nil {
		err = errors.Wrapf(err, "updating number of milestones for repository %d", accessor.repoID)
		return err
	}

	return nil
}

// GetCommitURL retrieves the URL for viewing a given commit in the current repository
func (accessor *DefaultAccessor) GetCommitURL(commitID string) string {
	repoURL := accessor.getUserRepoURL()
	return fmt.Sprintf("%s/commit/%s", repoURL, commitID)
}

// GetSourceURL retrieves the URL for viewing the latest version of a source file on a given branch of the current repository
func (accessor *DefaultAccessor) GetSourceURL(branchPath string, filePath string) string {
	repoURL := accessor.getUserRepoURL()
	return fmt.Sprintf("%s/src/branch/%s/%s", repoURL, branchPath, filePath)
}

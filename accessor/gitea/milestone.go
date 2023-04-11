// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/stevejefferson/trac2gitea/log"
	"gorm.io/gorm"
)

// GetMilestoneID gets the ID of a named milestone - returns NullID if no such milestone
func (accessor *DefaultAccessor) GetMilestoneID(milestoneName string) (int64, error) {
	var id int64 = NullID
	err := accessor.db.Model(&Milestone{}).
		Where("repo_id=? AND name=?", accessor.repoID, milestoneName).
		Limit(1).
		Pluck("id", &id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrapf(err, "retrieving id of milestone %s", milestoneName)
		return NullID, err
	}

	return id, nil
}

// updateMilestone updates an existing milestone
func (accessor *DefaultAccessor) updateMilestone(milestoneID int64, milestone *Milestone) error {
	milestone.RepoId = accessor.repoID
	milestone.ID = milestoneID

	if err := accessor.db.Save(&milestone).Error; err != nil {
		return errors.Wrapf(err, "updating milestone %s", milestone.Name)
	}

	log.Debug("updated milestone %s (id %d)", milestone.Name, milestoneID)

	return nil
}

// insertMilestone inserts a new milestone, returns milstone id.
func (accessor *DefaultAccessor) insertMilestone(milestone *Milestone) (int64, error) {
	milestone.RepoId = accessor.repoID

	if err := accessor.db.Create(&milestone).Error; err != nil {
		err = errors.Wrapf(err, "adding milestone %s", milestone.Name)
		return NullID, err
	}

	log.Debug("added milestone %s (id %d)", milestone.Name, milestone.ID)

	return milestone.ID, nil
}

// AddMilestone adds a milestone to Gitea, returns id of created milestone
func (accessor *DefaultAccessor) AddMilestone(milestone *Milestone) (int64, error) {
	milestoneID, err := accessor.GetMilestoneID(milestone.Name)
	if err != nil {
		return NullID, err
	}

	if milestoneID == NullID {
		return accessor.insertMilestone(milestone)
	}

	if accessor.overwrite {
		err = accessor.updateMilestone(milestoneID, milestone)
		if err != nil {
			return NullID, err
		}
	} else {
		log.Debug("milestone %s already exists - ignored", milestone.Name)
	}

	return milestoneID, nil
}

// GetMilestoneURL gets the URL for accessing a given milestone
func (accessor *DefaultAccessor) GetMilestoneURL(milestoneID int64) string {
	repoURL := accessor.getUserRepoURL()
	return fmt.Sprintf("%s/milestone/%d", repoURL, milestoneID)
}

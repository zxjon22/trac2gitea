// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"github.com/pkg/errors"
	"github.com/stevejefferson/trac2gitea/log"
	"gorm.io/gorm"
)

// getIssueAssigneeID retrieves the id of the given issue/assignee association, returns gitea.NullID if no such association
func (accessor *DefaultAccessor) getIssueAssigneeID(issueID int64, assigneeID int64) (int64, error) {
	var id int64 = NullID
	err := accessor.db.Model(&IssueAssignee{}).
		Where("issue_id=? AND assignee_id=?", issueID, assigneeID).
		Limit(1).
		Pluck("id", &id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrapf(err, "retrieving id for issue %d/assignee %d", issueID, assigneeID)
		return NullID, err
	}

	return id, nil
}

// updateIssueAssignee updates an existing issue assignee
func (accessor *DefaultAccessor) updateIssueAssignee(issueAssigneeID int64, issueID int64, assigneeID int64) error {
	if err := accessor.db.Model(&IssueAssignee{}).
		Where("id=?", issueAssigneeID).
		Updates(map[string]interface{}{"issue_id": issueID, "assignee_id": assigneeID}).
		Error; err != nil {

		return errors.Wrapf(err, "updating issue %d/assignee %d", issueID, assigneeID)
	}

	log.Debug("updated assignee %d for issue %d (id %d)", assigneeID, issueID, issueAssigneeID)

	return nil
}

// insertIssueAssignee adds a new assignee to a Gitea issue
func (accessor *DefaultAccessor) insertIssueAssignee(issueID int64, assigneeID int64) error {
	issueAssignee := IssueAssignee{IssueID: issueID, AssigneeId: assigneeID}

	if err := accessor.db.Create(&issueAssignee).Error; err != nil {
		err = errors.Wrapf(err, "adding user %d as assignee for issue id %d", assigneeID, issueID)
		return err
	}

	log.Debug("added assignee %d for issue %d", assigneeID, issueID)

	return nil
}

// AddIssueAssignee adds an assignee to a Gitea issue
func (accessor *DefaultAccessor) AddIssueAssignee(issueID int64, assigneeID int64) error {
	issueAssigneeID, err := accessor.getIssueAssigneeID(issueID, assigneeID)
	if err != nil {
		return err
	}

	if issueAssigneeID == NullID {
		return accessor.insertIssueAssignee(issueID, assigneeID)
	}

	if accessor.overwrite {
		err = accessor.updateIssueAssignee(issueAssigneeID, issueID, assigneeID)
		if err != nil {
			return err
		}
	} else {
		log.Debug("issue %d already has assignee %d - ignored", issueID, assigneeID)
	}

	return nil
}

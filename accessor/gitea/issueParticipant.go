// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"github.com/pkg/errors"
	"github.com/stevejefferson/trac2gitea/log"
	"gorm.io/gorm"
)

// getIssueParticipantID retrieves the id of the given issue participant, returns NullID if no such participant
func (accessor *DefaultAccessor) getIssueParticipantID(issueID int64, userID int64) (int64, error) {
	var id = NullID
	err := accessor.db.Model(&IssueUser{}).
		Where("issue_id=? AND uid=?", issueID, userID).
		Limit(1).
		Pluck("id", &id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrapf(err, "retrieving id for participant %d in issue %d", userID, issueID)
		return NullID, err
	}

	return id, nil
}

// updateIssueParticipant updates an existing issue participant
func (accessor *DefaultAccessor) updateIssueParticipant(issueParticipantID int64, issueID int64, userID int64) error {
	if err := accessor.db.Model(&IssueUser{}).
		Where("id=?", issueParticipantID).
		Updates(map[string]interface{}{"issue_id": issueID, "uid": userID}).
		Error; err != nil {

		return errors.Wrapf(err, "updating participant %d in issue %d", userID, issueID)
	}

	log.Debug("updated participant %d in issue %d (id %d)", userID, issueID, issueParticipantID)

	return nil
}

// insertIssueParticipant creates a new issue participant
func (accessor *DefaultAccessor) insertIssueParticipant(issueID int64, userID int64) error {
	issueUser := IssueUser{IssueId: issueID, UserId: userID, IsRead: true, IsMentioned: false}

	if err := accessor.db.Create(&issueUser).Error; err != nil {
		err = errors.Wrapf(err, "adding participant %d in issue %d", userID, issueID)
		return err
	}

	log.Debug("added participant %d in issue %d", userID, issueID)

	return nil
}

// AddIssueParticipant adds a participant to a Gitea issue.
func (accessor *DefaultAccessor) AddIssueParticipant(issueID int64, userID int64) error {
	issueParticipantID, err := accessor.getIssueParticipantID(issueID, userID)
	if err != nil {
		return err
	}

	if issueParticipantID == NullID {
		return accessor.insertIssueParticipant(issueID, userID)
	}

	if accessor.overwrite {
		err = accessor.updateIssueParticipant(issueParticipantID, issueID, userID)
		if err != nil {
			return err
		}
	} else {
		log.Debug("issue %d already has participant %d - ignored", issueID, userID)
	}

	return nil
}

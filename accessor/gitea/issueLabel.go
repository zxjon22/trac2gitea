// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"github.com/pkg/errors"
	"github.com/stevejefferson/trac2gitea/log"
	"gorm.io/gorm"
)

// getIssueLabelID retrieves the id of the given Gitea issue label - returns NullID if no such issue label.
func (accessor *DefaultAccessor) getIssueLabelID(issueID int64, labelID int64) (int64, error) {
	var id int64 = NullID
	err := accessor.db.Model(&IssueLabel{}).
		Where("issue_id=? AND label_id=?", issueID, labelID).
		Limit(1).
		Pluck("id", &id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrapf(err, "retrieving id of issue label for issue %d, label %d", issueID, labelID)
		return NullID, err
	}

	return id, nil
}

// insertIssueLabel adds a new label to a Gitea issue, returns id of created issue label.
func (accessor *DefaultAccessor) insertIssueLabel(issueID int64, labelID int64) (int64, error) {
	issueLabel := IssueLabel{IssueID: issueID, LabelID: labelID}

	if err := accessor.db.Create(&issueLabel).Error; err != nil {
		err = errors.Wrapf(err, "adding issue label for issue %d, label %d", issueID, labelID)
		return NullID, err
	}

	log.Debug("added label %d for issue %d (id %d)", labelID, issueID, issueLabel.LabelID)

	return issueLabel.ID, nil
}

// AddIssueLabel adds an issue label to Gitea, returns issue label ID
func (accessor *DefaultAccessor) AddIssueLabel(issueID int64, labelID int64) (int64, error) {
	issueLabelID, err := accessor.getIssueLabelID(issueID, labelID)
	if err != nil {
		return NullID, err
	}

	if issueLabelID == NullID {
		return accessor.insertIssueLabel(issueID, labelID)
	}

	// association between issue_id and label_id already exists - nothing to do
	return issueLabelID, nil
}

// UpdateLabelIssueCounts updates issue counts for all labels.
func (accessor *DefaultAccessor) UpdateLabelIssueCounts() error {
	err := accessor.db.Exec(`
		UPDATE label AS l SET
			num_issues = (
				SELECT COUNT(il1.issue_id)
				FROM issue_label il1
				WHERE l.id = il1.label_id
				GROUP BY il1.label_id),
			num_closed_issues = (
				SELECT COUNT(il2.issue_id)
				FROM issue_label il2, issue i
				WHERE l.id = il2.label_id
				AND il2.issue_id = i.id
				AND i.is_closed = 1
				GROUP BY il2.label_id)
		WHERE l.repo_id=?`, accessor.repoID).Error

	if err == nil {
		err = accessor.db.Exec(`
			UPDATE label SET
			num_issues = COALESCE(num_issues,0),
			num_closed_issues = COALESCE(num_closed_issues,0)
			WHERE repo_id=?`, accessor.repoID).Error
	}

	if err != nil {
		err = errors.Wrapf(err, "updating number of issues for labels")
		return err
	}

	return nil
}

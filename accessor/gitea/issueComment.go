// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/stevejefferson/trac2gitea/log"
)

// GetIssueCommentIDsByTime retrieves the IDs of all comments created at a given time for a given issue
func (accessor *DefaultAccessor) GetIssueCommentIDsByTime(issueID int64, createdTime int64) ([]int64, error) {
	var commentIDs = []int64{}
	err := accessor.db.Model(&IssueComment{}).
		Select("id").
		Where("issue_id=? AND created_unix=?", issueID, createdTime).
		Find(&commentIDs).
		Error

	if err != nil {
		err = errors.Wrapf(err, "retrieving ids of comments created at \"%s\" for issue %d", time.Unix(createdTime, 0), issueID)
		return []int64{}, err
	}

	return commentIDs, nil
}

// updateIssueComment updates an existing issue comment
func (accessor *DefaultAccessor) updateIssueComment(issueCommentID int64, issueID int64, comment *IssueComment) error {
	comment.ID = issueCommentID
	comment.IssueID = issueID
	comment.CreatedTime = comment.Time

	if err := accessor.db.Save(&comment).Error; err != nil {
		return errors.Wrapf(err, "updating comment on issue %d timed at %s", issueID, time.Unix(comment.Time, 0))
	}

	log.Debug("updated issue comment at %s for issue %d (id %d)", time.Unix(comment.Time, 0), issueID, issueCommentID)

	return nil
}

// insertIssueComment adds a new comment to a Gitea issue, returns id of created comment.
func (accessor *DefaultAccessor) insertIssueComment(issueID int64, comment *IssueComment) (int64, error) {
	comment.IssueID = issueID
	comment.CreatedTime = comment.Time

	if err := accessor.db.Create(&comment).Error; err != nil {
		err = errors.Wrapf(err, "adding comment \"%s\" for issue %d", comment.Text, issueID)
		return NullID, err
	}

	log.Debug("added issue comment at %s for issue %d (id %d)", time.Unix(comment.Time, 0), issueID, comment.ID)

	return comment.ID, nil
}

var prevIssueID = NullID
var prevCommentTime = int64(0)
var issueCommentIDIndex = 0
var issueCommentIDs = []int64{}

// AddIssueComment adds a comment on a Gitea issue, returns id of created comment
func (accessor *DefaultAccessor) AddIssueComment(issueID int64, comment *IssueComment) (int64, error) {
	var err error

	// HACK:
	// Timestamps associated with Gitea comments are not necessarily unique for comments originating from Trac
	// because Trac stores timestamps to a greater precision which we have to round to Gitea's precision.
	// Unfortunately timestamp is the best key we have for identifying whether a particular issue comment already exists
	// (and hence whether we need to insert or update it).
	// We get round this by observing that comments are always added consecutively for a given issue so we can
	// cache all comment IDs for our current issue and timestamp and extract the subsequent entries from that list.
	if issueID != prevIssueID || comment.Time != prevCommentTime {
		prevIssueID = issueID
		prevCommentTime = comment.Time
		issueCommentIDIndex = 0
		issueCommentIDs, err = accessor.GetIssueCommentIDsByTime(issueID, comment.Time)
		if err != nil {
			return NullID, err
		}
	}

	if issueCommentIDIndex >= len(issueCommentIDs) {
		// should only happen where no issue comments for timestamp
		return accessor.insertIssueComment(issueID, comment)
	}

	issueCommentID := issueCommentIDs[issueCommentIDIndex]
	issueCommentIDIndex++

	if accessor.overwrite {
		err := accessor.updateIssueComment(issueCommentID, issueID, comment)
		if err != nil {
			return NullID, err
		}
	} else {
		log.Debug("issue %d already has comment timed at %s - ignored", issueID, time.Unix(comment.Time, 0))
	}

	return issueCommentID, nil
}

// GetIssueCommentURL retrieves the URL for viewing a Gitea comment for a given issue.
func (accessor *DefaultAccessor) GetIssueCommentURL(issueID int64, commentID int64) string {
	repoURL := accessor.getUserRepoURL()
	return fmt.Sprintf("%s/issues/%d#issuecomment-%d", repoURL, issueID, commentID)
}

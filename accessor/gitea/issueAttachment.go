// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/stevejefferson/trac2gitea/log"
	"gorm.io/gorm"
)

// getIssueAttachmentIDandUUID retrieves the id and UUID of the given issue attachment, returns id of gitea.NullID if no such attachment
func (accessor *DefaultAccessor) getIssueAttachmentIDandUUID(issueID int64, fileName string) (int64, string, error) {
	issueAttachment := IssueAttachment{ID: NullID}

	err := accessor.db.
		Where("issue_id=? AND name=?", issueID, fileName).
		First(&issueAttachment).
		Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrapf(err, "retrieving id for attachment %s for issue %d", fileName, issueID)
		return NullID, "", err
	}

	return issueAttachment.ID, issueAttachment.UUID, nil
}

// GetIssueAttachmentUUID returns the UUID for a named attachment of a given issue - returns empty string if cannot find issue/attachment.
func (accessor *DefaultAccessor) GetIssueAttachmentUUID(issueID int64, fileName string) (string, error) {
	_, uuid, err := accessor.getIssueAttachmentIDandUUID(issueID, fileName)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

// getAttachmentPath returns the path at which to store an attachment with a given UUID
func (accessor *DefaultAccessor) getAttachmentPath(UUID string) (string, error) {
	attachmentsRootDir := accessor.GetStringConfig("attachment", "PATH")
	if attachmentsRootDir == "" {
		attachmentsRootDir = filepath.Join(accessor.rootDir, "data", "attachments")
	}

	if _, err := os.Stat(attachmentsRootDir); os.IsNotExist(err) {
		return "", errors.Wrap(err, "Attachments folder does not exist - aborting")
	}

	d1 := UUID[0:1]
	d2 := UUID[1:2]
	dir := filepath.Join(attachmentsRootDir, d1, d2)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(dir, UUID), nil
}

// copyAttachment copies a given attachment file to the Gitea attachment with the given UUID
func (accessor *DefaultAccessor) copyAttachment(filePath string, UUID string) error {
	attachmentPath, err := accessor.getAttachmentPath(UUID)

	if err != nil {
		return err
	}

	return copyFile(filePath, attachmentPath)
}

// deleteAttachment deletes the Gitea attachment with the given UUID
func (accessor *DefaultAccessor) deleteAttachment(UUID string) error {
	attachmentPath, err := accessor.getAttachmentPath(UUID)

	if err != nil {
		return err
	}

	return deleteFile(attachmentPath)
}

// updateIssueAttachment updates an existing issue attachment
func (accessor *DefaultAccessor) updateIssueAttachment(issueAttachmentID int64, issueID int64, attachment *IssueAttachment, filePath string) error {
	attachment.ID = issueAttachmentID
	attachment.IssueID = issueID

	if err := accessor.db.Save(&attachment).Error; err != nil {
		return errors.Wrapf(err, "updating attachment %s for issue %d", attachment.FileName, issueID)
	}

	log.Debug("updated attachment %s for issue %d (id %d)", attachment.UUID, issueID, issueAttachmentID)

	return nil
}

// insertIssueAttachment creates a new attachment to a Gitea issue, returns id of created attachment
func (accessor *DefaultAccessor) insertIssueAttachment(issueID int64, attachment *IssueAttachment, filePath string) (int64, error) {
	attachment.IssueID = issueID

	if err := accessor.db.Create(&attachment).Error; err != nil {
		err = errors.Wrapf(err, "adding attachment %s for issue %d", attachment.FileName, issueID)
		return NullID, err
	}

	log.Debug("added attachment %s for issue %d", attachment.FileName, issueID)

	return attachment.ID, nil
}

// AddIssueAttachment adds a new attachment to an issue using the provided file - returns id of created attachment
func (accessor *DefaultAccessor) AddIssueAttachment(issueID int64, attachment *IssueAttachment, filePath string) (int64, error) {
	issueAttachmentID, issueAttachmentUUID, err := accessor.getIssueAttachmentIDandUUID(issueID, attachment.FileName)
	if err != nil {
		return NullID, err
	}

	if issueAttachmentID == NullID {
		issueAttachmentID, err = accessor.insertIssueAttachment(issueID, attachment, filePath)
		if err != nil {
			return NullID, err
		}
	} else if accessor.overwrite {
		err = accessor.updateIssueAttachment(issueAttachmentID, issueID, attachment, filePath)
		if err != nil {
			return NullID, err
		}

		err = accessor.deleteAttachment(issueAttachmentUUID)
	} else {
		if attachment.UUID != issueAttachmentUUID {
			log.Warn("attachment %s already exists for issue %d but under UUID %s (expecting UUID %s)", attachment.FileName, issueID, issueAttachmentUUID, attachment.UUID)
		} else {
			log.Debug("issue %d already has attachment %s - ignored", issueID, attachment.FileName)
		}
		return issueAttachmentID, nil
	}

	err = accessor.copyAttachment(filePath, attachment.UUID)
	if err != nil {
		return NullID, err
	}

	return issueAttachmentID, nil
}

// GetIssueAttachmentURL retrieves the URL for viewing a Gitea attachment
func (accessor *DefaultAccessor) GetIssueAttachmentURL(issueID int64, uuid string) string {
	baseURL := accessor.getUserRepoURL()
	return fmt.Sprintf("%s/attachments/%s", baseURL, uuid)
}

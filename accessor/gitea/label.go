// Copyright 2020 Steve Jefferson. All rights reserved.
// Use of this source code is governed by a GPL-style
// license that can be found in the LICENSE file.

package gitea

import (
	"github.com/pkg/errors"
	"github.com/stevejefferson/trac2gitea/log"
	"gorm.io/gorm"
)

// GetLabelID retrieves the id of the given label, returns NullID if no such label
func (accessor *DefaultAccessor) GetLabelID(labelName string) (int64, error) {
	var id int64 = NullID
	err := accessor.db.Model(&Label{}).
		Where("repo_id=? AND name=?", accessor.repoID, labelName).
		Limit(1).
		Pluck("id", &id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		err = errors.Wrapf(err, "retrieving id of label %s", labelName)
		return NullID, err
	}

	return id, nil
}

// updateLabel updates an existing label
func (accessor *DefaultAccessor) updateLabel(labelID int64, label *Label) error {
	label.ID = labelID
	label.RepoId = accessor.repoID

	if err := accessor.db.Save(&label).Error; err != nil {
		return errors.Wrapf(err, "updating label %s", label.Name)
	}

	log.Debug("updated label %s, color %s (id %d)", label.Name, label.Color, labelID)

	return nil
}

// insertLabel inserts a new label, returns label id.
func (accessor *DefaultAccessor) insertLabel(label *Label) (int64, error) {
	label.RepoId = accessor.repoID

	if err := accessor.db.Create(&label).Error; err != nil {
		err = errors.Wrapf(err, "adding label %s", label.Name)
		return NullID, err
	}

	log.Debug("added label %s, color %s (id %d)", label.Name, label.Color, label.ID)

	return label.ID, nil
}

// AddLabel adds a label to Gitea, returns label id.
func (accessor *DefaultAccessor) AddLabel(label *Label) (int64, error) {
	labelID, err := accessor.GetLabelID(label.Name)
	if err != nil {
		return NullID, err
	}

	if labelID == NullID {
		return accessor.insertLabel(label)
	}

	if accessor.overwrite {
		err = accessor.updateLabel(labelID, label)
		if err != nil {
			return NullID, err
		}
	} else {
		log.Debug("label %s already exists - ignored", label.Name)
	}

	return labelID, nil
}

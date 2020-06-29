// Copyright 2020 Google LLC. All Rights Reserved.

package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	rpc "apigov.dev/registry/rpc"
	ptypes "github.com/golang/protobuf/ptypes"
)

// LabelEntityName is used to represent labels in the datastore.
const LabelEntityName = "Label"

// LabelsRegexp returns a regular expression that matches collection of labels.
func LabelsRegexp() *regexp.Regexp {
	return regexp.MustCompile(
		"^projects/" + nameRegex +
			"(/products/" + nameRegex +
			"(/versions/" + nameRegex +
			"(/specs/" + nameRegex +
			")?" +
			")?" +
			")?" +
			"/labels$")
}

// LabelRegexp returns a regular expression that matches a Label resource name.
func LabelRegexp() *regexp.Regexp {
	return regexp.MustCompile("^projects/" + nameRegex +
		"(/products/" + nameRegex +
		"(/versions/" + nameRegex +
		"(/specs/" + nameRegex +
		")?" +
		")?" +
		")?" +
		"/labels/" + nameRegex + "$")
}

// Label ...
type Label struct {
	ProjectID  string    // Project associated with label (required).
	ProductID  string    // Product associated with label (if appropriate).
	VersionID  string    // Version associated with label (if appropriate).
	SpecID     string    // Spec associated with label (if appropriate).
	LabelID    string    // Label identifier (required).
	CreateTime time.Time // Creation time.
	UpdateTime time.Time // Time of last change.
	Subject    string    // Subject of the label.
}

// NewLabelFromParentAndLabelID returns an initialized label for a specified parent and labelID.
func NewLabelFromParentAndLabelID(parent string, labelID string) (*Label, error) {
	// Return an error if the labelID is invalid.
	if err := validateID(labelID); err != nil {
		return nil, err
	}
	// Match regular expressions to identify the parent of this label.
	var m [][]string
	// Is the parent a project?
	m = ProjectRegexp().FindAllStringSubmatch(parent, -1)
	if m != nil {
		return &Label{
			ProjectID: m[0][1],
			LabelID:   labelID,
			Subject:   parent,
		}, nil
	}
	// Is the parent a product?
	m = ProductRegexp().FindAllStringSubmatch(parent, -1)
	if m != nil {
		return &Label{
			ProjectID: m[0][1],
			ProductID: m[0][2],
			LabelID:   labelID,
			Subject:   parent,
		}, nil
	}
	// Is the parent a version?
	m = VersionRegexp().FindAllStringSubmatch(parent, -1)
	if m != nil {
		return &Label{
			ProjectID: m[0][1],
			ProductID: m[0][2],
			VersionID: m[0][3],
			LabelID:   labelID,
			Subject:   parent,
		}, nil
	}
	// Is the parent a spec?
	m = SpecRegexp().FindAllStringSubmatch(parent, -1)
	if m != nil {
		return &Label{
			ProjectID: m[0][1],
			ProductID: m[0][2],
			VersionID: m[0][3],
			SpecID:    m[0][4],
			LabelID:   labelID,
			Subject:   parent,
		}, nil
	}
	// Return an error for an unrecognized parent.
	return nil, fmt.Errorf("invalid parent '%s'", parent)
}

// NewLabelFromResourceName parses resource names and returns an initialized label.
func NewLabelFromResourceName(name string) (*Label, error) {
	// split name into parts
	parts := strings.Split(name, "/")
	if parts[len(parts)-2] != "labels" {
		return nil, fmt.Errorf("invalid label name '%s'", name)
	}
	// build label from parent and labelID
	parent := strings.Join(parts[0:len(parts)-2], "/")
	labelID := parts[len(parts)-1]
	return NewLabelFromParentAndLabelID(parent, labelID)
}

// ResourceName generates the resource name of a label.
func (label *Label) ResourceName() string {
	switch {
	case label.SpecID != "":
		return fmt.Sprintf("projects/%s/products/%s/versions/%s/specs/%s/labels/%s",
			label.ProjectID, label.ProductID, label.VersionID, label.SpecID, label.LabelID)
	case label.VersionID != "":
		return fmt.Sprintf("projects/%s/products/%s/versions/%s/labels/%s",
			label.ProjectID, label.ProductID, label.VersionID, label.LabelID)
	case label.ProductID != "":
		return fmt.Sprintf("projects/%s/products/%s/labels/%s",
			label.ProjectID, label.ProductID, label.LabelID)
	case label.ProjectID != "":
		return fmt.Sprintf("projects/%s/labels/%s",
			label.ProjectID, label.LabelID)
	default:
		return "UNKNOWN"
	}
}

// Message returns a message representing a label.
func (label *Label) Message() (message *rpc.Label, err error) {
	message = &rpc.Label{}
	message.Name = label.ResourceName()
	message.Subject = label.Subject
	message.CreateTime, err = ptypes.TimestampProto(label.CreateTime)
	message.UpdateTime, err = ptypes.TimestampProto(label.UpdateTime)
	return message, err
}

// Update modifies a label using the contents of a message.
func (label *Label) Update(message *rpc.Label) error {
	label.Subject = message.GetSubject()
	label.UpdateTime = label.CreateTime
	return nil
}
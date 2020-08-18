// Code generated by MockGen. DO NOT EDIT.
// Source: stevejefferson.co.uk/trac2gitea/accessor/gitea (interfaces: Accessor)

// Package mock_gitea is a generated GoMock package.
package mock_gitea

import (
	sql "database/sql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccessor is a mock of Accessor interface
type MockAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockAccessorMockRecorder
}

// MockAccessorMockRecorder is the mock recorder for MockAccessor
type MockAccessorMockRecorder struct {
	mock *MockAccessor
}

// NewMockAccessor creates a new mock instance
func NewMockAccessor(ctrl *gomock.Controller) *MockAccessor {
	mock := &MockAccessor{ctrl: ctrl}
	mock.recorder = &MockAccessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccessor) EXPECT() *MockAccessorMockRecorder {
	return m.recorder
}

// AddAttachment mocks base method
func (m *MockAccessor) AddAttachment(arg0 string, arg1, arg2 int64, arg3, arg4 string, arg5 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddAttachment", arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddAttachment indicates an expected call of AddAttachment
func (mr *MockAccessorMockRecorder) AddAttachment(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAttachment", reflect.TypeOf((*MockAccessor)(nil).AddAttachment), arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddComment mocks base method
func (m *MockAccessor) AddComment(arg0, arg1 int64, arg2 string, arg3 int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int64)
	return ret0
}

// AddComment indicates an expected call of AddComment
func (mr *MockAccessorMockRecorder) AddComment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockAccessor)(nil).AddComment), arg0, arg1, arg2, arg3)
}

// AddIssue mocks base method
func (m *MockAccessor) AddIssue(arg0 int64, arg1 string, arg2 int64, arg3 string, arg4 sql.NullString, arg5 string, arg6 bool, arg7 string, arg8 int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIssue", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	ret0, _ := ret[0].(int64)
	return ret0
}

// AddIssue indicates an expected call of AddIssue
func (mr *MockAccessorMockRecorder) AddIssue(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIssue", reflect.TypeOf((*MockAccessor)(nil).AddIssue), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// AddIssueLabel mocks base method
func (m *MockAccessor) AddIssueLabel(arg0 int64, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddIssueLabel", arg0, arg1)
}

// AddIssueLabel indicates an expected call of AddIssueLabel
func (mr *MockAccessorMockRecorder) AddIssueLabel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIssueLabel", reflect.TypeOf((*MockAccessor)(nil).AddIssueLabel), arg0, arg1)
}

// AddLabel mocks base method
func (m *MockAccessor) AddLabel(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddLabel", arg0, arg1)
}

// AddLabel indicates an expected call of AddLabel
func (mr *MockAccessorMockRecorder) AddLabel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLabel", reflect.TypeOf((*MockAccessor)(nil).AddLabel), arg0, arg1)
}

// AddMilestone mocks base method
func (m *MockAccessor) AddMilestone(arg0, arg1 string, arg2 bool, arg3, arg4 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddMilestone", arg0, arg1, arg2, arg3, arg4)
}

// AddMilestone indicates an expected call of AddMilestone
func (mr *MockAccessorMockRecorder) AddMilestone(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMilestone", reflect.TypeOf((*MockAccessor)(nil).AddMilestone), arg0, arg1, arg2, arg3, arg4)
}

// CloneWiki mocks base method
func (m *MockAccessor) CloneWiki() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloneWiki")
}

// CloneWiki indicates an expected call of CloneWiki
func (mr *MockAccessorMockRecorder) CloneWiki() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneWiki", reflect.TypeOf((*MockAccessor)(nil).CloneWiki))
}

// CopyFileToWiki mocks base method
func (m *MockAccessor) CopyFileToWiki(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CopyFileToWiki", arg0, arg1)
}

// CopyFileToWiki indicates an expected call of CopyFileToWiki
func (mr *MockAccessorMockRecorder) CopyFileToWiki(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFileToWiki", reflect.TypeOf((*MockAccessor)(nil).CopyFileToWiki), arg0, arg1)
}

// GetAttachmentURL mocks base method
func (m *MockAccessor) GetAttachmentURL(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttachmentURL", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAttachmentURL indicates an expected call of GetAttachmentURL
func (mr *MockAccessorMockRecorder) GetAttachmentURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachmentURL", reflect.TypeOf((*MockAccessor)(nil).GetAttachmentURL), arg0)
}

// GetAttachmentUUID mocks base method
func (m *MockAccessor) GetAttachmentUUID(arg0 int64, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttachmentUUID", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAttachmentUUID indicates an expected call of GetAttachmentUUID
func (mr *MockAccessorMockRecorder) GetAttachmentUUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachmentUUID", reflect.TypeOf((*MockAccessor)(nil).GetAttachmentUUID), arg0, arg1)
}

// GetCommentID mocks base method
func (m *MockAccessor) GetCommentID(arg0 int64, arg1 string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentID", arg0, arg1)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetCommentID indicates an expected call of GetCommentID
func (mr *MockAccessorMockRecorder) GetCommentID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentID", reflect.TypeOf((*MockAccessor)(nil).GetCommentID), arg0, arg1)
}

// GetCommentURL mocks base method
func (m *MockAccessor) GetCommentURL(arg0, arg1 int64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentURL", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCommentURL indicates an expected call of GetCommentURL
func (mr *MockAccessorMockRecorder) GetCommentURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentURL", reflect.TypeOf((*MockAccessor)(nil).GetCommentURL), arg0, arg1)
}

// GetCommitURL mocks base method
func (m *MockAccessor) GetCommitURL(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitURL", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCommitURL indicates an expected call of GetCommitURL
func (mr *MockAccessorMockRecorder) GetCommitURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitURL", reflect.TypeOf((*MockAccessor)(nil).GetCommitURL), arg0)
}

// GetDefaultAssigneeID mocks base method
func (m *MockAccessor) GetDefaultAssigneeID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultAssigneeID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetDefaultAssigneeID indicates an expected call of GetDefaultAssigneeID
func (mr *MockAccessorMockRecorder) GetDefaultAssigneeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultAssigneeID", reflect.TypeOf((*MockAccessor)(nil).GetDefaultAssigneeID))
}

// GetDefaultAuthorID mocks base method
func (m *MockAccessor) GetDefaultAuthorID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultAuthorID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetDefaultAuthorID indicates an expected call of GetDefaultAuthorID
func (mr *MockAccessorMockRecorder) GetDefaultAuthorID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultAuthorID", reflect.TypeOf((*MockAccessor)(nil).GetDefaultAuthorID))
}

// GetIssueID mocks base method
func (m *MockAccessor) GetIssueID(arg0 int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueID", arg0)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetIssueID indicates an expected call of GetIssueID
func (mr *MockAccessorMockRecorder) GetIssueID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueID", reflect.TypeOf((*MockAccessor)(nil).GetIssueID), arg0)
}

// GetIssueURL mocks base method
func (m *MockAccessor) GetIssueURL(arg0 int64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssueURL", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetIssueURL indicates an expected call of GetIssueURL
func (mr *MockAccessorMockRecorder) GetIssueURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssueURL", reflect.TypeOf((*MockAccessor)(nil).GetIssueURL), arg0)
}

// GetMilestoneID mocks base method
func (m *MockAccessor) GetMilestoneID(arg0 string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMilestoneID", arg0)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetMilestoneID indicates an expected call of GetMilestoneID
func (mr *MockAccessorMockRecorder) GetMilestoneID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMilestoneID", reflect.TypeOf((*MockAccessor)(nil).GetMilestoneID), arg0)
}

// GetMilestoneURL mocks base method
func (m *MockAccessor) GetMilestoneURL(arg0 int64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMilestoneURL", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMilestoneURL indicates an expected call of GetMilestoneURL
func (mr *MockAccessorMockRecorder) GetMilestoneURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMilestoneURL", reflect.TypeOf((*MockAccessor)(nil).GetMilestoneURL), arg0)
}

// GetSourceURL mocks base method
func (m *MockAccessor) GetSourceURL(arg0, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSourceURL", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourceURL indicates an expected call of GetSourceURL
func (mr *MockAccessorMockRecorder) GetSourceURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceURL", reflect.TypeOf((*MockAccessor)(nil).GetSourceURL), arg0, arg1)
}

// GetStringConfig mocks base method
func (m *MockAccessor) GetStringConfig(arg0, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringConfig", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStringConfig indicates an expected call of GetStringConfig
func (mr *MockAccessorMockRecorder) GetStringConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringConfig", reflect.TypeOf((*MockAccessor)(nil).GetStringConfig), arg0, arg1)
}

// GetUserEMailAddress mocks base method
func (m *MockAccessor) GetUserEMailAddress(arg0 int64) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserEMailAddress", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUserEMailAddress indicates an expected call of GetUserEMailAddress
func (mr *MockAccessorMockRecorder) GetUserEMailAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserEMailAddress", reflect.TypeOf((*MockAccessor)(nil).GetUserEMailAddress), arg0)
}

// GetUserID mocks base method
func (m *MockAccessor) GetUserID(arg0 string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserID", arg0)
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetUserID indicates an expected call of GetUserID
func (mr *MockAccessorMockRecorder) GetUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserID", reflect.TypeOf((*MockAccessor)(nil).GetUserID), arg0)
}

// GetWikiAttachmentRelPath mocks base method
func (m *MockAccessor) GetWikiAttachmentRelPath(arg0, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWikiAttachmentRelPath", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWikiAttachmentRelPath indicates an expected call of GetWikiAttachmentRelPath
func (mr *MockAccessorMockRecorder) GetWikiAttachmentRelPath(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWikiAttachmentRelPath", reflect.TypeOf((*MockAccessor)(nil).GetWikiAttachmentRelPath), arg0, arg1)
}

// GetWikiFileURL mocks base method
func (m *MockAccessor) GetWikiFileURL(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWikiFileURL", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWikiFileURL indicates an expected call of GetWikiFileURL
func (mr *MockAccessorMockRecorder) GetWikiFileURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWikiFileURL", reflect.TypeOf((*MockAccessor)(nil).GetWikiFileURL), arg0)
}

// GetWikiHtdocRelPath mocks base method
func (m *MockAccessor) GetWikiHtdocRelPath(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWikiHtdocRelPath", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWikiHtdocRelPath indicates an expected call of GetWikiHtdocRelPath
func (mr *MockAccessorMockRecorder) GetWikiHtdocRelPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWikiHtdocRelPath", reflect.TypeOf((*MockAccessor)(nil).GetWikiHtdocRelPath), arg0)
}

// SetIssueUpdateTime mocks base method
func (m *MockAccessor) SetIssueUpdateTime(arg0, arg1 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIssueUpdateTime", arg0, arg1)
}

// SetIssueUpdateTime indicates an expected call of SetIssueUpdateTime
func (mr *MockAccessorMockRecorder) SetIssueUpdateTime(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIssueUpdateTime", reflect.TypeOf((*MockAccessor)(nil).SetIssueUpdateTime), arg0, arg1)
}

// TranslateWikiPageName mocks base method
func (m *MockAccessor) TranslateWikiPageName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TranslateWikiPageName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// TranslateWikiPageName indicates an expected call of TranslateWikiPageName
func (mr *MockAccessorMockRecorder) TranslateWikiPageName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TranslateWikiPageName", reflect.TypeOf((*MockAccessor)(nil).TranslateWikiPageName), arg0)
}

// UpdateRepoIssueCount mocks base method
func (m *MockAccessor) UpdateRepoIssueCount(arg0, arg1 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateRepoIssueCount", arg0, arg1)
}

// UpdateRepoIssueCount indicates an expected call of UpdateRepoIssueCount
func (mr *MockAccessorMockRecorder) UpdateRepoIssueCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRepoIssueCount", reflect.TypeOf((*MockAccessor)(nil).UpdateRepoIssueCount), arg0, arg1)
}

// WikiCommit mocks base method
func (m *MockAccessor) WikiCommit(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WikiCommit", arg0, arg1, arg2)
}

// WikiCommit indicates an expected call of WikiCommit
func (mr *MockAccessorMockRecorder) WikiCommit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WikiCommit", reflect.TypeOf((*MockAccessor)(nil).WikiCommit), arg0, arg1, arg2)
}

// WikiComplete mocks base method
func (m *MockAccessor) WikiComplete() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WikiComplete")
}

// WikiComplete indicates an expected call of WikiComplete
func (mr *MockAccessorMockRecorder) WikiComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WikiComplete", reflect.TypeOf((*MockAccessor)(nil).WikiComplete))
}

// WriteWikiPage mocks base method
func (m *MockAccessor) WriteWikiPage(arg0, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteWikiPage", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// WriteWikiPage indicates an expected call of WriteWikiPage
func (mr *MockAccessorMockRecorder) WriteWikiPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteWikiPage", reflect.TypeOf((*MockAccessor)(nil).WriteWikiPage), arg0, arg1)
}

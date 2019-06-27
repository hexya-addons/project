package project

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/pool/h"
)

//vars

var (
)


//rights
func init() {
	h.ProjectProject().Methods().Load().AllowGroup(GroupProjectUser)
	h.ProjectProject().Methods().AllowAllToGroup(GroupProjectManager)
	h.Analytic.ModelAccountAnalyticAccount().Methods().Load().AllowGroup(GroupProjectUser)
	h.Analytic.ModelAccountAnalyticAccount().Methods().AllowAllToGroup(GroupProjectManager)
	h.ProjectTaskType().Methods().Load().AllowGroup(base.GroupUser)
	h.ProjectTaskType().Methods().Load().AllowGroup(GroupProjectUser)
	h.ProjectTaskType().Methods().AllowAllToGroup(GroupProjectManager)
	h.ProjectTask().Methods().AllowAllToGroup(GroupProjectUser)
	h.ReportProjectTaskUser().Methods().AllowAllToGroup(GroupProjectManager)
	h.Base.ModelResPartner().Methods().Load().AllowGroup(GroupProjectUser)
	h.ProjectTask().Methods().Load().AllowGroup(base.GroupUser)
	h.ProjectProject().Methods().Load().AllowGroup(base.GroupUser)
	h.Analytic.ModelAccountAnalyticLine().Methods().AllowAllToGroup(GroupProjectManager)
	h.Resource.ModelResourceCalendar().Methods().Load().AllowGroup(GroupProjectUser)
	h.Resource.ModelResourceCalendarAttendance().Methods().Load().AllowGroup(GroupProjectUser)
	h.Resource.ModelResourceCalendarLeaves().Methods().AllowAllToGroup(GroupProjectUser)
	h.ProjectTags().Methods().Load().AllowGroup(security.GroupEveryone)
	h.ProjectTags().Methods().AllowAllToGroup(GroupProjectManager)
	h.Mail.ModelMailAlias().Methods().AllowAllToGroup(GroupProjectManager)
}

package project

import (
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.WebPlanner().DeclareModel()

	h.WebPlanner().Methods().GetPlannerApplication().DeclareMethod(
		`GetPlannerApplication`,
		func(rs m.WebPlannerSet) {
			//        planner = super(PlannerProject, self)._get_planner_application()
			//        planner.append(['planner_project', 'Project Planner'])
			//        return planner
		})
	h.WebPlanner().Methods().PreparePlannerProjectData().DeclareMethod(
		`PreparePlannerProjectData`,
		func(rs m.WebPlannerSet) {
			//        return {
			//            'timesheet_menu': self.env.ref('hr_timesheet_sheet.menu_act_hr_timesheet_sheet_form_my_current', raise_if_not_found=False),
			//        }
		})
}

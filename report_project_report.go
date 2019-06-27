package project

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.ReportProjectTaskUser().DeclareModel()

	h.ReportProjectTaskUser().AddFields(map[string]models.FieldDefinition{
		"Name": models.CharField{
			String:   "Task Title",
			ReadOnly: true,
		},
		"UserId": models.Many2OneField{
			RelationModel: h.User(),
			String:        "Assigned To",
			ReadOnly:      true,
		},
		"DateStart": models.DateTimeField{
			String:   "Assignation Date",
			ReadOnly: true,
		},
		"NoOfDays": models.IntegerField{
			String:   "# Working Days",
			ReadOnly: true,
		},
		"DateEnd": models.DateTimeField{
			String:   "Ending Date",
			ReadOnly: true,
		},
		"DateDeadline": models.DateField{
			String:   "Deadline",
			ReadOnly: true,
		},
		"DateLastStageUpdate": models.DateTimeField{
			String:   "Last Stage Update",
			ReadOnly: true,
		},
		"ProjectId": models.Many2OneField{
			RelationModel: h.ProjectProject(),
			String:        "Project",
			ReadOnly:      true,
		},
		"ClosingDays": models.FloatField{
			String: "# Days to Close",
			//digits=(16, 2)
			ReadOnly: true,
			//group_operator="avg"
			Help: "Number of Days to close the task",
		},
		"OpeningDays": models.FloatField{
			String: "# Days to Assign",
			//digits=(16, 2)
			ReadOnly: true,
			//group_operator="avg"
			Help: "Number of Days to Open the task",
		},
		"DelayEndingsDays": models.FloatField{
			String: "# Days to Deadline",
			//digits=(16, 2)
			ReadOnly: true,
		},
		"Nbr": models.IntegerField{
			String:   "# of Tasks",
			ReadOnly: true,
		},
		"Priority": models.SelectionField{
			Selection: types.Selection{
				"0": "Low",
				"1": "Normal",
				"2": "High",
			},
			Size:     1,
			ReadOnly: true,
		},
		"State": models.SelectionField{
			Selection: types.Selection{
				"normal":  "In Progress",
				"blocked": "Blocked",
				"done":    "Ready for next stage",
			},
			String:   "Kanban State",
			ReadOnly: true,
		},
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			ReadOnly:      true,
		},
		"PartnerId": models.Many2OneField{
			RelationModel: h.Partner(),
			String:        "Contact",
			ReadOnly:      true,
		},
		"StageId": models.Many2OneField{
			RelationModel: h.ProjectTaskType(),
			String:        "Stage",
			ReadOnly:      true,
		},
	})
	h.ReportProjectTaskUser().Methods().Select().DeclareMethod(
		`Select`,
		func(rs m.ReportProjectTaskUserSet) {
			//        select_str = """
			//             SELECT
			//                    (select 1 ) AS nbr,
			//                    t.id as id,
			//                    t.date_start as date_start,
			//                    t.date_end as date_end,
			//                    t.date_last_stage_update as date_last_stage_update,
			//                    t.date_deadline as date_deadline,
			//                    abs((extract('epoch' from (t.write_date-t.date_start)))/(3600*24))  as no_of_days,
			//                    t.user_id,
			//                    t.project_id,
			//                    t.priority,
			//                    t.name as name,
			//                    t.company_id,
			//                    t.partner_id,
			//                    t.stage_id as stage_id,
			//                    t.kanban_state as state,
			//                    (extract('epoch' from (NULLIF(t.date_end, t.write_date)-t.create_date)))/(3600*24)  as closing_days,
			//                    (extract('epoch' from (t.date_start-t.create_date)))/(3600*24)  as opening_days,
			//                    (extract('epoch' from (t.date_deadline-(now() at time zone 'UTC'))))/(3600*24)  as delay_endings_days
			//        """
			//        return select_str
		})
	h.ReportProjectTaskUser().Methods().GroupBy().Extend(
		`GroupBy`,
		func(rs m.ReportProjectTaskUserSet) {
			//        group_by_str = """
			//                GROUP BY
			//                    t.id,
			//                    create_date,
			//                    write_date,
			//                    date_start,
			//                    date_end,
			//                    date_deadline,
			//                    date_last_stage_update,
			//                    t.user_id,
			//                    t.project_id,
			//                    t.priority,
			//                    name,
			//                    t.company_id,
			//                    t.partner_id,
			//                    stage_id
			//        """
			//        return group_by_str
		})
	h.ReportProjectTaskUser().Methods().Init().DeclareMethod(
		`Init`,
		func(rs m.ReportProjectTaskUserSet) {
			//        tools.drop_view_if_exists(self._cr, self._table)
			//        self._cr.execute("""
			//            CREATE view %s as
			//              %s
			//              FROM project_task t
			//                WHERE t.active = 'true'
			//                %s
			//        """ % (self._table, self._select(), self._group_by()))
		})
}

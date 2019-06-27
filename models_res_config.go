package project

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/pool/h"
)

func init() {
	h.ProjectConfigSettings().DeclareModel()

	h.ProjectConfigSettings().AddFields(map[string]models.FieldDefinition{
		"CompanyId": models.Many2OneField{
			RelationModel: h.Company(),
			String:        "Company",
			Required:      true,
			Default:       func(env models.Environment) interface{} { return env.Uid().company_id },
		},
		"ProjectTimeModeId": models.Many2OneField{
			Related: `CompanyId.ProjectTimeModeId`,
			String:  "Project Time Unit *",
		},
		"ModulePad": models.SelectionField{
			Selection: types.Selection{
				"": "Task description is plain text",
				"": "Collaborative rich text on task description",
			},
			String: "Pads",
			Help: "Lets the company customize which Pad installation should" +
				"be used to link to new pads (for example: http://ietherpad.com/)." +
				"-This installs the module pad.",
		},
		"ModuleRatingProject": models.SelectionField{
			Selection: types.Selection{
				"": "No customer rating",
				"": "Track customer satisfaction on tasks",
			},
			String: "Rating on task",
			Help:   "This allows customers to give rating on provided services",
		},
		"GenerateProjectAlias": models.SelectionField{
			Selection: types.Selection{
				"": "Do not create an email alias automatically",
				"": "Automatically generate an email alias at the project creation",
			},
			String: "Project Alias",
			Help: "Odoo will generate an email alias at the project creation" +
				"from project name.",
		},
		"ModuleProjectForecast": models.BooleanField{
			String: "Forecasts, planning and Gantt charts",
		},
	})
	h.ProjectConfigSettings().Methods().SetDefaultGenerateProjectAlias().DeclareMethod(
		`SetDefaultGenerateProjectAlias`,
		func(rs m.ProjectConfigSettingsSet) {
			//        check = self.env.user.has_group('base.group_system')
			//        Values = check and self.env['ir.values'].sudo(
			//        ) or self.env['ir.values']
			//        for config in self:
			//            Values.set_default('project.config.settings',
			//                               'generate_project_alias', config.generate_project_alias)
		})
}

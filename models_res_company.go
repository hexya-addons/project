package project

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.Company().AddFields(map[string]models.FieldDefinition{
		"ProjectTimeModeId": models.Many2OneField{
			RelationModel: h.ProductUom(),
			String:        "Project Time Unit",
			Help: "This will set the unit of measure used in projects and tasks." +
				"If you use the timesheet linked to projects, don't forget" +
				"to setup the right unit of measure in your employees.",
		},
	})
}

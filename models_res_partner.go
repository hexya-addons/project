package project

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.Partner().AddFields(map[string]models.FieldDefinition{
		"TaskIds": models.One2ManyField{
			RelationModel: h.ProjectTask(),
			ReverseFK:     "",
			String:        "Tasks",
		},
		"TaskCount": models.IntegerField{
			Compute: h.Partner().Methods().ComputeTaskCount(),
			String:  "# Tasks",
		},
	})
	h.Partner().Methods().ComputeTaskCount().DeclareMethod(
		`ComputeTaskCount`,
		func(rs h.PartnerSet) h.PartnerData {
			//        for partner in self:
			//            partner.task_count = len(partner.task_ids)
		})
}

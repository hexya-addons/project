package project

	import (
		"net/http"

		"github.com/hexya-erp/hexya/src/controllers"
		"github.com/hexya-erp/hexya/src/models"
		"github.com/hexya-erp/hexya/src/models/types"
		"github.com/hexya-erp/hexya/src/models/types/dates"
		"github.com/hexya-erp/pool/h"
		"github.com/hexya-erp/pool/q"
	)
	
func init() {
h.ProjectTaskType().DeclareModel()



h.ProjectTaskType().Methods().GetMailTemplateIdDomain().DeclareMethod(
`GetMailTemplateIdDomain`,
func(rs m.ProjectTaskTypeSet)  {
//        return [('model', '=', 'project.task')]
})
h.ProjectTaskType().Methods().GetDefaultProjectIds().DeclareMethod(
`GetDefaultProjectIds`,
func(rs m.ProjectTaskTypeSet)  {
//        default_project_id = self.env.context.get('default_project_id')
//        return [default_project_id] if default_project_id else None
})
h.ProjectTaskType().AddFields(map[string]models.FieldDefinition{
"Name": models.CharField{
String: "Stage Name",
Required: true,
Translate: true,
},
"Description": models.TextField{
Translate: true,
},
"Sequence": models.IntegerField{
Default: models.DefaultValue(1),
},
"ProjectIds": models.Many2ManyField{
RelationModel: h.ProjectProject(),
M2MLinkModelName: "",
M2MOurField: "",
M2MTheirField: "",
String: "Projects",
Default: models.DefaultValue(_get_default_project_ids),
},
"LegendPriority": models.CharField{
String: "Priority Management Explanation",
Translate: true,
Help: "Explanation text to help users using the star and priority" + 
"mechanism on stages or issues that are in this stage.",
},
"LegendBlocked": models.CharField{
String: "Kanban Blocked Explanation",
Translate: true,
Help: "Override the default value displayed for the blocked state" + 
"for kanban selection, when the task or issue is in that stage.",
},
"LegendDone": models.CharField{
String: "Kanban Valid Explanation",
Translate: true,
Help: "Override the default value displayed for the done state" + 
"for kanban selection, when the task or issue is in that stage.",
},
"LegendNormal": models.CharField{
String: "Kanban Ongoing Explanation",
Translate: true,
Help: "Override the default value displayed for the normal state" + 
"for kanban selection, when the task or issue is in that stage.",
},
"MailTemplateId": models.Many2OneField{
RelationModel: h.MailTemplate(),
String: "Email Template",
Filter: q.,
Help: "If set an email will be sent to the customer when the task" + 
"or issue reaches this step.",
},
"Fold": models.BooleanField{
String: "Folded in Kanban",
Help: "This stage is folded in the kanban view when there are" + 
"no records in that stage to display.",
},
})
h.ProjectProject().DeclareModel()
h.ProjectProject().AddSQLConstraint("project_date_greater", "check(date >= date_start)", "Error! project start-date must be lower than project end-date.")





//    _period_number = 5
h.ProjectProject().Methods().GetAliasModelName().DeclareMethod(
`GetAliasModelName`,
func(rs m.ProjectProjectSet, vals interface{})  {
//        return vals.get('alias_model', 'project.task')
})
h.ProjectProject().Methods().GetAliasValues().DeclareMethod(
`GetAliasValues`,
func(rs m.ProjectProjectSet)  {
//        values = super(Project, self).get_alias_values()
//        values['alias_defaults'] = {'project_id': self.id}
//        return values
})
h.ProjectProject().Methods().Unlink().Extend(
`Unlink`,
func(rs m.ProjectProjectSet)  {
//        analytic_accounts_to_delete = self.env['account.analytic.account']
//        for project in self:
//            if project.tasks:
//                raise UserError(
//                    _('You cannot delete a project containing tasks. You can either delete all the project\'s tasks and then delete the project or simply deactivate the project.'))
//            if project.analytic_account_id and not project.analytic_account_id.line_ids:
//                analytic_accounts_to_delete |= project.analytic_account_id
//        res = super(Project, self).unlink()
//        analytic_accounts_to_delete.unlink()
//        return res
})
h.ProjectProject().Methods().ComputeAttachedDocsCount().DeclareMethod(
`ComputeAttachedDocsCount`,
func(rs m.ProjectProjectSet)  {
//        Attachment = self.env['ir.attachment']
//        for project in self:
//            project.doc_count = Attachment.search_count([
//                '|',
//                '&',
//                ('res_model', '=', 'project.project'), ('res_id', '=', project.id),
//                '&',
//                ('res_model', '=', 'project.task'), ('res_id',
//                                                     'in', project.task_ids.ids)
//            ])
})
h.ProjectProject().Methods().ComputeTaskCount().DeclareMethod(
`ComputeTaskCount`,
func(rs m.ProjectProjectSet)  {
//        for project in self:
//            project.task_count = len(project.task_ids)
})
h.ProjectProject().Methods().ComputeTaskNeedactionCount().DeclareMethod(
`ComputeTaskNeedactionCount`,
func(rs m.ProjectProjectSet)  {
//        projects_data = self.env['project.task'].read_group([
//            ('project_id', 'in', self.ids),
//            ('message_needaction', '=', True)
//        ], ['project_id'], ['project_id'])
//        mapped_data = {project_data['project_id'][0]: int(project_data['project_id_count'])
//                       for project_data in projects_data}
//        for project in self:
//            project.task_needaction_count = mapped_data.get(project.id, 0)
})
h.ProjectProject().Methods().GetAliasModels().DeclareMethod(
` Overriden in project_issue to offer more options `,
func(rs m.ProjectProjectSet)  {
//        return [('project.task', "Tasks")]
})
h.ProjectProject().Methods().AttachmentTreeView().DeclareMethod(
`AttachmentTreeView`,
func(rs m.ProjectProjectSet)  {
//        self.ensure_one()
//        domain = [
//            '|',
//            '&', ('res_model', '=', 'project.project'), ('res_id', 'in', self.ids),
//            '&', ('res_model', '=', 'project.task'), ('res_id', 'in', self.task_ids.ids)]
//        return {
//            'name': _('Attachments'),
//            'domain': domain,
//            'res_model': 'ir.attachment',
//            'type': 'ir.actions.act_window',
//            'view_id': False,
//            'view_mode': 'kanban,tree,form',
//            'view_type': 'form',
//            'help': _('''<p class="oe_view_nocontent_create">
//                        Documents are attached to the tasks and issues of your project.</p><p>
//                        Send messages or log internal notes with attachments to link
//                        documents to your project.
//                    </p>'''),
//            'limit': 80,
//            'context': "{'default_res_model': '%s','default_res_id': %d}" % (self._name, self.id)
//        }
})
h.ProjectProject().Methods().ActivateSampleProject().DeclareMethod(
` Unarchives the sample project 'project.project_project_data' and
            reloads the project dashboard `,
func(rs m.ProjectProjectSet)  {
//        project = self.env.ref('project.project_project_data', False)
//        if project:
//            project.write({'active': True})
//        cover_image = self.env.ref('project.msg_task_data_14_attach', False)
//        cover_task = self.env.ref('project.project_task_data_14', False)
//        if cover_image and cover_task:
//            cover_task.write({'displayed_image_id': cover_image.id})
//        action = self.env.ref('project.open_view_project_all', False)
//        action_data = None
//        if action:
//            action.sudo().write({
//                "help": _('''<p class="oe_view_nocontent_create">Click to create a new project.</p>''')
//            })
//            action_data = action.read()[0]
//        return action_data
})
h.ProjectProject().Methods().ComputeIsFavorite().DeclareMethod(
`ComputeIsFavorite`,
func(rs m.ProjectProjectSet)  {
//        for project in self:
//            project.is_favorite = self.env.user in project.favorite_user_ids
})
h.ProjectProject().Methods().GetDefaultFavoriteUserIds().DeclareMethod(
`GetDefaultFavoriteUserIds`,
func(rs m.ProjectProjectSet)  {
//        return [(6, 0, [self.env.uid])]
})
h.ProjectProject().Methods().DefaultGet().Extend(
`DefaultGet`,
func(rs m.ProjectProjectSet, flds interface{})  {
//        result = super(Project, self).default_get(flds)
//        result['use_tasks'] = True
//        return result
})
h.ProjectProject().Methods().AliasModels().DeclareMethod(
`AliasModels`,
func(rs m.ProjectProjectSet)  {
//    def _alias_models(self): return self._get_alias_models()
})
h.ProjectProject().AddFields(map[string]models.FieldDefinition{
"Active": models.BooleanField{
Default: models.DefaultValue(true),
Help: "If the active field is set to False, it will allow you" + 
"to hide the project without removing it.",
},
"Sequence": models.IntegerField{
Default: models.DefaultValue(10),
Help: "Gives the sequence order when displaying a list of Projects.",
},
"AnalyticAccountId": models.Many2OneField{
RelationModel: h.AccountAnalyticAccount(),
String: "Contract/Analytic",
Help: "Link this project to an analytic account if you need financial" + 
"management on projects. It enables you to connect projects" + 
"with budgets, planning, cost and revenue analysis, timesheets" + 
"on projects, etc.",
OnDelete: `cascade`,
Required: true,

},
"FavoriteUserIds": models.Many2ManyField{
RelationModel: h.User(),
M2MLinkModelName: "",
M2MOurField: "",
M2MTheirField: "",
Default: models.DefaultValue(_get_default_favorite_user_ids),
String: "Members",
},
"IsFavorite": models.BooleanField{
Compute: h.ProjectProject().Methods().ComputeIsFavorite(),
String: "Show Project on dashboard",
Help: "Whether this project should be displayed on the dashboard or not",
},
"LabelTasks": models.CharField{
String: "Use Tasks as",
Default: models.DefaultValue("Tasks"),
Help: "Gives label to tasks on project's kanban view.",
},
"Tasks": models.One2ManyField{
RelationModel: h.ProjectTask(),
ReverseFK: "",
String: "Task Activities",
},
"ResourceCalendarId": models.Many2OneField{
RelationModel: h.ResourceCalendar(),
String: "Working Time",
Help: "Timetable working hours to adjust the gantt diagram report",
},
"TypeIds": models.Many2ManyField{
RelationModel: h.ProjectTaskType(),
M2MLinkModelName: "",
M2MOurField: "",
M2MTheirField: "",
String: "Tasks Stages",
},
"TaskCount": models.IntegerField{
Compute: h.ProjectProject().Methods().ComputeTaskCount(),
String: "Tasks",
},
"TaskNeedactionCount": models.IntegerField{
Compute: h.ProjectProject().Methods().ComputeTaskNeedactionCount(),
String: "Tasks",
},
"TaskIds": models.One2ManyField{
RelationModel: h.ProjectTask(),
ReverseFK: "",
String: "Tasks",
Filter: q.StageId().Fold().Equals(False).Or().StageId().Equals(False),
},
"Color": models.IntegerField{
String: "Color Index",
},
"UserId": models.Many2OneField{
RelationModel: h.User(),
String: "Project Manager",
Default: func (env models.Environment) interface{} { return env.Uid() },
},
"AliasId": models.Many2OneField{
RelationModel: h.MailAlias(),
String: "Alias",
OnDelete: `restrict`,
Required: true,
Help: "Internal email associated with this project. Incoming emails" + 
"are automatically synchronized with Tasks (or optionally" + 
"Issues if the Issue Tracker module is installed).",
},
"AliasModel": models.SelectionField{
Selection: _alias_models,
String: "Alias Model",
Index: true,
Required: true,
Default: models.DefaultValue("project.task"),
Help: "The kind of document created when an email is received" + 
"on this project's email alias",
},
"PrivacyVisibility": models.SelectionField{
Selection: types.Selection{
"followers": "",
"employees": "",
"portal": "",
},
String: "Privacy",
Required: true,
Default: models.DefaultValue("employees"),
Help: "Holds visibility of the tasks or issues that belong to" + 
"the current project:" + 
"- On invitation only: Employees may only see the followed" + 
"project, tasks or issues" + 
"- Visible by all employees: Employees may see all project," + 
"tasks or issues" + 
"- Visible by following customers: employees see everything;" + 
"   if website is activated, portal users may see project," + 
"tasks or issues followed by" + 
"   them or by someone of their company" + 
"",
},
"DocCount": models.IntegerField{
Compute: h.ProjectProject().Methods().ComputeAttachedDocsCount(),
String: "Number of documents attached",
},
"DateStart": models.DateField{
String: "Start Date",
},
"Date": models.DateField{
String: "Expiration Date",
Index: true,
//track_visibility='onchange'
},

})
h.ProjectProject().Methods().MapTasks().DeclareMethod(
` copy and map tasks from old to new project `,
func(rs m.ProjectProjectSet, new_project_id interface{})  {
//        tasks = self.env['project.task']
//        for task in self.tasks:
//            # preserve task name and stage, normally altered during copy
//            defaults = {'stage_id': task.stage_id.id,
//                        'name': task.name}
//            tasks += task.copy(defaults)
//        return self.browse(new_project_id).write({'tasks': [(6, 0, tasks.ids)]})
})
h.ProjectProject().Methods().Copy().Extend(
`Copy`,
func(rs m.ProjectProjectSet, defaultName models.RecordData)  {
//        if default is None:
//            default = {}
//        self = self.with_context(active_test=False)
//        if not default.get('name'):
//            default['name'] = _("%s (copy)") % (self.name)
//        project = super(Project, self).copy(default)
//        for follower in self.message_follower_ids:
//            project.message_subscribe(
//                partner_ids=follower.partner_id.ids, subtype_ids=follower.subtype_ids.ids)
//        if 'tasks' not in default:
//            self.map_tasks(project.id)
//        return project
})
h.ProjectProject().Methods().Create().Extend(
`Create`,
func(rs m.ProjectProjectSet, vals models.RecordData)  {
//        ir_values = self.env['ir.values'].get_default(
//            'project.config.settings', 'generate_project_alias')
//        if ir_values:
//            vals['alias_name'] = vals.get('alias_name') or vals.get('name')
//        self = self.with_context(
//            project_creation_in_progress=True, mail_create_nosubscribe=True)
//        return super(Project, self).create(vals)
})
h.ProjectProject().Methods().Write().Extend(
`Write`,
func(rs m.ProjectProjectSet, vals models.RecordData)  {
//        if vals.get('alias_model'):
//            vals['alias_model_id'] = self.env['ir.model'].search([
//                ('model', '=', vals.get('alias_model', 'project.task'))
//            ], limit=1).id
//        res = super(Project, self).write(vals)
//        if 'active' in vals:
//            # archiving/unarchiving a project does it on its tasks, too
//            self.with_context(active_test=False).mapped(
//                'tasks').write({'active': vals['active']})
//            # archiving/unarchiving a project implies that we don't want to use the analytic account anymore
//            self.with_context(active_test=False).mapped(
//                'analytic_account_id').write({'active': vals['active']})
//        return res
})
h.ProjectProject().Methods().ToggleFavorite().DeclareMethod(
`ToggleFavorite`,
func(rs m.ProjectProjectSet)  {
//        favorite_projects = not_fav_projects = self.env['project.project'].sudo(
//        )
//        for project in self:
//            if self.env.user in project.favorite_user_ids:
//                favorite_projects |= project
//            else:
//                not_fav_projects |= project
//        not_fav_projects.write({'favorite_user_ids': [(4, self.env.uid)]})
//        favorite_projects.write({'favorite_user_ids': [(3, self.env.uid)]})
})
h.ProjectProject().Methods().CloseDialog().DeclareMethod(
`CloseDialog`,
func(rs m.ProjectProjectSet)  {
//        return {'type': 'ir.actions.act_window_close'}
})
h.ProjectTask().DeclareModel()




//    _mail_post_access = 'read'

h.ProjectTask().Methods().DefaultGet().Extend(
` Set 'date_assign' if user_id is set. `,
func(rs m.ProjectTaskSet, field_list interface{})  {
//        result = super(Task, self).default_get(field_list)
//        if 'user_id' in result:
//            result['date_assign'] = fields.Datetime.now()
//        return result
})
h.ProjectTask().Methods().GetDefaultPartner().DeclareMethod(
`GetDefaultPartner`,
func(rs m.ProjectTaskSet)  {
//        if 'default_project_id' in self.env.context:
//            default_project_id = self.env['project.project'].browse(
//                self.env.context['default_project_id'])
//            return default_project_id.exists().partner_id
})
h.ProjectTask().Methods().GetDefaultStageId().DeclareMethod(
` Gives default stage_id `,
func(rs m.ProjectTaskSet)  {
//        project_id = self.env.context.get('default_project_id')
//        if not project_id:
//            return False
//        return self.stage_find(project_id, [('fold', '=', False)])
})
h.ProjectTask().Methods().ReadGroupStageIds().DeclareMethod(
`ReadGroupStageIds`,
func(rs m.ProjectTaskSet, stages interface{}, domain interface{}, order interface{})  {
//        search_domain = [('id', 'in', stages.ids)]
//        if 'default_project_id' in self.env.context:
//            search_domain = [
//                '|', ('project_ids', '=', self.env.context['default_project_id'])] + search_domain
//        stage_ids = stages._search(
//            search_domain, order=order, access_rights_uid=SUPERUSER_ID)
//        return stages.browse(stage_ids)
})
h.ProjectTask().AddFields(map[string]models.FieldDefinition{
"Active": models.BooleanField{
Default: models.DefaultValue(true),
},
"Name": models.CharField{
String: "Task Title",
//track_visibility='always'
Required: true,
Index: true,
},
"Description": models.HTMLField{
String: "Description",
},
"Priority": models.SelectionField{
Selection: types.Selection{
"0": "Normal",
"1": "High",
},
Default: models.DefaultValue("0"),
Index: true,
},
"Sequence": models.IntegerField{
String: "Sequence",
Index: true,
Default: models.DefaultValue(10),
Help: "Gives the sequence order when displaying a list of tasks.",
},
"StageId": models.Many2OneField{
RelationModel: h.ProjectTaskType(),
String: "Stage",
//track_visibility='onchange'
Index: true,
Default: models.DefaultValue(_get_default_stage_id),
//group_expand='_read_group_stage_ids'
Filter: q.ProjectIds().Equals(project_id),
NoCopy: true,
},
"TagIds": models.Many2ManyField{
RelationModel: h.ProjectTags(),
String: "Tags",
//oldname='categ_ids'
},
"KanbanState": models.SelectionField{
Selection: types.Selection{
"normal": "In Progress",
"done": "Ready for next stage",
"blocked": "Blocked",
},
String: "Kanban State",
Default: models.DefaultValue("normal"),
//track_visibility='onchange'
Required: true,
NoCopy: true,
Help: "A task's kanban state indicates special situations affecting it:" + 
" * Normal is the default situation" + 
" * Blocked indicates something is preventing the progress of this task" + 
" * Ready for next stage indicates the task is ready to" + 
"be pulled to the next stage",
},
"DateStart": models.DateTimeField{
String: "Starting Date",
Default: func (env models.Environment) interface{} { return dates.Now() },
Index: true,
NoCopy: true,
},
"DateEnd": models.DateTimeField{
String: "Ending Date",
Index: true,
NoCopy: true,
},
"DateAssign": models.DateTimeField{
String: "Assigning Date",
Index: true,
NoCopy: true,
ReadOnly: true,
},
"DateDeadline": models.DateField{
String: "Deadline",
Index: true,
NoCopy: true,
},
"DateLastStageUpdate": models.DateTimeField{
String: "Last Stage Update",
Default: func (env models.Environment) interface{} { return dates.Now() },
Index: true,
NoCopy: true,
ReadOnly: true,
},
"ProjectId": models.Many2OneField{
RelationModel: h.ProjectProject(),
String: "Project",
Default: func (env models.Environment) interface{} { return env.context.get() },
Index: true,
//track_visibility='onchange'
//change_default=True
},
"Notes": models.TextField{
String: "Notes",
},
"PlannedHours": models.FloatField{
String: "Initially Planned Hours",
Help: "Estimated time to do the task, usually set by the project" + 
"manager when the task is in draft state.",
},
"RemainingHours": models.FloatField{
String: "Remaining Hours",
//digits=(16, 2)
Help: "Total remaining time, can be re-estimated periodically" + 
"by the assignee of the task.",
},
"UserId": models.Many2OneField{
RelationModel: h.User(),
String: "Assigned to",
Default: func (env models.Environment) interface{} { return env.uid },
Index: true,
//track_visibility='always'
},
"PartnerId": models.Many2OneField{
RelationModel: h.Partner(),
String: "Customer",
Default: models.DefaultValue(_get_default_partner),
},
"ManagerId": models.Many2OneField{
RelationModel: h.User(),
String: "Project Manager",
Related: `ProjectId.UserId`,
ReadOnly: true,
},
"CompanyId": models.Many2OneField{
RelationModel: h.Company(),
String: "Company",
Default: func (env models.Environment) interface{} { return env["res.company"]._company_default_get() },
},
"Color": models.IntegerField{
String: "Color Index",
},
"UserEmail": models.CharField{
Related: `UserId.Email`,
String: "User Email",
ReadOnly: true,
},
"AttachmentIds": models.One2ManyField{
RelationModel: h.Attachment(),
ReverseFK: "",
Filter: q.ResModel().Equals("ProjectTask"),

String: "Attachments",
},
"DisplayedImageId": models.Many2OneField{
RelationModel: h.Attachment(),
Filter: q.ResModel().Equals("project.task").And().ResId().Equals(id).And().Mimetype().IContains("image"),
String: "Displayed Image",
},
"LegendBlocked": models.CharField{
Related: `StageId.LegendBlocked`,
String: "Kanban Blocked Explanation",
ReadOnly: true,
},
"LegendDone": models.CharField{
Related: `StageId.LegendDone`,
String: "Kanban Valid Explanation",
ReadOnly: true,
},
"LegendNormal": models.CharField{
Related: `StageId.LegendNormal`,
String: "Kanban Ongoing Explanation",
ReadOnly: true,
},
})
h.ProjectTask().Fields().CreateDate().setIndex( true,)
h.ProjectTask().Fields().WriteDate().setIndex( true,)
h.ProjectTask().Methods().OnchangeProject().DeclareMethod(
`OnchangeProject`,
func(rs m.ProjectTaskSet)  {
//        if self.project_id:
//            if self.project_id.partner_id:
//                self.partner_id = self.project_id.partner_id
//            self.stage_id = self.stage_find(
//                self.project_id.id, [('fold', '=', False)])
//        else:
//            self.stage_id = False
})
h.ProjectTask().Methods().OnchangeUser().DeclareMethod(
`OnchangeUser`,
func(rs m.ProjectTaskSet)  {
//        if self.user_id:
//            self.date_start = fields.Datetime.now()
})
h.ProjectTask().Methods().Copy().Extend(
`Copy`,
func(rs m.ProjectTaskSet, defaultName models.RecordData)  {
//        if default is None:
//            default = {}
//        if not default.get('name'):
//            default['name'] = _("%s (copy)") % self.name
//        if 'remaining_hours' not in default:
//            default['remaining_hours'] = self.planned_hours
//        return super(Task, self).copy(default)
})
h.ProjectTask().Methods().CheckDates().DeclareMethod(
`CheckDates`,
func(rs m.ProjectTaskSet)  {
//        if any(self.filtered(lambda task: task.date_start and task.date_end and task.date_start > task.date_end)):
//            raise ValidationError(
//                _('Error ! Task starting date must be lower than its ending date.'))
})
h.ProjectTask().Methods().FieldsViewGet().Extend(
`FieldsViewGet`,
func(rs m.ProjectTaskSet, view_id webdata.FieldsViewGetParams, view_type interface{}, toolbar interface{}, submenu interface{})  {
//        obj_tm = self.env.user.company_id.project_time_mode_id
//        tm = obj_tm and obj_tm.name or 'Hours'
//        res = super(Task, self).fields_view_get(view_id=view_id,
//                                                view_type=view_type, toolbar=toolbar, submenu=submenu)
//        obj_tm = self.env.user.company_id.project_time_mode_id
//        uom_hour = self.env.ref('product.product_uom_hour', False)
//        if not obj_tm or not uom_hour or obj_tm.id == uom_hour.id:
//            return res
//        eview = etree.fromstring(res['arch'])
//        def _check_rec(eview):
//            if eview.attrib.get('widget', '') == 'float_time':
//                eview.set('widget', 'float')
//            for child in eview:
//                _check_rec(child)
//            return True
//        _check_rec(eview)
//        res['arch'] = etree.tostring(eview)
//        for f in res['fields']:
//            # TODO this NOT work in different language than english
//            # the field 'Initially Planned Hours' should be replaced by 'Initially Planned Days'
//            # but string 'Initially Planned Days' is not available in translation
//            if 'Hours' in res['fields'][f]['string']:
//                res['fields'][f]['string'] = res['fields'][f]['string'].replace(
//                    'Hours', obj_tm.name)
//        return res
})
h.ProjectTask().Methods().GetEmptyListHelp().DeclareMethod(
`GetEmptyListHelp`,
func(rs m.ProjectTaskSet, help interface{})  {
//        self = self.with_context(
//            empty_list_help_id=self.env.context.get('default_project_id'),
//            empty_list_help_model='project.project',
//            empty_list_help_document_name=_("tasks")
//        )
//        return super(Task, self).get_empty_list_help(help)
})
h.ProjectTask().Methods().StageFind().DeclareMethod(
` Override of the base.stage method
            Parameter of the stage search taken from the lead:
            - section_id: if set, stages must belong to this section or
              be a default stage; if not set, stages must be default
              stages
        `,
func(rs m.ProjectTaskSet, section_id interface{}, domain interface{}, order interface{})  {
//        section_ids = []
//        if section_id:
//            section_ids.append(section_id)
//        section_ids.extend(self.mapped('project_id').ids)
//        search_domain = []
//        if section_ids:
//            search_domain = [('|')] * (len(section_ids) - 1)
//            for section_id in section_ids:
//                search_domain.append(('project_ids', '=', section_id))
//        search_domain += list(domain)
//        return self.env['project.task.type'].search(search_domain, order=order, limit=1).id
})
h.ProjectTask().Methods().Create().Extend(
`Create`,
func(rs m.ProjectTaskSet, vals models.RecordData)  {
//        context = dict(self.env.context, mail_create_nolog=True)
//        if vals.get('project_id') and not context.get('default_project_id'):
//            context['default_project_id'] = vals.get('project_id')
//        if vals.get('user_id'):
//            vals['date_assign'] = fields.Datetime.now()
//        task = super(Task, self.with_context(context)).create(vals)
//        return task
})
h.ProjectTask().Methods().Write().Extend(
`Write`,
func(rs m.ProjectTaskSet, vals models.RecordData)  {
//        now = fields.Datetime.now()
//        if 'stage_id' in vals:
//            vals['date_last_stage_update'] = now
//            # reset kanban state when changing stage
//            if 'kanban_state' not in vals:
//                vals['kanban_state'] = 'normal'
//        if vals.get('user_id'):
//            vals['date_assign'] = now
//        result = super(Task, self).write(vals)
//        return result
})
h.ProjectTask().Methods().TrackTemplate().DeclareMethod(
`TrackTemplate`,
func(rs m.ProjectTaskSet, tracking interface{})  {
//        res = super(Task, self)._track_template(tracking)
//        test_task = self[0]
//        changes, tracking_value_ids = tracking[test_task.id]
//        if 'stage_id' in changes and test_task.stage_id.mail_template_id:
//            res['stage_id'] = (test_task.stage_id.mail_template_id, {
//                               'composition_mode': 'mass_mail'})
//        return res
})
h.ProjectTask().Methods().TrackSubtype().DeclareMethod(
`TrackSubtype`,
func(rs m.ProjectTaskSet, init_values interface{})  {
//        self.ensure_one()
//        if 'kanban_state' in init_values and self.kanban_state == 'blocked':
//            return 'project.mt_task_blocked'
//        elif 'kanban_state' in init_values and self.kanban_state == 'done':
//            return 'project.mt_task_ready'
//        elif 'user_id' in init_values and self.user_id:  # assigned -> new
//            return 'project.mt_task_new'
//        elif 'stage_id' in init_values and self.stage_id and self.stage_id.sequence <= 1:  # start stage -> new
//            return 'project.mt_task_new'
//        elif 'stage_id' in init_values:
//            return 'project.mt_task_stage'
//        return super(Task, self)._track_subtype(init_values)
})
h.ProjectTask().Methods().NotificationRecipients().DeclareMethod(
` Handle project users and managers recipients that can convert assign
        tasks and create new one directly from notification emails. `,
func(rs m.ProjectTaskSet, message interface{}, groups interface{})  {
//        groups = super(Task, self)._notification_recipients(message, groups)
//        self.ensure_one()
//        if not self.user_id:
//            take_action = self._notification_link_helper('assign')
//            project_actions = [{'url': take_action, 'title': _('I take it')}]
//        else:
//            new_action_id = self.env.ref('project.action_view_task').id
//            new_action = self._notification_link_helper(
//                'new', action_id=new_action_id)
//            project_actions = [{'url': new_action, 'title': _('New Task')}]
//        new_group = (
//            'group_project_user', lambda partner: bool(partner.user_ids) and any(user.has_group('project.group_project_user') for user in partner.user_ids), {
//                'actions': project_actions,
//            })
//        return [new_group] + groups
})
h.ProjectTask().Methods().MessageGetReplyTo().DeclareMethod(
` Override to get the reply_to of the parent project. `,
func(rs m.ProjectTaskSet, res_ids interface{}, defaultName interface{})  {
//        tasks = self.sudo().browse(res_ids)
//        project_ids = tasks.mapped('project_id').ids
//        aliases = self.env['project.project'].message_get_reply_to(
//            project_ids, default=default)
//        return {task.id: aliases.get(task.project_id.id, False) for task in tasks}
})
h.ProjectTask().Methods().EmailSplit().DeclareMethod(
`EmailSplit`,
func(rs m.ProjectTaskSet, msg interface{})  {
//        email_list = tools.email_split(
//            (msg.get('to') or '') + ',' + (msg.get('cc') or ''))
//        aliases = self.mapped('project_id.alias_name')
//        return filter(lambda x: x.split('@')[0] not in aliases, email_list)
})
h.ProjectTask().Methods().MessageNew().DeclareMethod(
` Override to updates the document according to the email. `,
func(rs m.ProjectTaskSet, msg interface{}, custom_values interface{})  {
//        if custom_values is None:
//            custom_values = {}
//        defaults = {
//            'name': msg.get('subject'),
//            'planned_hours': 0.0,
//            'partner_id': msg.get('author_id')
//        }
//        defaults.update(custom_values)
//        res = super(Task, self).message_new(msg, custom_values=defaults)
//        task = self.browse(res)
//        email_list = task.email_split(msg)
//        partner_ids = filter(None, task._find_partner_from_emails(
//            email_list, force_create=False))
//        task.message_subscribe(partner_ids)
//        return res
})
h.ProjectTask().Methods().MessageUpdate().DeclareMethod(
` Override to update the task according to the email. `,
func(rs m.ProjectTaskSet, msg interface{}, update_vals interface{})  {
//        if update_vals is None:
//            update_vals = {}
//        maps = {
//            'cost': 'planned_hours',
//        }
//        for line in msg['body'].split('\n'):
//            line = line.strip()
//            res = tools.command_re.match(line)
//            if res:
//                match = res.group(1).lower()
//                field = maps.get(match)
//                if field:
//                    try:
//                        update_vals[field] = float(res.group(2).lower())
//                    except (ValueError, TypeError):
//                        pass
//        email_list = self.email_split(msg)
//        partner_ids = filter(None, self._find_partner_from_emails(
//            email_list, force_create=False))
//        self.message_subscribe(partner_ids)
//        return super(Task, self).message_update(msg, update_vals=update_vals)
})
h.ProjectTask().Methods().MessageGetSuggestedRecipients().DeclareMethod(
`MessageGetSuggestedRecipients`,
func(rs m.ProjectTaskSet)  {
//        recipients = super(Task, self).message_get_suggested_recipients()
//        for task in self.filtered('partner_id'):
//            reason = _('Customer Email') if task.partner_id.email else _(
//                'Customer')
//            task._message_add_suggested_recipient(
//                recipients, partner=task.partner_id, reason=reason)
//        return recipients
})
h.ProjectTask().Methods().MessageGetEmailValues().DeclareMethod(
`MessageGetEmailValues`,
func(rs m.ProjectTaskSet, notif_mail interface{})  {
//        res = super(Task, self).message_get_email_values(notif_mail=notif_mail)
//        headers = {}
//        if res.get('headers'):
//            try:
//                headers.update(safe_eval(res['headers']))
//            except Exception:
//                pass
//        if self.project_id:
//            current_objects = filter(None, headers.get(
//                'X-Odoo-Objects', '').split(','))
//            current_objects.insert(
//                0, 'project.project-%s, ' % self.project_id.id)
//            headers['X-Odoo-Objects'] = ','.join(current_objects)
//        if self.tag_ids:
//            headers['X-Odoo-Tags'] = ','.join(self.tag_ids.mapped('name'))
//        res['headers'] = repr(headers)
//        return res
})
h.AccountAnalyticAccount().DeclareModel()


h.AccountAnalyticAccount().AddFields(map[string]models.FieldDefinition{
"UseTasks": models.BooleanField{
String: "Use Tasks",
Help: "Check this box to manage internal activities through this project",
},
"CompanyUomId": models.Many2OneField{
RelationModel: h.ProductUom(),
Related: `CompanyId.ProjectTimeModeId`,
String: "Company UOM",
},
"ProjectIds": models.One2ManyField{
RelationModel: h.ProjectProject(),
ReverseFK: "",
String: "Projects",
},
"ProjectCount": models.IntegerField{
Compute: h.AccountAnalyticAccount().Methods().ComputeProjectCount(),
String: "Project Count",
},
})
h.AccountAnalyticAccount().Methods().ComputeProjectCount().DeclareMethod(
`ComputeProjectCount`,
func(rs h.AccountAnalyticAccountSet) h.AccountAnalyticAccountData {
//        for account in self:
//            account.project_count = len(
//                account.with_context(active_test=False).project_ids)
})
h.AccountAnalyticAccount().Methods().TriggerProjectCreation().DeclareMethod(
`
        This function is used to decide if a project needs
to be automatically created or not when an analytic account
is created. It returns True if it needs to be so, False otherwise.
        `,
func(rs m.AccountAnalyticAccountSet, vals interface{})  {
//        return vals.get('use_tasks') and 'project_creation_in_progress' not in self.env.context
})
h.AccountAnalyticAccount().Methods().ProjectCreate().DeclareMethod(
`
        This function is called at the time of analytic
account creation and is used to create a project automatically
linked to it if the conditions are meet.
        `,
func(rs m.AccountAnalyticAccountSet, vals interface{})  {
//        self.ensure_one()
//        Project = self.env['project.project']
//        project = Project.with_context(active_test=False).search(
//            [('analytic_account_id', '=', self.id)])
//        if not project and self._trigger_project_creation(vals):
//            project_values = {
//                'name': vals.get('name'),
//                'analytic_account_id': self.id,
//                'use_tasks': True,
//            }
//            return Project.create(project_values).id
//        return False
})
h.AccountAnalyticAccount().Methods().Create().Extend(
`Create`,
func(rs m.AccountAnalyticAccountSet, vals models.RecordData)  {
//        analytic_account = super(AccountAnalyticAccount, self).create(vals)
//        analytic_account.project_create(vals)
//        return analytic_account
})
h.AccountAnalyticAccount().Methods().Write().Extend(
`Write`,
func(rs m.AccountAnalyticAccountSet, vals models.RecordData)  {
//        vals_for_project = vals.copy()
//        for account in self:
//            if not vals.get('name'):
//                vals_for_project['name'] = account.name
//            account.project_create(vals_for_project)
//        return super(AccountAnalyticAccount, self).write(vals)
})
h.AccountAnalyticAccount().Methods().Unlink().Extend(
`Unlink`,
func(rs m.AccountAnalyticAccountSet)  {
//        projects = self.env['project.project'].search(
//            [('analytic_account_id', 'in', self.ids)])
//        has_tasks = self.env['project.task'].search_count(
//            [('project_id', 'in', projects.ids)])
//        if has_tasks:
//            raise UserError(
//                _('Please remove existing tasks in the project linked to the accounts you want to delete.'))
//        return super(AccountAnalyticAccount, self).unlink()
})
h.AccountAnalyticAccount().Methods().NameSearch().Extend(
`NameSearch`,
func(rs m.AccountAnalyticAccountSet, name webdata.NameSearchParams, args interface{}, operator interface{}, limit interface{})  {
//        if args is None:
//            args = []
//        if self.env.context.get('current_model') == 'project.project':
//            return self.search(args + [('name', operator, name)], limit=limit).name_get()
//        return super(AccountAnalyticAccount, self).name_search(name, args=args, operator=operator, limit=limit)
})
h.AccountAnalyticAccount().Methods().ProjectsAction().DeclareMethod(
`ProjectsAction`,
func(rs m.AccountAnalyticAccountSet)  {
//        projects = self.with_context(active_test=False).mapped('project_ids')
//        result = {
//            "type": "ir.actions.act_window",
//            "res_model": "project.project",
//            "views": [[False, "tree"], [False, "form"]],
//            "domain": [["id", "in", projects.ids]],
//            "context": {"create": False},
//            "name": "Projects",
//        }
//        if len(projects) == 1:
//            result['views'] = [(False, "form")]
//            result['res_id'] = projects.id
//        else:
//            result = {'type': 'ir.actions.act_window_close'}
//        return result
})
h.ProjectTags().DeclareModel()
h.ProjectTags().AddSQLConstraint("name_uniq", "unique (name)", "Tag name already exists !")


h.ProjectTags().AddFields(map[string]models.FieldDefinition{
"Name": models.CharField{
Required: true,
},
"Color": models.IntegerField{
String: "Color Index",
},

})
}
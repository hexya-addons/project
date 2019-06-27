package project

import (
	"github.com/hexya-addons/web/controllers"
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "project"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})
	controllers.BackendJS = append(controllers.BackendJS,
		"/static/project/src/js/project.js",
		"/static/project/src/js/tour.js",
		"/static/project/src/js/web_planner_project.js",
	)
	controllers.Backend = append(controllers.Backend,
		"/static/project/src/css/project.css",
		"/static/project/src/less/project_dashboard.less",
	)

}

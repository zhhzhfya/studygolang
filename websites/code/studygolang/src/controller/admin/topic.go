package admin

import (
	"filter"
	"model"
	"net/http"
)

// 所有帖子（分页）
func TopicsHandler(rw http.ResponseWriter, req *http.Request) {
	user, _ := filter.CurrentUser(req)
	// 设置内容模板
	req.Form.Set(filter.CONTENT_TPL_KEY, "/template/admin/users.html")
	filter.SetData(req, map[string]interface{}{"user": user})
}

// 所有节点（分页）
func NodesHandler(rw http.ResponseWriter, req *http.Request) {
	user, _ := filter.CurrentUser(req)
	// 设置内容模板
	req.Form.Set(filter.CONTENT_TPL_KEY, "/template/admin/nodes.html")
	filter.SetData(req, map[string]interface{}{"user": user, "nodes": model.AllNode})
}

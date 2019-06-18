package middleware

import (
	// "fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/sisimengchen/serverme/controllers"
	"github.com/sisimengchen/serverme/models"
	"github.com/sisimengchen/serverme/utils"
	"net/http"
)

var enforcer = casbin.NewEnforcer("configs/rbac_model.conf", "configs/rbac_policy.csv")

// 权限认证中间件
func Casbin( /*enforcer *casbin.Enforcer*/ ) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		contextUser, _ := controllers.GetContextUser(ctx)
		userRoles, _ := models.GetUserRoles(&models.User{ID: contextUser.ID})
		var isPermission = false
		for _, userRole := range *userRoles {
			roleCode := userRole.Role.Code
			if ok, err := enforcer.EnforceSafe(roleCode, path, method); err != nil { // 进入这里属于服务器错误的范畴，直接停止中间件
				ctx.Abort()
				return
			} else if ok {
				isPermission = true
				break
			}
		}
		if !isPermission {
			if !utils.IsApiRequest(ctx) {
				ctx.Redirect(http.StatusFound, "/page/unauthorized")
			} else {
				ctx.JSON(controllers.ResponseResource(http.StatusForbidden, "无权限", nil))
			}
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

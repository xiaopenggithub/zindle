package logic

import (
	"backend/common/errorx"
	"context"
	"encoding/json"
	"fmt"

	"backend/service/backend/cmd/api/internal/svc"
	"backend/service/backend/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuLogic {
	return &GetMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuLogic) GetMenu() (resp *types.AdminReply, err error) {
	data := make(map[string]interface{})
	roleIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("roleId")))
	roleId, _ := roleIdNumber.Int64()
	data["menus"] = l.getAdminsMenu(roleId)
	return nil, errorx.NewCodeError(200, "success", data)
}

// get admin's menu
func (l *GetMenuLogic) getAdminsMenu(roleId int64) interface{} {
	list, err := l.svcCtx.SystemRoleMenusModel.ListByRoleId(roleId)
	if err != nil {
		return nil
	}
	menuIds := ""
	for _, v := range list {
		if menuIds == "" {
			menuIds += fmt.Sprintf("%v", v.MenuId)
		} else {
			menuIds += fmt.Sprintf(",%v", v.MenuId)
		}
	}
	// 防止查询所有
	if menuIds == "" {
		menuIds = "0"
	}
	adminsMenu, _ := l.svcCtx.SystemMenusModel.Tree(menuIds)

	// 构造菜单树
	var topMenu []interface{}
	for _, v := range adminsMenu {
		// 一级
		if v.ParentId == 0 {
			// 结构体转json
			sJson, err := json.Marshal(v)
			if err != nil {
				fmt.Println(err)
			}
			// json 转map
			m := make(map[string]interface{})
			json.Unmarshal(sJson, &m)
			meta := make(map[string]interface{})
			meta["title"] = v.Title
			delete(m, "title")
			meta["icon"] = v.Icon
			delete(m, "icon")
			m["meta"] = meta

			delete(m, "created_at")
			delete(m, "updated_at")
			delete(m, "deleted_at")

			delete(m, "id")
			delete(m, "parent_id")
			delete(m, "sort")

			// 二级
			var children []interface{}
			//children := make([]interface{}, 0) //空切片
			for _, subv := range adminsMenu {
				if subv.ParentId > 0 && v.Id == subv.ParentId {
					// 结构体转json
					subJson, err := json.Marshal(subv)
					if err != nil {
						fmt.Println(err)
					}
					// json 转map
					subMenu := make(map[string]interface{})
					json.Unmarshal(subJson, &subMenu)

					subMeta := make(map[string]interface{})
					subMeta["title"] = subv.Title
					subMeta["icon"] = subv.Icon
					subMenu["meta"] = subMeta

					delete(subMenu, "title")
					delete(subMenu, "icon")
					delete(subMenu, "created_at")
					delete(subMenu, "updated_at")
					delete(subMenu, "deleted_at")

					delete(subMenu, "id")
					delete(subMenu, "parent_id")
					delete(subMenu, "sort")

					children = append(children, subMenu)
				}
			}
			fmt.Println("================children-len=========", len(children))
			m["children"] = children
			topMenu = append(topMenu, m)
		}
	}

	return topMenu
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Test is the golang structure of table hg_test for DAO operations like Where/Data.
type Test struct {
	g.Meta    `orm:"table:hg_test, do:true"`
	Id        interface{} //
	App       interface{} // app
	Owner     interface{} //
	Repo      interface{} // repo
	Version   interface{} // 版本号
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}

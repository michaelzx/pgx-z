// Code generated by pgx-z. DO NOT EDIT.
// Code generated by pgx-z. DO NOT EDIT.
// Code generated by pgx-z. DO NOT EDIT.
// Code generated at 2023-08-19 14:01:12.1067663 +0800 CST m=+0.169408301

package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// Domain mapped from table <user>
type User struct {
	No          string              `db:"no" udt:"varchar" json:"no"`                       // 编号
	DingId      string              `db:"ding_id" udt:"varchar" json:"ding_id"`             // 钉钉的UserID
	DeptIds     []int64             `db:"dept_ids" udt:"_int8" json:"dept_ids"`             // 所在部门ID
	WorkNo      string              `db:"work_no" udt:"varchar" json:"work_no"`             // 工号
	RealName    string              `db:"real_name" udt:"varchar" json:"real_name"`         // 姓名
	AvatarUrl   string              `db:"avatar_url" udt:"varchar" json:"avatar_url"`       // 头像地址
	Email       string              `db:"email" udt:"varchar" json:"email"`                 // 邮箱
	Phone       string              `db:"phone" udt:"varchar" json:"phone"`                 // 手机
	TeamNo      string              `db:"team_no" udt:"varchar" json:"team_no"`             // 所属团队编号
	Leader      bool                `db:"leader" udt:"bool" json:"leader"`                  // 是否是领导角色
	LeadExtra   []string            `db:"lead_extra" udt:"_varchar" json:"lead_extra"`      // 额外领导的下属（附加）
	LeadOmitted []string            `db:"lead_omitted" udt:"_varchar" json:"lead_omitted"`  // 额外领导的下属（排除）
	DomainScope []string            `db:"domain_scope" udt:"_varchar" json:"domain_scope"`  // 领域范围
	TimeCreate  pgtype.Timestamptz  `db:"time_create" udt:"timestamptz" json:"time_create"` // 创建时间
	TimeLogin   *pgtype.Timestamptz `db:"time_login" udt:"timestamptz" json:"time_login"`   // 登录时间
	Settings    []byte              `db:"settings" udt:"jsonb" json:"settings"`             // 用户个性化设置（比如时区）
	AuthStatus  bool                `db:"auth_status" udt:"bool" json:"auth_status"`        // 授权状态
}

const TableNameUser = "user"

// TableName User's table name
func (User) TableName() string {
	return TableNameUser
}

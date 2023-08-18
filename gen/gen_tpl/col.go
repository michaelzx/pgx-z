package gen_tpl

//
// import "github.com/michaelzx/pgx-z/gen/tmp/model"
//
// func User(m *model.User) user {
// 	mapping := make(user)
// 	if m != nil {
// 		mapping.setIfNotNil("dept_ids", m.DeptIds)
// 		mapping.setIfNotNil("work_no", m.WorkNo)
// 	}
// 	return mapping
// }
//
// type user map[string]any
//
// func (u user) setIfNotNil(fn string, v any) {
// 	if v != nil {
// 		u[fn] = v
// 	}
// }
//
// func (r user) Names() []string {
// 	return []string{
// 		"no", "ding_id",
// 	}
// }
//
// func (r user) No(No string) user {
// 	r["no"] = No
// 	return r
// }
// func (r user) DingID(DingID string) user {
// 	r["ding_id"] = DingID
// 	return r
// }

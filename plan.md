# Plan

## 1. gen model code for tabled
- [ ] status
```go
package model

type User struct {
	No     string `db:"no" udt:"xxx" json:"no"`           // comment1
	DingID string `db:"ding_id" udt:"xxx" json:"ding_id"` // comment2
}

```
## 2. gen cols code for table
dependence on model
- [ ] status
```go
package cols


func User(m *model.User) user {
	mapping := make(user)
	if m != nil {
		mapping.setIfNotNil("dept_ids", m.DeptIds)
		mapping.setIfNotNil("work_no", m.WorkNo)
	}
	return mapping
}

type user map[string]any

func (u user) setIfNotNil(fn string, v any) {
	if v != nil {
		u[fn] = v
	}
}

func (r user) Names() []string {
	return []string{
		"no", "ding_id",
	}
}

func (r user) No(No string) user {
	r["no"] = No
	return r
}
func (r user) DingID(DingID string) user {
	r["ding_id"] = DingID
	return r
}

```
## 3. usage
todo
```go
package repo

func UpdateByDingId(dingId string, user *model.User) error {
	updateColumns := cols.User(nil).
		RealName("xxxx").
		Email(zutil.P("xxx"))
	err := pg.Update[model.User](r.db).Set(updateColumns).Where("id=?", dingId).Commit()
	return err
}
```
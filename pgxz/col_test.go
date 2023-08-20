package pgxz

import "testing"

func Benchmark(b *testing.B) {

	keys := []string{
		"no", "ding_id", "dept_ids", "work_no", "real_name", "avatar_url", "email", "phone", "team_no", "leader", "lead_extra", "lead_omitted", "domain_scope", "time_create", "time_login", "settings", "auth_status",
	}
	keyMap := map[string]any{
		"no":           struct{}{},
		"ding_id":      struct{}{},
		"dept_ids":     struct{}{},
		"work_no":      struct{}{},
		"real_name":    struct{}{},
		"avatar_url":   struct{}{},
		"email":        struct{}{},
		"phone":        struct{}{},
		"team_no":      struct{}{},
		"leader":       struct{}{},
		"lead_extra":   struct{}{},
		"lead_omitted": struct{}{},
		"domain_scope": struct{}{},
		"time_create":  struct{}{},
		"time_login":   struct{}{},
		"settings":     struct{}{},
		"auth_status":  struct{}{},
	}
	var find1 = func(fk string) bool {
		for _, key := range keys {
			if key == fk {
				return true
			}
		}
		return false
	}
	var find2 = func(fk string) bool {
		_, exists := keyMap[fk]
		return exists
	}
	benchmarks := []struct {
		name     string
		findKey  string
		findFunc func(fk string) bool
	}{
		{"list-no", "no", find1},
		{"list-team_no", "team_no", find1},
		{"list-auth_status", "auth_status", find1},
		{"map-no", "no", find2},
		{"map-team_no", "team_no", find2},
		{"map-auth_status", "auth_status", find2},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.findFunc(bm.findKey)
			}
		})
	}
}

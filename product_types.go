package defectdojo

type ProductTypes struct {
	Count    int `json:"count"`
	Next     any `json:"next"`
	Previous any `json:"previous"`
	Results  []struct {
		ID                  int    `json:"id"`
		Name                string `json:"name"`
		Description         any    `json:"description"`
		CriticalProduct     bool   `json:"critical_product"`
		KeyProduct          bool   `json:"key_product"`
		Updated             any    `json:"updated"`
		Created             any    `json:"created"`
		Members             []int  `json:"members"`
		AuthorizationGroups []any  `json:"authorization_groups"`
	} `json:"results"`

	Prefetch struct {
	} `json:"prefetch"`
}

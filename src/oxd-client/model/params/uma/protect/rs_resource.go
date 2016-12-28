package protect

type RsResource struct {
	Path string `json:"path"`
	Conditions []Condition `json:"conditions"`
}

package uma

type RpGetGatRequestParams struct {
	OxdId string `json:"oxd_id"`
	Username []string `json:"scopes"`
}

type RpGetGatResponseParams struct {
	Rpt string `json:"rpt"`
}

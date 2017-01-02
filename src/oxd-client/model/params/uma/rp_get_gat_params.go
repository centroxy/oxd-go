//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package uma

type RpGetGatRequestParams struct {
	OxdId string `json:"oxd_id"`
	Scopes []string `json:"scopes"`
}

type RpGetGatResponseParams struct {
	Rpt string `json:"rpt"`
}

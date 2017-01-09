//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 09/01/17
//
package uma

type RpAuthorizeRptRequestParams struct {
	OxdId string `json:"oxd_id"`
	Rpt string `json:"rpt"`
	Ticket string `json:"ticket"`
}

type RpAuthorizeRptResponseParams struct {
	OxdId string `json:"oxd_id"`
}

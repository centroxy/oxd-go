//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package uma

type RpGetRptRequestParams struct {
	OxdId string `json:"oxd_id"`
	ForceNew bool `json:"force_new"`
}

type RpGetRptResponseParams struct {
	Rpt string `json:"rpt"`
}


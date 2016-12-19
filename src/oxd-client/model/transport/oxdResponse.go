package transport

import "oxd-client/constants"

type OxdResponse struct {
	Status constants.Status `json:"status"`
	Params Param `json:"data"`
}

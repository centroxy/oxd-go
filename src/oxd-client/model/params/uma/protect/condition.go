package protect

type Condition struct {
	HttpMethods []string `json:"httpMethods"`
	Scopes []string `json:"scopes"`
	TicketScopes []string `json:"ticketScopes"`
}

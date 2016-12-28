package constants

type CommandType string

const (
	// Register
	REGISTER_SITE CommandType = "register_site"
	UPDATE_SITE   CommandType = "update_site_registration"

	// Connect (stateful)
	GET_AUTHORIZATION_URL  CommandType = "get_authorization_url"
	GET_AUTHORIZATION_CODE CommandType = "get_authorization_code"
	GET_TOKENS_BY_CODE     CommandType = "get_tokens_by_code"
	GET_USER_INFO          CommandType = "get_user_info"
	GET_LOGOUT_URI         CommandType = "get_logout_uri"

	CHECK_ID_TOKEN     CommandType = "id_token_status"
	CHECK_ACCESS_TOKEN CommandType = "access_token_status"
	LICENSE_STATUS     CommandType = "license_status"

	// UMA
	RS_PROTECT       CommandType = "uma_rs_protect"
	RS_CHECK_ACCESS  CommandType = "uma_rs_check_access"
	RP_GET_RPT       CommandType = "uma_rp_get_rpt"
	RP_GET_GAT       CommandType = "uma_rp_get_gat"
	RP_AUTHORIZE_RPT CommandType = "uma_rp_authorize_rpt"

	// stateless
	AUTHORIZATION_CODE_FLOW CommandType = "authorization_code_flow"
	IMPLICIT_FLOW           CommandType = "implicit_flow"
)

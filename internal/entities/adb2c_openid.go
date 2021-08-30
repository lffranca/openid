package entities

type ADB2COpenID struct {
	TokenEndpoint                     string   `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported []string `json:"token_endpoint_auth_methods_supported"`
	JwksURI                           string   `json:"jwks_uri"`
	ResponseModesSupported            []string `json:"response_modes_supported"`
	SubjectTypesSupported             []string `json:"subject_types_supported"`
	IDTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
	ResponseTypesSupported            []string `json:"response_types_supported"`
	ScopesSupported                   []string `json:"scopes_supported"`
	Issuer                            string   `json:"issuer"`
	RequestURIParameterSupported      bool     `json:"request_uri_parameter_supported"`
	UserinfoEndpoint                  string   `json:"userinfo_endpoint"`
	AuthorizationEndpoint             string   `json:"authorization_endpoint"`
	DeviceAuthorizationEndpoint       string   `json:"device_authorization_endpoint"`
	HTTPLogoutSupported               bool     `json:"http_logout_supported"`
	FrontchannelLogoutSupported       bool     `json:"frontchannel_logout_supported"`
	EndSessionEndpoint                string   `json:"end_session_endpoint"`
	ClaimsSupported                   []string `json:"claims_supported"`
	KerberosEndpoint                  string   `json:"kerberos_endpoint"`
	TenantRegionScope                 string   `json:"tenant_region_scope"`
	CloudInstanceName                 string   `json:"cloud_instance_name"`
	CloudGraphHostName                string   `json:"cloud_graph_host_name"`
	MsgraphHost                       string   `json:"msgraph_host"`
	RbacURL                           string   `json:"rbac_url"`
}

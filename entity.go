package lumada

type LoginRequest struct {
	GtantType string
	ClientId  string
	Username  string
	Password  string
	scope     string
}

type LoginResponse struct {
	AccessToken  string              `json:"access_token"`
	TokenType    string              `json:"token_type"`
	RefreshToken string              `json:"refresh_token"`
	ExpiresIn    int                 `json:"expires_in"`
	Scope        string              `json:"scope"`
	SessionId    string              `json:"session_id"`
	Entity       LoginResponseEntity `json:"entity"`
	Jti          string              `json:"jti"`
}

type LoginResponseEntity struct {
	EntityId             string   `json:"entityId"`
	State                string   `json:"state"`
	EntityType           string   `json:"entityType"`
	EntityValidationType string   `json:"entityValidationType"`
	EntityRole           string   `json:"entityRole"`
	EntityPrivileges     []string `json:"entityPrivileges"`
	EntityValue          string   `json:"entityValue"`
}

type Asset struct {
	Id          string            `json:"id"`
	Version     int               `json:"version"`
	Name        string            `json:"name"`
	GatewayId   string            `json:"gatewayId"`
	AssetTypeId string            `json:"assetTypeId"`
	Properties  []AssetProperties `json:"properties"`
	Created     uint64            `json:"created"`
	Modified    uint64            `json:"modified"`
}

type AssetProperties struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

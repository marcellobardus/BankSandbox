package login

type LoginAccountDrt struct {
	SessionToken *string `json:"token"`
	ErrorCode    int     `json:"errorCode"`
	Error        bool    `json:"error"`
}

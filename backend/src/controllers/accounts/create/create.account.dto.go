package create

type CreateAccountDto struct {
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	SocialInsuranceID string `json:"socialInsuranceID"`
	PhoneNumber       string `json:"phoneNumber"`
	Mail              string `json:"mail"`

	PasswordHash string `json:"passwordHash"`
}

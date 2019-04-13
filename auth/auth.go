package auth

type PexelSessionObj struct {
	API_KEY string `json:"apikey"`
}

var PexelSession *PexelSessionObj = &PexelSessionObj{}

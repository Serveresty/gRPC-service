package models

type ConnectionData struct {
	IP   string `json:"IP"`
	Port string `json:"port"`
}

type ConnectionArr struct {
	ConData ConnectionData `json:"connectionData"`
}

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthArr struct {
	AuData AuthData `json:"authData"`
}

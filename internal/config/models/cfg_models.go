package models

type ConnectionData struct {
	IP   string `json:"IP"`
	Port string `json:"port"`
}

type ServerConnection struct {
	ConData        ConnectionData `json:"serverConnectionData"`
	MaxWorkerCount int            `json:"max-worker-count"`
	MaxQueueSize   int            `json:"max-queue-size"`
}

type ClientConnection struct {
	ConData ConnectionData `json:"clientConnectionData"`
}

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthArr struct {
	AuData AuthData `json:"authData"`
}

type SecretKey struct {
	SecretKey string `json:"secret-key"`
}

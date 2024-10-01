package types

type key string

type VWS struct {
	Env      string
	BotToken string
}

const InitDataKey key = "initdatakey"

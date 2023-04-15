package models

type Profile struct {
	ETHAddress []string
	Username   string
	Metadata   map[string]string
	Tags       []string
}

package model

type RepoQueryUtils struct {
	MapRoleTableOmitField      map[string][]string          `json:"mapRoleTableOmitField"`
	MapRoleTableValidateColumn map[string]map[string]string `json:"mapRoleTableValidateColumn"`
	MapModelTypeFieldPreload   map[string]map[string]string `json:"mapModelTypeFieldPreload"`
	ListOmitField              []string                     `json:"listOmitField"`
}

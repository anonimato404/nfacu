package cmd

type NFACUConfigObject struct {
	Path     string            `json:"path"`
	Settings map[string]string `json:"settings"`
}

type NFACUConfig []NFACUConfigObject

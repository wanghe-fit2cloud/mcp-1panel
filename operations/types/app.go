package types


type AppInstallCreate struct {
	AppDetailID uint                   `json:"appDetailId"`
	Params      map[string]interface{} `json:"params"`
	Name        string                 `json:"name"`
}

type AppRes struct {
	Response
	Data App `json:"data"`
}

type App struct {
	ID       uint   `json:"id"`
	Versions []string `json:"versions"`
}

type AppDetailRes struct {
	Response
	Data AppDetail `json:"data"`
}

type AppDetail struct {
	ID       uint   `json:"id"`
}
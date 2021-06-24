package godashboard

type Signal struct {
	ItemName string           `json:"item_name" bson:"item_name"`
	Params   map[string]Param `json:"params" bson:"params"`
}

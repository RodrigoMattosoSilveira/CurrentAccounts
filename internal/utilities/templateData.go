package utilities

type T_Data struct {
	Data map[string]any
}

func NewTemplateData() *T_Data {
	return &T_Data{Data: map[string]any{
		"Tenant": "MC",
		"Host":   "Madrone Logistics",
	}}
}

func (tdata *T_Data) AddData(key string, value any) {
	tdata.Data[key] = value
}

func (tdata *T_Data) GetAllData() map[string]any {
	return tdata.Data
}

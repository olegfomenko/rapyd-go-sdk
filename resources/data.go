package resources

type Data map[string]interface{}

func (d Data) GetId() string {
	if id, ok := d["id"]; ok {
		return id.(string)
	} else {
		return ""
	}
}

func (d Data) GetString(key string) string {
	if id, ok := d[key]; ok {
		return id.(string)
	} else {
		return ""
	}
}

func (d Data) GetInt64(key string) int64 {
	if id, ok := d[key]; ok {
		return id.(int64)
	} else {
		return 0
	}
}

func (d Data) GetData(key string) map[string]interface{} {
	if id, ok := d[key]; ok {
		return id.(Data)
	} else {
		return make(map[string]interface{})
	}
}

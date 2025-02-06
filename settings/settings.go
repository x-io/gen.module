package settings

var settings map[string]interface{}

func init() {
	settings = make(map[string]interface{})
}

//Set ...
func Set(name string, value interface{}) error {
	if value != nil {
		settings[name] = value
	}
	return nil
}

//Get ...
func Get(name string, defaults ...interface{}) interface{} {
	if v, ok := settings[name]; ok {
		return v
	}
	if len(defaults) > 0 {
		return defaults[0]
	}

	return nil
}

//GetString ...
func GetString(name string, defaults ...string) string {
	if v, ok := settings[name]; ok {
		if vv, ok := v.(string); ok && vv != "" {
			return vv
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}

	return ""
}

//GetMapString ...
func GetMapString(name, key string, defaults ...string) string {
	if v, ok := settings[name]; ok {
		if vv, ok := v.(map[string]string); ok && vv != nil {
			return vv[key]
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}

	return ""
}

//GetBool ...
func GetBool(name string, defaults ...bool) bool {
	if v, ok := settings[name]; ok {
		if vv, ok := v.(bool); ok {
			return vv
		}
	}
	if len(defaults) > 0 {
		return defaults[0]
	}

	return false
}

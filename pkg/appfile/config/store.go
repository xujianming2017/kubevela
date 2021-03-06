package config

import "errors"

// Store will get config data
type Store interface {
	GetConfigData(configName, envName string) ([]map[string]string, error)
	Type() string
	Namespace(envName string) (string, error)
}

// TypeFake is a fake type
const TypeFake = "fake"

var _ Store = &Fake{}

// Fake is a fake implementation of config store, help for test
type Fake struct {
	Data []map[string]string
}

// GetConfigData get data
func (f *Fake) GetConfigData(_ string, _ string) ([]map[string]string, error) {
	return f.Data, nil
}

// Type return the type
func (Fake) Type() string {
	return TypeFake
}

// Namespace return the Namespace
func (Fake) Namespace(_ string) (string, error) {
	return "", nil
}

// EncodeConfigFormat will encode key-value to config{name: key, value: value} format
func EncodeConfigFormat(key, value string) map[string]string {
	return map[string]string{
		"name":  key,
		"value": value,
	}
}

// DecodeConfigFormat will decode config{name: key, value: value} format to key-value mode.
func DecodeConfigFormat(data []map[string]string) (map[string]string, error) {
	var res = make(map[string]string)
	for _, d := range data {
		key, ok := d["name"]
		if !ok {
			return nil, errors.New("invalid data format, no 'name' found")
		}
		value, ok := d["value"]
		if !ok {
			return nil, errors.New("invalid data format, no 'value' found")
		}
		res[key] = value
	}
	return res, nil
}

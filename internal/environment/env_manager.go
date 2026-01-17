package environment

import (
	"fmt"
	"os"
	"strconv"
)

type EnvManager struct{}

func NewEnvManager() *EnvManager {
	return &EnvManager{}
}

func (e *EnvManager) Get(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("env var %s not set", key)
	}
	return val, nil
}

func (e *EnvManager) GetInt(key string) (int, error) {
	val, err := e.Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

func (e *EnvManager) GetBool(key string) (bool, error) {
	val, err := e.Get(key)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(val)
}

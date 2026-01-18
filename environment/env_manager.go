package environment

import (
	"fmt"
	"os"
	"strconv"
)

func Get(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("env var %s not set", key)
	}
	return val, nil
}

func GetInt(key string) (int, error) {
	val, err := Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(val)
}

func GetBool(key string) (bool, error) {
	val, err := Get(key)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(val)
}

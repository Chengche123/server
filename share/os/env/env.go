package env

import (
	"fmt"
	"log"
	"os"
)

func GetEnvOrDefault(env string, def string) string {
	v, ok := os.LookupEnv(env)
	if !ok {
		log.Printf("cannot find %s, use default value %s\n", env, def)
		return def
	}
	log.Printf("%s: %s\n", env, v)

	return v
}

func FormatEnvOrDefault(formatter string, env string, def string) string {
	v := GetEnvOrDefault(env, def)
	return fmt.Sprintf(formatter, v)
}

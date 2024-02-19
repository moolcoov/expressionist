package lib

import "os"

func Getenv(key string, defaultValue string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return defaultValue
	}
	return value
}

// «Если заблудился в коде - иди домой»
//                  (c) Джейсон Стэтхэм

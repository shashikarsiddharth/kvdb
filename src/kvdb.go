package kvdb

import "strconv"

var db = make(map[string]string)

// Set value in KV DB
func Set(key string, value string) bool {
	db[key] = value
	return true
}

// Get value in KV DB
func Get(key string) (string, bool) {
	value, found := db[key]
	return value, found
}

// Del value in KV DB
func Del(key string) bool {
	if _, found := Get(key); found {
		delete(db, key)
		return true
	}
	return false
}

// Incr value by 1 in KV DB
func Incr(key string) bool {
	if value, found := Get(key); found {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			strValue := strconv.Itoa(intValue + 1)
			return Set(key, strValue)

		}
	}
	return Set(key, "1")
}

// IncrBy by given value in KV DB
func IncrBy(key string, incrValue int) bool {
	if value, found := Get(key); found {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			strValue := strconv.Itoa(intValue + incrValue)
			return Set(key, strValue)
		}
	}
	return false
}

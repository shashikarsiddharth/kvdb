package kvdb

import (
	"strconv"
	"strings"
)

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

// Compact view of KV DB
func Compact() map[string]string {
	return db
}

// Multi will create a buffer and will execute elements based on it
func Multi(queue []string) []string {
	var res []string
	for _, cmd := range queue {
		args := strings.Split(strings.Trim(cmd, "\n"), " ")
		if args[0] == "SET" {
			res = append(res, strconv.FormatBool(Set(args[1], args[2])))
		} else if args[0] == "GET" {
			value, found := Get(args[1])
			if found == false {
				res = append(res, strconv.FormatBool(found))
			} else {
				res = append(res, value)
			}
			res = append(res)
		} else if args[0] == "DEL" {
			res = append(res, strconv.FormatBool(Del(args[1])))
		} else if args[0] == "INCR" {
			res = append(res, strconv.FormatBool(Incr(args[1])))
		} else if args[0] == "INCRBY" {
			incrValue, _ := strconv.Atoi(args[2])
			res = append(res, strconv.FormatBool(IncrBy(args[1], incrValue)))
		}

	}
	return res
}

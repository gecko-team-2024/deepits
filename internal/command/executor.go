package command

import (
	"deepits/internal/database"
	"strconv"
	"strings"
)

// xu li lenh tu client
func ExecuteCommand(db *database.Database, command string) string {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "ERR: Empty command"
	}

	switch strings.ToUpper(parts[0]) {
	case "SET":
		if len(parts) < 3 {
			return "ERR: Usage: SET key value [EX seconds]"
		}
		ttl := int64(0)
		if len(parts) == 5 && strings.ToUpper(parts[3]) == "EX" {
			t, err := strconv.Atoi(parts[4])
			if err == nil {
				ttl = int64(t)
			}
		}
		db.Set(parts[1], parts[2], ttl)
		return "OK"
	case "GET":
		if len(parts) != 2 {
			return "ERR: Usage: GET key"
		}
		value, exists := db.Get(parts[1])
		if !exists {
			return "(nil)"
		}
		return value
	case "DEL":
		if len(parts) != 2 {
			return "ERR: Usage: DEL key"
		}
		db.Delete(parts[1])
		return "OK"
	default:
		return "ERR: Unknown command"
	}
}

package permission

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rxdn/gdl/objects/member"
	"strconv"
	"time"
)

const timeout = time.Minute * 5

func GetPermissionLevel(redis *redis.Client, member member.Member, guildId uint64) PermissionLevel {
	key := fmt.Sprintf("permissions:%d:%d", guildId, member.User.Id)
	res, err := redis.Get(key).Result(); if err != nil {
		return Everyone
	}

	parsed, err := strconv.Atoi(res); if err != nil {
		return Everyone
	}

	return PermissionLevel(parsed)
}

func SetPermissionLevel(redis *redis.Client, member member.Member, guildId uint64, level PermissionLevel) error {
	key := fmt.Sprintf("permissions:%d:%d", guildId, member.User.Id)
	return redis.Set(key, level.Int(), timeout).Err()
}
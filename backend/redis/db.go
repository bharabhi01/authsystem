package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

func IsUsernamePresentInCache(username string) (bool, error) {
	log.Println("Checking username in cache:", username)
	if client == nil {
		return false, errors.New("redis client not initialized")
	}

	key := fmt.Sprintf("username:%s", username)
	ctx := context.Background()

	_, err := client.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func StoreUsernameInCache(username string) error {
	log.Println("Storing username in cache:", username)
	if client == nil {
		return errors.New("redis client not initialized")
	}

	key := fmt.Sprintf("username:%s", username)
	ctx := context.Background()

	return client.Set(ctx, key, "1", 24*time.Hour).Err()
}

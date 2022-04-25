package store

import (
	"errors"
	"log"
	"regexp"

	"github.com/go-redis/redis"
)

type StoreRedis struct {
	redis *redis.Client
	r     *regexp.Regexp
}

type EntityStore interface {
	Persist(e Entity, prefix string) error
	GetById(id string, prefix string) (string, error)
	GetAll(prefix string) ([]string, error)
}

func NewStoreRedis() (*StoreRedis, error) {
	sr := &StoreRedis{
		redis: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
		r: regexp.MustCompile("[a-zA-Z]+"),
	}

	return sr, sr.Ping()
}

func (r *StoreRedis) GetAll(prefix string) ([]string, error) {
	if ok := r.r.MatchString(prefix); ok {
		keys, err := r.redis.Keys(prefix + "-*").Result()
		if err != nil {
			return nil, err
		}
		array := make([]string, 0)
		for _, k := range keys {
			value, err := r.redis.Get(k).Result()
			if err != nil {
				log.Println("couldnt get key", err.Error())
				continue
			}
			array = append(array, value)
		}

		return array, nil
	}

	return nil, errors.New("wrong prefix")
}

func (r *StoreRedis) GetById(id string, prefix string) (string, error) {
	key := id
	if ok := r.r.MatchString(prefix); ok {
		key = prefix + "-" + key
	}

	value, err := r.redis.Get(key).Result()

	if err != nil {
		return "", err
	}

	return value, nil
}

func (r *StoreRedis) Persist(entity Entity, prefix string) error {
	json, err := entity.ConvertToJson()

	if err != nil {
		return err
	}
	key := entity.GetId()
	if ok := r.r.MatchString(prefix); ok {
		key = prefix + "-" + key
	}
	log.Println("save to redis with key:" + key)
	return r.Put(key, string(json))
}

func (r *StoreRedis) Put(key string, value string) error {
	return r.redis.Set(key, value, 0).Err()
}

func (r *StoreRedis) Gut(key string, value string) (string, error) {
	return r.redis.Get(key).Result()
}

func (r *StoreRedis) Ping() error {
	return r.redis.Ping().Err()
}

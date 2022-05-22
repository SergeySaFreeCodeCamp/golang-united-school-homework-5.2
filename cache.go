package cache

import "time"

type Value struct {
	value    string
	deadline time.Time
}

type Cache struct {
	cashe map[string]Value
}

func NewCache() Cache {
	return Cache{cashe: make(map[string]Value)}
}

func (c Cache) Get(key string) (string, bool) {
	v, ok := c.cashe[key]
	if v.deadline.IsZero() || time.Until(v.deadline) > 0 {
		return v.value, ok
	}
	if !v.deadline.IsZero() && time.Until(v.deadline) <= 0 {
		delete(c.cashe, key)
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.cashe[key] = Value{value: value}
}

func (c Cache) Keys() []string {
	res := make([]string, 0)
	for k, v := range c.cashe {
		if v.deadline.IsZero() || time.Until(v.deadline) > 0 {
			res = append(res, k)
		}
		if !v.deadline.IsZero() && time.Until(v.deadline) <= 0 {
			delete(c.cashe, k)
		}
	}
	return res
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.cashe[key] = Value{value: value, deadline: deadline}
}

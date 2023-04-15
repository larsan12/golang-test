package cache

import (
	"sync"
	"time"
)

type Cache[T any] interface {
	Get(key string) (T, bool)
	Set(key string, value T)
	Remove(key string)
}

type StatsReceiver interface {
	SettingInc()
	HitInc()
	ReqInc()
	ErrorInc()
}

// одна запись кэша
type unit[T any] struct {
	key       string
	value     T
	expiredAt time.Time // время экспирации
	alive     bool      // true если не удалён
	next      *unit[T]  // ссылка на созданный следующую созданную единицу кэша (если есть ttl иначе nil)
}

type implementation[T any] struct {
	mx            sync.RWMutex        // мютекс для потокобезопасности
	m             map[string]*unit[T] // сам кэш
	ttl           int                 // TTL в милисекундах
	statsReceiver StatsReceiver       // для метрик или статитстики
	oldestUnit    *unit[T]            // самая старая единица кэша (если есть ttl иначе nil)
	lastUnit      *unit[T]            // последняя созданная единица кэша (если есть ttl иначе nil)
}

func New[T any](ttl int, statsReceiver StatsReceiver) Cache[T] {
	cache := &implementation[T]{
		m:             make(map[string]*unit[T]),
		ttl:           ttl,
		statsReceiver: statsReceiver,
	}
	if ttl > 0 {
		go cache.runExpireLoop() // бесконечный цикл проверки TTL
	}
	return cache
}

// в бесконечном циуле проверяем условие экспирации для самого старого созданного кэша
func (c *implementation[T]) runExpireLoop() {
	for {
		if c.oldestUnit != nil {
			if c.oldestUnit.alive {
				if time.Now().After(c.oldestUnit.expiredAt) { //  проверяем TTL
					c.Remove(c.oldestUnit.key)
					c.oldestUnit = c.oldestUnit.next
				}
			} else if c.oldestUnit.next != nil {
				c.oldestUnit = c.oldestUnit.next
			}
		}
	}
}

func (c *implementation[T]) Get(key string) (T, bool) {
	c.statsReceiver.ReqInc()
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	if !ok {
		var res T
		return res, false
	}

	c.statsReceiver.HitInc()
	return val.value, ok
}

func (c *implementation[T]) Set(key string, value T) {
	c.statsReceiver.SettingInc()
	c.mx.Lock()
	defer c.mx.Unlock()
	newUnit := &unit[T]{
		key:   key,
		value: value,
		alive: true,
	}
	oldVal, ok := c.m[key]
	if ok {
		oldVal.alive = false
	}
	c.m[key] = newUnit

	if c.ttl > 0 {
		newUnit.expiredAt = time.Now().Add(time.Duration(c.ttl) * time.Millisecond)
		if c.oldestUnit == nil {
			c.oldestUnit = newUnit
		}
		if c.lastUnit != nil {
			c.lastUnit.next = newUnit
		}
		c.lastUnit = newUnit
	}
}

func (c *implementation[T]) Remove(key string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.m[key]
	if ok {
		val.alive = false
	} else {
		c.statsReceiver.ErrorInc()
	}
	delete(c.m, key)
}

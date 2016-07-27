package latency

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// EtcdHTTPHandlerConfig is a HTTPHandlerConfig configed through etcd. It uses an
// exponential distribution to choose latency.
// The distribution is a number between min and max.
type EtcdHTTPHandlerConfig struct {
	key    string
	keyAPI client.KeysAPI

	store *Store
}

// NewEtcdHTTPHandlerConfig returns a new EtcdHTTPHandlerConfig which is a type of HTTPHandlerConfig
func NewEtcdHTTPHandlerConfig(key string, min, max time.Duration) EtcdHTTPHandlerConfig {
	cfg := client.Config{
		Endpoints: []string{ConfigEtcdURL()},
		Transport: client.DefaultTransport,
	}

	c, err := client.New(cfg)
	if err != nil {
		panic("Error in etcd config. Could not setup etcd client.")
	}

	keyAPI := client.NewKeysAPI(c)

	handler := EtcdHTTPHandlerConfig{
		key:    key,
		keyAPI: keyAPI,
		store:  NewStore(min, max),
	}

	handler.setup()

	ch := handler.watch()

	go func() {
		for {
			node := <-ch
			handler.selfUpdate(node)
			fmt.Println(node.Key, node.Value)
		}
	}()

	return handler
}

// GetLatency returns a number between min/max using exponential distribution
func (e EtcdHTTPHandlerConfig) GetLatency() *time.Duration {
	min := math.Max(e.store.GetMin().Seconds(), rand.ExpFloat64())
	actual := math.Min(min, e.store.GetMax().Seconds())
	dur := time.Duration(actual) * time.Second
	return &dur
}

func (e EtcdHTTPHandlerConfig) setup() {
	res, err := e.keyAPI.Get(context.Background(), e.key, nil)

	if err != nil && err.(client.Error).Code == client.ErrorCodeKeyNotFound {
		e.register()
	} else if err != nil {
		panic(err)
	} else {
		for _, n := range res.Node.Nodes {
			e.selfUpdate(*n)
		}
	}
}

func (e EtcdHTTPHandlerConfig) selfUpdate(n client.Node) {
	switch parseKey(n.Key) {
	case "max":
		val, err := strconv.ParseInt(n.Value, 10, 64)
		if err != nil {
			panic(err)
		}
		e.store.SetMax(millis2Nano(val))
	case "min":
		val, err := strconv.ParseInt(n.Value, 10, 64)
		if err != nil {
			panic(err)
		}
		e.store.SetMin(millis2Nano(val))
	}
	fmt.Println(n.Key, n.Value)
}

func parseKey(key string) string {
	spl := strings.Split(key, "/")
	return spl[len(spl)-1]
}

func (e EtcdHTTPHandlerConfig) register() {
	_, err := e.keyAPI.Create(context.Background(),
		fmt.Sprintf("%v/min", e.key), strconv.FormatInt(millis(e.store.GetMin()), 10))

	if err != nil {
		panic(err)
	}

	_, err = e.keyAPI.Create(context.Background(),
		fmt.Sprintf("%v/max", e.key), strconv.FormatInt(millis(e.store.GetMax()), 10))

	if err != nil {
		panic(err)
	}
}

func millis(n time.Duration) int64 {
	return n.Nanoseconds() / 1000000
}

func millis2Nano(millis int64) time.Duration {
	return time.Duration(millis*1000000) * time.Nanosecond
}

func (e EtcdHTTPHandlerConfig) watch() <-chan client.Node {
	c := make(chan client.Node)
	watcher := e.keyAPI.Watcher(e.key, &client.WatcherOptions{Recursive: true})

	go func(watcher client.Watcher) {
		for {
			res, err := watcher.Next(context.Background())
			if err != nil {
				// Perform backoff later
				panic(err)
			}

			c <- *res.Node
		}
	}(watcher)

	return c
}

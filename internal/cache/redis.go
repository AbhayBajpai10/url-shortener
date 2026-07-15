package cache
import 
(
	"context"
	"fmt"
	"github.com/AbhayBajpai10/url-shortener/config"
	"github.com/redis/go-redis/v9"
)
var Ctx= context.Background()
func ConnectRedis (cfg *config.Config) *redis.Client{
	addr:=fmt.Sprintf(
		"%s:%s",
		cfg.RedisHost,
		cfg.RedisPort,
	)
	client:=redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_,err:=client.Ping(Ctx).Result()
	if err!=nil {
		panic(err)
	}
	fmt.Println("Connected to Redis")
	return client
}
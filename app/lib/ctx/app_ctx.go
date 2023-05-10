package ctx

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"github.com/bwmarrin/snowflake"
	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
	"gomock/app/lib/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type AppContext struct {
	MysqlClient   *sql.DB
	HttpClient    *resty.Client
	InMemoryCache *bigcache.BigCache
	RedisClient   *redis.Client
	IdGen         *snowflake.Node
}

var (
	datetimePrecision = 2
	AppCtx            = &AppContext{}
)

func NewAppContext(config config.Config) {
	setUpMysql(config)
	setUpCache()
	setUpIdGen()
	//setUpRedis()
	return
}

func setUpMysql(config config.Config) {
	db, err := sql.Open("mysql", config.Mysql.Dsn)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	AppCtx.MysqlClient = db
}

func setUpCache() {
	cache, e := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	if e != nil {
		panic(e)
	}

	AppCtx.InMemoryCache = cache
}

func setUpIdGen() {
	// todo
	node, err := snowflake.NewNode(1)

	if err != nil {
		return
	}

	AppCtx.IdGen = node
}

func (ctx *AppContext) DestroyAppCtx() {
	ctx.MysqlClient.Close()
	fmt.Println("destroy ok")
}

package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    "github.com/amichelins/amsrtl"
    storage_redis "github.com/amichelins/amsrtl/storage/redis"
    "github.com/redis/go-redis/v9"
)

func main() {
    os.Setenv("LIMITER_MAX", "5")
    os.Setenv("LIMITER_BLOCK_DURATION", "30")
    os.Setenv("LIMITER_TOKENS", `[{"token":"TOKENA","limit": 10},{"token":"TOKENB","limit": 10}]`)

    redisCli := redis.NewClient(&redis.Options{Addr: "redis:6379"})
    stdRedis := storage_redis.NewRedisStorage(redisCli, true)
    rateLimit := amsrtl.NewEnvLimiter(stdRedis)

    http.Handle("/", amsrtl.Handle(rateLimit, http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            var out any
            w.Header().Set("Content-Type", "application/json")
            _ = json.Unmarshal([]byte(`{"msg": "Chamada realizada com sucesso   "}`), &out)
            _ = json.NewEncoder(w).Encode(out)
        }),
    ))

    log.Println("listening on port :8080")

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

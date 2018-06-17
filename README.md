# go-clash

Clash Royale internal API bindings for Go.

## Usage

```
base, _ := url.Parse("Base URI")

client := clash.Client{
    BaseURL:   base,
    Bearer:    "Your bearer key",
    UserAgent: "StatsRoyale",
}

chests, _ := client.GetPlayerUpcomingChests("9PLJLPQ8G")
fmt.Println(chests)

log, _ := client.GetPlayerBattleLog("9PLJLPQ8G")
fmt.Println(log)
```

## Error handling

Any issues with HTTP transport or response codes >=400 will be reflected in the returned error.
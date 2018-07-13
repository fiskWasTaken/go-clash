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

chests, _ := client.Player().UpcomingChests("9PLJLPQ8G") // "#9PLJLPQ8G" is also fine
fmt.Println(chests)

log, _ := client.Player().BattleLog("9PLJLPQ8G")
fmt.Println(log)
```

## Error handling

Any issues with HTTP transport or response codes >=400 will be reflected in the returned error.
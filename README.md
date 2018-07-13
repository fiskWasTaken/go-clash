# go-clash

API bindings for the public Clash Royale API, for Go.

Provides a fluent API

https://developer.clashroyale.com/#/documentation

## Usage

```
base, _ := url.Parse("Base URI")

client := clash.Client{
    BaseURL:   base,
    Bearer:    "Your bearer key",
    UserAgent: "StatsRoyale",
}

player := client.Player("9PLJLPQ8G") // "#9PLJLPQ8G" is also fine

chests, _ := player.UpcomingChests()
fmt.Println(chests)

log, _ := player.BattleLog()
fmt.Println(log)
```

## Error handling

Any issues with HTTP transport or response codes >=400 will be reflected in the returned error.

Cast the error to an API error structure. If casting wasn't successful, the error must be from net/http.

```
if err, ok := err.(clash.APIError); ok {
    // handle a clash API error response
}
```
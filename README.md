# timeutil

The util library to manipulate go time.

## `StartOf` 

Setting time to the start of a unit of time.

```go
now    = time.Date(2022, 10, 6, 21, 4, 34, 0, time.UTC)
StartOf(now, Y) // 2022-01-01 00:00:00
StartOf(now, M) // 2022-10-01 00:00:00
StartOf(now, D) // 2022-10-06 00:00:00
StartOf(now, H) // 2022-10-06 21:00:00
```

## `EndOf`

Setting time to the start of a unit of time.

```go
now    = time.Date(2022, 10, 6, 21, 4, 34, 0, time.UTC)
EndOf(now, Y) // 2022-12-31 23:59:59
EndOf(now, M) // 2022-10-31 23:59:59
EndOf(now, D) // 2022-10-06 23:59:59
EndOf(now, H) // 2022-10-06 21:59:59
```
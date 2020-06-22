# redis-healthcheck

## Usage

Compile

```bash
go get -d .
go build redis-healthcheck.go
```

Run the healthcheck

```bash
./redis-healthcheck:1323 <IP>:<PORT> <PASSWORD>
```

NOAUTH

```bash
./redis-healthcheck:1323 127.0.0.1:6390 ""
```

With AUTH

```bash
./redis-healthcheck:1323 127.0.0.1:6390 xyz123
OR
./redis-healthcheck:1323 127.0.0.1:6390 "xyz123"
```

## Response

200 (OK)

```bash
Redis OK
```

500 (Error)

```bash
Redis ERROR
```

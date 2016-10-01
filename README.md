# torus

Simple shareable torrent server.

Not production ready. Not even alpha software.

## Configuration

```toml
listen_addr = ":8000"
database_path = "/static/torrents.db"
data_path = "/static/data"
htpasswd_path = "/static/htpasswd"
download_token = 5
torrent_port = 50007

[mailgun]
domain = "example.com"
secret_key = "XXXXX"
public_key = "XXXXX"
email = "torus@example.com"
```

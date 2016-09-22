# torus

Simple shareable torrent server.

Not production ready. Not even alpha software.

## Configuration

```toml
database_path = "/static/torrents.db"
data_path = "/static/data"
htpasswd_path = "/static/htpasswd"
download_token = 5
torrent_port = 50007
web_port = 8000

[mailgun]
domain = "example.com"
secret_key = "XXXXX"
public_key = "XXXXX"
email = "torus@example.com"
```

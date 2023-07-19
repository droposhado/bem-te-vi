# bem-te-vi

Birthday notification based on contacts stored in CardDAV server.

The message is simple format so it can be bridged and delivered on IRC too.

![notification.png](notification.png)

## Environment variables

| Variable | Description | Required | Default |
| --------------- | --------------- | --------------- | --------------- |
| BTV_DEBUG | Show sensitive info about connection and other relevant things | N | |
| WEBDAV_SERVER | Server domain or subdomain url | Y | |
| WEBDAV_ADRESSBOOK | Path representing the address book in the DAV service | Y | |
| WEBDAV_USERNAME | Username used in DAV service | Y | |
| WEBDAV_PASSWORD | Password used in DAV service | Y | |
| WEBHOOK_URL | Slack compatible webhook | Y | |
| DATE_LAYOUT | Date parse mask | N | 2006-01-02T15:04:05Z |
| BOT_NAME | The name of the bot on chat | N | Bot defined |
| ICON_URL | Avatar of bot | N | Bot defined or service default user icon |


## Usage

### GitHub Actions

Simple use [bem-te-vi action](https://github.com/droposhado/bem-te-vi-action).

### Local

You need to export the variables with the values for use in the execution of the program

```
$ export VARIABLE=value
```

And run:

```
$ go run main.go
```

## License

See [LICENSE](LICENSE.md)

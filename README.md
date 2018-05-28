# Service blocks [![CircleCI](https://circleci.com/gh/remoteview/service-blocks.svg?style=svg)](https://circleci.com/gh/remoteview/service-blocks)

## Healthcheck
```http
GET /_health
```

##### Response `200 OK`
```js
{
  version: "0.0.0",
  status: "up"
}
```

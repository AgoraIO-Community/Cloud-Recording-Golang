# Agora Cloud Recording Backend Service

## Features
* RTC and RTM tokens
* Cloud Recording
* Fetch Recordings URLs for given Channel Name


[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://dashboard.heroku.com/new?template=https://github.com/AgoraIO-Community/Cloud-Recording-Golang/tree/main)

## Routes
Start call recording

`POST /api/start/call`

Stop call recording

`POST /api/stop/call`

Query status of recording

`POST /api/status/call `

Get list of files for channel name

`GET /api/get/list/<channelName>`

Get presigned url for file

`GET /api/get/file/<S3FileKey>`

Get URLs for m3u8 files for channel name

`GET /api/get/recordingUrls/<channelName>`

Get RTC token for channel name

`GET /api/get/rtc/<channelName>`

Get RTM token for UID

`GET /api/get/rtm/<uid>`

Get RTC and RTM token for channel

`GET /api/tokens/<channelName>`

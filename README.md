# FamPay Task

#### Make an API to fetch latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response.

# Requirements

1. Golang 1.18+
2. Docker

# API's

- ## GET
   - `GET    /v1/get/data?genere=football&page=1&limit=10`
   ![](assets/get.png)

   - `GET    /v1/search?title=hello&description=world`
   ![](assets/search.png)

- ## POST
   - `POST   /v1/create/cron`
   ```
   Body
   {
  "name": "shikhar-01",
  "query": "football",
  "interval": "10s",
  "api_keys": [
    "AIzaSyD-9tSrke72PouQMnMX-a7eZSW0jkFMBWY"
  ]
   }
  ```
   ![](assets/cron_create.png)

- ## DELETE
   - `DELETE /v1/delete/cron`
   ```
   Body
   {
      "id": "2"
   }
   ```
   ![](assets/cron_delete.png)


## Error Log
**YouTube Data API v3 has not been used in project 742952353666 before or it is disabled**

   ![](assets/youtube_error.png)

   ```
   ERRO[0194] failed to get video metadata from youtube     Type=Platform error="googleapi: Error 403: YouTube Data API v3 has not been used in project 742952353666 before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/youtube.googleapis.com/overview?project=742952353666 then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.\nDetails:\n[\n  {\n    \"@type\": \"type.googleapis.com/google.rpc.Help\",\n    \"links\": [\n      {\n        \"description\": \"Google developers console API activation\",\n        \"url\": \"https://console.developers.google.com/apis/api/youtube.googleapis.com/overview?project=742952353666\"\n      }\n    ]\n  },\n  {\n    \"@type\": \"type.googleapis.com/google.rpc.ErrorInfo\",\n    \"domain\": \"googleapis.com\",\n    \"metadata\": {\n      \"consumer\": \"projects/742952353666\",\n      \"service\": \"youtube.googleapis.com\"\n    },\n    \"reason\": \"SERVICE_DISABLED\"\n  }\n]\n, accessNotConfigured" service=main
   ```


# Features

- Scheduling a cron job with given query, interval & multiple API keys

- Get the latest videos sorted in reverse chronological order of their publishing date-time from YouTube for a given tag/search query in a paginated response

- Search for a video by title & description

- Delete a cron job

# Database

We are using MySQL as our database, models can be found in `models` folder. `pkg/store/relational/models.go`

# Running the server

1. Can use `make run` to run the server
2. Can use `docker-compose up` to run the server
3. Can build the image then run too.
## Alternatives

1.   `docker run --name=mysql1 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
`

2. `docker exec -it mysql1 mysql -uroot -p` Enter 123456 as password

3. `go run cmd/main.go`

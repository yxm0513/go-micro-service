
curl -XPUT "http://localhost:8080/api/feed/create_feed" -d '{"id": 100, "user_id": 123, "content": "hello world"}'
curl -XPUT "http://localhost:8080/api/feed/create_feed" -d '{"id": 101, "user_id": 123, "content": "goodbye!"}'
curl -XGET "http://localhost:8080/api/feed/get_feeds?user_id=123&&size=2" 

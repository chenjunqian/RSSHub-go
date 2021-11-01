curl -X PUT "localhost:9200/rss_item?pretty" -u "elastic:qazxsw" -H 'Content-Type: application/json' -d'
{
    "mappings": {
        "properties": {
            "id": {
                "type": "text",
                "index": false
            },
            "channel_id": {
                "type": "text",
                "index": false
            },
            "title": {
                "type": "text",
                "analyzer": "ik_smart",
                "search_analyzer": "ik_smart"
            },
            "channel_desc": {
                "type": "text",
                "analyzer": "ik_smart",
                "search_analyzer": "ik_smart"
            },
            "link": {
                "type": "text",
                "index": false,
                "analyzer": "ik_smart",
                "search_analyzer": "ik_smart"
            },
            "date": {
                "type": "date",
                "index": false,
                "format" : "yyyy-MM-dd HH:mm:ss"
            },
            "author": {
                "type": "text",
                "analyzer": "ik_smart",
                "search_analyzer": "ik_smart"
            },
            "input_date": {
                "type": "date",
                "index": false,
                "format" : "yyyy-MM-dd HH:mm:ss"
            },
            "channel_title": {
                "type": "text",
                "index": false
            },
            "channel_image_link": {
                "type": "text",
                "index": false
            },
            "channel_link": {
                "type": "text",
                "index": false
            }
        }
    }
}
'

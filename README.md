# API仕様

### / (GET) 回答一覧を取得する。
#### リクエスト
```shell
curl localhost:8080
```
#### レスポンス

```json
{
    "status": "OK",
    "result": [
        {
            "Id": 1,
            "Ans1": "1",
            "Ans2": "",
            "Ans3": ""
        },
        {
            "Id": 2,
            "Ans1": "1",
            "Ans2": "2",
            "Ans3": "3"
        }
    ]
}

```

### /insert (POST)新しい回答を追加する。

#### リクエスト

```shell
curl -XPOST -d 'ans1=1&ans2=2&ans3=3' localhost:8080/insert 
```


#### レスポンス
```json
{
    "status": "OK",
    "result": null
}
```
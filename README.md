# tips
## これなに
色々サンプルお試し用

## sdk-sample  
aws-sdk-goを利用したAWSのリソース操作用  

### 準備
必要パッケージのインストール  
```
go get "github.com/aws/aws-sdk-go/aws"
go get "github.com/aws/aws-sdk-go/aws/session"
go get "github.com/aws/aws-sdk-go/service/ec2"
go get "github.com/olekukonko/tablewriter"
```

### 使い方
* EC2でインスタンスの作成とインスタンス一覧情報の表示ができる
* パラメータは固定値なので実行するだけで動きます
```
go run main.go
```


base on bxcodec/go-clean-arch edit by JR.
https://github.com/bxcodec/go-clean-arch


### 資料夾結構：

├── Makefile
├── README.md
├── administrator
│   ├── delivery                  # 通訊層
│   │   └── http                  # HTTP也可以新增Websocket資料夾
│   │       └── handler.go
│   ├── mocks                     # 假資料 (測試用)
│   │   ├── repository.go
│   │   └── usecase.go
│   ├── repository                # 資料層
│   │   ├── mysql_repository.go
│   │   └── redis_repository.go
│   ├── repository.go             # 資料層介面
│   ├── usecase                   # 核心邏輯層
│   │   ├── usercase.go
│   │   └── usercase_test.go
│   └── usecase.go                # 核心邏輯層介面
├── bin
│   └── golangci-lint             # golang 的 linter執行檔
├── go.mod                        # 套件模組管理的檔案 (自動生成的)
├── go.sum                        # 套件模組管理的檔案 (自動生成的)
├── main.go                       # 程式進入點
├── models
    └── administrator.go          # 資料狀態結構


### 推薦套件：
config可以用viper

log可以用lumberjack

cmd可以用cobra

mongodb client我沒在用

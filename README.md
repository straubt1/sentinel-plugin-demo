# sentinel-plugin-demo


## Setup

```sh
# touch main.go
go mod init sentinel-plugin-demo
go get github.com/hashicorp/sentinel-sdk
go get github.com/hashicorp/sentinel-sdk/framework
go get github.com/hashicorp/sentinel-sdk/rpc
```

## Sentinel Plugin

```
import "plugin-demo" as pd

pd.envs
```


## References

- https://developer.hashicorp.com/sentinel/docs/extending/plugins
- https://github.com/hashicorp/sentinel-sdk


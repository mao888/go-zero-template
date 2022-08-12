# go-zero-template

go zero 框架使用示例

- [x] 集成gorm
- [x] 请求参数校验
- [x] 增加错误类型
- [x] 统一返回响应格式
- [ ] JWT集成

### 准备工作
#### 安装 goctl
- [安装文档](https://go-zero.dev/cn/docs/goctl/installation)
#### 修改本地handler模板（如果使用远程仓库生成，跳过这一步）
- 执行生成默认模版：`goctl template init`
- 备份初始handler模板： `cp ~/.goctl/${goctl版本号}/api/handler.tpl ~/.goctl/${goctl版本号}/api/handler_init.tpl`
- 执行：`vim ~/.goctl/${goctl版本号}/api/handler.tpl`, 添加以下内容
```
package handler

import (
    "net/http"
    "go-zero-template/greet/common/response"
    {{.ImportPackages}}
)

func {{.HandlerName}}(ctx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        {{if .HasRequest}}var req types.{{.RequestType}}
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }{{end}}

        l := logic.New{{.LogicType}}(r.Context(), ctx)
        {{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        {{if .HasResp}}response.Response(w, resp, err){{else}}response.Response(w, nil, err){{end}}
            
    }
}

```
### 运行
`cd ./greet && go run greet.go`
### 构建handler 
`goctl api go -api greet.api -dir . -style gozero --remote http://gitlab.ziggurat.cn/caict-dna/go-zero-goctl-template.git --branch go-zero-template`


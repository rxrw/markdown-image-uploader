# MARKDOWN 图片自动上传图床工具

可以将 markdown 中的本地图片，远端图片自动修改为自定义图床。

目前仅支持阿里云 oss

## cli

[下载最新版本](https://github.com/rxrw/markdown-image-uploader/releases/latest) 随意放

```bash
./uploader /usr/docs
```

就 ok 了

## 环境变量

CLIENT_NAME aliyun/qiniu

ENDPOINT 阿里云为对应endpoint，七牛为地域名，如Huabei/Beimei/Xinjiapo

ACCESS_KEY 对应的ak

ACCESS_SECRET 对应的sk

BUCKET_NAME 对应的 bucket 名

VISIT_URL 一般为cdn 的 url。很重要！用于判断文件存在等。具体格式为完整的url。如： `https://example.org`

## Github Actions

本脚本支持 Github Actions 自动部署

### 用法

在你的仓库 Settings -> Secrets 中创建上述变量（当然像不敏感的可以直接写）

在你的 actions.yml 中添加：

  ```yml
  -uses: rxrw/markdown-image-uploader@v2.0.1
    with:
      content_path: "content"
      client_name: "qiniu"
      access_key: ${{ secrets.ACCESS_KEY }}
      access_secret: ${{ secrets.ACCESS_SECRET }}
      endpoint: ${{ secrets.ENDPOINT }}
      bucket_name: ${{ secrets.BUCKET_NAME }}
      visit_url: ${{ secrets.VISIT_URL }}
  ```

继续编译其他内容，此脚本输出的就是你需要的 markdown 了

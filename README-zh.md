# MARKDOWN 图片自动上传图床工具

可以将 markdown 中的本地图片，远端图片自动修改为自定义图床。

目前仅支持阿里云 oss

## cli

[下载最新版本](https://github.com/rxrw/markdown-image-uploader/releases/latest) 随意放

```bash
./uploader /usr/docs
```

就 ok 了

## 参数

通过环境变量设置

ALIYUN_OSS_ENDPOINT

ALIYUN_OSS_ACCESSKEY_ID

ALIYUN_OSS_ACCESSKEY_SECRET

ALIYUN_OSS_BUCKET_NAME

ALIYUN_OSS_VISIT_URL

## Github Actions

本脚本支持 Github Actions 自动部署

### 用法

在你的仓库 Settings -> Secrets 中创建上述变量（当然像不敏感的可以直接写）

在你的 actions.yml 中添加：

  ```yml
  - uses: rxrw/markdown-image-uploader@v1.1.2
    with:
      content_path: "content"
      aliyun_oss_accesskey_id: ${{ secrets.ALIYUN_OSS_ACCESSKEY_ID }}
      aliyun_oss_accesskey_secret: ${{ secrets.ALIYUN_OSS_ACCESSKEY_SECRET }}
      aliyun_oss_endpoint: ${{ secrets.ALIYUN_OSS_ENDPOINT }}
      aliyun_oss_bucket: ${{ secrets.ALIYUN_OSS_BUCKET_NAME }}
      visit_url: ${{ secrets.VISIT_URL }}
  ```

继续编译其他内容，此脚本输出的就是你需要的 markdown 了

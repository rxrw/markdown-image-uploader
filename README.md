# MARKDOWN Picture automatic upload image

For Chinese Version: [Here](https://github.com/rxrw/markdown-image-uploader/README-zh.md)

You can automatically modify the local pictures and remote pictures in markdown to custom pictures.

Currently only supports Alibaba Cloud OSS

## cli

[Download the latest version](https://github.com/rxrw/markdown-image-uploader/releases/latest) Feel free to put

```bash
./uploader /usr/docs
```

It's ok

## Parameters

Set through environment variables

ALIYUN_OSS_ENDPOINT

ALIYUN_OSS_ACCESSKEY_ID

ALIYUN_OSS_ACCESSKEY_SECRET

ALIYUN_OSS_BUCKET_NAME

ALIYUN_OSS_VISIT_URL

## Github Actions

This script supports automatic deployment of Github Actions

### Usage

Create the above variables in your warehouse Settings -> Secrets (of course, you can write directly like insensitive ones)

Add in your actions.yml:

  ```yml
  -uses: rxrw/markdown-image-uploader-actions@v1.1
    with:
      content_path: "content"
      aliyun_oss_accesskey_id: ${{ secrets.ALIYUN_OSS_ACCESSKEY_ID }}
      aliyun_oss_accesskey_secret: ${{ secrets.ALIYUN_OSS_ACCESSKEY_SECRET }}
      aliyun_oss_endpoint: ${{ secrets.ALIYUN_OSS_ENDPOINT }}
      aliyun_oss_bucket: ${{ secrets.ALIYUN_OSS_BUCKET_NAME }}
      visit_url: ${{ secrets.VISIT_URL }}
  ```

Continue to compile other content, the output of this script is the markdown you need
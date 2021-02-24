# MARKDOWN Picture automatic upload image

For Chinese Version: [Here](https://github.com/rxrw/markdown-image-uploader/blob/master/README-zh.md)

You can automatically modify the local pictures and remote pictures in markdown to custom pictures.

Currently only supports Alibaba Cloud OSS

## cli

[Download the latest version](https://github.com/rxrw/markdown-image-uploader/releases/latest) Feel free to put

```bash
./uploader /usr/docs
```

It's ok

## Environments

CLIENT_NAME aliyun/qiniu

ENDPOINT endpoint value for aliyun, and zone for qiniu，such as Huabei/Beimei/Xinjiapo

ACCESS_KEY ak

ACCESS_SECRET sk

BUCKET_NAME bucket name

VISIT_URL the url of visit. **IMPORTANT!** example: `https://example.org`

## Github Actions

This script supports automatic deployment of Github Actions

### Usage

Create the above variables in your warehouse Settings -> Secrets (of course, you can write directly like insensitive ones)

Add in your actions.yml:

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

Continue to compile other content, the output of this script is the markdown you need

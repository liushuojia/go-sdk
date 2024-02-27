package minioConn

import (
	"context"
	"errors"
	"github.com/liushuojia/go-sdk/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/tags"
	"io"
	"log"
	"net/url"
	"os"
	"time"
)

// https://min.io/docs/minio/linux/developers/go/API.html#ListBuckets

func New() *Conn {
	return &Conn{
		useSSL: false,
		ctx:    context.Background(),
	}
}

type Conn struct {
	addresses string
	accessKey string
	secretKey string
	useSSL    bool
	client    *minio.Client
	ctx       context.Context

	makeBucketOptions   minio.MakeBucketOptions
	listObjectsOptions  minio.ListObjectsOptions
	statObjectOptions   minio.StatObjectOptions
	getObjectOptions    minio.GetObjectOptions
	putObjectOptions    minio.PutObjectOptions
	removeObjectOptions minio.RemoveObjectOptions
	reqParams           url.Values
}

func (c *Conn) SetAddresses(addresses string) *Conn {
	c.addresses = addresses
	return c
}
func (c *Conn) SetAccessKey(accessKey string) *Conn {
	c.accessKey = accessKey
	return c
}
func (c *Conn) SetSecretKey(secretKey string) *Conn {
	c.secretKey = secretKey
	return c
}
func (c *Conn) SetUseSSL(useSSL bool) *Conn {
	c.useSSL = useSSL
	return c
}

func (c *Conn) SetMakeBucketOptions(makeBucketOptions minio.MakeBucketOptions) *Conn {
	c.makeBucketOptions = makeBucketOptions
	return c
}
func (c *Conn) SetListObjectsOptions(listObjectsOptions minio.ListObjectsOptions) *Conn {
	c.listObjectsOptions = listObjectsOptions
	return c
}
func (c *Conn) SetStatObjectOptions(statObjectOptions minio.StatObjectOptions) *Conn {
	c.statObjectOptions = statObjectOptions
	return c
}
func (c *Conn) SetGetObjectOptions(getObjectOptions minio.GetObjectOptions) *Conn {
	c.getObjectOptions = getObjectOptions
	return c
}
func (c *Conn) SetPutObjectOptions(putObjectOptions minio.PutObjectOptions) *Conn {
	c.putObjectOptions = putObjectOptions
	return c
}
func (c *Conn) SetRemoveObjectOptions(removeObjectOptions minio.RemoveObjectOptions) *Conn {
	c.removeObjectOptions = removeObjectOptions
	return c
}
func (c *Conn) SetRequestParams(reqParams url.Values) *Conn {
	c.reqParams = reqParams
	return c
}

func (c *Conn) Connect() *Conn {
	log.Println("[minio]", "connect minio")

	// Initialize minio client object.
	minioClient, _ := minio.New(c.addresses, &minio.Options{
		Creds:  credentials.NewStaticV4(c.accessKey, c.secretKey, ""),
		Secure: c.useSSL,
	})
	c.client = minioClient
	return c
}
func (c *Conn) Client() *minio.Client {
	return c.Connect().client
}

func (c *Conn) ListBuckets() ([]minio.BucketInfo, error) {
	return c.Client().ListBuckets(c.ctx)
}
func (c *Conn) MakeBucket(bucketName string) (err error) {
	// 创建这个bucket
	f, err := c.BucketExists(bucketName)
	if err != nil {
		return
	}
	if f == true {
		return errors.New("bucket exists")
	}
	return c.Client().MakeBucket(c.ctx, bucketName, c.makeBucketOptions)
}
func (c *Conn) BucketExists(bucketName string) (bool, error) {
	return c.Client().BucketExists(c.ctx, bucketName)
}
func (c *Conn) RemoveBucket(bucketName string) error {
	return c.Client().RemoveBucket(c.ctx, bucketName)
}

func (c *Conn) ListObjects(bucketName string) <-chan minio.ObjectInfo {
	return c.Client().ListObjects(c.ctx, bucketName, c.listObjectsOptions)
}

func (c *Conn) SetBucketTagging(bucketName string, tagMap map[string]string) (err error) {
	tagData, err := tags.NewTags(tagMap, false)
	if err != nil {
		return err
	}
	return c.Client().SetBucketTagging(c.ctx, bucketName, tagData)
}
func (c *Conn) GetBucketTagging(bucketName string) (map[string]string, error) {
	tagData, err := c.Client().GetBucketTagging(c.ctx, bucketName)
	if err != nil {
		return nil, err
	}
	return tagData.ToMap(), nil
}
func (c *Conn) RemoveBucketTagging(bucketName string) error {
	return c.Client().RemoveBucketTagging(c.ctx, bucketName)
}

// ListIncompleteUploads(ctx context.Context, bucketName, prefix string, recursive bool) <- chan ObjectMultipartInfo

func (c *Conn) StatObject(bucketName string, objectName string) (minio.ObjectInfo, error) {
	return c.Client().StatObject(c.ctx, bucketName, objectName, c.statObjectOptions)
}
func (c *Conn) GetObject(bucketName string, objectName string) (*minio.Object, error) {
	return c.Client().GetObject(c.ctx, bucketName, objectName, c.getObjectOptions)
}
func (c *Conn) GetObjectContent(bucketName string, objectName string) ([]byte, error) {
	object, err := c.Client().GetObject(c.ctx, bucketName, objectName, c.getObjectOptions)
	if err != nil {
		return nil, err
	}
	/*
		//写入文件
		out, err := os.Create(savePath)
		defer out.Close()
		if err != nil {
			return err
		}
		_, err = io.Copy(out, object)
		if err != nil {
			return err
		}
	*/
	return io.ReadAll(object)
}
func (c *Conn) FGetObject(bucketName string, objectName string, savePath string) error {
	return c.Client().FGetObject(c.ctx, bucketName, objectName, savePath, c.getObjectOptions)
}
func (c *Conn) PutObject(bucketName string, objectName string, filePath string) (info minio.UploadInfo, err error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return
	}
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	c.putObjectOptions.ContentType, _ = utils.GetContentType(filePath)
	return c.Client().PutObject(c.ctx, bucketName, objectName, f, fileInfo.Size(), c.putObjectOptions)
}
func (c *Conn) FPutObject(bucketName string, objectName string, filePath string) (info minio.UploadInfo, err error) {
	return c.Client().FPutObject(c.ctx, bucketName, objectName, filePath, c.putObjectOptions)
}
func (c *Conn) CopyObjects(srcBucketName, srcObjectName, dstBucketName, dstObjectName string) (info minio.UploadInfo, err error) {
	srcOpts := minio.CopySrcOptions{
		Bucket: srcBucketName,
		Object: srcObjectName,
	}

	// Destination object
	dstOpts := minio.CopyDestOptions{
		Bucket: dstBucketName,
		Object: dstObjectName,
	}

	// copy
	return c.Client().CopyObject(c.ctx, dstOpts, srcOpts)
}
func (c *Conn) RemoveObject(bucketName string, objectName string) (err error) {
	//opts := minio.RemoveObjectOptions{
	//	GovernanceBypass: true,
	//	VersionID:        "myversionid",
	//}
	//opts.GovernanceBypass = true
	return c.Client().RemoveObject(c.ctx, bucketName, objectName, c.removeObjectOptions)
}

func (c *Conn) PutObjectTagging(bucketName, objectName string, tagMap map[string]string) (err error) {
	tagData, err := tags.NewTags(tagMap, false)
	if err != nil {
		return err
	}
	return c.Client().PutObjectTagging(c.ctx, bucketName, objectName, tagData, minio.PutObjectTaggingOptions{})
}
func (c *Conn) GetObjectTagging(bucketName, objectName string) (map[string]string, error) {
	tagData, err := c.Client().GetObjectTagging(c.ctx, bucketName, objectName, minio.GetObjectTaggingOptions{})
	if err != nil {
		return nil, err
	}
	return tagData.ToMap(), nil
}
func (c *Conn) RemoveObjectTagging(bucketName, objectName string) error {
	return c.Client().RemoveObjectTagging(c.ctx, bucketName, objectName, minio.RemoveObjectTaggingOptions{})
}

// PresignedGetObject
// expiry	time.Second*24*60*60
// reqParams
//
//	 	make(url.Values)
//		reqParams.Set("response-content-disposition", "attachment; filename=\"your-filename.txt\"")
func (c *Conn) PresignedGetObject(bucketName, objectName string, expiry time.Duration) (*url.URL, error) {
	return c.Client().PresignedGetObject(c.ctx, bucketName, objectName, expiry, c.reqParams)
}

// PresignedPutObject
// PUT url
// Body is file data
func (c *Conn) PresignedPutObject(bucketName, objectName string, expiry time.Duration) (*url.URL, error) {
	return c.Client().PresignedPutObject(c.ctx, bucketName, objectName, expiry)
}

// PresignedPostPolicy
// Post url
// Body is formdata  file=@/filePath
func (c *Conn) PresignedPostPolicy(bucketName, objectName string, expiry time.Duration) (*url.URL, map[string]string, error) {
	// Initialize policy condition config.
	policy := minio.NewPostPolicy()

	// Apply upload policy restrictions:
	if err := policy.SetBucket(bucketName); err != nil {
		return nil, nil, err
	}
	if err := policy.SetKey(objectName); err != nil {
		return nil, nil, err
	}
	if err := policy.SetExpires(time.Now().UTC().Add(expiry)); err != nil {
		return nil, nil, err
	}

	// Only allow 'png' images.
	//policy.SetContentType("image/png")

	// Only allow content size in range 1KB to 1MB.
	//policy.SetContentLengthRange(1024, 1024*1024)

	// Add a user metadata using the key "custom" and value "user"
	//policy.SetUserMetadata("custom", "user")

	// Get the POST form key/value object:
	return c.Client().PresignedPostPolicy(c.ctx, policy)
}

/*
	在MinIO中,PutObject和FPutObject主要有以下几点区别:
	操作对象方式不同:
	PutObject上传对象采用单个对象上传方式,一次请求上传一个对象。
	FPutObject上传对象采用分片上传方式,可以在多个请求中上传同一个对象。

	适用场景不同:
	PutObject适用于小对象的上传。
	FPutObject适用于大对象的上传,可以支持超过5GB的对象上传。

	上传方式不同:
	PutObject直接上传对象。
	FPutObject需要首先调用InitiateMultipartUpload初始化分片上传,然后调用UploadPart上传每个分片,最后调用CompleteMultipartUpload完成上传。

	错误处理不同:
	PutObject上传错误需要重新上传整个对象。
	FPutObject上传错误只需要重新上传出错的分片。

	上传事务不同:
	PutObject上传具有原子性,一个对象要么完全上传,要么失败。
	FPutObject上传可以暂停和恢复,但多个分片上传完成后才视为成功。

	所以简单来说,对于小对象可以直接使用PutObject,对于大对象建议使用FPutObject分片上传。


	前端使用（暂未考虑大文件问题）
		PresignedGetObject 获取前端访问文件路径带时效
		PresignedPutObject 前端put上传文件， body为文件
		PresignedPostPolicy 前端post form方式 file为文件， 返回的map必须在form中
*/

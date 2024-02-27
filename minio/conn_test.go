package minioConn

import (
	"fmt"
	"testing"
	"time"
)

func TestMinio(t *testing.T) {

	m := New().SetAddresses("localhost:9000").
		SetAccessKey("wtXQGX8jOXQ0r3wMpf4Q").
		SetSecretKey("LmOTOUDJJuetuKzctlO0x4dRrRVWvz7RGz8A51cV").
		SetUseSSL(false).Connect()

	bucketName := "test"
	objectName := "2.jpg"

	// Get the POST form key/value object:
	url, formData, err := m.PresignedPostPolicy(bucketName, "myobject", time.Hour)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(url.String())
	fmt.Println(formData)

	// POST your content from the command line using `curl`
	fmt.Printf("curl ")
	for k, v := range formData {
		fmt.Printf("-F %s=%s ", k, v)
	}
	fmt.Printf("-F file=@/Volumes/work/app-aliyun/docker/README.md ")
	fmt.Printf("%s\n", url)

	return
	//if err := m.MakeBucket(bucketName, minio.MakeBucketOptions{}); err != nil {
	//	log.Fatalln(err)
	//}

	//info, err := m.PutObject(bucketName, "go/conn.go", "./conn.go")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(info)

	u1, _ := m.PresignedPutObject(bucketName, objectName, time.Hour)

	fmt.Println(u1.String())
	//fmt.Println(
	//	m.StatObject(bucketName, "conn.go"),
	//)

	//err := m.SetRemoveObjectOptions(minio.RemoveObjectOptions{
	//	ForceDelete: true,
	//}).RemoveObject(bucketName, "conn.go")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//m.SetBucketTagging("mypic", map[string]string{
	//	"a1": "3",
	//	"b1": "4",
	//})
	//return

	//base64, info, err := m.GetObjectsBase64("mypic", "1.jpg", minio.GetObjectOptions{})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//fmt.Println(base64)
	//fmt.Println(info)
	//m.SaveObjects("mypic", "credentials", minio.GetObjectOptions{}, "./credentials")

	//minioClient, err := minio.New("127.0.0.1:9000", &minio.Options{
	//	Creds:  credentials.NewStaticV4("VtKTJPtTY74Q6g596dHX", "cKWZbOcNhCURoKGRRqp21DtQPSR5Ztqg1tjg7Zc4", ""),
	//	Secure: false,
	//})
	//
	//buckets, err := minioClient.ListBuckets(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//return
	//// bucket名称
	//bucketName := "test001"
	//
	//// 创建这个bucket
	//if err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{}); err != nil {
	//	// 检测这个bucket是否已经存在
	//	exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
	//	if errBucketExists == nil && exists {
	//		log.Printf("We already own %s\n", bucketName)
	//	} else {
	//		log.Fatalln(err)
	//	}
	//} else {
	//	log.Printf("Successfully created %s\n", bucketName)
	//}
	//
	////Upload the test file
	////Change the value of filePath if the file is in another location
	//objectName := "credentials"
	//filePath := "/Users/liushuojia/Downloads/credentials.json"
	//contentType := "application/octet-stream"
	//
	//// Upload the test file with FPutObject
	//info, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	// 需要上传文件的基本信息
	//objectName := "README.md"
	//filePath := "/Volumes/work/app-aliyun/sdk"
	//contentType := "multipart/form-data"
	//fPath := filepath.Join(filePath, objectName)
	//fileInfo, err := os.Stat(fPath)
	//if err == os.ErrNotExist {
	//	log.Printf("%s目标文件不存在", fPath)
	//}
	//
	//f, err := os.Open(fPath)
	//if err != nil {
	//	fmt.Println("file open err", err)
	//	return
	//}
	//
	//uploadInfo, err := minioClient.PutObject(context.Background(), bucketName, objectName, f, fileInfo.Size(), minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, uploadInfo.Size)
	//return

	// Make a new bucket called testbucket.
	//bucketName := "testbucket"
	//location := "us-east-1"
	//
	//err := es.GetClient().MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	//if err != nil {
	//	// Check to see if we already own this bucket (which happens if you run this twice)
	//	exists, errBucketExists := es.GetClient().BucketExists(ctx, bucketName)
	//	if errBucketExists == nil && exists {
	//		log.Printf("We already own %s\n", bucketName)
	//	} else {
	//		log.Fatalln(err)
	//	}
	//} else {
	//	log.Printf("Successfully created %s\n", bucketName)
	//}

	//buckets, err := es.GetClient().ListBuckets(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, bucket := range buckets {
	//	fmt.Println(bucket)
	//}
	//
	//return

	// Upload the test file
	// Change the value of filePath if the file is in another location
	//objectName := "credentials"
	//filePath := "/Users/liushuojia/Downloads/credentials.json"
	//contentType := "application/octet-stream"
	//
	//// Upload the test file with FPutObject
	//info, err := es.GetClient().FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

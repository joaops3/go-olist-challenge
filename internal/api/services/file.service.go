package services

import (
	"context"
	"errors"
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joaops3/go-olist-challenge/internal/data/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BucketName string

const BUCKET_GENERIC BucketName = "playerimg"

const (timeout = time.Millisecond *10000)


type FileService struct {
	
}

type FileServiceInterface interface {
	GenerateFile(dto *multipart.FileHeader, bucketName BucketName, fileType models.FileTypes) (*models.S3FileModel, error)
	DeleteFile(bucketName string, key string) (error)
}

func NewFileService() FileServiceInterface {
	return &FileService{}
}

func (s *FileService)GenerateFile(fileHeader *multipart.FileHeader, bucketName BucketName, fileType models.FileTypes) (*models.S3FileModel, error){
	ctx := context.Background()


	file, err := fileHeader.Open()

	if err != nil {
		return nil, err
	}

	defer file.Close()

	
	mimeType, err := s.getMimeType(file)

	if err != nil {
		return nil, err
	}

	key, err := s.generateKey(mimeType)

	if err != nil {
		return nil, err
	}
	path := s.generatePath(bucketName, key)

	newFile := models.NewS3FileModel(key, fileHeader.Filename, path, mimeType, fileType, fileHeader.Size)

  	var cancelFn func()
  	if timeout > 0 {
  		ctx, cancelFn = context.WithTimeout(ctx, timeout)
  	}
  
	
	if cancelFn != nil {
		defer cancelFn()
  	}
	
	sess := s.getSession()

	if sess == nil {
		return nil, errors.New("invalid session")
	}

	svc := s3.New(sess, &aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	})

	

	_, err = svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(string(bucketName)),
		Key:    aws.String(key),
		Body:   file,
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			return nil, fmt.Errorf("upload canceled due to timeout %v\n", err)
		} else {
			return nil, fmt.Errorf("failed to upload object %v\n", err)
		}
	}

	return newFile, nil
}

func (s *FileService)DeleteFile(bucketName string, key string) (error){
	ctx := context.Background()
	var cancelFn func()
  	if timeout > 0 {
  		ctx, cancelFn = context.WithTimeout(ctx, timeout)
  	}
  
	
	if cancelFn != nil {
		defer cancelFn()
  	}
	
	sess := s.getSession()

	if sess == nil {
		return errors.New("invalid session")
	}

	svc := s3.New(sess, &aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	})

	_, err := svc.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(string(bucketName)),
		Key:    aws.String(key),
	})

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			return fmt.Errorf("upload canceled due to timeout %v\n", err)
		} else {
			return fmt.Errorf("failed to upload object %v\n", err)
		}
	}
	return  nil
}

func (s *FileService) generateKey(mimeType string) (string, error){
	id := primitive.NewObjectID().Hex()
	ext, err := mime.ExtensionsByType(mimeType)
	if err != nil || len(ext) == 0 {
		return "", fmt.Errorf("failed to get file extension for MIME type: %s", mimeType)
	}
	key := fmt.Sprintf("%s%d%s", id, time.Now().Unix(), ext[0])
	return key, nil
}

func (s *FileService) generatePath(bucket BucketName, key string) (string){

	path := fmt.Sprintf(`https://%s.s3-us-east-1.amazonaws.com/%s`, bucket, key)
	return path
}

func  (s *FileService) getMimeType(file multipart.File)(string, error){
	buf := make([]byte, 512)
    _, err := file.Read(buf)
    if err != nil {
        return "", err
    }
	mimeType := http.DetectContentType(buf)
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}
	return mimeType, nil
}

	
func (s *FileService) getSession() *session.Session{
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	}))

	return sess
}
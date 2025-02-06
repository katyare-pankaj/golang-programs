
package main  
import (  
    "context"
    "fmt"
    "log"
    "os"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

var s3client *s3.S3

// Function Handler
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    //Get key from request querystring
    key := request.QueryStringParameters["key"]

    //Create a bucket to store data
    bucket := "your-bucket-name"

    // Defer statement to ensure the bucket is deleted after the function returns
    defer deleteBucket(bucket)

    //Create a new session with the S3 service
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")}, // Set your desired region
    )
    if err != nil {
        log.Println("Error creating session:", err)
        return events.APIGatewayProxyResponse{
            StatusCode: 500,
            Body:       "Error creating session",
        }, nil
    }

    // Create an S3 client
    s3client = s3.New(sess)

    //Create the bucket
    createBucketErr := createBucket(bucket)
    if createBucketErr != nil {
        log.Println("Error creating bucket:", createBucketErr)
        return events.APIGatewayProxyResponse{
            StatusCode: 500,
            Body:       "Error creating bucket",
        }, nil
    }

    // Put the object to the bucket
    putObjectErr := putObject(bucket, key, "Hello, World!")
    if putObjectErr != nil {
        log.Println("Error putting object:", putObjectErr)
        return events.APIGatewayProxyResponse{
            StatusCode: 500,
            Body:       "Error putting object",
        }, nil
    }

    // Get the object from the bucket
    getObjectOutput, getObjectErr := getObject(bucket, key)
    if getObjectErr != nil {
        log.Println("Error getting object:", getObjectErr)
        return events.APIGatewayProxyResponse{
            StatusCode: 500,
            Body:       "Error getting object",
        }, nil
    }

    // Return the response
    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       fmt.Sprintf("Data retrieved: %s", string(getObjectOutput.Body)),
    }, nil
}

func main() {  
    lambda.Start(Handler)
}

func createBucket(bucket string) error {
    _, err := s3client.CreateBucket(&s3.CreateBucketInput{
        Bucket: aws.String(bucket),
    })
    return err
}

func putObject(bucket, key, data string) error {
    _, err := s3client.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
        Body:   aws.ReadSeekCloser(strings.NewReader(data)),
    })
    return err
}

func getObject(bucket, key string) (*s3.GetObjectOutput, error) {
    return s3client.GetObject(&s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    })
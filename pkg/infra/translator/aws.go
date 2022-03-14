package translator

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
	"os"
)

var DefaultAwsClient Client

func init() {
	defaultRegion := os.Getenv("AWS_REGION")
	if len(defaultRegion) == 0 {
		defaultRegion = "ap-northeast-1"
	}

	config := &aws.Config{
		Region:     aws.String(defaultRegion),
	}

	sess := session.Must(session.NewSession())
	svc := translate.New(sess, config)

	DefaultAwsClient = &awsClient{
		svc:    svc,
		source: "en",
		target: "ja",
	}
}

type awsClient struct {
	svc    *translate.Translate
	source string
	target string
}

func NewAwsClient(config *aws.Config, source, target string) Client {
	svc := translate.New(session.Must(session.NewSession(config)))

	return &awsClient{svc,source, target}
}

func (t *awsClient) Do(text string) (string, error) {
	result, err := t.svc.Text(&translate.TextInput{
		SourceLanguageCode: aws.String(t.source),
		TargetLanguageCode: aws.String(t.target),
		Text:               aws.String(text),
	})
	if err != nil {
		return "", fmt.Errorf("translator.Do(%s): %s", text, err.Error())
	}

	return *result.TranslatedText, nil
}

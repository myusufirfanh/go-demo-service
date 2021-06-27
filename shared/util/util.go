package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/fgrosse/goldi"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/myusufirfanh/go-demo-service/models"
	"github.com/myusufirfanh/go-demo-service/shared/config"

	"github.com/nats-io/nats.go"
	"github.com/streadway/amqp"
)

type (
	qoalaResponse models.QoalaResponsePattern

	// CustomApplicationContext return qoala custom application context
	CustomApplicationContext struct {
		echo.Context
		Container     *goldi.Container
		SharedConf    config.ImmutableConfigInterface
		RedisSession  *redis.Client
		MysqlSession  *gorm.DB
		RabbitSession *amqp.Connection
		NatsSession   *nats.EncodedConn
		SqsService    *sqs.SQS
		S3Service     *s3.S3
		SESService    *ses.SES
		UserJWT       *models.UserJWT
	}
)

// CustomResponse is a method that returns custom object response
func (c *CustomApplicationContext) CustomResponse(status string, data interface{}, message string, code int, meta *models.QoalaResponsePatternMeta, sytemMessage string) error {
	resp := &qoalaResponse{
		Status:  status,
		Data:    data,
		Message: message,
		Code:    code,
	}
	if meta != nil {
		resp.Meta = meta
	}
	if data != nil {
		// to print out what response give to client
		// no need to catch the marshal error
		respStr, _ := json.Marshal(&resp)
		if respStr != nil {
			fmt.Printf("%s -- Response: %s", getCallerMethod(), string(respStr))
		}
	}
	return c.JSON(code, resp)
}

func getCallerMethod() string {
	var source string
	if pc, _, _, ok := runtime.Caller(2); ok {
		var funcName string
		if fn := runtime.FuncForPC(pc); fn != nil {
			funcName = fn.Name()
			if i := strings.LastIndex(funcName, "."); i != -1 {
				funcName = funcName[i+1:]
			}
		}

		source = path.Base(funcName)
	}
	return source
}

//ReadJSON is a function to read json in path into pointer
func ReadJSON(path string, pointer interface{}) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(byteValue), pointer)
	return nil

}

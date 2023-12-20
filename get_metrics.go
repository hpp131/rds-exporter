// call ali-openapi
package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	// "os"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"

	// "github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
)

// 加入`json: "xxx"`增加unmarshel()的稳定性
type RdsResponse struct {
	MetricName   string
	InstanceId   string
	Maximum      float32
	InstanceName string
	Tag          map[string]string
}

// 代码可复用，所有云监控产品都使用该阿里云接口
func getMetrics(m string) []*RdsResponse {
	// defer wg.Done()
	config := sdk.NewConfig()
	endTime := time.Now()
	startTime := time.Now().Add(-1 * time.Minute)

	// Please ensure that the environment variables ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set.
	credential := credentials.NewAccessKeyCredential(AccessKey, SecretKey)
	client, err := cms.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}

	request := cms.CreateDescribeMetricListRequest()

	request.Scheme = "https"

	// 是否可以用形参传递代替？
	request.Namespace = "acs_rds_dashboard"
	request.MetricName = m
	request.StartTime = fmt.Sprint(startTime.UnixMilli())
	request.EndTime = fmt.Sprint(endTime.UnixMilli())

	response, err := client.DescribeMetricList(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	s := trimBackSlash(response.Datapoints)
	result := decode(s)
	fillTag(result, getNameAndTag(), m, instanceidName)
	return result

}

func trimBackSlash(s string) string {
	// trim '\' from s. The "\\" is escape
	r := strings.ReplaceAll(s, "\\", "")
	return r
}

// optimize: custome type and bind it with func
func decode(s string) []*RdsResponse {
	var r []*RdsResponse
	err := json.Unmarshal([]byte(s), &r)
	if err != nil {
		fmt.Printf("parse string to []RdsResponse failed : %v", err)
	}
	return r
}

// fill RdsResponse.Tag and RdsResponse.MetricName and RdsResponse.InstanceName
func fillTag(r []*RdsResponse, t map[string]map[string]string, n string, name map[string]string) {
	for _, res := range r {
		res.Tag = t[res.InstanceId]
		res.MetricName = n
		if _, ok := name[res.InstanceId];ok{
			res.InstanceName = name[res.InstanceId]
		}else{
			res.InstanceName = ""
		}		
	}
}

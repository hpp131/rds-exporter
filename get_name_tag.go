package main

//get instance name and tag

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
)

// type  DBTagMap map[string]map[string]string

// 是否可以把map[string]map[string]string替换成*map[string]map[string]string？
// 该函数代码不通用，需要根据云产品对应的接口进行修改
func getNameAndTag() map[string]map[string]string {
	config := sdk.NewConfig()
	// Please ensure that the environment variables ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set.
	credential := credentials.NewAccessKeyCredential(AccessKey, SecretKey)

	/* use STS Token
	credential := credentials.NewStsTokenCredential(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"), os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"), os.Getenv("ALIBABA_CLOUD_SECURITY_TOKEN"))
	*/
	client, err := rds.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}

	request := rds.CreateDescribeDBInstanceByTagsRequest()
	request.Scheme = "https"

	response, err := client.DescribeDBInstanceByTags(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	// fmt.Printf("response is %#v\n", response.Items.DBInstanceTag)
	// r := make(map[string]map[string]string)
	r := responseToMap(response.Items.DBInstanceTag)
	return r
}

// if neccessary struct -> map[string]map[string]string ?
// 是否可以把[]rds.DBInstanceTag替换为*[]rds.DBInstanceTag?
func responseToMap(r []rds.DBInstanceTag) map[string]map[string]string {
	m := make(map[string]map[string]string)
	for _, dbInstanceTag := range r {
		// iterate
		temmap := make(map[string]string)
		for _, tag := range dbInstanceTag.Tags.Tag {
			temmap[tag.TagKey] = tag.TagValue
		}
		m[dbInstanceTag.DBInstanceId] = temmap
	}
	// fmt.Printf("responseToMap is %v\n", m)
	return m
}



func rdsName() {
	config := sdk.NewConfig()
	// Please ensure that the environment variables ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set.
	credential := credentials.NewAccessKeyCredential(AccessKey, SecretKey)
	client, err := rds.NewClientWithOptions("cn-hangzhou", config, credential)
	if err != nil {
		panic(err)
	}
	request := rds.CreateDescribeDBInstancesRequest()
	request.Scheme = "https"
	request.PageSize = requests.NewInteger(100)
	response, err := client.DescribeDBInstances(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	for _, instance := range response.Items.DBInstance {
		instanceidName[instance.DBInstanceId] = instance.DBInstanceDescription
	}
	
}

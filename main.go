package main

import (
	"bytes"
	"common/nacos"
	"fmt"
	"gopkg.in/yaml.v3"
)

type Item struct {
	Host string `yaml:"Host"`
}

// 定义信息的结构体
type FactoryYaml struct {
	DataSource string `yaml:"DataSource"`
	Cache      []Item `yaml:"Cache"`
}

func main() {
	bootstrapConfig := nacos.BootstrapConfig{
		NacosConfig: nacos.NacosConfig{
			DataId: "application-dev.yaml",
			Group:  "DEFAULT_GROUP",
			ServerConfigs: []nacos.NacosServerConfig{
				nacos.NacosServerConfig{
					IpAddr: "127.0.0.1",
					Port:   8848,
				},
			},
			ClientConfig: nacos.NacosClientConfig{
				NamespaceId: "f49daff5-dbf9-437d-abb5-1431b8dcfc45",
			},
		},
	}
	factory := nacos.NacosFactory(bootstrapConfig)
	var factory_yaml FactoryYaml
	instance := factory.CreateConfigClient(func(data string) {
		body := bytes.TrimPrefix([]byte(data), []byte("\xef\xbb\xbf"))
		err := yaml.Unmarshal(body, &factory_yaml)
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println("获取的配置信息")
		fmt.Println(factory_yaml.DataSource)
		for _, r := range factory_yaml.Cache {
			fmt.Println(r)

		}
	})
	body := bytes.TrimPrefix([]byte(instance), []byte("\xef\xbb\xbf"))
	err := yaml.Unmarshal(body, &factory_yaml)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("获取的配置信息")
	fmt.Println(factory_yaml.DataSource)
	for _, r := range factory_yaml.Cache {
		fmt.Println(r)

	}
	for true {

	}
}

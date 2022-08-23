# go_nacos

#### 介绍
包括服务的注册和服务的发现和获取在线的配置信息


#### 使用说明

1.  ServiceRegistration 进行服务的注册和获取配置文件的服务链接
2.  FindInstance  查询指定服务的信息
3.  CreateConfigClient  获取的是配置信息

#### 获取配置使用


```
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
	//factory.ServiceRegistration("test", "81.69.174.251", 8080)
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

```



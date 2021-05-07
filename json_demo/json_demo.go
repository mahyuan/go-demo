package json_demo

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"server_name"` // `json:"key"` 转化成json后的key是后面的key，映射关系
	ServerIp string `json:"server_ip"`
	ServerPort int `json:"server_port"`
}

// SerializeStruct 结构体转化json
func SerializeStruct()  {
	server := new(Server)
	server.ServerName = "Demo for json"
	server.ServerIp = "127.0.0.1"
	server.ServerPort = 8080

	// 序列化成json字节数组
	b, err := json.Marshal(server)

	if err != nil {
		fmt.Println("Marshal struct err:", err)
		return
	}

	fmt.Println("Marshal struct:", string(b))
}


//SerializeMap 序列化map成json
func SerializeMap() {
	server := make(map[string]interface{})
	server["ServerName"] = "Demo for json"
	server["ServerIp"] = "127.0.0.1"
	server["ServerPort"] = 8080

	// 序列化成json字节数组
	b, err := json.Marshal(server)

	if err != nil {
		fmt.Println("Marshal map err:", err)
		return
	}

	fmt.Println("Marshal map:", string(b))
}

// DeSerializeStruct 反序列化结构体
func DeSerializeStruct()  {
	jsonString := `{"server_name":"Demo for json","server_ip":"127.0.0.1","server_port":8080}`
	server := new(Server)
	err := json.Unmarshal([]byte(jsonString), &server)

	if err != nil {
		fmt.Println("Unmarshal json string to struct err: ", err.Error())
		return
	}
	fmt.Println("Unmarshal json string to struct: ", server)
}

// DeSerializeMap 反序列化map
func DeSerializeMap()  {
	jsonString := `{"server_name":"Demo for json","server_ip":"127.0.0.1","server_port":8080}`
	server := new(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &server)

	if err != nil {
		fmt.Println("Unmarshal json string to map err: ")
		return
	}
	fmt.Println("Unmarshal json string to map: ", server)
}

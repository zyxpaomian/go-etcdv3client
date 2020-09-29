package main

import(
    "github.com/zyxpaomian/etcdv3client"
    "fmt"    
)

var endPoints []string
endPoints[0] = "192.168.159.133:2379"
err := etcdv3client.ClientInit(10, 10, 300, endPoints)
if err != nil {
    fmt.Println("etcdClient Init Error")
    panic(err)
}

err = etcdv3client.Etcdclient.Put("putTest", "ok")
if err != nil {
    fmt.Println("etcdClient Put Error")
    panic(err)
}

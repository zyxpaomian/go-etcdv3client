package main

import(
    mec "github.com/zyxpaomian/etcdV3Client"
    "fmt"    
)

var endPoints []string
endPoints[0] = "192.168.159.133:2379"
err := mec.ClientInit(10, 10, 300, endPoints)
if err != nil {
    fmt.Println("etcdClient Init Error")
    panic(err)
}

err = mec.Etcdclient.Put("putTest", "ok")
if err != nil {
    fmt.Println("etcdClient Put Error")
    panic(err)
}
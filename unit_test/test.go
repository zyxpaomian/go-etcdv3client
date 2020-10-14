package main

import(
    mce "github.com/zyxpaomian/etcdv3client"
    "fmt"    
    "time"
)


func main() {
    var endPoints  = []string{"192.168.159.133:2379"}
    err := mce.ClientInit(10, 10, endPoints)
    if err != nil {
        fmt.Println("etcdClient Init Error")
        panic(err)
    }


    err = mce.Etcdclient.SetLease(15)
    if err != nil {
        fmt.Println("set lease error")
    }

    errChan := make(chan error)
    err = mce.Etcdclient.DataRegister("/test", "aaa", errChan)
    if err != nil {
        fmt.Println("set error chan error")
    }

    go func() {
        for {
            time.Sleep(time.Duration(5)*time.Second)
            errdata := <- errChan
            fmt.Println(errdata.Error())
        }
    }()
    select {}

    /*keyChan := make(chan string)
    valueChan := make(chan string)
    typeChan := make(chan string)
    go mce.Etcdclient.WatchPrefix("/server/", keyChan, valueChan, typeChan)

    go func() {
        for {
            key := <- keyChan
            value := <- valueChan
            etype := <- typeChan
            fmt.Println(key)
            fmt.Println(etype)
            fmt.Println(value)
        }
    }()

    select {}*/

    
    /*
    // put test
    err = mce.Etcdclient.Put("putTest/mykey", "ok")
    if err != nil {
        fmt.Println("etcdClient Put Error")
        panic(err)
    }

    // put test
    err = mce.Etcdclient.Put("putTest/mykey2", "ok")
    if err != nil {
        fmt.Println("etcdClient Put Error")
        panic(err)
    }    

    // get test
    myvalue, err := mce.Etcdclient.Get("putTest/mykey")
    if err != nil {
        fmt.Println("etcdClient Get Error")
        panic(err)
    }
    fmt.Println(myvalue)

    // get prefix
    myprefix, err := mce.Etcdclient.GetPrefix("putTest")
    if err != nil {
        fmt.Println("etcdClient GetPrefix Error")
        panic(err)
    }
    fmt.Println(myprefix) 
    
    // del test
    err = mce.Etcdclient.Del("putTest/mykey")
    if err != nil {
        fmt.Println("etcdClient Del Error")
        panic(err)
    }

    // del test
    err = mce.Etcdclient.DelPrefix("putTest")
    if err != nil {
        fmt.Println("etcdClient DelPrefix Error")
        panic(err)
    }

    select {}
    */
}
# API Protocol

Table of Contents

* [Service ProtoDoc](#service-protodoc)
    * [Method ProtoDoc.Convert](#method-protodocconvert)
* [Service Hello](#service-hello)
    * [Method Hello.Hello](#method-hellohello)
* [Enums](#enums)
* [Objects](#objects)




## Service ProtoDoc



### Method ProtoDoc.Convert

> POST /protodoc/ProtoDoc/Convert <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

把 proto3 定义的协议转换成文档

Request parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| proto | string | proto3 定义的协议内容 |
| filename | string | 协议文件名，用于解析时的附加信息 |

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| html | string | HTML 格式的文档 |





## Service Hello



### Method Hello.Hello

> POST /protodoc/Hello/Hello <br/>
> Content-Type: application/json <br/>
> Authorization: Bearer (token) <br/>

用于性能测试

Request is empty

Response parameters

|   Name    |   Type    |  Description |
| --------- | --------- | ------------ |
| data | string |  |





## Enums

## Objects

# API Protocol

Table of Contents

* [Service ProtoDoc](#service-protodoc)
    * [Method ProtoDoc.Convert](#method-protodocconvert)
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
| markdown | string | markdown 格式的文档 |
| html | string | HTML 格式的文档 |





## Enums

## Objects

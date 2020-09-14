# face-go-sdk
百度人脸识别SDK

- [x] 人脸库管理  
- [x] 在线活体检测 
- [x] 人脸搜索 
- [x] 人脸对比 
- [x] 身份验证 
- [x] 手势识别 

## 示例
具体用法请参照测试用例
```go
var f = New(APP_KEY, APP_SECRET)
if res, err := f.GetUser("1", "demo1"); err != nil {
	fmt.Printf("%+v",err)
} else {
	fmt.Printf("%+v", res)
}
```

> 错误码表 https://cloud.baidu.com/doc/FACE/s/xk25rddsw

### 具体请查阅API文档 https://cloud.baidu.com/doc/FACE/s/rk25rddle

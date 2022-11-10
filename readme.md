# Hyperledger Fabric Chaincode with golang

This is a demo chaincode with golang


## how to run
~~~sh
git clone git@github.com:tekintian/myfcc-go.git
cd myfcc-go
# 去掉不需要的依赖和增加依赖引用
go mod tidy

# 下载依赖项目到本地的vendor目录
go mod vendor
~~~

## Package Chaincode


~~~sh
# 打包chaincode 注意你需要先配置全局的fabric bin目录的PATH才能在当前目录下使用 peer命令
peer lifecycle chaincode package myfcc_v1.tar.gz --path $(pwd) --lang golang --label myfcc_v1

# 在未安装之前可以先计算你的包的ID, 这个ID非常重要，是你的chaincode的唯一标识，后面会用到
peer lifecycle chaincode calculatepackageid myfcc_v1.tar.gz

# 在需要的peer上面安装 chaincode
peer lifecycle chaincode install myfcc_v1.tar.gz
~~~

- 参数说明：
--path 你写好的chaincode文件的路径，$(pwd) 为当前路径
--lang 指定你自己写的chaincode的语言，当前为 golang
--label 包的标签名称，最好是包含名称和版本 









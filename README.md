go mod init github.com/CarireLi/go-zore.git

git@github.com:CarireLi/go-zore.git


<!-- 配置用户 -->
git config --global user.email "2459400814@qq.com"
git config --global user.name "CarireLi"
  
  
git config --global user.name
git config --global user.email



goctl rpc protoc *.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
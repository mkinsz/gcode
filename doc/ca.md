# SSL

## 捷径

使用mkcert

## CA与自签名
CA 是权威机构才能做的，并且如果该机构达不到安全标准就会被浏览器厂商“封杀”，前不久的沃通、StartSSL 就被 Mozilla、Chrome 封杀了。不过这并不影响我们进行双向认证配置，因为我们是自建 CA 的..

### 制作CA私钥
openssl genrsa -out ca.key 2048
### 制作 CA 根证书（公钥）
openssl req -new -x509 -days 3650 -key ca.key -out ca.crt

## 服务器端证书

### 制作服务端私钥
openssl genrsa -out server.pem 1024
openssl rsa -in server.pem -out server.key

### 生成签发请求
openssl req -new -key server.pem -out server.csr

### 用CA签发
openssl x509 -req -sha256 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 -out server.crt


## Tips

openssl req -x509 -nodes -days 36500 -newkey rsa:2048 -keyout nginx.key -out nginx.crt

req: 配置参数-x509指定使用 X.509证书签名请求管理(certificate signing request (CSR))."X.509" 是一个公钥代表that SSL and TLS adheres to for its key and certificate management.
-nodes: 告诉OpenSSL生产证书时忽略密码环节.(因为我们需要Nginx自动读取这个文件，而不是以用户交互的形式)。
-days 36500: 证书有效期，100年
-newkey rsa:2048: 同时产生一个新证书和一个新的SSL key(加密强度为RSA 2048)
-keyout:SSL输出文件名
-out:证书生成文件名
它会问一些问题。需要注意的是在common name中填入网站域名，如wiki.xby1993.net即可生成该站点的证书，同时也可以使用泛域名如*.xby1993.net来生成所有二级域名可用的网站证书。
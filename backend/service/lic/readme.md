# 项目说明

## 一、代码功能概述
本项目主要实现了通过获取服务器的通用唯一信息（包括CPU信息、主板序列号、网卡信息等），结合RSA密钥对生成及加密解密操作，来实现license授权的唯一性和有效性验证。具体功能如下：
- 生成指定长度（示例中为2048位）的RSA密钥对。
- 获取服务器在不同环境（包括Docker容器内）下的通用唯一信息，如CPU信息、主板序列号、网卡信息等，并将这些信息与license一同保存到文件中。
- 能够从文件中加载私钥、提取license和服务器信息，并进行对比验证，确保授权的一致性。
- 实现了使用RSA公钥对数据进行加密以及使用私钥对加密数据进行解密的功能。

## 二、使用流程

### （一）生成密钥对、获取服务器信息并保存相关数据
1. **生成RSA密钥对**：
在`main`函数中，首先通过调用`generateRSAKeyPair`函数生成一个2048位的RSA密钥对。示例代码如下：
```go
privateKey, err := generateRSAKeyPair(2048)
if err!= nil {
    fmt.Println("Error generating RSA key pair:", err)
    return
}
```
2. **获取服务器通用唯一信息**：
接着调用`getServerUniqueInfo`函数获取服务器的通用唯一信息，该函数会尝试获取CPU信息、主板序列号（示例中的获取方式可能因系统不同需调整）、网卡信息等。如果程序运行在Docker容器内，还会尝试获取主机的相关信息（如主机CPU信息、主机网卡信息等）。示例代码如下：
```go
serverInfo := getServerUniqueInfo()
```
3. **定义license并保存相关数据**：
定义一个`license`字符串（示例中简单设置为"license"，实际应用中应为有效的授权码），然后通过调用`savePrivateKeyAndLicenseAndServerInfo`函数将生成的私钥、license以及获取到的服务器信息保存到名为"privateKey.pem"的文件中。示例代码如下：
```go
license := "license"
err = savePrivateKeyAndLicenseAndServerInfo(privateKey, license, serverInfo, "privateKey.pem")
if err!= nil {
    fmt.Println("Error saving private key:", err)
    return
}
```

### （二）加载数据并进行验证
1. **从文件加载相关数据**：
通过调用`loadPrivateKeyAndLicenseAndServerInfo`函数从"privateKey.pem"文件中加载私钥、提取license和服务器信息。示例代码如下：
```go
loadedPrivateKey, loadedLicense, loadedServerInfo, err := loadPrivateKeyAndLicenseAndServerInfo("privateKey.pem")
if err!= nil {
    fmt.Println("Error loading private key:", err)
    return
}
```
2. **验证license和服务器信息**：
将加载出来的license和服务器信息与之前保存时的原始信息进行对比验证，如果不匹配则输出错误信息并终止程序。示例代码如下：
```go
if license!= loadedLicense || serverInfo!= loadedServerInfo {
    fmt.Println("License or server info mismatch!")
    return
}
```

### （三）数据加密与解密操作
1. **加密数据**：
定义要加密的数据（示例中为"Hello, RSA!"），然后通过调用`encryptWithPublicKey`函数使用生成的RSA公钥对数据进行加密，并输出加密后的结果。示例代码如下：
```go
data := "Hello, RSA!"
encryptedData, err := encryptWithPublicKey([]byte(data), &privateKey.PublicKey)
if err!= nil {
    fmt.Println("Error encrypting data:", err)
    return
}
fmt.Println("Encrypted data:", encryptedData)
```
2. **解密数据**：
最后通过调用`decryptWithPrivateKey`函数使用加载出来的私钥对加密数据进行解密，并输出解密后的结果。示例代码如下：
```go
decryptedData, err := decryptWithPrivateKey(encryptedData, loadedPrivateKey)
if err!= nil {
    fmt.Println("Error decrypting data:", err)
    return
}
fmt.Println("Decrypted data:", string(decryptedData))
```

## 三、不同环境下文件挂载及获取主机信息相关说明

### （一）Docker环境
1. **判断是否在Docker容器内**：
代码通过`isInDocker`函数来判断程序是否运行在Docker容器内。该函数通过读取`/proc/1/cgroup`文件内容，并检查是否包含"docker"字符串来进行判断。示例代码如下：
```go
if isInDocker() {
    // 在Docker容器内的相关操作
}
```
2. **在Docker容器内读取主机文件内容**：
如果程序在Docker容器内，`readHostFileInDocker`函数可用于读取主机文件内容。它通过将容器内的路径与`/host`进行拼接来获取主机文件的实际路径，例如读取主机的`/proc/cpuinfo`文件内容可通过如下方式：
```go
hostCpuInfo, err := readHostFileInDocker("/proc/cpuinfo")
if err == nil {
    // 对读取到的主机CPU信息进行处理
}
```
3. **获取Docker容器内虚拟网卡对应的主机网卡信息**：
`getHostNetworkInterfacesInDocker`函数用于获取Docker容器内虚拟网卡对应的主机网卡信息。在示例代码中，它假设通过挂载`/host/etc/hosts`这样的路径可以访问主机文件系统，然后通过读取该文件内容并进行解析（示例中只是简单示意，实际需完善解析逻辑）来获取主机网卡信息。示例代码如下：
```go
hostIfaces, err := getHostNetworkInterfacesInDocker()
if err!= nil {
    fmt.Println("Error getting host network interfaces in Docker:", err)
} else {
    for _, hostIface := range hostIfaces {
        // 对获取到的主机网卡信息进行处理
    }
}
```
4. **Docker文件挂载写法示例**：
在启动Docker容器时，需要进行适当的文件挂载以便容器内程序能够访问主机文件系统获取相关信息。以下是一个简单的`docker run`命令示例，假设当前目录下有一个名为`app`的可执行文件（对应本项目编译后的程序），并且要挂载主机的`/proc`和`/etc`目录到容器内的`/host/proc`和`/host/etc`目录：
```bash
docker run -v /proc:/host/proc -v /etc:/host/etc -it --name myapp app
```
在上述命令中，`-v`参数用于指定文件挂载，格式为`主机目录:容器内目录`。通过这样的挂载，容器内程序就可以按照代码中定义的方式（如`/host/proc/cpuinfo`等路径）来访问主机文件获取相关信息。

### （二）Docker Compose环境
1. **配置文件挂载**：
在`docker-compose.yml`文件中，可以通过`volumes`字段来配置文件挂载。以下是一个简单的示例配置，假设要实现与上述`docker run`命令类似的挂载效果：
```yaml
version: '3'
services:
    myapp:
        image: your_image_name
        volumes:
            - /proc:/host/proc
            - /etc:/host/etc
        command: ["your_app_command"]
```
在上述配置中，`volumes`字段下的每一项同样按照`主机目录:容器内目录`的格式来指定文件挂载。这里的`myapp`是服务名称，`image`指定了要使用的镜像，`command`则是容器启动后要执行的命令（需替换为实际对应本项目程序的启动命令）。
2. **获取主机信息**：
一旦按照上述方式进行了文件挂载配置，容器内运行的本项目程序就可以按照在Docker环境下相同的方式（即通过代码中`readHostFileInDocker`、`getHostNetworkInterfacesInDocker`等函数）来获取主机信息。

### （三）K8S环境
1. **配置Volume挂载**：
在Kubernetes（K8S）环境中，需要通过创建`Volume`和`VolumeMount`资源来实现文件挂载。以下是一个简单的示例，假设要将主机的`/proc`和`/etc`目录挂载到Pod内的容器中以便获取主机信息：
首先，创建一个`Volume`资源，示例如下：
```yaml
apiVersion: v1
kind: Volume
metadata:
    name: host-volumes
spec:
    hostPath:
        path: /proc
---
apiVersion: v1
kind: Volume
metadata:
        name: host-volumes-etc
spec:
    hostPath:
        path: /etc
```
上述代码创建了两个`Volume`资源，分别用于挂载主机的`/proc`和`/etc`目录，名称分别为`host-volumes`和`host-volumes-etc`。

然后，在Pod的定义中配置`VolumeMount`来挂载这些`Volume`资源到容器内。示例如下：
```yaml
apiVersion: v1
kind: Pod
metadata:
    name: myapp-pod
spec:
    containers:
        - name: myapp-container
            image: your_image_name
            volumeMounts:
                - name: host-volumes
                    mountPath: /host/proc
                - name: host-volumes-etc
                    mountPath: /host/etc
    volumes:
        - name: host-volumes
            volume: &host-volumes
                hostPath:
                    path: /proc
        - name: host-volumes-etc
            volume: &host-volumes-etc
                hostPath:
                    path: /etc
```
在上述Pod定义中，`volumeMounts`字段用于指定容器内的挂载路径，`volumes`字段则用于引用之前创建的`Volume`资源。通过这样的配置，容器内运行的本项目程序就可以按照在Docker环境下类似的方式（通过代码中相关函数）来获取主机信息。
2. **获取主机信息**：
同样，在K8S环境下，一旦完成上述文件挂载配置，容器内运行的本项目程序就可以通过代码中定义的函数（如`readHostFileInDocker`、`getHostNetworkInterfacesInDocker`等函数，虽然名称中有Docker，但在K8S环境下通过类似的挂载方式也可实现类似功能）来获取主机信息。

通过上述在不同环境下的文件挂载及相关操作，本项目代码能够准确获取到主机信息，从而确保在部署后不易被伪造主机信息，达到license授权的唯一性、有效性。
todo:
1.生成主机信息文件
2.上传license文件
package main

import (
    "bufio"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "io/ioutil"
    "net"
    "os"
    "path/filepath"
    "strings"

    "golang.org/x/sys/unix"
)

// 生成RSA密钥对
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err!= nil {
        return nil, err
    }
    return privateKey, nil
}

// 获取服务器通用唯一信息
func getServerUniqueInfo() string {
    var info strings.Builder

    // 获取 CPU 信息
    cpuInfo, err := ioutil.ReadFile("/proc/cpuinfo")
    if err == nil {
        info.Write(cpuInfo)
    }

    // 获取主板序列号（这里只是示例，不同系统获取方式不同，可能不准确）
    var serialNumber [64]byte
    unix.SysctlRaw("hw.model", serialNumber[:])
    info.Write(serialNumber[:])

    // 获取网卡信息
    ifaces, err := net.Interfaces()
    if err!= nil {
        fmt.Println("Error getting network interfaces:", err)
    } else {
        for _, iface := range ifaces {
            info.WriteString(iface.Name + " ")
            mac := iface.HardwareAddr
            if mac!= nil {
                info.WriteString(mac.String() + " ")
            }
            addrs, err := iface.Addrs()
            if err!= nil {
                fmt.Println("Error getting addresses for interface", iface.Name, err)
            } else {
                for _, addr := format.Sprintf("Error getting addresses for interface %s: %s", iface.Name, err)
            } else {
                for _, addr := range addrs {
                    info.WriteString(addr.String() + " ")
                }
            }
        }
    }

    // 判断是否在 Docker 容器内，如果是，尝试获取主机信息
    if isInDocker() {
        hostCpuInfo, err := readHostFileInDocker("/proc/cpuinfo")
        if err == nil {
            info.Write(hostCpuInfo)
        }
        // 这里还可以添加获取主机主板序列号等其他信息的逻辑，与上面类似，通过读取主机相关文件

        // 获取Docker容器内虚拟网卡对应的主机网卡信息
        hostIfaces, err := getHostNetworkInterfacesInDocker()
        if err!= nil {
            fmt.Println("Error getting host network interfaces in Docker:", err)
        } else {
            for _, hostIface := range hostIfaces {
                info.WriteString(hostIface.Name + " ")
                mac := hostIface.HardwareAddr
                if mac!= nil {
                    info.WriteString(mac.String() + " ")
                }
                addrs, err := hostIface.Addrs()
                if err!= nil {
                    fmt.Println("Error getting addresses for host interface", hostIface.Name, err)
                } else {
                    for _, addr := range addrs {
                        info.WriteString(addr.String() + " ")
                    }
                }
            }
        }
    }

    return info.String()
}

// 判断是否在 Docker 容器内
func isInDocker() bool {
    file, err := os.Open("/proc/1/cgroup")
    if err!= nil {
        return false
    }
    defer file.Close()

    scanner := bufio.Scanner(file)
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), "docker") {
            return true
        }
    }
    return false
}

// 在 Docker 容器内读取主机文件内容
func readHostFileInDocker(path string) ([]byte, error) {
    hostPath := filepath.Join("/host", path)
    return ioutil.ReadFile(hostPath)
}

// 获取Docker容器内虚拟网卡对应的主机网卡信息
func getHostNetworkInterfacesInDocker() ([]net.Interface, error) {
    // 这里假设Docker容器内可以通过 /host/etc/hosts 这样的挂载路径访问主机文件系统
    // 实际情况可能因Docker配置而异，需要根据具体环境进行调整
    hostEtcPath := filepath.Join("/host/etc")
    etcHostsPath := filepath.Join(hostEtcPath, "hosts")

    // 读取 /host/etc/hosts 文件内容，从中解析出主机网卡信息（这里只是简单示例，实际可能更
    // 复杂）
    hostsContent, err := ioutil.ReadFile(etcHostsPath)
    if err!= nil {
        return nil, err
    }

    // 这里可以根据解析 hostsContent 的结果，进一步获取主机网卡信息
    // 例如，通过查找特定的IP地址关联的网卡等方式
    // 以下只是简单返回一个空的接口切片作为示例，实际需要完善此逻辑
    return []net.Interface{}, nil
}

// 保存私钥、license和服务器信息到文件
func savePrivateKeyAndLicenseAndServerInfo(privateKey *rsa.PrivateKey, license string, serverInfo string, filename string) error {
    privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
    combinedData := append(privateKeyBytes, []byte(license+serverInfo)...)
    privateKeyBlock := &pem.Block{
        Type:  "RSA PRIVATE KEY WITH LICENSE AND SSERVER INFO",
        Bytes: combinedData,
    }
    file, err := os.Create(filename)
    if err!= nil {
        return err
    }
    err = pem.Encode(file, privateKeyBlock)
    if err!= nil {
        file.Close()
        return err
    }
    file.Close()
    return nil
}

// 从文件加载私钥、提取license和服务器信息
func loadPrivateKeyAndLicenseAndServerInfo(filename string) (*rsa.PrivateKey, string, string, error) {
    file, err := os.ReadFile(filename)
    if err!= nil {
    return nil, "", "", err
    }
    privateKeyBlock, _ := pem.Decode(file)
    if privateKeyBlock == nil {
        return nil, "", "", fmt.Errorf("failed to decode PEM block containing the key")
    }
    privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes[:len(privateKeyBlock.Bytes)-len("license")-len(getServerUniqueInfo())])
    if err!= nil {
        return nil, "", "", err
    }
    license := string(privateKeyBlock.Bytes[len(privateKeyBlock.Bytes)-len("license")-len(getServerUniqueInfo()): len(privateKeyBlock.Bytes)-len(getServerUniqueInfo())])
    serverInfo := string(privateKeyBlock.Bytes[len(privateKeyBlock.Bytes)-len(getServerUniqueInfo()):])
    return privateKey, license, serverInfo, nil
}

// 使用公钥加密
func encryptWithPublicKey(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
    encryptedBytes, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
    if err!= nil {
        return nil, err
    }
    return encryptedBytes, nil
}

// 使用私钥解密
func decryptWithPrivateKey(encryptedBytes []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
    plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedBytes)
    if err!=nil 
    return plaintext, nil
}

// 检查license文件是否有效，license文件是一个简单文本文件，包含一个特定授权码
func isLicenseValid(licenseFilePath string, expectedLicenseCode string) bool {
    licenseContent, err := ioutil.ReadFile(licenseFilePath)
    if err!= nil {
        fmt.Println("Error reading license file:", err)
        return false
    }

    actualLicenseCode := strings.TrimSpace(string(licenseContent))
    return actualLicenseCode == expectedLicenseCode
}

func main() {
    // 生成密钥对
    privateKey, err := generateRSAKeyPair(2048)
    if err!= nil {
        fmt.Println("Error generating RSA key pair:", err)
        return
    }

    license := "license"
    serverInfo := getServerUniqueInfo()
    // 保存私钥、license和服务器信息
    err = savePrivateKeyAndLicenseAndServerInfo(privateKey, license, serverInfo, "privateKey.pem")
    if err!= nil {
        fmt.Println("Error saving private key:", err)
        return
    }

    // 加载私钥、license和服务器信息
    loadedPrivateKey, loadedLicense, loadedServerInfo, err := loadPrivateKeyAndLicenseAndServerInfo("privateKey.pem")
    if err!= nil {
        fmt.Println("Error loading private key:", err)
        return
    }

    // 对比license和服务器信息
    if license!= loadedLicense || serverInfo!= loadedServerInfo {
        fmt.Println("License or server info mismatch!")
        return
    }

    // 要加密的数据
    data := "Hello, RSA!"

    // 使用公钥加密数据
    encryptedData, err := encryptWithPublicKey([]byte(data), &privateKey.PublicKey)
    if err!= nil {
        fmt.Println("Error encrypting data:", err)
        return
    }
    fmt.Println("Encrypted data:", encryptedData)

    // 使用私钥解密数据
    decryptedData, err := decryptWithPrivateKey(encryptedData, loadedPrivateKey)
    if err!= nil {
        fmt.Println("Error decrypting data:", err)
        return
    }
    fmt.Println("Decrypted data:", string(decryptedData))
}
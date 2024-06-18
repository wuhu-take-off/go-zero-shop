package file_load

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"strconv"
	"time"
)

type (
	FileLoad interface {
		DownloadFile(path string) ([]byte, error)
		UpFile(data []byte, filename string) error
		Close() error
	}
	SFTP struct {
		Host       string
		Port       int
		User       string
		Password   string
		ServerPath string
	}
	defaultFileLoad struct {
		sftpClient *sftp.Client
		SFTP
	}
)

func NewFileLoad(s SFTP) (FileLoad, error) {
	connect, err := SftpConnect(s.Host, s.User, s.Password, s.Port)
	if err != nil {
		return nil, err
	}
	return &defaultFileLoad{
		sftpClient: connect,
		SFTP:       s,
	}, nil
}

func (m *defaultFileLoad) DownloadFile(path string) ([]byte, error) { //远程文件路径，返回二进制流
	path = m.ServerPath + path
	srcFile, err := m.sftpClient.Open(path)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(srcFile)
	if err != nil {
		return nil, err
	}
	//if err != nil {
	//	return err
	//}
	//err = os.WriteFile("restored_file.png", bytes, 0644)
	//if err != nil {
	//	return nil, err
	//}
	return bytes, err
}

func (m *defaultFileLoad) UpFile(data []byte, filename string) error { //参数：本地文件位置，新文件名称
	filename = m.ServerPath + filename
	remoteFile, _ := m.sftpClient.Create(filename)
	_, err := remoteFile.Write(data)
	if err != nil {
		return err
	}
	return nil
}
func (m *defaultFileLoad) Close() error {
	return m.sftpClient.Close()
}

func SftpConnect(host, user, password string, port int) (sftpClient *sftp.Client, err error) { //参数: 远程服务器用户名, 密码, ip, 端口
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig := &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := host + ":" + strconv.Itoa(port)
	sshClient, err := ssh.Dial("tcp", addr, clientConfig) //连接ssh
	if err != nil {
		fmt.Println("连接ssh失败", err)
		return
	}

	if sftpClient, err = sftp.NewClient(sshClient); err != nil { //创建客户端
		fmt.Println("创建客户端失败", err)
		return
	}
	return
}

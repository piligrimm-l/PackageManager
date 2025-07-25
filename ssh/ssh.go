package ssh

import (
    "github.com/pkg/sftp"
    "golang.org/x/crypto/ssh"
)

func UploadFile(client *sftp.Client, localPath, remotePath string) error {
    srcFile, err := os.Open(localPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    dstFile, err := client.Create(remotePath)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    return err
}

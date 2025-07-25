package utils

import (
    "archive/tar"
    "compress/gzip"
    "os"
)

func CreateTarGzArchive(filePaths []string, archiveName string) error {
    fw, err := os.Create(archiveName)
    if err != nil {
        return err
    }
    defer fw.Close()

    gzw := gzip.NewWriter(fw)
    defer gzw.Close()

    tw := tar.NewWriter(gzw)
    defer tw.Close()

    for _, filePath := range filePaths {
        err := addToTar(tw, filePath)
        if err != nil {
            return err
        }
    }

    return nil
}

func addToTar(tw *tar.Writer, path string) error {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return err
    }
    header, err := tar.FileInfoHeader(fileInfo, "")
    if err != nil {
        return err
    }

    if err := tw.WriteHeader(header); err != nil {
        return err
    }

    switch mode := fileInfo.Mode(); {
    case mode.IsDir():
        dirFiles, _ := ioutil.ReadDir(path)
        for _, fi := range dirFiles {
            addToTar(tw, filepath.Join(path, fi.Name()))
        }
    case mode.IsRegular():
        data, err := ioutil.ReadFile(path)
        if err != nil {
            return err
        }
        if _, err := tw.Write(data); err != nil {
            return err
        }
    default:
        return fmt.Errorf("unsupported file type %v for %q", mode, path)
    }

    return nil
}

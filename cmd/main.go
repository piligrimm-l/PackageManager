package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "github.com/spf13/viper"
    "github.com/piligrimm-l/PackageManager/internal/jsonproc"
    "github.com/piligrimm-l/PackageManager/internal/ssh"
    "github.com/piligrimm-l/PackageManager/internal/utils"
)

var configPath string

func init() {
    flag.StringVar(&configPath, "c", "", "Configuration file")
    viper.SetConfigType("yaml") // Поддержка yaml конфигурации
}

func main() {
    flag.Parse()
    args := flag.Args()

    if len(args) < 1 || args[0] == "" {
        log.Fatalf("Command is required.")
        os.Exit(1)
    }

    command := args[0]
    switch command {
    case "create":
        createCmd(os.Stdout, args[1:])
    case "update":
        updateCmd(os.Stdout, args[1:])
    default:
        fmt.Printf("Unknown command '%s'\n", command)
        os.Exit(1)
    }
}

func createCmd(out io.Writer, args []string) {
    if len(args) < 1 {
        log.Fatal("No package configuration provided")
    }

    configPath := args[0]
    packet, err := pkg.LoadPackage(configPath)
    if err != nil {
        log.Fatalf("Error loading package configuration: %v\n", err)
    }

    // Создаем архив, отправляем на сервер
    filesForArchiving := packet.CollectFiles()
    archiveName := fmt.Sprintf("%s-%s.tar.gz", packet.Name, packet.Version)

    err = utils.CreateTarGzArchive(filesForArchiving, archiveName)
    if err != nil {
        log.Fatalf("Failed to create archive: %v\n", err)
    }

    client, err := ssh.Connect()
    if err != nil {
        log.Fatalf("SSH connection failed: %v\n", err)
    }
    defer client.Close()

    err = ssh.UploadFile(client, archiveName, "/remote/path/"+archiveName)
    if err != nil {
        log.Fatalf("Uploading archive failed: %v\n", err)
    }

    fmt.Fprintf(out, "Package created and uploaded successfully.\n")
}

func updateCmd(out io.Writer, args []string) {
    if len(args) < 1 {
        log.Fatal("No packages specification provided")
    }

    specPath := args[0]
    packages, err := pkg.LoadPackages(specPath)
    if err != nil {
        log.Fatalf("Error loading packages specification: %v\n", err)
    }

    // Проверяем наличие каждого пакета, загружаем необходимые пакеты
    for _, p := range packages {
        err := processPackage(p)
        if err != nil {
            log.Fatalf("Processing package %s failed: %v\n", p.Name, err)
        }
    }

    fmt.Fprintf(out, "All packages updated successfully.\n")
}

func processPackage(packetRef pkg.PacketRef) error {
    // Логика проверки наличия пакета на сервере,
    // скачивания и распаковки нужного пакета
    return nil
}

package jsonyamlproc

import (
    "io"
)

type Package struct {
    Name   string     `json:"name"`
    Version string    `json:"ver"`
    Targets []Target  `json:"targets"` // Пути к целевым файлам
    Packets []PacketRef `json:"packets"` // Зависимости пакетов
}

// Target описывает путь к файлу или каталогу с опциональным исключением файлов
type Target struct {
    Path    string   `json:"path"`
    Exclude []string `json:"exclude,omitempty"`
}

// PacketRef ссылается на внешний пакет с указанием версии
type PacketRef struct {
    Name    string `json:"name"`
    Version string `json:"ver"`
}

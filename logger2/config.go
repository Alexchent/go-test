package logger

import (
	_ "embed"
	"strings"
)

const (
	OutputTypeConsole = "console"
	OutputTypeFile    = "file"
)

const DefaultCategoryName = "default"
const DefaultConfigFileName = "logger.toml"

//go:embed logger.default.toml
var configTemplate string

func GetConfigTemplate() string {
	return configTemplate
}

type EnvKey string

// 配置文件中可使用的环境变量，使用形式为${varName}
const (
	EnvProgramName EnvKey = "programName"
	EnvLogDir      EnvKey = "logDir"
)

func (k EnvKey) Ref() string {
	return "${" + string(k) + "}"
}

type Output interface {
	GetType() string
	GetCategories() []string
}

type OutputBase struct {
	Type     string `json:"type" toml:"type" mapstructure:"type"`             // 类型，可选值为"console", "file"
	Category string `json:"category" toml:"category" mapstructure:"category"` // 分类名称, 默认分类名称为"default"
}

func (ob OutputBase) GetType() string {
	return ob.Type
}

func (ob OutputBase) GetCategories() []string {
	if len(ob.Category) == 0 {
		return []string{DefaultCategoryName}
	}
	fields := strings.Split(ob.Category, ",")
	categories := make([]string, 0, len(fields))
	for _, field := range fields {
		str := strings.TrimSpace(field)
		if len(str) > 0 {
			categories = append(categories, str)
		}
	}
	return categories

}

// ConsoleOutput 类型为console时的配置结构
type ConsoleOutput struct {
	OutputBase `mapstructure:",squash"`
	Colored    bool `json:"colored" toml:"colored" mapstructure:"colored"`
}

// FileOutput 类型为file时的配置结构
type FileOutput struct {
	OutputBase `mapstructure:",squash"`
	FileName   string `json:"fileName" toml:"fileName" mapstructure:"fileName"`
	MaxSize    int    `json:"maxSize" toml:"maxSize" mapstructure:"maxSize"`
	MaxDays    int    `json:"maxDays" toml:"maxDays" mapstructure:"maxDays"`
	MaxBackups int    `json:"maxBackups" toml:"maxBackups" mapstructure:"maxBackups"`
}

// Config logger的配置文件结构，其中Outputs需要二次解析（根据type字段解析为ConsoleOutput或FileOutput)
type Config struct {
	Outputs []map[string]interface{} `json:"outputs" toml:"outputs"`
	Levels  map[string]string        `json:"levels" toml:"levels"`
	Level   string                   `json:"level" toml:"level"`
	LogDir  string                   `json:"logDir" toml:"logDir"`
	Format  string                   `json:"format" toml:"format"`
}

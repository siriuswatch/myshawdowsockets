package cmd
import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	configPath string
)

type Config struct {
	ListenAddr string `json:"listen"`
	RemoteAddr string `json:"remote"`
	Password   string `json:"password"`
}

func init() {
	home, _ := homedir.Dir()
	// 默认的配置文件名称
	configFilename := ".myshawdowsocks.json"
	// 如果用户有传配置文件，就使用用户传入的配置文件
	if len(os.Args) == 2 {
		configFilename = os.Args[1]
	}
	configPath = path.Join(home, configFilename)
}

// 保存配置到配置文件
func (config *Config) SaveConfig() {
	configJson, _ := json.MarshalIndent(config, "", "	")
	err := ioutil.WriteFile(configPath, configJson, 0644)
	if err != nil {
		fmt.Errorf("save the config file to %s ||error: %s", configPath, err)
	}
	log.Printf("save config file to the path %s successfully! \n", configPath)
}

func (config *Config) ReadConfig() {
	// 如果配置文件存在，就读取配置文件中的配置 assign 到 config
	if _, err := os.Stat(configPath); !os.IsNotExist(err) {
		log.Printf("reading config from %s and waiting...\n", configPath)
		file, err := os.Open(configPath)
		if err != nil {
			log.Fatalf("open the config file: %s ||error:%s", configPath, err)
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(config)
		if err != nil {
			log.Fatalf("format is invalid in your config JSON:\n%s", file.Name())
		}
	}
}
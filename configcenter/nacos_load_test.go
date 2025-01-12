package configcenter

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/apache/dubbo-go-pixiu/pkg/logger"
	. "github.com/smartystreets/goconvey/convey"
)

// isNacosRunning checks whether the Nacos server is running.
// It returns true if Nacos is running, otherwise false.
func isNacosRunning(t *testing.T) bool {
	t.Helper()
	_, err := getNacosConfigClient(getBootstrap())
	return err == nil
}

// TestNewNacosConfig tests the creation of a new Nacos configuration.
// If Nacos is not running, the test is skipped.
func TestNewNacosConfig(t *testing.T) {
	if !isNacosRunning(t) {
		t.Skip("Nacos is not running, skipping the test.")
		return
	}

	Convey("Test NewNacosConfig", t, func() {
		cfg := getBootstrap()

		// Test successful creation of NacosConfig.
		_, err := NewNacosConfig(cfg)
		So(err, ShouldBeNil)

		// Test creation failure when Nacos server configurations are missing.
		cfg.Nacos.ServerConfigs = nil
		_, err = NewNacosConfig(cfg)
		So(err, ShouldNotBeNil)
	})
}

// TestNacosConfig_onChange tests the onChange method of NacosConfig.
func TestNacosConfig_onChange(t *testing.T) {
	Convey("TestNacosConfig_onChange", t, func() {
		cfg := getBootstrap()
		c, err := NewNacosConfig(cfg)
		So(err, ShouldBeNil)

		client, ok := c.(*NacosConfig)
		So(ok, ShouldBeTrue)

		// Verify the current working directory.
		wd, err := os.Getwd()
		So(err, ShouldBeNil)

		paths := strings.Split(wd, "/")
		So(paths[len(paths)-1], ShouldEqual, "configcenter")

		// Open the configuration file for testing.
		file, err := os.Open(fmt.Sprintf("/%s/configs/conf.yaml", path.Join(paths[:len(paths)-1]...)))
		So(err, ShouldBeNil)
		defer func() { So(file.Close(), ShouldBeNil) }()

		conf, err := io.ReadAll(file)
		So(err, ShouldBeNil)

		Convey("Test onChange with valid input", func() {
			So(client.remoteConfig, ShouldBeNil)
			client.onChange(Namespace, Group, DataId, string(conf))
			So(client.remoteConfig, ShouldNotBeNil)
		})

		Convey("Test onChange with empty input", func() {
			// Suppress logs during this test.
			logger.SetLoggerLevel("fatal")

			client.remoteConfig = nil
			client.onChange(Namespace, Group, DataId, "")
			So(client.remoteConfig, ShouldBeNil)

			// Restore the logger level.
			logger.SetLoggerLevel("info")
		})
	})
}

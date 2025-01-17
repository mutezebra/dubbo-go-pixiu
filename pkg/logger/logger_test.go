/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestInitLog(t *testing.T) {
	var (
		err  error
		path string
	)

	err = InitLog("")
	assert.EqualError(t, err, "log configure file name is nil")

	path, err = filepath.Abs("./log.xml")
	assert.NoError(t, err)
	err = InitLog(path)
	assert.EqualError(t, err, "log configure file name "+path+" suffix must be .yml")

	path, err = filepath.Abs("./logger.yml")
	assert.NoError(t, err)
	err = InitLog(path)
	var errMsg string
	if runtime.GOOS == "windows" {
		errMsg = fmt.Sprintf("open %s: The system cannot find the file specified.", path)
	} else {
		errMsg = fmt.Sprintf("open %s: no such file or directory", path)
	}
	assert.EqualError(t, err, fmt.Sprintf("os.ReadFile file:%s, error:%s", path, errMsg))

	err = InitLog("./log.yml")
	assert.NoError(t, err)

	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	Debugf("%s", "debug")
	Infof("%s", "info")
	Warnf("%s", "warn")
	Errorf("%s", "error")
}

func TestSetLoggerLevel(t *testing.T) {
	assert.NotNil(t, control, "control should not be nil")

	assert.True(t, SetLoggerLevel("info"), "when pass info to SetLoggerLevel, result should be true")
	assert.True(t, SetLoggerLevel("debug"), "when pass debug to SetLoggerLevel, result should be true")
	assert.True(t, SetLoggerLevel("error"), "when pass error to SetLoggerLevel, result should be true")
	assert.True(t, SetLoggerLevel("panic"), "when pass panic to SetLoggerLevel, result should be true")
	assert.True(t, SetLoggerLevel("INFO"), "when pass INFO to SetLoggerLevel, result should be true")
	assert.True(t, SetLoggerLevel("DEbug"), "when pass DEbug to SetLoggerLevel, result should be true")
	assert.True(t, SetLoggerLevel("ErRor"), "when pass ErRor to SetLoggerLevel, result should be true")
	assert.True(t, SetLoggerLevel("WaRN"), "when pass WaRN to SetLoggerLevel, result should be true")

	assert.False(t, SetLoggerLevel("i"), "when pass i to SetLoggerLevel, result should be false")
	assert.False(t, SetLoggerLevel(""), "when pass nothing to SetLoggerLevel, result should be false")
}

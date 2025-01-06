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
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logController controls log output or configuration changes throughout the project
type logController struct {
	mu sync.RWMutex

	logger *logger
}

// setLoggerLevel concurrent secure change log level
func (c *logController) setLoggerLevel(level string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	lvl := c.parseLevel(level)
	if lvl == nil {
		return false
	}

	c.logger.config.Level = *lvl
	l, _ := c.logger.config.Build()
	c.logger = &logger{SugaredLogger: l.Sugar(), config: c.logger.config}
	return true
}

// updateLogger concurrent change logger object
func (c *logController) updateLogger(l *logger) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.logger = l
}

func (c *logController) debug(args ...interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.logger.Debug(args...)
}

func (c *logController) info(args ...interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.logger.Info(args...)
}

func (c *logController) warn(args ...interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.logger.Warn(args...)
}

func (c *logController) error(args ...interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.logger.Error(args...)
}

func (c *logController) debugf(fmt string, args ...interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.logger.Debugf(fmt, args...)
}

func (c *logController) infof(fmt string, args ...interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.logger.Infof(fmt, args...)
}

func (c *logController) warnf(fmt string, args ...interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.logger.Warnf(fmt, args...)
}

func (c *logController) errorf(fmt string, args ...interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.logger.Errorf(fmt, args...)
}

// parseLevel parses the level of logs
func (c *logController) parseLevel(level string) *zap.AtomicLevel {
	var lvl zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = zapcore.DebugLevel
	case "info":
		lvl = zapcore.InfoLevel
	case "warn":
		lvl = zapcore.WarnLevel
	case "error":
		lvl = zapcore.ErrorLevel
	case "panic":
		lvl = zapcore.PanicLevel
	case "fatal":
		lvl = zapcore.FatalLevel
	default:
		return nil
	}

	al := zap.NewAtomicLevelAt(lvl)
	return &al
}

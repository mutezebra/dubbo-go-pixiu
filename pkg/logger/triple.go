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

// TripleLogger does not add user-defined logger fields so that there is no pointer reference after it is updated
// make all references point to control
type TripleLogger struct {
}

func GetTripleLogger() *TripleLogger {
	return &TripleLogger{}
}

func (l *TripleLogger) Info(args ...interface{}) {
	control.info(args...)
}

func (l *TripleLogger) Warn(args ...interface{}) {
	control.warn(args...)
}

func (l *TripleLogger) Error(args ...interface{}) {
	control.error(args...)
}

func (l *TripleLogger) Debug(args ...interface{}) {
	control.debug(args...)
}

func (l *TripleLogger) Infof(fmt string, args ...interface{}) {
	control.infof(fmt, args...)
}

func (l *TripleLogger) Warnf(fmt string, args ...interface{}) {
	control.warnf(fmt, args...)
}

func (l *TripleLogger) Errorf(fmt string, args ...interface{}) {
	control.errorf(fmt, args...)
}

func (l *TripleLogger) Debugf(fmt string, args ...interface{}) {
	control.debugf(fmt, args...)
}

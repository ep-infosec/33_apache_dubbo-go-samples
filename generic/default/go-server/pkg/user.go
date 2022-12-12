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

package pkg

import (
	"time"
)

var userMap = map[string]*User{
	"001": {"001", "other-zhangsan", 23, time.Date(1998, 1, 2, 3, 4, 5, 0, time.Local)},
	"002": {"002", "other-lisi", 25, time.Date(1996, 1, 2, 3, 4, 5, 0, time.Local)},
	"003": {"003", "other-lily", 28, time.Date(1993, 1, 2, 3, 4, 5, 0, time.Local)},
	"004": {"004", "other-lisa", 36, time.Date(1985, 1, 2, 3, 4, 5, 0, time.Local)},
}

type User struct {
	ID   string
	Name string
	Age  int32
	Time time.Time
}

func (u *User) JavaClassName() string {
	return "org.apache.dubbo.samples.User"
}

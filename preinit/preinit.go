// Copyright (c) 2019-2022 The Gim.Cool developers
// All Rights Reserved.
// NOTICE: All information contained herein is, and remains
// the property of Gim.Cool and its suppliers,
// if any. The intellectual and technical concepts contained
// herein are proprietary to Gim.Cool
// Dissemination of this information or reproduction of this materia
// is strictly forbidden unless prior written permission is obtained
// from Gim.Cool.

package preinit

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	ConfigurationFile = "conf/mockserver.app.conf"
)

func init() {
	logs.Info("LoadAppConfig %s", ConfigurationFile)
	err := beego.LoadAppConfig("ini", ConfigurationFile)

	if err != nil {
		logs.Error("LoadAppConfig,An error occurred:", err)
		os.Exit(1)
	}

}

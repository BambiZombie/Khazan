package main

import (
	"flag"
)

var (
	arch = flag.String("a", "", "运行的平台架构:\n [*] x86\n [*] x64")
	url  = flag.String("u", "", "监听的URL地址")
)

func Init() {
	banner := `
:::    ::::::    :::    :::    :::::::::    :::    ::::    ::: 
:+:   :+: :+:    :+:  :+: :+:       :+:   :+: :+:  :+:+:   :+: 
+:+  +:+  +:+    +:+ +:+   +:+     +:+   +:+   +:+ :+:+:+  +:+ 
+#++:++   +#++:++#+++#++:++#++:   +#+   +#++:++#++:+#+ +:+ +#+ 
+#+  +#+  +#+    +#++#+     +#+  +#+    +#+     +#++#+  +#+#+# 
#+#   #+# #+#    #+##+#     #+# #+#     #+#     #+##+#   #+#+# 
###    ######    ######     ###############     ######    ####


`
	print(banner)
}
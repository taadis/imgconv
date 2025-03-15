# imgconv

imgconv 这个名字参考标准库中 strconv,有 go 的命名调性

## 为什么造轮子?

有那么一天在做象棋游戏时,需要一些 .gif 图片需要转为 .png 格式, 有 `brew install imagemagick` 但是下载很慢，并且明显感觉依赖很大

而我只需要一个简单直接的工具、为什么不直接用 go 写一个呢?

最开始只有 .git to .png 的功能(满足我的最初需求),但可以预见是这个工具会有很多图片类型转化的需求，因此独立维护此库

- [ ] 作为一个 go 语言独立库,以便大家 go 项目可以直接导入依赖和集成使用
- [ ] 在独立库的集成上，实现了一套命令行工具，以便开发人员作为工具使用，比如纯终端场景
- [ ] 在独立库的基础上，实现一套桌面应用程序，以便大家使用，满足可视化界面操作

## imgconv 命令行工具

可以通过 go install 直接获取并安装 imgconv 命令行工具

```
go install github.com/taadis/imgconv/cmd/imgconv@latest
```

安装成功后，在你的 `GOBIN` 目录下会出现一个 imgconv 工具

```
ls $(go env GOPATH)/bin
# or
ls $(go env GOBIN)
```

接下来就可以使用 imgconv 命令行工具进行图片转换操作了

把某一个 .gif 格式图片转为 .png 格式

```
imgconv -input your/path/xxx.gif -output your/path/xxx.png
```

把指定目录下的所有 .gif 格式图片都转为 .png 格式

```
imgconv -input your/path/*.gif -output your/path/*.png
```

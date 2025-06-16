下面是对你当前 Markdown 文档的详细完善，面向开发者与使用者，加入了更多背景信息、功能说明、参数解释、工作流程、常见问题、以及贡献指南，方便后续开发或参与。

---

# 🧠 中英AI翻译工具（终端剪贴板增强）

## 简介

这是一个基于本地大模型（如 Ollama 的 `gemma3:4b`）的中英文翻译小工具。结合剪贴板内容与弹窗反馈，可通过自定义快捷键快速触发翻译操作，适合开发者、阅读者、写作者在 Ubuntu 环境下使用。

> 当前仅在 **Ubuntu 系统** 下测试通过，欢迎提交其它系统的兼容性反馈。

---

## 🛠 安装与部署

### 1. 安装依赖项

```bash
sudo apt install ollama xclip zenity
```

* `ollama`：用于运行本地大语言模型（LLM）
* `xclip`：获取当前剪贴板内容
* `zenity`：弹窗提示翻译结果

### 2. 启动模型（后台运行）

```bash
ollama run gemma3:4b
```

确保服务运行在 `http://localhost:11434`

### 3. 下载并部署可执行文件

```bash
git clone https://github.com/crrall/translate.git
cd translate
sudo mv ./translate /usr/bin/translate
```

> 将主程序移动到系统路径中，便于后续通过快捷命令调用。

### 4. 创建快捷命令（软链接）

```bash
sudo ln -s /usr/bin/translate /usr/bin/ten  # 翻译成英文
sudo ln -s /usr/bin/translate /usr/bin/tzh  # 翻译成中文
```

你现在可以通过 `ten` 或 `tzh` 命令来直接翻译剪贴板内容。

---

## ⚡ 快捷键设置（推荐）

在 GNOME 桌面环境下配置：

> `Settings > Keyboard > Keyboard Shortcuts > View and Customize Shortcuts > Customize Shortcuts > Add`

### 示例设置：

* Name: `translate(ten)`
* Command: `ten`
* Shortcut: `Ctrl+Alt+E`

---

## 🔍 使用方式

命令行支持以下方式：

```bash
ten "This is a test."       # 翻译成英文（ten）
tzh "这是一个测试。"        # 翻译成中文（tzh）

# 不加参数时，默认从剪贴板读取文本进行翻译
ten
tzh
```

也可以直接使用原始命令：

```bash
translate -t "内容" -l en -m gemma3:4b
```

---

## 📦 代码逻辑简介

### 翻译流程：

1. 获取执行命令名（ten 或 tzh），自动判断目标语言
2. 优先从命令行参数中获取翻译内容；若无参数，则读取剪贴板
3. 构造 prompt：要求模型 **仅输出翻译句子**，无解释、无 emoji
4. 向本地 Ollama 模型发起 HTTP 请求，获取翻译
5. 使用 `zenity` 弹窗展示翻译结果

---

## 📁 主程序结构（main.go）

### 参数说明

```go
-t : 要翻译的文本（Text）
-l : 目标语言，默认根据命令名自动选择（en / zh）
-m : 模型名称，默认使用 gemma3:4b
```

### 重要函数

* `BuildPrompt(lang, text)`：构造翻译请求的 prompt
* `opts.GetText()`：从剪贴板读取内容
* `http.NewRequest(...)`：向 Ollama 本地服务发起请求
* `zenity --info`：弹出翻译结果窗口

---

## 🤔 常见问题

### Q1: 翻译无响应？

* 确保 Ollama 模型服务已启动：

  ```bash
  curl http://localhost:11434
  ```
* 检查网络请求是否成功
* 检查剪贴板是否为空

### Q2: 如何换用别的模型？

* 修改启动命令：

  ```bash
  ollama run mistral
  ```
* 使用 `-m` 参数指定新模型名称：

  ```bash
  ten -m mistral
  ```

### Q3: 如何添加更多语言支持？

* 修改 `main.go` 的 `BuildPrompt`，添加语言提示
* 增加更多软链接，如：

  ```bash
  sudo ln -s /usr/bin/translate /usr/bin/tfr  # 翻译成法语
  ```

---

## 🧩 TODO & 开发建议

* ✅ 基于命令名判断语言（ten, tzh 等）
* ✅ 剪贴板自动读取
* ✅ zenity 弹窗提示
* 🟨 支持更多语言（如日语、韩语、德语）
* 🟨 加入 GUI 前端或系统托盘图标
* ⬜ 支持翻译历史记录查看
* ⬜ 支持 stream 流式输出翻译结果（进度显示）

---

## 🧑‍💻 参与贡献

欢迎 issue / PR！

* 提需求：功能建议、使用场景
* 报 bug：翻译结果异常、错误提示
* 贡献代码：新语言支持、CLI 参数优化、模型兼容性扩展等

---

## 📜 License

本项目遵循 MIT 许可证，详见 [LICENSE](./LICENSE)。

---

如需进一步优化或扩展功能，请随时提交 [issue](https://github.com/crrall/translate/issues)。

---


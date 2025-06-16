# 中英AI翻译

## 暂时只是在ubuntu系统下测试

# 使用方法

```base

// 安装对应的依赖(ollama:模型,xclip:剪贴板,zenity:弹窗)
sudo apt install ollama xclip zenity
// 运行模型
ollama run gemma3:4b
// 拉起仓库
git clone https://github.com/crrall/translate.git
// 进入文件
cd translate
// 移动执行文件
sudo mv ./translate /usr/bin/translate
// 创建快捷软链ten
sudo ln -s /usr/bin/translate /usr/bin/ten
// 创建快捷软链tzh
sudo ln -s /usr/bin/translate /usr/bin/tzh

```

## 可以自定义快捷键调用方式
### Settings>Keyboard>Keyboard Shortcuts>View and Customize Shortcuts>Customize Shortcuts>Add>{"Name":"translate(ten)","Command":"ten","Shortcut":"按键设置"}

# 🚀 Go Language Lab (golang-lab)

## 📝 Description

This project is a Go language lab, mainly used for testing and experimenting with Go language features. 🎉 Currently, it implements a text client that can generate front-end code (such as TailwindCSS style code) through an API. 💻

## ✅ Implemented Features

- Load configuration file. 📄
- Create a text client and set parameters. 🔧
- Handle API responses. 🔄

## ❌ Unfinished Parts

- Image processing functionality has not been implemented. 🖼️

## 🛠️ Configuration

The project configuration file `config.yaml` includes the following content:

```yaml
api:
  apiKey: 'sk-xxx'
  apiUrl: 'https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions'

server:
  host: 'localhost'
  port: 8080
```

## 🚀 Start the Project

Ensure the configuration file is correct, then run the following command to start the project:

```bash
go run main.go
```

## 🙏 Acknowledgments

Thank you for using this project! We hope it helps you in your Go language experiments. 😊

## 📜 Open Source License

This project is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.html). 📜

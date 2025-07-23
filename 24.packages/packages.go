package main

import (
	"fmt"

	"github.com/MdZunaed/go_package/auth"
	"github.com/MdZunaed/go_package/user"
	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredential("Vincenzo", "Pass1234")

	session := auth.GetSession()
	fmt.Println("session:", session)

	user := user.User{
		Email: "cassano@email.com",
		Name:  "Vincenzo Cassano",
	}
	fmt.Println(user.Name)
	color.Cyan(user.Email) // by external package
}

// promt to create package:
// go mod init github.com/MdZunaed/go_package

// promt to install external package:
// go get package_name

// To auto fix warning of package:
// go mod tidy


// Here's an expanded list of highly useful third-party Go packages, organized by category—for web, database, testing, CLI, utilities, and more.

// 🌐 Web & API Frameworks
// Package	Description
// gin-gonic/gin	🔥 Fast and minimal web framework (most popular)
// labstack/echo	Fast, full-featured alternative to gin
// gorilla/mux	Powerful router (but Gorilla project is archived now)
// fiber	Web framework inspired by Express.js (very fast)
// chi	Lightweight, idiomatic HTTP router (great for microservices)

// 🔐 Authentication & Security
// Package	Description
// dgrijalva/jwt-go	JWT token generation and parsing
// golang-jwt/jwt	Community-maintained fork of jwt-go
// golang.org/x/crypto/bcrypt	Password hashing
// securecookie	Secure, tamper-proof cookie encoding

// 🛢️ Database & ORM
// Package	Description
// go-sql-driver/mysql	MySQL driver
// lib/pq	PostgreSQL driver
// jmoiron/sqlx	Enhanced SQL with struct mapping
// gorm	Full-featured ORM for MySQL/PostgreSQL/SQLite
// ent	Entity framework from Facebook (code generation-based ORM)
// go.mongodb.org/mongo-driver/bson	🔧 BSON serialization/deserialization
// github.com/kamva/mgm	ODM wrapper (like Mongoose for Go)
// github.com/qiniu/qmgo	Mongo driver with active development, API similar to mgo

// ⚙️ Configuration / Env
// Package	Description
// spf13/viper	Advanced configuration with JSON/YAML/env support
// joho/godotenv	Loads .env files (like Node.js dotenv)

// 🧪 Testing & Mocking
// Package	Description
// stretchr/testify	Assertion + mocking for unit tests
// ginkgo	BDD testing framework
// go-playground/assert	Lightweight alternative to testify
// mockery	Interface mocking tool for Go tests

// 🧰 Utilities
// Package	Description
// go-playground/validator	Struct and field validation (used in gin)
// pterm	Beautiful terminal output (colors, tables, progress bars)
// mitchellh/mapstructure	Decode maps into Go structs
// hashicorp/go-hclog	Structured logger from HashiCorp
// sirupsen/logrus	Popular structured logging framework
// uber-go/zap	⚡ High-performance logger (faster than logrus)

// 💬 Real-Time & WebSocket
// Package	Description
// gorilla/websocket	Most-used WebSocket package
// nhooyr/websocket	Minimal and modern WebSocket library
// golang-socketio	Socket.IO (Node.js-style) support

// 🧱 CLI Apps
// Package	Description
// spf13/cobra	CLI framework used in kubectl, hugo, etc.
// urfave/cli	Easy CLI development (less complex than cobra)

// 📦 Package Management (Go Modules mostly handles this now)
// Tool	Use
// go mod	Default package manager since Go 1.11
// gopls	Language server for editor support
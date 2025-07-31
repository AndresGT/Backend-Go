package utils

import (
	"log"

	"github.com/fatih/color"
)

// Funciones de estilo para colores
var (
	Success = color.New(color.FgGreen).SprintFunc()
	Info    = color.New(color.FgCyan).SprintFunc()
	Warn    = color.New(color.FgYellow).SprintFunc()
	Error   = color.New(color.FgRed).SprintFunc()
)

// Logging semántico
func LogSuccess(msg string) {
	log.Println(Success("✅ " + msg))
}

func LogInfo(msg string) {
	log.Println(Info("ℹ️ " + msg))
}

func LogWarn(msg string) {
	log.Println(Warn("⚠️ " + msg))
}

func LogError(msg string) {
	log.Println(Error("❌ " + msg))
}

func LogFatal(msg string) {
	log.Fatal(Error("❌ " + msg)) // Finaliza el programa
}

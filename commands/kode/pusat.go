package kode

import (
	"bill/commands/fitur/github"
	fitur "bill/commands/fitur/help"
	"bill/commands/fitur/laravel"
	"bill/commands/fitur/system"
	"fmt"
	"os"
)

func Jalankan() {
	// Jika dipanggil tanpa argumen (bisa jadi Double-Click)
	if len(os.Args) < 2 {
		system.ShowWelcomeInInteractiveMode(prosesPerintah)
		return
	}

	prosesPerintah(os.Args[1:])
}

func prosesPerintah(args []string) {
	if len(args) == 0 {
		return
	}

	subCommand := args[0]

	switch subCommand {
	case "buat":
		if len(args) < 3 || args[1] != "laravel" {
			fmt.Println("    ❌ Format salah. Gunakan: bill buat laravel [nama]")
			return
		}
		laravel.BuatLaravel(args[2])

	case "cek":
		if len(args) < 2 || args[1] != "dev" {
			fmt.Println("    ❌ Format salah. Gunakan: bill cek dev")
			return
		}
		laravel.CekDev()

	case "setup":
		if len(args) < 2 || args[1] != "dev" {
			fmt.Println("    ❌ Format salah. Gunakan: bill setup dev")
			return
		}
		laravel.SetupDev()

	case "push":
		if len(args) < 3 {
			fmt.Println("    ❌ Format salah. Gunakan:")
			fmt.Println("       - bill push [link_github] [pesan]")
			fmt.Println("       - bill push update [pesan]")
			return
		}
		github.PushToGithub(args[1], args[2])

	case "install":
		system.Install()

	case "help":
		fitur.PrintHelp()

	case "exit", "keluar":
		fmt.Println("    👋 Sampai jumpa!")
		os.Exit(0)

	default:
		fmt.Printf("    ❌ Perintah '%s' tidak dikenali.\n", subCommand)
		fitur.PrintHelp()
	}
}

package kode

import (
	"bill/commands/fitur"
	"bill/commands/fitur/github"
	"bill/commands/fitur/laravel"
	"fmt"
	"os"
)

func Jalankan() {
	if len(os.Args) < 2 {
		fitur.PrintHelp()
		return
	}

	subCommand := os.Args[1]

	switch subCommand {
	case "buat":
		if len(os.Args) < 4 || os.Args[2] != "laravel" {
			fmt.Println("❌ Format salah. Gunakan: bill buat laravel [nama]")
			return
		}
		laravel.BuatLaravel(os.Args[3])

	case "cek":
		if len(os.Args) < 3 || os.Args[2] != "dev" {
			fmt.Println("❌ Format salah. Gunakan: bill cek dev")
			return
		}
		laravel.CekDev()

	case "setup":
		if len(os.Args) < 3 || os.Args[2] != "dev" {
			fmt.Println("❌ Format salah. Gunakan: bill setup dev")
			return
		}
		laravel.SetupDev()

	case "push":
		if len(os.Args) < 4 {
			fmt.Println("❌ Format salah. Gunakan:")
			fmt.Println("   - bill push [link_github] [pesan] (Pertama kali)")
			fmt.Println("   - bill push update [pesan]        (Repository lama)")
			return
		}
		github.PushToGithub(os.Args[2], os.Args[3])

	case "help":
		fitur.PrintHelp()

	default:
		fmt.Printf("❌ Perintah '%s' tidak dikenali.\n", subCommand)
		fitur.PrintHelp()
	}
}

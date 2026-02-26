package fitur

import (
	"fmt"
)

func PrintHelp() {
	// Variabel warna (Native ANSI)
	reset := "\033[0m"
	cyan := "\033[36m"
	yellow := "\033[33m"
	green := "\033[32m"
	bold := "\033[1m"

	// Pakai String Literal (Backticks) biar rapi di kode
	helpText := `
    %s╔═════════════════════════════════════════════════════╗
    ║                                                     ║
    ║   🚀  %sBILLAHI ROBBY COMMANDS%s  🚀               ║
    ║        Asisten CLI Terkeren kamu!               ║
    ║                                                     ║
    ╚═════════════════════════════════════════════════════╝%s

    ┌─── ✨ %sFITUR UTAMA%s ──────────────────────────────────┐
    │                                                     │
    │  %s🔨 LARAVEL FRAMEWORK%s                               │
    │  - bill buat laravel [nama]       → Project Baru    │
    │  - bill cek laravel              → Status Env       │
    │  - bill setup lingkungan laravel  → Pasang Tools     │
    │  - bill setup laravel           → Auto-setup Clone  │
    │  - bill ser [port]              → Nyalain Server    │
    │  - bill cleanup                 → Hapus Cache/Log   │
    │                                                     │
    │  %s🐙 GITHUB INTEGRATION%s                              │
    │  - bill push [link] [pesan]  → Push Pertama         │
    │  - bill push update [pesan]  → Push Update          │
    │                                                     │
    │  %s⚙️  SYSTEM TOOLS%s                                    │
    │  - bill help                 → Menu Bantuan         │
    │  - bill build-win [nama]     → Build Windows (Dev Only) │
    │  - bill build-linux [nama]   → Build Linux (Dev Only)   │
    │  - bill build-mac [nama]     → Build Mac (Dev Only)     │
    │  - bill info                 → Info Env             │
    │                                                     │
    └─────────────────────────────────────────────────────┘

    %s✨ Tips: Jalankan 'bill install' dulu ya!%s
    %s🌐 Info: https://github.com/Billy-dev12%s
`

	// Print dengan format warna
	fmt.Printf(helpText,
		cyan, bold, cyan, reset,
		yellow, reset,
		green, reset,
		green, reset,
		green, reset,
		yellow, reset,
		cyan, reset,
	)
}

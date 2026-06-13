package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gookit/color"
)

func main() {
	// Railway က ပေးမယ့် Port ကို ဖတ်မည်၊ မရှိပါက ပုံမှန် 8080 ကို သုံးမည်
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ပင်မစာမျက်နှာ (/) အတွက် ရိုးရှင်းသော အလုပ်တစ်ခု သတ်မှတ်ခြင်း
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Nova Proxy Web is running successfully!")
	})

	color.Green.Printf("Nova Proxy Web is starting successfully on port %s...\n", port)

	// Web Server ကို စတင် လည်ပတ်ခြင်း
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

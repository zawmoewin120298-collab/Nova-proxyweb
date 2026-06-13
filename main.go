package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gookit/color"
)

// Server အချက်အလက်နှင့် Config စာသားများကို သိမ်းဆည်းမည့် ပုံစံ
type VPNNode struct {
	Name      string
	Ping      string
	Protocol  string
	Transport string
	Config    string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ------------------------------------------------------------------
	// ⚠️ ဤနေရာတွင် အစ်ကို့ရဲ့ VPN Config (Connection Strings) များကို လဲလှယ်ထည့်သွင်းပါ
	// ------------------------------------------------------------------
	nodes := []VPNNode{
		{
			Name:      "🇹🇭 AIS Online Free 20ms", // display name formatting
			Ping:      "15 ms",
			Protocol:  "VLESS",
			Transport: "WebSocket (WS)",
			Config:    "vless://your-real-ais-thailand-config-data-here@127.0.0.1:443?path=%2F&security=tls&encryption=none#Ais%20online%2020ms", // display name formatted
		},
		{
			Name:      "🇹🇭 DTAC Unlimited Premium", // bug host testing for Dtac
			Ping:      "18 ms",
			Protocol:  "VLESS",
			Transport: "gRPC",
			Config:    "vless://your-real-dtac-config-data-here@127.0.0.1:443?encryption=none&security=tls&type=grpc#Dtac%20Premium",
		},
		{
			Name:      "🇸🇬 Singapore High-Speed VIP",
			Ping:      "28 ms",
			Protocol:  "Trojan",
			Transport: "WebSocket (WS)",
			Config:    "trojan://your-real-singapore-trojan-config-here@127.0.0.1:443?peer=sg.domain.com#Singapore%20VIP",
		},
	}

	// ပင်မစာမျက်နှာ Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		
		tmpl, err := template.New("dashboard").Parse(htmlTemplate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		// ၎င်း Nodes Data များကို UI ဘက်သို့ လှမ်းပို့ပေးခြင်း
		tmpl.Execute(w, nodes)
	})

	color.Green.Printf("Nova Proxy Config Site is active on port %s...\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// Dynamic HTML Template
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nova Proxy - Premium VIP Configs</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <style>
        :root {
            --bg-color: #0b0f19;
            --card-bg: #151f32;
            --text-main: #f8fafc;
            --text-muted: #64748b;
            --accent-green: #10b981;
            --accent-blue: #3b82f6;
            --border-line: #22314d;
        }

        * { box-sizing: border-box; margin: 0; padding: 0; font-family: 'Segoe UI', system-ui, sans-serif; }
        body { background-color: var(--bg-color); color: var(--text-main); padding: 15px; display: flex; justify-content: center; align-items: center; min-height: 100vh; }
        .container { width: 100%; max-width: 550px; background: var(--card-bg); border-radius: 20px; box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5); border: 1px solid var(--border-line); overflow: hidden; }
        
        .header { background: linear-gradient(135deg, #1e40af, #6366f1); padding: 30px 20px; text-align: center; }
        .header h1 { font-size: 24px; font-weight: 700; letter-spacing: 1px; margin-bottom: 8px; }
        
        .status-badge { display: inline-flex; align-items: center; background: rgba(16, 185, 129, 0.15); color: var(--accent-green); padding: 6px 14px; border-radius: 30px; font-size: 13px; font-weight: 600; border: 1px solid rgba(16, 185, 129, 0.3); }
        .status-dot { width: 8px; height: 8px; background-color: var(--accent-green); border-radius: 50%; margin-right: 8px; animation: pulse 1.8s infinite; }
        
        .content { padding: 20px; }
        .search-box { position: relative; margin-bottom: 20px; }
        .search-box i { position: absolute; left: 15px; top: 50%; transform: translateY(-50%); color: var(--text-muted); }
        .search-input { width: 100%; padding: 12px 15px 12px 45px; background: rgba(11, 15, 25, 0.8); border: 1px solid var(--border-line); border-radius: 12px; color: var(--text-main); font-size: 14px; }
        .search-input:focus { outline: none; border-color: var(--accent-blue); }

        .node-list { display: flex; flex-direction: column; gap: 15px; }
        .card { background: rgba(11, 15, 25, 0.4); border: 1px solid var(--border-line); border-radius: 14px; padding: 18px; transition: all 0.3s; }
        .card:hover { border-color: var(--accent-blue); }
        
        .card-header { display: flex; justify-content: space-between; align-items: center; margin-bottom: 12px; }
        .node-title { font-size: 15px; font-weight: 600; display: flex; align-items: center; gap: 8px; }
        .ping-tag { font-size: 12px; color: var(--accent-green); background: rgba(16, 185, 129, 0.1); padding: 3px 8px; border-radius: 6px; font-weight: 600; }
        
        .card-info { font-size: 13px; color: #94a3b8; margin-bottom: 15px; }
        .card-info span { color: var(--text-main); font-weight: 500; }
        
        .btn-copy { width: 100%; background: linear-gradient(135deg, #2563eb, #1d4ed8); color: white; border: none; padding: 11px; border-radius: 8px; font-size: 13px; font-weight: 600; cursor: pointer; display: flex; justify-content: center; align-items: center; gap: 8px; }
        .btn-copy:active { transform: scale(0.98); }

        .support-section { margin-top: 25px; padding-top: 20px; border-top: 1px solid var(--border-line); text-align: center; }
        .support-title { font-size: 13px; color: var(--text-muted); margin-bottom: 12px; }
        .support-buttons { display: flex; gap: 10px; justify-content: center; }
        .btn-social { padding: 8px 16px; border-radius: 8px; font-size: 12px; font-weight: 600; text-decoration: none; color: white; display: flex; align-items: center; gap: 6px; }
        .btn-messenger { background: #006aff; }

        .footer { text-align: center; padding: 15px; color: var(--text-muted); font-size: 12px; background: rgba(0,0,0,0.15); }
        .toast { position: fixed; bottom: 20px; left: 50%; transform: translateX(-50%) translateY(100px); background: var(--accent-green); color: white; padding: 10px 24px; border-radius: 30px; font-size: 14px; font-weight: 600; box-shadow: 0 4px 12px rgba(0,0,0,0.3); transition: transform 0.3s; z-index: 1000; }
        .toast.show { transform: translateX(-50%) translateY(0); }
        @keyframes pulse { 0% { transform: scale(0.95); opacity: 0.5; } 50% { transform: scale(1.15); opacity: 1; } 100% { transform: scale(0.95); opacity: 0.5; } }
    </style>
</head>
<body>

    <div class="container">
        <div class="header">
            <h1>NOVA PREMIUM CONFIGS</h1>
            <div class="status-badge"><div class="status-dot"></div>Servers Active</div>
        </div>

        <div class="content">
            <div class="search-box">
                <i class="fa-solid fa-magnifying-glass"></i>
                <input type="text" id="searchInput" class="search-input" placeholder="Search location or network...">
            </div>

            <div class="node-list" id="nodeList">
                {{range .}}
                <div class="card">
                    <div class="card-header">
                        <div class="node-title">{{.Name}}</div>
                        <div class="ping-tag">{{.Ping}}</div>
                    </div>
                    <div class="card-info">
                        <div>Protocol: <span>{{.Protocol}}</span></div>
                        <div>Transport: <span>{{.Transport}}</span></div>
                    </div>
                    <button class="btn-copy" onclick="copyConfig('{{.Config}}')">
                        <i class="fa-regular fa-copy"></i> Copy Config V2Ray
                    </button>
                </div>
                {{end}}
            </div>

            <div class="support-section">
                <div class="support-title">Contact Admin for Support</div>
                <div class="support-buttons">
                    <a href="https://m.me/your-facebook-profile" target="_blank" class="btn-social btn-messenger">
                        <i class="fa-brands fa-facebook-messenger"></i> Admin Support
                    </a>
                </div>
            </div>
        </div>

        <div class="footer">&copy; 2026 Nova Proxy. Auto-Distribution Web.</div>
    </div>

    <div id="toast" class="toast">Config successfully copied to clipboard!</div>

    <script>
        function copyConfig(configText) {
            navigator.clipboard.writeText(configText).then(() => {
                const toast = document.getElementById('toast');
                toast.classList.add('show');
                setTimeout(() => { toast.classList.remove('show'); }, 2000);
            }).catch(err => { console.error('Error: ', err); });
        }

        document.getElementById('searchInput').addEventListener('input', function(e) {
            const searchText = e.target.value.toLowerCase();
            const cards = document.querySelectorAll('.card');
            cards.forEach(card => {
                const title = card.querySelector('.node-title').textContent.toLowerCase();
                if(title.includes(searchText)) { card.style.display = 'block'; } else { card.style.display = 'none'; }
            });
        });
    </script>
</body>
</html>
`


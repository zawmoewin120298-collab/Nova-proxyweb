package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gookit/color"
)

func main() {
	// Railway Port ဖတ်ခြင်း
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ပင်မစာမျက်နှာအတွက် HTML Design Response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, htmlTemplate)
	})

	color.Green.Printf("Nova Proxy Web is starting beautifully on port %s...\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// Dashboard UI အတွက် HTML & CSS ပုံစံအလှ
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nova Proxy Web - Dashboard</title>
    <style>
        :root {
            --bg-color: #0f172a;
            --card-bg: #1e293b;
            --text-main: #f8fafc;
            --text-muted: #94a3b8;
            --accent-green: #10b981;
            --accent-blue: #3b82f6;
            --border-line: #334155;
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
        }

        body {
            background-color: var(--bg-color);
            color: var(--text-main);
            padding: 20px;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
        }

        .container {
            width: 100%;
            max-width: 900px;
            background: var(--card-bg);
            border-radius: 16px;
            box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);
            border: 1px solid var(--border-line);
            overflow: hidden;
        }

        /* Header Section */
        .header {
            background: linear-gradient(135deg, #1e3a8a, #3b82f6);
            padding: 30px;
            text-align: center;
            position: relative;
        }

        .header h1 {
            font-size: 24px;
            font-weight: 600;
            letter-spacing: 1px;
            margin-bottom: 8px;
        }

        .status-badge {
            display: inline-flex;
            align-items: center;
            background: rgba(16, 185, 129, 0.2);
            color: var(--accent-green);
            padding: 6px 16px;
            border-radius: 20px;
            font-size: 14px;
            font-weight: bold;
            border: 1px solid var(--accent-green);
        }

        .status-dot {
            width: 8px;
            height: 8px;
            background-color: var(--accent-green);
            border-radius: 50%;
            margin-right: 8px;
            animation: pulse 1.5s infinite;
        }

        /* Server Grid Section */
        .content {
            padding: 30px;
        }

        .section-title {
            font-size: 18px;
            color: var(--text-muted);
            margin-bottom: 20px;
            font-weight: 500;
        }

        .grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
            gap: 20px;
        }

        .card {
            background: rgba(15, 23, 42, 0.6);
            border: 1px solid var(--border-line);
            border-radius: 12px;
            padding: 20px;
            transition: transform 0.2s, border-color 0.2s;
        }

        .card:hover {
            transform: translateY(-2px);
            border-color: var(--accent-blue);
        }

        .card-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 15px;
        }

        .server-name {
            font-size: 16px;
            font-weight: 600;
        }

        .ping {
            font-size: 13px;
            color: var(--accent-green);
            background: rgba(16, 185, 129, 0.1);
            padding: 2px 8px;
            border-radius: 4px;
        }

        .card-body p {
            font-size: 14px;
            color: var(--text-muted);
            margin-bottom: 8px;
        }

        .card-body span {
            color: var(--text-main);
            font-weight: 500;
        }

        /* Footer */
        .footer {
            text-align: center;
            padding: 20px;
            color: var(--text-muted);
            font-size: 13px;
            border-top: 1px solid var(--border-line);
            background: rgba(0,0,0,0.1);
        }

        @keyframes pulse {
            0% { transform: scale(0.9); opacity: 0.6; }
            50% { transform: scale(1.1); opacity: 1; }
            100% { transform: scale(0.9); opacity: 0.6; }
        }
    </style>
</head>
<body>

    <div class="container">
        <div class="header">
            <h1>NOVA PROXY WEB</h1>
            <div class="status-badge">
                <div class="status-dot"></div>
                System Core Online
            </div>
        </div>

        <div class="content">
            <div class="section-title">Available Premium Nodes</div>
            
            <div class="grid">
                <div class="card">
                    <div class="card-header">
                        <div class="server-name">🇹🇭 Thailand Premium 01</div>
                        <div class="ping">15 ms</div>
                    </div>
                    <div class="card-body">
                        <p>Protocol: <span>VLESS / Trojan</span></p>
                        <p>Status: <span>Active (Normal)</span></p>
                    </div>
                </div>

                <div class="card">
                    <div class="card-header">
                        <div class="server-name">🇸🇬 Singapore High-Speed</div>
                        <div class="ping">28 ms</div>
                    </div>
                    <div class="card-body">
                        <p>Protocol: <span>VLESS / VMess</span></p>
                        <p>Status: <span>Active (Normal)</span></p>
                    </div>
                </div>

                <div class="card">
                    <div class="card-header">
                        <div class="server-name">🇯🇵 Japan Streaming Node</div>
                        <div class="ping">65 ms</div>
                    </div>
                    <div class="card-body">
                        <p>Protocol: <span>VLESS / ShadowDNS</span></p>
                        <p>Status: <span>Active (Normal)</span></p>
                    </div>
                </div>
            </div>
        </div>

        <div class="footer">
            &copy; 2026 Nova Proxy Web. Admin Dashboard System.
        </div>
    </div>

</body>
</html>
`


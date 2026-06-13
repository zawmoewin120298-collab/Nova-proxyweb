package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gookit/color"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, htmlTemplate)
	})

	color.Green.Printf("Nova Premium Dashboard is running on port %s...\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Nova Proxy Web - Premium VIP Dashboard</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <style>
        :root {
            --bg-color: #0b0f19;
            --card-bg: #151f32;
            --text-main: #f8fafc;
            --text-muted: #64748b;
            --accent-green: #10b981;
            --accent-blue: #3b82f6;
            --accent-purple: #8b5cf6;
            --border-line: #22314d;
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
            font-family: 'Segoe UI', system-ui, sans-serif;
        }

        body {
            background-color: var(--bg-color);
            color: var(--text-main);
            padding: 15px;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
        }

        .container {
            width: 100%;
            max-width: 600px; /* Mobile Friendly Width */
            background: var(--card-bg);
            border-radius: 20px;
            box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
            border: 1px solid var(--border-line);
            overflow: hidden;
        }

        /* Modern Gradient Header */
        .header {
            background: linear-gradient(135deg, #1e40af, #6366f1);
            padding: 35px 20px;
            text-align: center;
            position: relative;
        }

        .header h1 {
            font-size: 26px;
            font-weight: 700;
            letter-spacing: 1.5px;
            margin-bottom: 10px;
            text-shadow: 0 2px 4px rgba(0,0,0,0.3);
        }

        .status-badge {
            display: inline-flex;
            align-items: center;
            background: rgba(16, 185, 129, 0.15);
            color: var(--accent-green);
            padding: 6px 14px;
            border-radius: 30px;
            font-size: 13px;
            font-weight: 600;
            border: 1px solid rgba(16, 185, 129, 0.3);
        }

        .status-dot {
            width: 8px;
            height: 8px;
            background-color: var(--accent-green);
            border-radius: 50%;
            margin-right: 8px;
            animation: pulse 1.8s infinite;
        }

        /* Interactive Controls */
        .content {
            padding: 20px;
        }

        .search-box {
            position: relative;
            margin-bottom: 20px;
        }

        .search-box i {
            position: absolute;
            left: 15px;
            top: 50%;
            transform: translateY(-50%);
            color: var(--text-muted);
        }

        .search-input {
            width: 100%;
            padding: 12px 15px 12px 45px;
            background: rgba(11, 15, 25, 0.8);
            border: 1px solid var(--border-line);
            border-radius: 12px;
            color: var(--text-main);
            font-size: 14px;
            transition: all 0.3s;
        }

        .search-input:focus {
            outline: none;
            border-color: var(--accent-blue);
            box-shadow: 0 0 8px rgba(59, 130, 246, 0.3);
        }

        /* Node List & Cards */
        .node-list {
            display: flex;
            flex-direction: column;
            gap: 15px;
        }

        .card {
            background: rgba(11, 15, 25, 0.4);
            border: 1px solid var(--border-line);
            border-radius: 14px;
            padding: 18px;
            position: relative;
            transition: all 0.3s ease;
        }

        .card:hover {
            border-color: var(--accent-blue);
            background: rgba(11, 15, 25, 0.6);
        }

        .card-header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 12px;
        }

        .node-title {
            font-size: 16px;
            font-weight: 600;
            display: flex;
            align-items: center;
            gap: 10px;
        }

        .ping-tag {
            font-size: 12px;
            color: var(--accent-green);
            background: rgba(16, 185, 129, 0.1);
            padding: 3px 10px;
            border-radius: 6px;
            font-weight: 600;
        }

        .card-info {
            font-size: 13px;
            color: #94a3b8;
            margin-bottom: 15px;
            line-height: 1.6;
        }

        .card-info span {
            color: var(--text-main);
            font-weight: 500;
        }

        /* Copy Button Actions */
        .btn-copy {
            width: 100%;
            background: linear-gradient(135deg, #2563eb, #1d4ed8);
            color: white;
            border: none;
            padding: 10px;
            border-radius: 8px;
            font-size: 13px;
            font-weight: 600;
            cursor: pointer;
            display: flex;
            justify-content: center;
            align-items: center;
            gap: 8px;
            transition: background 0.2s;
        }

        .btn-copy:active {
            transform: scale(0.99);
        }

        /* Toast Notification */
        .toast {
            position: fixed;
            bottom: 20px;
            left: 50%;
            transform: translateX(-50%) translateY(100px);
            background: var(--accent-green);
            color: white;
            padding: 10px 24px;
            border-radius: 30px;
            font-size: 14px;
            font-weight: 600;
            box-shadow: 0 4px 12px rgba(0,0,0,0.3);
            transition: transform 0.3s ease;
            z-index: 1000;
        }

        .toast.show {
            transform: translateX(-50%) translateY(0);
        }

        /* Community Support Buttons */
        .support-section {
            margin-top: 25px;
            padding-top: 20px;
            border-top: 1px solid var(--border-line);
            text-align: center;
        }

        .support-title {
            font-size: 13px;
            color: var(--text-muted);
            margin-bottom: 12px;
            text-transform: uppercase;
            letter-spacing: 0.5px;
        }

        .support-buttons {
            display: flex;
            gap: 10px;
            justify-content: center;
        }

        .btn-social {
            padding: 8px 16px;
            border-radius: 8px;
            font-size: 12px;
            font-weight: 600;
            text-decoration: none;
            color: white;
            display: flex;
            align-items: center;
            gap: 6px;
        }

        .btn-messenger { background: #006aff; }
        .btn-telegram { background: #0088cc; }

        .footer {
            text-align: center;
            padding: 15px;
            color: var(--text-muted);
            font-size: 12px;
            background: rgba(0,0,0,0.15);
        }

        @keyframes pulse {
            0% { transform: scale(0.95); opacity: 0.5; }
            50% { transform: scale(1.15); opacity: 1; }
            100% { transform: scale(0.95); opacity: 0.5; }
        }
    </style>
</head>
<body>

    <div class="container">
        <div class="header">
            <h1>NOVA PROXY WEB</h1>
            <div class="status-badge">
                <div class="status-dot"></div>
                Premium Core Engine V2
            </div>
        </div>

        <div class="content">
            <div class="search-box">
                <i class="fa-solid fa-magnifying-glass"></i>
                <input type="text" id="searchInput" class="search-input" placeholder="Search location or protocol...">
            </div>

            <div class="node-list" id="nodeList">
                
                <div class="card">
                    <div class="card-header">
                        <div class="node-title">🇹🇭 Thailand Online 20ms</div>
                        <div class="ping-tag">15 ms</div>
                    </div>
                    <div class="card-info">
                        <div>Protocol: <span>VLESS / Trojan</span></div>
                        <div>Transport: <span>WebSocket (WS)</span></div>
                    </div>
                    <button class="btn-copy" onclick="copyConfig('vless://example-thailand-premium-node-config-string-data')">
                        <i class="fa-regular fa-copy"></i> Copy Configuration
                    </button>
                </div>

                <div class="card">
                    <div class="card-header">
                        <div class="node-title">🇸🇬 Singapore High-Speed</div>
                        <div class="ping-tag">28 ms</div>
                    </div>
                    <div class="card-info">
                        <div>Protocol: <span>VLESS / VMess</span></div>
                        <div>Transport: <span>gRPC / WS</span></div>
                    </div>
                    <button class="btn-copy" onclick="copyConfig('vless://example-singapore-premium-node-config-string-data')">
                        <i class="fa-regular fa-copy"></i> Copy Configuration
                    </button>
                </div>

                <div class="card">
                    <div class="card-header">
                        <div class="node-title">🇯🇵 Japan Video Streaming</div>
                        <div class="ping-tag">65 ms</div>
                    </div>
                    <div class="card-info">
                        <div>Protocol: <span>VLESS / SlowDNS</span></div>
                        <div>Transport: <span>Direct / DNS</span></div>
                    </div>
                    <button class="btn-copy" onclick="copyConfig('vless://example-japan-premium-node-config-string-data')">
                        <i class="fa-regular fa-copy"></i> Copy Configuration
                    </button>
                </div>

            </div>

            <div class="support-section">
                <div class="support-title">Contact Support & Purchase VIP</div>
                <div class="support-buttons">
                    <a href="https://m.me/your-facebook-profile" target="_blank" class="btn-social btn-messenger">
                        <i class="fa-brands fa-facebook-messenger"></i> Admin Messenger
                    </a>
                    <a href="https://t.me/your-telegram-channel" target="_blank" class="btn-social btn-telegram">
                        <i class="fa-brands fa-telegram"></i> Telegram Group
                    </a>
                </div>
            </div>
        </div>

        <div class="footer">
            &copy; 2026 Nova Proxy. Designed beautifully for Admin Community.
        </div>
    </div>

    <div id="toast" class="toast">Configuration copied to clipboard!</div>

    <script>
        // Copy စနစ် အလုပ်လုပ်ပုံ 
        function copyConfig(configText) {
            navigator.clipboard.writeText(configText).then(() => {
                const toast = document.getElementById('toast');
                toast.classList.add('show');
                setTimeout(() => {
                    toast.classList.remove('show');
                }, 2000);
            }).catch(err => {
                console.error('Could not copy text: ', err);
            });
        }

        // Live Search Filter စနစ် အလုပ်လုပ်ပုံ
        document.getElementById('searchInput').addEventListener('input', function(e) {
            const searchText = e.target.value.toLowerCase();
            const cards = document.querySelectorAll('.card');
            
            cards.forEach(card => {
                const title = card.querySelector('.node-title').textContent.toLowerCase();
                const info = card.querySelector('.card-info').textContent.toLowerCase();
                
                if(title.includes(searchText) || info.includes(searchText)) {
                    card.style.display = 'block';
                } else {
                    card.style.display = 'none';
                }
            });
        });
    </script>
</body>
</html>
`


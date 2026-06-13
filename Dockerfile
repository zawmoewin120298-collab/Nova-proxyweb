# Step 1: Build stage
FROM golang:1.22-alpine AS builder

# Build လုပ်ဖို့ လိုအပ်တဲ့ git ရော၊ final ဆာဗာအတွက်ပါ သုံးမည့် ca-certificates ကိုပါ တစ်ခါတည်းသွင်းမယ်
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Source code အားလုံးကို အရင် copy ကူးယူခြင်း
COPY . .

# dependency များကို အလိုအလျောက် ညှိပေးပြီး static binary အဖြစ် build လုပ်ခြင်း
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# ==========================================
# Step 2: Final minimal stage (အပေါ့ပါးဆုံးနှင့် အတည်ငြိမ်ဆုံးပုံစံ)
# ==========================================
FROM alpine:latest

WORKDIR /root/

# Builder ထဲကနေ binary ရော၊ သေချာပေါက်အလုပ်လုပ်မည့် certificates တွေကိုပါ တစ်ခါတည်း ဆွဲယူမယ်
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main .

# Proxy app အတွက် လိုအပ်သော အခြား file/folder များရှိလျှင် (ဥပမာ config.json သို့မဟုတ် .env) 
# ၎င်းတို့ကိုပါ တစ်ခါတည်း ကူးယူသွားပါမည်
COPY --from=builder /app /root/

# Railway အတွက် internal port ဖွင့်ပေးခြင်း
EXPOSE 8080

CMD ["./main"]


# Client-Server Dosya Transferi

Bu proje, Go diliyle bir istemci (client) ve sunucu (server) arasında dosya transferi yapmayı sağlar.

## Proje Dosyaları
- **client.go**: Dosyayı sunucuya gönderen istemci uygulaması.
- **server.go**: Sunucuyu başlatan ve dosyayı alan uygulama.
- **cn.txt**: Gönderilecek örnek dosya.

## Gereksinimler
- Go yüklü olmalı.
- Aynı ağda bulunan bir sunucuya bağlanmalısınız.

## Kurulum ve Kullanım

### 1. Sunucu Başlatma
1. `server.go` dosyasının bulunduğu dizine gidin.
2. Sunucuyu başlatmak için terminalde:

   ```bash
   go run server.go

Sunucu 9000 portunda çalışmaya başlayacak.

### 2. İstemci Başlatma
1-client.go dosyasındaki SERVER IP kısmını sunucunun IP adresiyle güncelleyin:

SERVER_IP := "***.***.*.***"  // Sunucunun IP adresini buraya girin("*" bir sayı karakterine karşılık gelir)

2-İstemciyi çalıştırın:

go run client.go

3-Dosya yolunu girin (örneğin cn.txt)

### 3. Dosya Transferi
İstemci, dosyayı sunucuya gönderecek ve sunucu dosyayı alacak.








# --- Tahap 1: Build ---
# Gunakan image Go resmi sebagai basis build. Pilih versi yang sesuai.
# Alpine lebih kecil, tapi kadang perlu penyesuaian CGO/linking.
FROM golang:1.21-alpine AS builder

# Set environment variables untuk build
ENV CGO_ENABLED=0   # Disable CGO for static linking (penting untuk Alpine/Scratch)
ENV GOOS=linux      # Target OS Linux
ENV GOARCH=amd64    # Target architecture amd64

# Set working directory di dalam container build
WORKDIR /app

# Copy file go.mod dan go.sum terlebih dahulu
# Ini memanfaatkan Docker layer caching. Jika file ini tidak berubah,
# dependensi tidak akan di-download ulang pada build berikutnya.
COPY go.mod go.sum ./

# Download dependensi
RUN go mod download
RUN go mod verify # Verifikasi integritas dependensi

# Copy seluruh source code aplikasi ke working directory
COPY . .

# Build aplikasi Go
# -ldflags="-s -w": Hapus debug symbols & DWARF info -> binary lebih kecil
# -o /myapp: Nama dan path output binary di dalam container build
RUN go build -ldflags="-s -w" -o /myapp ./cmd/api/main.go
# Ganti ./cmd/api/main.go jika struktur proyek Anda berbeda (misal hanya main.go di root)


# --- Tahap 2: Run ---
# Gunakan image dasar yang sangat kecil untuk runtime.
# 'alpine' adalah pilihan bagus, punya shell dan package manager dasar.
# 'scratch' adalah image paling kecil (kosong), tapi tanpa shell/tools,
#   memerlukan binary yang benar-benar statically linked.
FROM alpine:latest
# FROM scratch

# Set working directory di container runtime
WORKDIR /app

# (Opsional tapi sering diperlukan) Install sertifikat CA
# Dibutuhkan jika aplikasi Anda perlu membuat koneksi HTTPS keluar (misal ke API lain, DB over SSL)
# 'scratch' tidak bisa menjalankan ini. Jika pakai scratch, sertifikat harus di-copy manual.
RUN apk --no-cache add ca-certificates

# Copy binary yang sudah di-build dari tahap 'builder'
COPY --from=builder /myapp /myapp

# (Opsional) Copy file statis, template, atau file konfigurasi
# Jika tidak di-embed ke dalam binary.
# Pastikan path source (--from=builder /app/...) dan destination (./...) benar.
# COPY --from=builder /app/public ./public
# COPY --from=builder /app/views ./views
# COPY --from=builder /app/config.yaml ./config.yaml

# Expose port yang didengarkan oleh aplikasi Fiber di dalam container
# Ini hanya metadata, Anda masih perlu mapping port saat 'docker run' (-p)
EXPOSE 3000

# Set environment variable default untuk container runtime
# Ini bisa di-override saat menjalankan container dengan 'docker run -e'
ENV APP_ENV=production
ENV LISTEN_ADDR=:3000
# ENV DATABASE_URL= # Sebaiknya diset saat runtime
# ENV JWT_SECRET_KEY= # Sebaiknya diset saat runtime

# User untuk menjalankan aplikasi (opsional tapi lebih aman daripada root)
# RUN addgroup -S appgroup && adduser -S appuser -G appgroup
# USER appuser

# Perintah untuk menjalankan aplikasi saat container dimulai
# ENTRYPOINT memastikan binary adalah proses utama container.
ENTRYPOINT ["/myapp"]
# CMD bisa digunakan untuk memberikan argumen default ke ENTRYPOINT.
# Contoh: CMD ["--config", "./config.yaml"]

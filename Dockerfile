# Build image chính thức từ Golang
FROM golang:1.23

# Đặt thư mục làm việc
WORKDIR /app

# Tải module
COPY go.mod go.sum ./
RUN go mod download

# Copy toàn bộ source
COPY . .

# Tải script wait-for-it để đợi MySQL sẵn sàng
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

# Build ứng dụng
RUN go build -o app ./cmd/main.go

# Khởi động app sau khi DB sẵn sàng
CMD ["./wait-for-it.sh", "db:3306", "--", "./app"]

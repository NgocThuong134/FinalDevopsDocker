# Hướng Dẫn Tạo Docker Compose Cho Dự Án Golang và React

Dự án này hướng dẫn cách tạo một server bằng Golang sử dụng Gin và MongoDB, cùng với một client bằng React, và cấu hình Docker Compose để chạy cả hai ứng dụng.

## Nội Dung

1. [Tạo Server bằng Golang](#tạo-server-bằng-golang)
2. [Viết Dockerfile cho Server](#viết-dockerfile-cho-server)
3. [Tạo Client bằng React](#tạo-client-bằng-react)
4. [Viết Dockerfile cho Client](#viết-dockerfile-cho-client)
5. [Viết file docker-compose.yml](#viết-file-docker-composeyml)
6. [Hướng Dẫn Clone và Chạy Dự Án](#hướng-dẫn-clone-và-chạy-dự-án)

## Tạo Server bằng Golang

1. Tạo một tệp `main.go` cho server.
2. Cài đặt Gin và MongoDB.
3. Cấu hình CORS cho server và kết nối với MongoDB.
4. Định nghĩa các route cần thiết, ví dụ như route `/api`.

## Viết Dockerfile cho Server

1. Tạo một tệp `Dockerfile` trong thư mục server.
2. Sử dụng hình ảnh `golang` làm base image.
3. Sao chép mã nguồn vào container và cài đặt các phụ thuộc.
4. Xây dựng ứng dụng và chỉ định cổng mà server sẽ chạy.

## Tạo Client bằng React

1. Sử dụng Create React App để khởi tạo một ứng dụng React mới.
2. Tạo các component và cấu hình ứng dụng như mong muốn.

## Viết Dockerfile cho Client

1. Tạo một tệp `Dockerfile` trong thư mục client.
2. Sử dụng hình ảnh `node` làm base image.
3. Sao chép `package.json` và cài đặt các phụ thuộc.
4. Xây dựng ứng dụng và chỉ định cổng mà client sẽ chạy.

## Viết file docker-compose.yml

1. Tạo một tệp `docker-compose.yml` trong thư mục gốc của dự án.
2. Định nghĩa các service cho MongoDB, server và client.
3. Chỉ định các cổng và phụ thuộc giữa các service.

## Hướng Dẫn Clone và Chạy Dự Án

1. **Clone dự án từ GitHub:**

   ```bash
   git clone https://github.com/NgocThuong134/FinalDevopsDocker.git
   cd FinalDevopsDocker
2. **Chạy Docker Compose:**

   Sử dụng lệnh sau để xây dựng và chạy tất cả các container được định nghĩa trong `docker-compose.yml`:

   ```bash
   docker-compose up --build

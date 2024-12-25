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

Để tạo server bằng Golang với Gin và MongoDB, bạn có thể làm theo các bước sau:

1. **Tạo một tệp `main.go` cho server.**
   - Tạo một thư mục mới cho dự án server của bạn, ví dụ: `server`.
   - Trong thư mục này, tạo một tệp có tên `main.go`. Đây sẽ là tệp chính của ứng dụng.

2. **Cài đặt Gin và MongoDB.**
   - Mở terminal và điều hướng đến thư mục dự án server của bạn.
   - Khởi tạo một module Go mới:

     ```bash
     go mod init server
     ```

   - Cài đặt các gói cần thiết cho Gin và MongoDB:

     ```bash
     go get -u github.com/gin-gonic/gin
     go get -u go.mongodb.org/mongo-driver/mongo
     go get -u github.com/gin-contrib/cors
     ```

3. **Cấu hình CORS cho server và kết nối với MongoDB.**
   - Trong tệp `main.go`, bắt đầu bằng việc nhập các gói cần thiết:

     ```go
     package main

     import (
         "context"
         "github.com/gin-gonic/gin"
         "github.com/gin-contrib/cors"
         "go.mongodb.org/mongo-driver/mongo"
         "go.mongodb.org/mongo-driver/mongo/options"
         "net/http"
     )
     ```

   - Tiếp theo, trong hàm `main`, cấu hình Gin và CORS:

     ```go
     func main() {
         router := gin.Default()

         // Cấu hình CORS
         router.Use(cors.Default())
     ```

   - Kết nối đến MongoDB:

     ```go
         // Kết nối MongoDB
         clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
         client, err := mongo.Connect(context.TODO(), clientOptions)
         if err != nil {
             panic(err)
         }
         defer client.Disconnect(context.TODO())
     ```

4. **Định nghĩa các route cần thiết, ví dụ như route `/hello`.**
   - Tạo một route đơn giản để kiểm tra server hoạt động:

     ```go
         router.GET("/hello", func(c *gin.Context) {
             c.JSON(http.StatusOK, gin.H{"message": "Hello from Go server!"})
         })
     ```

   - Cuối cùng, khởi động server trên cổng 6080:

     ```go
         router.Run(":6080")
     }
     ```

### Kết Luận
Sau khi hoàn tất các bước trên, bạn sẽ có một server cơ bản hoạt động với Gin và MongoDB. Bạn có thể mở terminal và chạy lệnh `go run main.go` để khởi động server. Khi server đang chạy, bạn có thể truy cập [http://localhost:6080/hello](http://localhost:6080/hello) để kiểm tra phản hồi từ server.

## Viết Dockerfile cho Server

1. Tạo một tệp `Dockerfile` trong thư mục server.
2. Sử dụng hình ảnh `golang` làm base image.
3. Sao chép mã nguồn vào container và cài đặt các phụ thuộc.
4. Xây dựng ứng dụng và chỉ định cổng mà server sẽ chạy.

## Tạo Client bằng React

Để tạo client bằng React, bạn có thể làm theo các bước sau:

1. **Khởi tạo ứng dụng React mới.**
   - Mở terminal và điều hướng đến thư mục mà bạn muốn tạo ứng dụng.
   - Sử dụng Create React App để khởi tạo một ứng dụng mới bằng lệnh:

     ```bash
     npx create-react-app client
     ```

   - Lệnh này sẽ tạo ra một thư mục `client` với cấu trúc cơ bản cho một ứng dụng React.

2. **Tạo các component và cấu hình ứng dụng.**
   - **Chỉnh sửa cấu trúc ứng dụng**: Mở thư mục `client/src` và bắt đầu tạo các component mới trong thư mục `components`. 
   - **Tạo giao diện**: Sử dụng các component để xây dựng giao diện người dùng. Bạn có thể tạo các component như Header, Footer, và các phần khác của ứng dụng.
   - **Kết nối với server**: Sử dụng `fetch` hoặc thư viện như Axios để gửi yêu cầu đến server Golang. Đảm bảo rằng bạn xử lý dữ liệu nhận được và cập nhật trạng thái của ứng dụng.

3. **Chạy ứng dụng.**
   - Sau khi cấu hình xong, bạn có thể chạy ứng dụng bằng lệnh:

     ```bash
     npm start
     ```

   - Ứng dụng sẽ tự động mở trong trình duyệt tại [http://localhost:3000], nơi bạn có thể thấy giao diện của ứng dụng.

### Kết Luận
Sau khi hoàn tất các bước trên, bạn sẽ có một client React cơ bản có thể giao tiếp với server Golang. Bạn có thể tiếp tục mở rộng và tùy chỉnh ứng dụng theo nhu cầu của mình.

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
   ```bash
   cd FinalDevopsDocker
2. **Chạy Docker Compose:**

   Sử dụng lệnh sau để xây dựng và chạy tất cả các container được định nghĩa trong `docker-compose.yml`:

   ```bash
   docker-compose up --build

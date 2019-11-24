# Sao lưu dữ liệu
Hiện nay để đảm bảo an toàn cho dữ liệu, tốt nhất nên đưa dữ liệu lên lưu trữ trên các dịch vụ trên đám mây (cloud). Các dịch vụ này sẽ đảm bảo dữ liệu của chúng ta không mất bởi chúng sẽ được sao lưu ở nhiều nơi khác nhau. Khi không có cloud hoặc muốn tạo một bản sao an toàn thì làm thế nào? Có rất nhiều công cụ hỗ trợ việc đồng bộ lưu trữ các file. Ở đây tôi giới thiệu với các bạn Rsync (Remote Sync), một dịch vụ đồng bộ file nổi tiếng trên môi trường Unix và Linux. Trên Windows có thể dùng Cwrsync.

Tính năng nổi bật của [Rsync](https://rsync.samba.org/) là:
- Rsync hỗ trợ sao chép giữ nguyên thông số của files/folder như liên kết, quyền, thời gian, sở hữu và nhóm.
- Rsync chỉ trao đổi những dữ liệu thay đổi nên thời gian đồng bộ nhanh.
- Rsync tiết kiệm băng thông do sử dụng phương pháp nén và giải nén khi đồng bộ.
- Rsync không yêu cầu quyền super-user.

Chúng ta có thể lấy file cài đặt từ trang download chính của Rsync, hiện có bản rpm cho Linux và cwRsync cho Windows. Các distro Linux có thể dùng câu lệnh cài đặt vì rsync có sẵn trong các bản phân phối Linux:
```bash
sudo apt-get install rsync    // Dùng cho debian, ubuntu, và tương tự
yum install rsync             // Dùng cho Fedora, Redhat, CentOS, và tương tự
rpm -ivh rsync                // Dùng cho Fedora, Redhat, CentOS, và tương tự
```

Ngoài ra chúng ta có thể cài đặt thông qua việc biên dịch mã nguồn, được cung cấp sẵn trên trang download.

Câu lệnh của Rsync:
rsync <tham số> <dữ liệu nguồn> <dữ liệu đích>
Các tham số phổ biến:
> -v: hiển thị trạng thái kết quả.

> -r: copy dữ liệu một cách đệ quy, nhưng không đảm bảo thông số của file và thư mục.

> -a: cho phép copy dữ liệu một cách đệ quy, đồng thời giữ nguyên được tất cả các thông số của thư mục và file.

> -z: nén dữ liệu khi đồng bộ, tiết kiệm băng thông tuy nhiên tốn thêm một chút thời gian.

> -h: xuất kết quả dễ đọc.

> --delete: xóa dữ liệu ở đích nếu bên nguồn không tồn tại dữ liệu đó.

> --exclude: loại trừ ra những dữ liệu không muốn truyền đi, nếu bạn cần loại ra nhiều file hoặc folder ở nhiều đường dẫn khác nhau thì mỗi cái bạn phải thêm --exclude tương ứng.

Tất cả các tham số có thể tham khảo [ở đây](https://download.samba.org/pub/rsync/rsync.html). 

Ví dụ đồng bộ từ thư mục /home/test/Documents/ebook/ đến /media/test/bkdisk/ebook ta thực hiện:
```bash
rsync -avzh /home/test/Documents/ebook /media/test/bkdisk/
```

Thường để sao lưu dữ liệu trên server chúng ta sẽ dùng rsync để đồng bộ về máy khác, lúc này ta có thể dùng câu lệnh như sau:
```bash
rsync -avz /home/test/Documents/ebook --password-file=/etc/rsyncd.secrets root@192.168.0.101:/media/test/bkdisk/
```

Tham số --password-file khai báo thông tin mật khẩu. Việc này giúp chúng ta không phải nhập mật khẩu mỗi lần chạy rsync. Tạo file chứa mật khẩu như sau:
```bash
echo '<tên truy cập>: <Mật khẩu>' > /etc/rsyncd.secrets
chmod 600 /etc/rsyncd.secrets
```

Thường đồng bộ sao lưu sẽ chạy tự động nhưng rsync không có chức năng này nên cần kết hợp với dịch vụ kiểu như crontab trên Linux/Mac OS hay Task Scheduler trên Windows. 

## Sao lưu cơ sở dữ liệu
Ở đây tôi bàn về sao lưu MySQL, cơ sở dữ liệu phổ phiến nhất cho các ứng dụng dịch vụ web vừa và nhỏ. Cách phổ biến để sao lưu cơ sở dữ liệu MySQL là sử dụng mysqldump để sao lưu ra file.

Câu lệnh đơn giản của nó như sau: 
```bash
mysqldump -u <tên truy cập> -p<mật khẩu> --<tùy chọn> <database cần sao lưu> > <tên file sao lưu>.sql
```

Ví dụ cần sao lưu nhiều database, ta dùng tùy chọn --databases, câu lệnh sẽ như sau:
```bash
mysqldump -u <tên truy cập> -p<mật khẩu> --databases <database 1> <database 2> > <tên file sao lưu>.sql
```

> Với --all-databases, mọi database sẽ được sao lưu. 

Lưu ý ở các câu lệnh trên là mật khẩu viết ngay sau -p, không có khoảng trắng. Nếu không có mật khẩu thì sau khi thực thi lệnh hệ thống yêu cầu nhập mật khẩu trước khi thực hiện sao lưu. 

Các tham số khác:
> --add-drop-table: Thêm câu lệnh DROP TABLE trước mỗi lệnh CREATE TABLE trong file sao lưu.

> --no-data: Chỉ sao lưu cấu trúc database, không sao lưu dữ liệu.

Để tạo file sao lưu nén chúng ta có thể dùng pipeline theo câu lệnh sau:
```bash
mysqldump -u <tên truy cập> -p<mật khẩu> --<tùy chọn> <database cần sao lưu> | gzip -9 > <tên file sao lưu>.sql.gz
```

Cách phục hồi lại như sau: 
```bash
mysql -u <tên truy cập> -p<mật khẩu> <database cần sao lưu> < <tên file sao lưu>.sql
```

Với database đã tồn tại, để phục hồi chúng ta dùng: 
```bash
mysqlimport -u <tên truy cập> -p<mật khẩu> <database cần sao lưu> <tên file sao lưu>.sql
```

Để phục hồi với file đã nén, chúng ta sử dụng câu lệnh: 
```bash
gunzip < [<tên file sao lưu>.sql.gz] | mysql -u <tên truy cập> -p<mật khẩu> <database cần sao lưu> 
```

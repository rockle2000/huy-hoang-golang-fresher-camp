**GoEx15: Trong trường hợp tạo cột đếm thì làm sao để update cột đó? Làm sao để API chính không bị block vì phải update số đếm?**
- Tạo thêm function trong storage để update mỗi khi insert hoặc delete 1 bản ghi.
- Lệnh update sử dụng gorm expression để tăng/ giảm giá trị cột đếm.
- Tạo thêm 1 store cho tầng business để gọi hàm tăng, giảm giá trị cột đếm.

**Làm sao để API chính không bị block vì phải update số đếm?**
- Đặt hàm update cột đếm vào 1 go routine khác và dùng cơ chế recover để nếu
xảy ra lỗi thì không bị crash chương trình.

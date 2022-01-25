**Ex4: Vì sao trong khoá học này các bạn được khuyên không nên dùng khoá ngoại (FK), điểm yếu của khoá ngoại là gì?**

- Không tách nhỏ được service.
- Khi dữ liệu lớn foreign key sẽ ảnh hưởng đến thời gian thực hiện insert/update/delete 
vì foreign key tạo sự ràng buộc đến các bảng khác nên khi insert/update thì phải kiểm tra
dữ liệu tương ứng có tồn tại ở bảng parent thì mới insert/update đc ở bảng con.
- Khi delete trên bảng parent thì sẽ phải kiểm tra xem có bao nhiêu bảng con
tham chiếu tới (bắt buộc delete từ bảng con rồi mới delete ở bảng parent hoặc delete on cascade thì sẽ delete hết).
- Chỉ nên sử dụng trong dịch vụ Enterprise khi hệ thống yêu cầu tính nhất quán dữ liệu cao.

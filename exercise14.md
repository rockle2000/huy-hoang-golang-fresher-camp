**GoEx14: Khi nào cần tạo các cột số đếm ngay trên table dữ liệu (VD: liked_count trên restaurants)?**
- Khi API yêu cầu tính toán, cần tổng hợp dữ liệu từ nhiều bảng chịu tải lớn.
- VD: Bảng restaurant_likes có nhiều bản ghi => việc truy xuất và lấy dữ liệu được tổng hợp, tính toán từ bảng này sẽ tốn nhiều thời gian hơn khi lấy danh sách restaurant.
=> Nên lưu cột số đếm (mang tính chất caching) ở bảng restaurant tên là liked_count sẽ giúp tăng tốc độ đọc dữ liệu từ database tránh việc phải tính toán, tổng hợp dữ liệu.

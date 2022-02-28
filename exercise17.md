**- Vì sao hệ thống lại cần pubsub và queue ?**
- Thay vì phải gửi các message trực tiếp đến cho nhau thì các modules chỉ cần gửi vào broker (trung gian).
- Giúp các modules không phải gọi trực tiếp lẫn nhau tránh việc gọi chồng chéo các modules gây phức tạp cho hệ thống.
- Queue có chức năng tương tự load balancing.

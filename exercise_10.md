**GOEX10** Vì sao không nên chứa file upload vào ngay chính bên trong service mà nên dùng Cloud ? Vì sao không chứa ảnh binary vào DB ? 

**1. Vì sao không nên chứa file upload vào ngay chính bên trong service mà nên dùng Cloud ?**
- Lưu file upload vào ngay bên trong service sẽ khiến service tốn dung lượng để lưu trữ, chứa các file vật lý.
- Dùng Cloud sẽ giảm chi phí cho các thiết bị lưu trữ vật lý.
- Dùng Cloud sẽ tăng cường bảo mật, giảm nguy cơ mất dữ liệu, có thể truy cập mọi lúc mọi nơi với nhiều loại thiết bị khác nhau và dễ mở rộng.

**2. Vì sao không chứa ảnh binary vào DB ?** 
- DB sẽ yêu cầu nhiều dung lượng để lưu trữ hơn và sẽ làm ảnh hưởng đến việc back-up, truy xuất dữ liệu.
- Lưu trữ binary data lớn trong DB sẽ ảnh hưởng đến hiệu năng của DB.
- Nếu có lỗi trong binary data sẽ khó phát hiện.


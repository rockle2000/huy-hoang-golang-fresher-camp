**Ex 5 Khoá chính (PK) trong table DB có công dụng gì? Vì sao dùng ID là số tự tăng? Khi nào một table dùng khoá chính trên nhiều cột?**

- Khóa chính PK trong bảng chứa giá trị duy nhất, không trùng nhau và không được là NULL, dùng để phân biệt từng bản ghi trong bảng ,khóa chính có thể bao gồm 1 hoặc nhiều cột trong bảng.
- Dùng Id tự động tăng giúp chúng ta không cần truyền giá trị cũng như check các Id trùng nhau khi thêm bản ghi mới.
- Khi cần liên kết 2 Entity có mối quan hệ (n-n) thì sẽ cần có bảng trung gian để liên kết, bảng đó sẽ có khóa chính trên nhiều cột 
dùng để liên kết 2 Entity trên và đảm bảo không có bản ghi nào có cặp khóa trùng nhau trong bảng.

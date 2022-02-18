**GoEx13: Nếu chúng ta có nhiều hơn 1 module, làm sao để giao tiếp với nhau.
Giả sử module "Restaurant" cần data số lượt like từ module "Like Restaurant" 
thì sẽ truy xuất như thế nào?**

- Để các module giao tiếp với nhau thì ta cần tạo thêm 1 storage trong business của 1 module
trong đó chứa hàm đã được định nghĩa ở storage của module còn lại và thực thi hàm này để lấy 
dữ liệu.
- VD: Tạo thêm storage trong business của module restaurant chứa hàm GetRestaurantLike
được định nghĩa ở tầng storage trong module restaurantlike và thực thi để lấy dữ liệu.

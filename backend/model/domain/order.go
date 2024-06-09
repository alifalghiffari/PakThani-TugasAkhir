package domain

// Order mendefinisikan sebuah pesanan
type Order struct {
	Id            int           // ID unik untuk pesanan
	UserId        int           // ID pengguna yang membuat pesanan
	TotalItems    int           // Jumlah total item dalam pesanan
	TotalPrice    int           // Harga total dari pesanan
	OrderStatus   OrderStatus   // Status pesanan (misalnya: pending, diproses, selesai, dll.)
	PaymentStatus PaymentStatus // Status pembayaran pesanan (misalnya: pending, berhasil, gagal, dll.)
	OrderItems    []OrderItems  // Item dalam pesanan
}

// OrderStatus mendefinisikan status pesanan
type OrderStatus string

const (
	Pending    OrderStatus = "pending"
	Processing OrderStatus = "processing"
	Shipped    OrderStatus = "shipped"
	Delivered  OrderStatus = "delivered"
	Cancelled  OrderStatus = "cancelled"
)

// PaymentStatus mendefinisikan status pembayaran pesanan
type PaymentStatus string

const (
	PaymentPending PaymentStatus = "pending"
	PaymentSuccess PaymentStatus = "success"
	PaymentFailed  PaymentStatus = "failed"
)

# Micro Payment Service
Aplikasi ini merupakan gateway untuk biller. Aplikasi ini mengirim request sesuai dengan code produk yang dikirim oleh user. Aplikasi ini menggunakan GRPC untuk berkomunikasi dengan micro service lainnya.

## GRPC Service
```proto3
syntax = "proto3";

package payment;

import "proto/google/type/datetime.proto";

option go_package = "github.com/fbriansyah/micro-payment-proto/protogen/go/payment";

message InquiryRequest {
    string user_id=1; [json_name="user_id"]
    string bill_number=2; [json_name="bill_number"]
    string product_code=3; [json_name="product_code"]
}

message InquiryResponse {
    string inq_id=1; [json_name="inq_id"]
    string bill_number=2; [json_name="bill_number"]
    string product_code=3; [json_name="product_code"]
    string name=4; [json_name="name"]
    double total_amount=5; [json_name="total_amount"]
}

message PaymentRequest {
    string user_id=1; [json_name="user_id"]
    string bill_number=2; [json_name="bill_number"]
    string product_code=3; [json_name="product_code"]
    double amount=4; [json_name="amount"]
}

message PaymentResponse {
    string bill_number=1; [json_name="bill_number"]
    string product_code=2; [json_name="product_code"]
    string name=3; [json_name="name"]
    double total_amount=4; [json_name="total_amount"]
    string refference_number=5; [json_name="refference_number"]
    google.type.DateTime transaction_datetime = 6; [json_name="transaction_datetime"]
}

service PaymentService {
    rpc Inquiry(InquiryRequest) returns InquiryResponse {}
    rpc Payment(PaymentRequest) returns PaymentResponse {}
}
```
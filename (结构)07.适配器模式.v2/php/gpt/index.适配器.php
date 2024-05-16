<?php
interface PaymentGateway {
    public function sendPayment($amount);
}
class Paypal {
    public function sendPaymentPaypal($amount) {
        // 发送支付到Paypal的逻辑
        echo "Paying via Paypal: {$amount}\n";
    }
}
class PaypalAdapter implements PaymentGateway {
    private $paypal;

    public function __construct(Paypal $paypal) {
        $this->paypal = $paypal;
    }

    public function sendPayment($amount) {
        // 使用适配的方法调用Paypal的支付处理
        $this->paypal->sendPaymentPaypal($amount);
    }
}
function processPayment(PaymentGateway $paymentGateway, $amount) {
    $paymentGateway->sendPayment($amount);
}

$paypal = new Paypal();
$paypalAdapter = new PaypalAdapter($paypal);

processPayment($paypalAdapter, 50); // 输出: Paying via Paypal: 50

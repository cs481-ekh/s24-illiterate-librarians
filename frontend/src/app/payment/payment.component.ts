import { Component } from '@angular/core';

@Component({
  selector: 'app-payment',
})
export class PaymentComponent {
  openNewTab() {
    window.open('/payment', 'https://commerce.cashnet.com/boisestateLiteracyLab');
  }
}
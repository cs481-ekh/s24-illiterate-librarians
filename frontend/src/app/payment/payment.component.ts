import { Component } from '@angular/core';

@Component({
  selector: 'app-payment',
  template: `
  <button (click)="openNewTab()">Pay Now</button>
  `,
  styleUrl: './payment.component.css'
})
export class PaymentComponent {
  openNewTab() {
    window.open('/payment', 'https://commerce.cashnet.com/boisestateLiteracyLab');
  }
}

/* import { Component } from '@angular/core';

@Component({
  selector: 'app-payment',
})
export class PaymentComponent {
  openNewTab() {
    window.open('/payment', 'https://commerce.cashnet.com/boisestateLiteracyLab');
  }
}
*/
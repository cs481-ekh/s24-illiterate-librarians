import { Component } from '@angular/core';

@Component({
  selector: 'app-payment',
  template: `
    <p>
      payment works!
    </p>
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
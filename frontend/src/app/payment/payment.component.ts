import { Component } from '@angular/core';

@Component({
  selector: 'app-payment',
  template: `
  <button mat-raised-button color="primary" (click)="openNewTab()">Pay Now</button>
  `,
  styleUrl: './payment.component.css'
})
export class PaymentComponent {
  openNewTab() {
    window.open('https://commerce.cashnet.com/boisestateLiteracyLab');
  }
}
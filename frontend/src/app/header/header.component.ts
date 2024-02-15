import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [CommonModule, RouterModule],
  template: `
    <header class="header">
      <a href="https://sdp.boisestate.edu/index.html">
        <img
          src="../assets/sdp-logo-3.png"
          alt="SDP Logo"
          width="100px"
          heigh="100px"
        />
      </a>
      <a [routerLink]="['/']">
        <span class="material-symbols-outlined"> menu </span>
      </a>

      <div id="links">
        <a [routerLink]="['/']">
          <div class="option">Home</div>
        </a>
        <a [routerLink]="['/dashboard']"><div class="option">Dashboard</div></a>
        <a [routerLink]="['/profile']"><div class="option">Profile</div></a>
        <a [routerLink]="['/login']"
          ><div class="option" id="login">Login or Register</div></a
        >
      </div>
    </header>
  `,
  styleUrl: './header.component.css',
})
export class HeaderComponent {
  title = 'Literacy Link';
}

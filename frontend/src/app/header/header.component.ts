import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [CommonModule],
  template: `
    <header class="header">
      <nav>
        <ul class="options">
          <li><a routerLink="/">Home</a></li>
          <li><a routerLink="/Dashboard">Dashboard</a></li>
          <li><a routerLink="/Profile">Profile</a></li>
        </ul>
      </nav>
    </header>
  `,
  styleUrl: './header.component.css',
})
export class HeaderComponent {}

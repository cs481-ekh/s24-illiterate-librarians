import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { HeaderComponent } from './header/header.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, CommonModule, HeaderComponent],
  template: `
    <main>
      <header>
        <app-header></app-header>
      </header>
    </main>
  `,

  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  title = 'frontend';
}

import { Component } from '@angular/core';
import { AuthService } from '../_services/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: 'header.component.html',
  styleUrl: './header.component.css',
})
export class HeaderComponent {
  title = 'Literacy Link';

  isLoggedIn = false;

  constructor(private authService: AuthService) {
    this.isLoggedIn = this.authService.isLoggedIn();
    console.log(this.isLoggedIn);
  }

  logout() {
    this.authService.logout();
    window.location.reload();
  }
}
import { Component } from '@angular/core';

@Component({
  selector: 'app-login',
  template: `
  <div class="login">
    <div class="login-container">
      <div class="login-title">
        <h1>Login</h1>
      </div>
      <div class="login-form">
        <form>
          <mat-form-field class="input">
            <input matInput placeholder="Username">
          </mat-form-field>
          <mat-form-field class="input">
            <input matInput placeholder="Password" type="password">
          </mat-form-field>
          <div class="login-button">

          <a routerLink="/register">
            <button mat-raised-button >Create Account</button>
          </a>
            <button mat-raised-button color="primary" type="submit">Login</button>
            <!-- Option to register -->
          </div>
        </form>
      </div>
    </div>
  </div>
  `,
  styleUrl: './login.component.css'
})
export class LoginComponent {

}

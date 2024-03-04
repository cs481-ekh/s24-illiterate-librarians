import { Component, inject } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { UserService } from '../user.service';

@Component({
  selector: 'app-login',
  template: `
    <div class="login">
      <div class="login-container">
        <div class="login-title">
          <h1>Login</h1>
        </div>

        <div class="login-form">
          <form [formGroup]="loginForm" (submit)="submit()">
            <mat-form-field class="input">
              <input
                matInput
                placeholder="Username"
                type="text"
                formControlName="username"
              />
            </mat-form-field>
            <div class="error-message">
            @if(username?.invalid && submitted){
              Username is required
            }
            </div>
            <mat-form-field class="input">
              <input
                matInput
                placeholder="Password"
                type="password"
                formControlName="password"
              />
            </mat-form-field>

            <div class="error-message">
            @if(password?.invalid && submitted){
              Password is required
            
            }
            </div>
            <div class="login-button">
              <a routerLink="/register">
                <button mat-raised-button>Create Account</button>
              </a>
              <button mat-raised-button color="primary" type="submit">
                Login
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  `,
  styleUrl: './login.component.css',
})
export class LoginComponent {
  loginForm = new FormGroup({
    username: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required),
  });

  submitted = false;

  userService: UserService = inject(UserService);

  get username() {
    return this.loginForm.get('username');
  }

  get password() {
    return this.loginForm.get('password');
  }

  async submit(): Promise<void>{
    this.submitted = true;
    if (this.loginForm.valid) {
      if (this.username && this.password) {
        const username = this.username.value;
        const password = this.password.value;
        if (username && password) {
          const response = await this.userService.login(username, password);
          console.log(response);
        }
      }
    }
  }
}

import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-register',
  template: `
    <div class="register">
      <div class="form-container">

      <div class="title">
        <h1>Register</h1>
      </div>
      <mat-stepper [linear]="false" #stepper class="stepper">
      <div class = "forms">
          <mat-step [stepControl]="nameForm">
            <form [formGroup]="nameForm" class="name">
              <ng-template matStepLabel
                >Name and Username</ng-template
              >
              <mat-form-field>
                <input
                  matInput
                  placeholder="Username"
                  formControlName="username"
                />
              </mat-form-field>
              <mat-form-field >
                <input
                  matInput
                  placeholder="First name"
                  formControlName="first_name"
                />
              </mat-form-field>
              <mat-form-field>
                <input
                  matInput
                  placeholder="Last name"
                  formControlName="last_name"
                />
              </mat-form-field>
              <button mat-button matStepperNext>Next</button>
            </form>
          </mat-step>
          <mat-step [stepControl]="locationForm" label="Fill out your address">
            <form [formGroup]="locationForm" class="addr">
              <ng-template matStepLabel>Address</ng-template>
              <mat-form-field>
                <input
                  matInput
                  placeholder="Address"
                  formControlName="address"
                />
              </mat-form-field>
              <mat-form-field>
                <input matInput placeholder="City" formControlName="city" />
              </mat-form-field>
              <mat-form-field>
                <input matInput placeholder="State" formControlName="state" />
              </mat-form-field>
              <mat-form-field>
                <input matInput placeholder="Zip" formControlName="zip" />
              </mat-form-field>
              <button mat-button matStepperNext>Next</button>
            </form>
          </mat-step>
          <mat-step
            [stepControl]="contactForm"
            label="Fill out your contact information"
          >
            <form [formGroup]="contactForm" class="contact">
              <ng-template matStepLabel>Contact Information</ng-template>
              <mat-form-field>
                <input matInput placeholder="Phone" formControlName="phone" />
              </mat-form-field>
              <mat-form-field>
                <input
                  matInput
                  placeholder="Method of contact"
                  formControlName="method_of_contact"
                />
              </mat-form-field>
              <button mat-button matStepperNext>Next</button>
            </form>
          </mat-step>
          <mat-step [stepControl]="passwordForm" label="Set a password">
            <form [formGroup]="passwordForm" class="password">
              <ng-template matStepLabel>Set a password</ng-template>
              <mat-form-field>
                <input
                  matInput
                  placeholder="Password"
                  formControlName="password"
                />
              </mat-form-field>
              <mat-form-field>
                <input
                  matInput
                  placeholder="Confirm password"
                  formControlName="confirmPassword"
                />
              </mat-form-field>
              <button mat-button matStepperNext>Submit</button>
            </form>
          </mat-step>
          <mat-step label="Done">
            <ng-template matStepLabel>Done</ng-template>
            You are now registered!
            <button mat-button matStepperPrevious>Return to Login</button>
          </mat-step>
        </div>
      </mat-stepper>
      </div>
    </div>
  `,
  styleUrl: './register.component.css',
})
export class RegisterComponent {

  nameForm = new FormGroup({
    first_name: new FormControl('', Validators.required),
    last_name: new FormControl('', Validators.required),
    username: new FormControl('', Validators.required),
  });

  locationForm = new FormGroup({
    address: new FormControl('', Validators.required),
    city: new FormControl('', Validators.required),
    state: new FormControl('', Validators.required),
    zip: new FormControl('', Validators.required),
  });

  contactForm = new FormGroup({
    phone: new FormControl('', Validators.required),
    method_of_contact: new FormControl('', Validators.required),
  });

  passwordForm = new FormGroup({
    password: new FormControl('', Validators.required),
    confirmPassword: new FormControl('', Validators.required),
  });
}

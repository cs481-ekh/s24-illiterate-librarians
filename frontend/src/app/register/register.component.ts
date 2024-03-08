import { Component } from '@angular/core';
import {
  FormControl,
  FormGroup,
  Validators,
  AbstractControl,
  ValidatorFn,
  ValidationErrors,
} from '@angular/forms';

export function phoneNumberValidator(): ValidatorFn {
  return (control: AbstractControl): { [key: string]: any } | null => {
    const rawNumbers = control.value.replace(/\D/g, ''); // Extract only digits

    // Check if the raw numbers meet your condition (e.g., 10 digits)
    if (rawNumbers.length === 10) {
      return null; // Validation passed
    } else {
      return { phoneNumberInvalid: true }; // Validation failed
    }
  };
}

@Component({
  selector: 'app-register',
  template: `
    <div class="register">
      <div class="form-container">
        <div class="title">
          <h1>Register</h1>
        </div>
        <mat-stepper [linear]="true" #stepper class="stepper">
          <div class="forms">
            <mat-step [stepControl]="nameForm">
              <form [formGroup]="nameForm" class="name">
                <ng-template matStepLabel>Name and Username</ng-template>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="Username"
                    formControlName="username"
                  />
                </mat-form-field>
                <div
                  *ngIf="isStep1Valid && username?.errors?.['required']"
                  class="error"
                >
                  Username is required
                </div>
                <div
                  *ngIf="isStep1Valid && username?.errors?.['minlength']"
                  class="error"
                >
                  Username must be at least 2 characters
                </div>
                <div
                  *ngIf="isStep1Valid && username?.errors?.['maxlength']"
                  class="error"
                >
                  Username must be at most 30 characters
                </div>
                <div
                  *ngIf="isStep1Valid && username?.errors?.['pattern']"
                  class="error"
                >
                  Invalid characters in username
                </div>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="First name"
                    formControlName="first_name"
                  />
                </mat-form-field>
                <div
                  *ngIf="isStep1Valid && first_name?.errors?.['required']"
                  class="error"
                >
                  First name is required
                </div>
                <div
                  *ngIf="isStep1Valid && first_name?.errors?.['minlength']"
                  class="error"
                >
                  First name must be at least 2 characters
                </div>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="Last name"
                    formControlName="last_name"
                  />
                </mat-form-field>
                <div
                  *ngIf="isStep1Valid && last_name?.errors?.['required']"
                  class="error"
                >
                  Last name is required
                </div>
                <div
                  *ngIf="isStep1Valid && last_name?.errors?.['minlength']"
                  class="error"
                >
                  Last name must be at least 2 characters
                </div>
                <button mat-button matStepperNext (click)="validateStep1()">
                  Next
                </button>
              </form>
            </mat-step>
            <mat-step
              [stepControl]="locationForm"
              label="Fill out your address"
            >
              <form [formGroup]="locationForm" class="addr">
                <ng-template matStepLabel>Address</ng-template>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="Address"
                    formControlName="address"
                  />
                </mat-form-field>
                <div *ngIf="isStep2Valid && address?.errors?.['required']" class="error">
                  Address is required
                </div>
                <div *ngIf="isStep2Valid && address?.errors?.['pattern']" class="error">
                  Invalid characters in address
                </div>
                <mat-form-field>
                  <input matInput placeholder="City" formControlName="city" />
                </mat-form-field>
                <div *ngIf="isStep2Valid && city?.errors?.['required']" class="error">
                  City is required
                </div>
                <div *ngIf="isStep2Valid && city?.errors?.['pattern']" class="error">
                  Invalid characters in city
                </div>

                <mat-form-field>
                  <mat-select
                    matNativeControl
                    placeholder="State"
                    formControlName="state"
                  >
                    <mat-option *ngFor="let state of usStates" [value]="state">
                      {{ state }}
                    </mat-option>
                  </mat-select>
                </mat-form-field>
                <div *ngIf="isStep2Valid && state?.errors?.['required']" class="error">
                  State is required
                </div>
                <mat-form-field>
                  <input matInput placeholder="Zip" formControlName="zip" />
                </mat-form-field>
                <div *ngIf="isStep2Valid && zip?.errors?.['required']" class="error">
                  Zip is required
                </div>
                <div *ngIf="isStep2Valid && zip?.errors?.['pattern']" class="error">
                  Invalid zip code
                </div>
                <button mat-button matStepperNext (click)="validateStep2()">Next</button>
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
                  <mat-select
                    matNativeControl
                    placeholder="Method of contact"
                    formControlName="method_of_contact"
                  >
                    <mat-option
                      *ngFor="let contact of contactOptions"
                      [value]="contact"
                    >
                      {{ contact }}
                    </mat-option>
                  </mat-select>
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
                    type="password"
                  />
                </mat-form-field>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="Confirm password"
                    formControlName="confirmPassword"
                    type="password"
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
  isStep1Valid: boolean = false;
  isStep2Valid: boolean = false;
  isStep3Valid: boolean = false;
  isStep4Valid: boolean = false;

  validateStep1() {
    this.isStep1Valid = true;
  }

  validateStep2() {
    this.isStep2Valid = this.locationForm.valid;
  }

  validateStep3() {
    this.isStep3Valid = this.contactForm.valid;
  }

  validateStep4() {
    this.isStep4Valid = this.passwordForm.valid;
  }

  usStates: string[] = [
    'Alabama',
    'Alaska',
    'Arizona',
    'Arkansas',
    'California',
    'Colorado',
    'Connecticut',
    'Delaware',
    'Florida',
    'Georgia',
    'Hawaii',
    'Idaho',
    'Illinois',
    'Indiana',
    'Iowa',
    'Kansas',
    'Kentucky',
    'Louisiana',
    'Maine',
    'Maryland',
    'Massachusetts',
    'Michigan',
    'Minnesota',
    'Mississippi',
    'Missouri',
    'Montana',
    'Nebraska',
    'Nevada',
    'New Hampshire',
    'New Jersey',
    'New Mexico',
    'New York',
    'North Carolina',
    'North Dakota',
    'Ohio',
    'Oklahoma',
    'Oregon',
    'Pennsylvania',
    'Rhode Island',
    'South Carolina',
    'South Dakota',
    'Tennessee',
    'Texas',
    'Utah',
    'Vermont',
    'Virginia',
    'Washington',
    'West Virginia',
    'Wisconsin',
    'Wyoming',
  ];

  contactOptions: string[] = ['Call', 'Email', 'Text'];

  nameForm = new FormGroup({
    first_name: new FormControl('', [
      Validators.required,
      Validators.minLength(2),
      Validators.maxLength(30),
      Validators.pattern('^[a-zA-Z]+$'),
    ]),
    last_name: new FormControl('', [
      Validators.required,
      Validators.minLength(2),
      Validators.maxLength(30),
      Validators.pattern('^[a-zA-Z]+$'),
    ]),
    username: new FormControl('', [
      Validators.required,
      Validators.minLength(2),
      Validators.maxLength(30),
      Validators.pattern('^[a-zA-Z0-9_!@#]+$'),
    ]),
  });

  get first_name() {
    return this.nameForm.get('first_name');
  }

  get last_name() {
    return this.nameForm.get('last_name');
  }

  get username() {
    return this.nameForm.get('username');
  }

  locationForm = new FormGroup({
    address: new FormControl('', [
      Validators.required,
      Validators.pattern('^[a-zA-Z0-9\\s,-.]+$'),
    ]),
    city: new FormControl('', [
      Validators.required,
      Validators.pattern('^[a-zA-Z\\s]+$'),
    ]),
    state: new FormControl('', Validators.required),
    zip: new FormControl('', [
      Validators.required,
      Validators.pattern('^[0-9]{5}$'),
    ]),
  });

  get address() {
    return this.locationForm.get('address');
  }

  get city() {
    return this.locationForm.get('city');
  }

  get state() {
    return this.locationForm.get('state');
  }

  get zip() {
    return this.locationForm.get('zip');
  }

  contactForm = new FormGroup({
    phone: new FormControl('', [Validators.required, phoneNumberValidator()]),
    method_of_contact: new FormControl('', Validators.required),
  });

  get phone() {
    return this.contactForm.get('phone');
  }

  get method_of_contact() {
    return this.contactForm.get('method_of_contact');
  }

  passwordMatchValidator(control: AbstractControl): ValidationErrors | null {
    const password = control.get('password');
    const confirmPassword = control.get('confirmPassword');

    return password &&
      confirmPassword &&
      password.value === confirmPassword.value
      ? null
      : { passwordMismatch: true };
  }

  passwordForm = new FormGroup(
    {
      password: new FormControl('', [
        Validators.required,
        Validators.minLength(8),
        Validators.pattern('^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).*$'),
      ]),
      confirmPassword: new FormControl('', [
        Validators.required,
        Validators.minLength(8),
        Validators.pattern('^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9]).*$'),
      ]),
    },
    { validators: this.passwordMatchValidator }
  );

  get password() {
    return this.passwordForm.get('password');
  }
}

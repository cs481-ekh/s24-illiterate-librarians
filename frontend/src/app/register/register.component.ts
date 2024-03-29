import { Component, ViewChild } from '@angular/core';
import {
  FormControl,
  FormGroup,
  Validators,
  AbstractControl,
  ValidatorFn,
  ValidationErrors,
} from '@angular/forms';
import { Router } from '@angular/router';
import { RegistrationService } from '../_services/register/registration.service';
import { NewAccount } from '../new-account';
import { MatStepper } from '@angular/material/stepper';
import { AuthService } from '../_services/auth.service';

export function phoneNumberValidator(): ValidatorFn {
  return (control: AbstractControl): { [key: string]: any } | null => {
    const rawNumbers = control.value.replace(/\D/g, '');
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
            <mat-step [stepControl]="userPassForm">
              <form [formGroup]="userPassForm" class="name">
                <ng-template matStepLabel>Username and Password</ng-template>

                <div *ngIf="duplicateUsername" class="error">
                  Username already exists, please choose another and try again
                </div>
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
                    placeholder="Username"
                    formControlName="username"
                  />
                </mat-form-field>

                <div class="pass-desc">
                  Password must be at least 8 characters and contain at least
                  one uppercase letter, one lowercase letter, and one number
                </div>
                <div
                  *ngIf="isStep1Valid && password?.errors?.['required']"
                  class="error"
                >
                  Password is required
                </div>
                <div
                  *ngIf="isStep1Valid && password?.errors?.['minlength']"
                  class="error"
                >
                  Password must be at least 8 characters
                </div>

                <div
                  *ngIf="isStep1Valid && password?.errors?.['pattern']"
                  class="error"
                >
                  Password must contain at least one uppercase letter, one
                  lowercase letter, and one number
                </div>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="Password"
                    formControlName="password"
                    type="password"
                  />
                </mat-form-field>
                <div
                  *ngIf="isStep1Valid && confirmPassword?.errors?.['required']"
                  class="error"
                >
                  Confirm password is required
                </div>
                <div
                  *ngIf="isStep1Valid && confirmPassword?.errors?.['minlength']"
                  class="error"
                >
                  Confirm password must be at least 8 characters
                </div>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="Confirm password"
                    formControlName="confirmPassword"
                    type="password"
                  />
                </mat-form-field>
                <div
                  *ngIf="isStep1Valid && userPassForm.errors?.['passwordMismatch']"
                  class="error"
                >
                  Passwords do not match
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
                <ng-template matStepLabel>Mailing Address</ng-template>

                <div
                  *ngIf="isStep2Valid && address?.errors?.['required']"
                  class="error"
                >
                  Address is required
                </div>
                <div
                  *ngIf="isStep2Valid && address?.errors?.['pattern']"
                  class="error"
                >
                  Invalid characters in address
                </div>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="Address"
                    formControlName="address"
                  />
                </mat-form-field>

                <div
                  *ngIf="isStep2Valid && city?.errors?.['required']"
                  class="error"
                >
                  City is required
                </div>

                <div
                  *ngIf="isStep2Valid && city?.errors?.['pattern']"
                  class="error"
                >
                  Invalid characters in city
                </div>
                <mat-form-field>
                  <input matInput placeholder="City" formControlName="city" />
                </mat-form-field>

                <div
                  *ngIf="isStep2Valid && state?.errors?.['required']"
                  class="error"
                >
                  State is required
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

                <div
                  *ngIf="isStep2Valid && zip?.errors?.['required']"
                  class="error"
                >
                  Zip is required
                </div>
                <div
                  *ngIf="isStep2Valid && zip?.errors?.['pattern']"
                  class="error"
                >
                  Invalid zip code
                </div>
                <mat-form-field>
                  <input matInput placeholder="Zip" formControlName="zip" />
                </mat-form-field>
                <button mat-button matStepperNext (click)="validateStep2()">
                  Next
                </button>
              </form>
            </mat-step>
            <mat-step
              [stepControl]="contactForm"
              label="Fill out your contact information"
            >
              <form [formGroup]="contactForm" class="contact">
                <ng-template matStepLabel>Contact Information</ng-template>

                <div
                  *ngIf="isStep3Valid && first_name?.errors?.['required']"
                  class="error"
                >
                  First name is required
                </div>
                <div
                  *ngIf="isStep3Valid && first_name?.errors?.['minlength']"
                  class="error"
                >
                  First name must be at least 2 characters
                </div>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="First name"
                    formControlName="first_name"
                  />
                </mat-form-field>

                <div
                  *ngIf="isStep3Valid && last_name?.errors?.['required']"
                  class="error"
                >
                  Last name is required
                </div>
                <div
                  *ngIf="isStep3Valid && last_name?.errors?.['minlength']"
                  class="error"
                >
                  Last name must be at least 2 characters
                </div>
                <mat-form-field>
                  <input
                    matInput
                    placeholder="Last name"
                    formControlName="last_name"
                  />
                </mat-form-field>

                <div *ngIf="duplicateEmail" class="error">
                  Email already exists, please choose another and try again
                </div>
                <div
                  *ngIf="isStep3Valid && email?.errors?.['required']"
                  class="error"
                >
                  Email is required
                </div>
                <div
                  *ngIf="isStep3Valid && email?.errors?.['email']"
                  class="error"
                >
                  Invalid email
                </div>
                <mat-form-field>
                  <input matInput placeholder="Email" formControlName="email" />
                </mat-form-field>
                <div *ngIf="duplicatePhone" class="error">
                  Phone number already exists, please choose another and try
                  again
                </div>
                <div
                  *ngIf="isStep3Valid && phone?.errors?.['required']"
                  class="error"
                >
                  Phone number is required
                </div>
                <div
                  *ngIf="isStep3Valid && phone?.errors?.['phoneNumberInvalid'] && !phone?.errors?.['required']"
                  class="error"
                >
                  Invalid phone number
                </div>

                <mat-form-field>
                  <input matInput placeholder="Phone" formControlName="phone" />
                </mat-form-field>
                <div
                  *ngIf="isStep3Valid && method_of_contact?.errors?.['required']"
                  class="error"
                >
                  Method of contact is required
                </div>
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
                <button mat-button matStepperNext (click)="validateStep3()">
                  Next
                </button>
              </form>
            </mat-step>
            <mat-step label="Done">
              <ng-template matStepLabel>Done</ng-template>
              <div class="done">
                Thank you for registering with Literacy Link!
                <button mat-button (click)="submit()">
                  Submit and Go to Dashboard
                </button>
              </div>
            </mat-step>
          </div>
        </mat-stepper>
      </div>
    </div>
  `,
  styleUrl: './register.component.css',
})
export class RegisterComponent {
  constructor(
    private regService: RegistrationService,
    private router: Router,
    private authService : AuthService
  ) {}

  @ViewChild('stepper') stepper!: MatStepper;

  isStep1Valid: boolean = false;
  isStep2Valid: boolean = false;
  isStep3Valid: boolean = false;

  validateStep1() {
    this.isStep1Valid = true;
  }

  validateStep2() {
    this.isStep2Valid = true;
  }

  validateStep3() {
    this.isStep3Valid = true;
  }

  duplicateUsername = false;
  duplicateEmail = false;
  duplicatePhone = false;

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

  userPassForm = new FormGroup(
    {
      username: new FormControl('', [
        Validators.required,
        Validators.minLength(2),
        Validators.maxLength(30),
        Validators.pattern('^[a-zA-Z0-9_!@#]+$'),
      ]),
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

  get username() {
    return this.userPassForm.get('username');
  }

  get password() {
    return this.userPassForm.get('password');
  }

  get confirmPassword() {
    return this.userPassForm.get('confirmPassword');
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

    email: new FormControl('', [Validators.required, Validators.email]),
  });

  get first_name() {
    return this.contactForm.get('first_name');
  }

  get last_name() {
    return this.contactForm.get('last_name');
  }

  get phone() {
    return this.contactForm.get('phone');
  }

  get method_of_contact() {
    return this.contactForm.get('method_of_contact');
  }

  get email() {
    return this.contactForm.get('email');
  }

  passwordMatchValidator(control: AbstractControl): ValidationErrors | null {
    const password = control.get('password');
    const confirmPassword = control.get('confirmPassword');

    if (
      password &&
      confirmPassword &&
      password.value === confirmPassword.value
    ) {
      return null;
    } else {
      return { passwordMismatch: true };
    }
  }

  submit() {
    if (
      !(
        this.userPassForm.valid &&
        this.locationForm.valid &&
        this.contactForm.valid
      )
    ) {
      return;
    }
    let formData = this.compileData();

    let contact = '';

    if (formData.method_of_contact === 'Call') {
      contact = 'C';
    } else if (formData.method_of_contact === 'Email') {
      contact = 'E';
    } else {
      contact = 'T';
    }

    let phone = formData.phone?.replace(/\D/g, '');

    let newAccount: NewAccount = {
      username: formData.username ?? '',
      password_hash: formData.password ?? '',
      email: formData.email ?? '',
      first_name: formData.first_name ?? '',
      last_name: formData.last_name ?? '',
      phone_number: phone ?? '',
      mailing_address: this.compileAddress() ?? '',
      pref_method_comm: contact,
    };

    this.regService.register(newAccount).subscribe(
      (response) => {
        console.log(response);
        this.authService.login(newAccount.username, newAccount.password_hash).subscribe
        (response => {
          console.log(response);
          this.router.navigate(['/dashboard']);
        },
        (error) => {
          console.log(error);
          this.router.navigate(['/login']);
        });
      },
      (error) => {
        const errorMessage = error.error.message;
        if (errorMessage.includes('username')) {
          this.duplicateUsername = true;
          this.stepper.selectedIndex = 0;
        }
        if (errorMessage.includes('email')) {
          this.duplicateEmail = true;

          this.stepper.selectedIndex = 2;
        }
        if (errorMessage.includes('phone')) {
          this.duplicatePhone = true;

          this.stepper.selectedIndex = 2;
        }
      }
    );
  }

  compileAddress() {
    // return string with address, city, state, zip
    return `${this.locationForm.value.address}, ${this.locationForm.value.city}, ${this.locationForm.value.state}, ${this.locationForm.value.zip}`;
  }

  compileData() {
    return {
      ...this.userPassForm.value,
      ...this.locationForm.value,
      ...this.contactForm.value,
    };
  }
}

import { Component } from '@angular/core';
import { AbstractControl, FormBuilder, FormControl, FormGroup, ValidatorFn, Validators } from '@angular/forms';
import { ProfileService } from '../../_services/profile/profile.service';
import { ProfileData } from './update-account';

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
  selector: 'app-parent',
  templateUrl: './parent.component.html',
  styleUrl: './parent.component.css'
})
export class ProfileParentComponent {
  formValid: boolean = false;


  constructor(
    private profileService: ProfileService
  ) {}

  profileForm = new FormGroup({
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
    email: new FormControl('', [
      Validators.required, 
      Validators.email
    ]),
    phone: new FormControl('', [
      Validators.required, 
      phoneNumberValidator()
    ]),
    street_address: new FormControl('', [
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

  get first_name() {
    return this.profileForm.get('first_name');
  }

  get last_name() {
    return this.profileForm.get('last_name');
  }

  get username() {
    return this.profileForm.get('username');
  }

  get email() {
    return this.profileForm.get('email');
  }

  get phone() {
    return this.profileForm.get('phone');
  }

  get street_address() {
    return this.profileForm.get('street_address');
  }

  get city() {
    return this.profileForm.get('city');
  }

  get state() {
    return this.profileForm.get('state');
  }

  get zip() {
    return this.profileForm.get('zip');
  }

  submit() {
    if (this.profileForm.valid) {
      this.formValid = true;
      const profileData: ProfileData = {
        firstName: this.profileForm.value.first_name ?? '',
        lastName: this.profileForm.value.last_name ?? '',
        username: this.profileForm.value.username ?? '',
        email: this.profileForm.value.email ?? '',
        phone: this.profileForm.value.phone ?? '',
        address: this.compileAddress() ?? '',
      };
      this.profileService.updateProfile(profileData).subscribe((response) => {
        console.log(response);
      },
      (error) => {
        console.log(error);
      }
      );
    }
  }

  ngOnInit() {
    const editButton = document.getElementById('edit');
    const saveButton = document.getElementById('save');
    const profileEditing = document.getElementById('profile-editing');
    const profileView = document.getElementById('profile-view');
    

    if (editButton && saveButton && profileEditing && profileView) {
      editButton.addEventListener('click', function () {
        editButton.style.display = 'none';
        profileEditing.style.display = 'block';
        saveButton.style.display = 'block';
        profileView.style.display = 'none';
      });

      saveButton.addEventListener('click', function () {
        editButton.style.display = 'block';
        profileEditing.style.display = 'none';
        saveButton.style.display = 'none';
        profileView.style.display = 'block';
      });
    }
  }

  compileAddress() {
    // return string with street_address, city, state, zip
    return this.profileForm.value.street_address + ', ' + this.profileForm.value.city + ', ' + this.profileForm.value.state + ' ' + this.profileForm.value.zip;
  }

  states = [
    { name: 'Alabama', value: 'AL' },
    { name: 'Alaska', value: 'AK' },
    { name: 'Arizona', value: 'AZ' },
    { name: 'Arkansas', value: 'AR' },
    { name: 'California', value: 'CA' },
    { name: 'Colorado', value: 'CO' },
    { name: 'Connecticut', value: 'CT' },
    { name: 'Delaware', value: 'DE' },
    { name: 'District Of Columbia', value: 'DC' },
    { name: 'Florida', value: 'FL' },
    { name: 'Georgia', value: 'GA' },
    { name: 'Hawaii', value: 'HI' },
    { name: 'Idaho', value: 'ID' },
    { name: 'Illinois', value: 'IL' },
    { name: 'Indiana', value: 'IN' },
    { name: 'Iowa', value: 'IA' },
    { name: 'Kansas', value: 'KS' },
    { name: 'Kentucky', value: 'KY' },
    { name: 'Louisiana', value: 'LA' },
    { name: 'Maine', value: 'ME' },
    { name: 'Maryland', value: 'MD' },
    { name: 'Massachusetts', value: 'MA' },
    { name: 'Michigan', value: 'MI' },
    { name: 'Minnesota', value: 'MN' },
    { name: 'Mississippi', value: 'MS' },
    { name: 'Missouri', value: 'MO' },
    { name: 'Montana', value: 'MT' },
    { name: 'Nebraska', value: 'NE' },
    { name: 'Nevada', value: 'NV' },
    { name: 'New Hampshire', value: 'NH' },
    { name: 'New Jersey', value: 'NJ' },
    { name: 'New Mexico', value: 'NM' },
    { name: 'New York', value: 'NY' },
    { name: 'North Carolina', value: 'NC' },
    { name: 'North Dakota', value: 'ND' },
    { name: 'Ohio', value: 'OH' },
    { name: 'Oklahoma', value: 'OK' },
    { name: 'Oregon', value: 'OR' },
    { name: 'Pennsylvania', value: 'PA' },
    { name: 'Rhode Island', value: 'RI' },
    { name: 'South Carolina', value: 'SC' },
    { name: 'South Dakota', value: 'SD' },
    { name: 'Tennessee', value: 'TN' },
    { name: 'Texas', value: 'TX' },
    { name: 'Utah', value: 'UT' },
    { name: 'Vermont', value: 'VT' },
    { name: 'Virginia', value: 'VA' },
    { name: 'Washington', value: 'WA' },
    { name: 'West Virginia', value: 'WV' },
    { name: 'Wisconsin', value: 'WI' },
    { name: 'Wyoming', value: 'WY' }
  ];
}

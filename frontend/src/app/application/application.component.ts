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
    const rawNumbers = control.value.replace(/\D/g, '');
    if (rawNumbers.length === 10) {
      return null; // Validation passed
    } else {
      return { phoneNumberInvalid: true }; // Validation failed
    }
  };
}

@Component({
  selector: 'app-application',
  templateUrl: './application.component.html',
  styleUrl: './application.component.css',
})



export class ApplicationComponent {

  isStep1Valid = false;
  isStep2Valid = false;
  isStep3Valid = false;

  validateStep1() {
    this.isStep1Valid = this.firstFormGroup.valid;
  }

  prelimFormGroup = new FormGroup({
  });
  firstFormGroup = new FormGroup({
    PolicyAgreement: new FormControl('', [Validators.required]),
    ChildDataConsent: new FormControl('', [Validators.required]),
    PhotoReleaseConsent: new FormControl('', [Validators.required]),
    ZoomConsent: new FormControl('', [Validators.required]),
    BookAcknowledgement: new FormControl('', [Validators.required]),
    NeedFinancialAssistance: new FormControl('', [Validators.required]),
  });

  get PolicyAgreement() {
    return this.firstFormGroup.get('PolicyAgreement');
  }

  get ChildDataConsent() {
    return this.firstFormGroup.get('ChildDataConsent');
  }

  get PhotoReleaseConsent() {
    return this.firstFormGroup.get('PhotoReleaseConsent');
  }

  get ZoomConsent() {
    return this.firstFormGroup.get('ZoomConsent');
  }

  get BookAcknowledgement() {
    return this.firstFormGroup.get('BookAcknowledgement');
  }

  get NeedFinancialAssistance() {
    return this.firstFormGroup.get('NeedFinancialAssistance');
  }


  secondFormGroup = new FormGroup({
    Guardian2FirstN: new FormControl(''),
    Guardian2LastN: new FormControl(''),
    Guardian2Phone: new FormControl('', phoneNumberValidator()),
    Guardian2Email: new FormControl(''),
    EmergencyConName: new FormControl('', [Validators.required]),
    EmergencyConRelation: new FormControl('', [Validators.required]),
    EmergencyConPhone: new FormControl('', [Validators.required, phoneNumberValidator()]),
  });

  get Guardian2FirstN() {
    return this.secondFormGroup.get('Guardian2FirstN');
  }

  get Guardian2LastN() {
    return this.secondFormGroup.get('Guardian2LastN');
  }

  get Guardian2Phone() {
    return this.secondFormGroup.get('Guardian2Phone');
  }

  get Guardian2Email() {
    return this.secondFormGroup.get('Guardian2Email');
  }

  get EmergencyConName() {
    return this.secondFormGroup.get('EmergencyConName');
  }

  get EmergencyConRelation() {
    return this.secondFormGroup.get('EmergencyConRelation');
  }

  get EmergencyConPhone() {
    return this.secondFormGroup.get('EmergencyConPhone');
  }

  thirdFormControlGroup = new FormGroup({
    PreviousChildParticipation: new FormControl('', [Validators.required]),
    WhatSSemester: new FormControl('', [Validators.required]),
    ChildCurrentSchool: new FormControl('', [Validators.required]),
    ListLanguagesSpoken: new FormControl('', [Validators.required]),
    ReceivedSpecialEd: new FormControl('', [Validators.required]),
    ListChallenges: new FormControl('', [Validators.required]),
    HowLongConcerned: new FormControl('', [Validators.required]),
    DescribeHopes: new FormControl('', [Validators.required]),
    ChildAllergyMeds: new FormControl('', [Validators.required]),
    MiscInfo: new FormControl('', [Validators.required]),
    HearAboutLitLab: new FormControl('', [Validators.required]),
  });

  get PreviousChildParticipation() {
    return this.thirdFormControlGroup.get('PreviousChildParticipation');
  }

  get WhatSSemester() {
    return this.thirdFormControlGroup.get('WhatSSemester');
  }

  get ChildCurrentSchool() {
    return this.thirdFormControlGroup.get('ChildCurrentSchool');
  }

  get ListLanguagesSpoken() {
    return this.thirdFormControlGroup.get('ListLanguagesSpoken');
  }

  get ReceivedSpecialEd() {
    return this.thirdFormControlGroup.get('ReceivedSpecialEd');
  }

  get ListChallenges() {
    return this.thirdFormControlGroup.get('ListChallenges');
  }

  get HowLongConcerned() {
    return this.thirdFormControlGroup.get('HowLongConcerned');
  }

  get DescribeHopes() {
    return this.thirdFormControlGroup.get('DescribeHopes');
  }

  get ChildAllergyMeds() {
    return this.thirdFormControlGroup.get('ChildAllergyMeds');
  }

  get MiscInfo() {
    return this.thirdFormControlGroup.get('MiscInfo');
  }

  get HearAboutLitLab() {
    return this.thirdFormControlGroup.get('HearAboutLitLab');
  }

  submit(){
    console.log(this.firstFormGroup.value);
    console.log(this.secondFormGroup.value);
    console.log(this.thirdFormControlGroup.value);
  }

  constructor() {}
}

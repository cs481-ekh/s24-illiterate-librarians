import { Component } from '@angular/core';
import {
  FormControl,
  FormGroup,
  Validators,
  AbstractControl,
  ValidatorFn,
  ValidationErrors,
} from '@angular/forms';
import { MatStepper } from '@angular/material/stepper';
// import { ApplicationService } from '../application.service';

@Component({
  selector: 'app-application',
  templateUrl: './application.component.html',
  styleUrl: './application.component.css',
})
export class ApplicationComponent {



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
  secondFormGroup = new FormGroup({
    Guardian2FirstN: new FormControl(''),
    Guardian2LastN: new FormControl(''),
    Guardian2Phone: new FormControl(''),
    Guardian2Email: new FormControl(''),
    EmergencyConName: new FormControl('', [Validators.required]),
    EmergencyConRelation: new FormControl('', [Validators.required]),
    EmergencyConPhone: new FormControl('', [Validators.required]),
  });
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

  submit(){
    console.log(this.firstFormGroup.value);
    console.log(this.secondFormGroup.value);
    console.log(this.thirdFormControlGroup.value);
  }

  constructor() {}
}

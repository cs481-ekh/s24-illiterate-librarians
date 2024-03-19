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
  template: `
  <div class="application">
    <div class="form-container">
        <div class="title">
            <h1>Application</h1>
        </div>
        <div class="form">
            <mat-stepper [linear]="false" #stepper class="stepper">
                <mat-step [stepControl]="firstFormGroup">
                    <form [formGroup]="firstFormGroup">
                        <ng-template matStepLabel>Personal Information</ng-template>
                        <mat-form-field>
                            <mat-label>Child Data Consent</mat-label>
                            <mat-select formControlName="ChildDataConsent">
                                <mat-option value="true">Yes</mat-option>
                                <mat-option value="false">No</mat-option>
                            </mat-select>
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Photo Release Consent</mat-label>
                            <mat-select formControlName="PhotoReleaseConsent">
                                <mat-option value="true">Yes</mat-option>
                                <mat-option value="false">No</mat-option>
                            </mat-select>
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Need Financial Assistance</mat-label>
                            <mat-select formControlName="NeedFinancialAssistance">
                                <mat-option value="true">Yes</mat-option>
                                <mat-option value="false">No</mat-option>
                            </mat-select>
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Guardian 2 First Name</mat-label>
                            <input matInput formControlName="Guardian2FirstN">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Guardian 2 Last Name</mat-label>
                            <input matInput formControlName="Guardian2LastN">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Guardian 2 Phone</mat-label>
                            <input matInput formControlName="Guardian2Phone">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Guardian 2 Email</mat-label>
                            <input matInput formControlName="Guardian2Email">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Emergency Contact Name</mat-label>
                            <input matInput formControlName="EmergencyConName">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Emergency Contact Relation</mat-label>
                            <input matInput formControlName="EmergencyConRelation">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Emergency Contact Phone</mat-label>
                            <input matInput formControlName="EmergencyConPhone">
                        </mat-form-field>
                        <div>
                            <button mat-button matStepperNext>Next</button>
                        </div>
                    </form>
                </mat-step>
                <mat-step [stepControl]="secondFormGroup">
                    <form [formGroup]="secondFormGroup">
                        <ng-template matStepLabel>Child Information</ng-template>
                        <mat-form-field>
                            <mat-label>Previous Child Participation</mat-label>
                            <mat-select formControlName="PreviousChildParticipation">
                                <mat-option value="true">Yes</mat-option>
                                <mat-option value="false">No</mat-option>
                            </mat-select>
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>What Semester</mat-label>
                            <input matInput formControlName="WhatSSemester">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Child Current School</mat-label>
                            <input matInput formControlName="ChildCurrentSchool">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>List Languages Spoken</mat-label>
                            <input matInput formControlName="ListLanguagesSpoken">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Received Special Ed</mat-label>
                            <input matInput formControlName="ReceivedSpecialEd">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>List Challenges</mat-label>
                            <input matInput formControlName="ListChallenges">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>How Long Concerned</mat-label>
                            <input matInput formControlName="HowLongConcerned">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Describe Hopes</mat-label>
                            <input matInput formControlName="DescribeHopes">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Child Allergy Meds</mat-label>
                            <input matInput formControlName="ChildAllergyMeds">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Misc Info</mat-label>
                            <input matInput formControlName="MiscInfo">
                        </mat-form-field>
                        <mat-form-field>
                            <mat-label>Hear About LitLab</mat-label>
                            <input matInput formControlName="HearAboutLitLab">
                        </mat-form-field>
                        <div>
                            <button mat-button (click)="stepper.previous()">Back</button>
                            <button mat-button (click)="stepper.next
                            ()">Next</button>
                        </div>
                    </form>
                </mat-step>

                <mat-step>
                    <ng-template matStepLabel>Review</ng-template>
                    <div>
                        <button mat-button (click)="stepper.previous()">Back</button>
                        <button mat-button (click)="stepper.next()">Next</button>
                    </div>
                </mat-step>
            </mat-stepper>
        </div>
    </div>
</div>
  `,
  styleUrl: './application.component.css',
})
export class ApplicationComponent {
  firstFormGroup = new FormGroup({
    ChildDataConsent: new FormControl('', [Validators.required]),
    PhotoReleaseConsent: new FormControl('', [Validators.required]),
    NeedFinancialAssistance: new FormControl('', [Validators.required]),
    Guardian2FirstN: new FormControl('', [Validators.required]),
    Guardian2LastN: new FormControl('', [Validators.required]),
    Guardian2Phone: new FormControl('', [Validators.required]),
    Guardian2Email: new FormControl('', [Validators.required]),
    EmergencyConName: new FormControl('', [Validators.required]),
    EmergencyConRelation: new FormControl('', [Validators.required]),
    EmergencyConPhone: new FormControl('', [Validators.required]),
  });
  secondFormGroup = new FormGroup({
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
}

// package model

// import (
// 	"time"
// )

// NOTE: not all of the things that will be asked in the application will
// be stored here / in the app_for_tutoring table. Some things like the days
// that parents will be unavailable will be stored in other tables. For a
// full list of questions in the application, look at the original applicaiton.

// TODO: list out in the comments here the questions that should be asked for the benefit of future developers

// type Application struct {
// 	AppForTutId       []byte    `gorm:"type:BINARY(16);default:UUID_TO_BIN(UUID(), 1);primaryKey" json:"app_for_tut_id"`
// 	ChildId           []byte    `gorm:"type:BINARY(16);not null" json:"child_id"`
// 	ParentId          []byte    `gorm:"type:BINARY(16);not null" json:"parent_id"`
// 	AppComplete       bool      `gorm:"default:false" json:"app_complete"`
// 	AppCompleteDate   time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"app_complete_date"`
// 	DesiredSemesterId []byte    `gorm:"type:BINARY(16);not null" json:"desired_semester_id"`

// 	ChildDataConsent        bool `gorm:"default:false" json:"child_data_consent"`
// 	PhotoReleaseConsent     bool `gorm:"default:false" json:"photo_release_consent"`
// 	NeedFinancialAssistance bool `gorm:"default:false" json:"need_financial_assistance"`

// 	Guardian2FirstN string `gorm:"type:VARCHAR(50)" json:"guardian2_first_n"`
// 	Guardian2LastN  string `gorm:"type:VARCHAR(50)" json:"guardian2_last_n"`
// 	Guardian2Phone  string `gorm:"type:VARCHAR(20)" json:"guardian2_phone"`
// 	Guardian2Email  string `gorm:"type:VARCHAR(255);NOT NULL" json:"guardian2_email"`

// 	EmergencyConName     string `gorm:"type:VARCHAR(50)" json:"emergency_con_name"`
// 	EmergencyConRelation string `gorm:"type:VARCHAR(255)" json:"emergency_con_relation"`
// 	EmergencyConPhone    string `gorm:"type:VARCHAR(20)" json:"emergency_con_phone"`

// 	Previous_child_participation bool   `gorm:"default:false"`
// 	WhatSSemester                string `gorm:"type:VARCHAR(50)" json:"what_semester"`
// 	ChildCurrentSchool           string `gorm:"type:VARCHAR(50)" json:"child_current_school"`
// 	ListLanguagesSpoken          string `gorm:"type:VARCHAR(255)" json:"list_languages_spoken"`
// 	ReceivedSpecialEd            string `gorm:"type:VARCHAR(50)" json:"received_special_ed"`
// 	ListChallenges               string `gorm:"type:VARCHAR(50)" json:"list_challenges"`
// 	HowLongConcerned             string `gorm:"type:VARCHAR(50)" json:"how_long_concerned"`
// 	DescribeHopes                string `gorm:"type:VARCHAR(50)" json:"describe_hopes"`
// 	ChildAllergyMeds             string `gorm:"type:VARCHAR(50)" json:"child_allergy_meds"`
// 	MiscInfo                     string `gorm:"type:VARCHAR(50)" json:"misc_info"`
// 	HearAboutLitLab              string `gorm:"type:VARCHAR(50)" json:"hear_about_litLab"`
// }

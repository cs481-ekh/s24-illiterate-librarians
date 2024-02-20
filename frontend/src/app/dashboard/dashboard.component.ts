import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CalendarModule } from 'angular-calendar';
@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule,
    CalendarModule],
  template: ` 
  <!-- <span>
    <mat-form-field>
      <button (click)="openDialog">Toggle View</button>
    </mat-form-field>
    </span> -->
  <div class="calendar">
  
  <mwl-calendar-week-view
  [viewDate]="viewDate"
  [dayStartHour]="dayStartHour"
  [dayEndHour]="dayEndHour"
  [weekStartsOn]="weekStartsOn"
  [excludeDays]="excludeDays"
  [hourSegments]="hourSegments"
  [events]="events"
 >
</mwl-calendar-week-view>
</div>
  `,
  styleUrls: ['./dashboard.component.css'],
})
export class DashboardComponent {
  
  dayStartHour = 12;
  dayEndHour = 17;
  viewDate: Date = new Date();
  weekStartsOn: number = 1;
  excludeDays: number[] = [0, 6];
  hourSegments: number = 1;
  events= [
    {
      start: new Date(),
      title: 'An event with no end date',
    },
    {
      start: new Date(),
      end: new Date(),
      title: 'An event with an end date',
    },
  ];
}
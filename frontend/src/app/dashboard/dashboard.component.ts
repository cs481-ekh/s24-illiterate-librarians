import { Component } from '@angular/core';
//use angular calendar to display events
import { CommonModule } from '@angular/common';
import { CalendarMonthModule } from 'angular-calendar';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [CommonModule, CalendarMonthModule],
  template: ` <p>dashboard works!</p> 
  <mwl-calendar-month-view
  [viewDate]="viewDate"
  [events]="events"
  ></mwl-calendar-month-view>
  `,
  styleUrls: ['./dashboard.component.css'],
})
export class DashboardComponent {
  viewDate: Date = new Date();
  events = [
    {
      start: new Date(),
      title: 'An event',
    },
  ];
}

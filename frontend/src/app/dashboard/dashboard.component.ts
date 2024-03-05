import { Component, inject } from '@angular/core';
import { CalendarOptions } from '@fullcalendar/core'; // useful for typechecking
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Inject } from '@angular/core';
import { Event } from '../event';
import { EventsService } from '../_services/events.service';

export interface eventData {
  title: string;
  start: Date;
  id: string;
}

@Component({
  selector: 'app-dashboard',
  template: `
  <div class="notifications">
    <app-notifications></app-notifications>
  </div>
  <div class="dashboard">
  <div class="topcomponents">
    <div class="sessions">
    <app-dashsessions [events]="events"></app-dashsessions>
    </div>
    <div class="calendar">
        <full-calendar [options]="calendarOptions"></full-calendar>
    </div>
  </div>
  
</div>
  `,
  styleUrl: './dashboard.component.css'
})
export class DashboardComponent {

  events: Event[] = [];
  eventService: EventsService = inject(EventsService);

  calendarOptions: CalendarOptions = {
    initialView: 'dayGridWeek',
    plugins: [dayGridPlugin, interactionPlugin], 
    eventClick: (info) => {
      this.openDialog(info.event.title, info.event.start, info.event.id);
    },
    events: [
      { title: 'Tutoring', start: '2024-02-20', id: 'a', time: '10:00' }, //we will use a service to pass events when endpoint is ready
    ],
    height: '100%',
  };

  constructor(public dialog: MatDialog) {
    this.events = this.eventService.getEvents();
  }

  openDialog(title: string, start: Date | null, id: string) {
    const dialogRef = this.dialog.open(SessionModal, {
      width: '500px',
      height: '500px',
      data: {title, start, id},
      position: { top: '7%'},
    });
  }
}


@Component({
  selector: 'session-modal',
  template: `
  <h1 mat-dialog-title>{{data.title}}</h1>
  <div mat-dialog-content>
    <p> This date of this session is {{data.start | date: 'fullDate'}} </p>
</div>
<div mat-dialog-actions>
  <button mat-button [mat-dialog-close]="data.id" cdkFocusInitial>Ok</button>
</div>

`,
})
export class SessionModal {

  constructor(
    public dialogRef: MatDialogRef<SessionModal>,
    @Inject(MAT_DIALOG_DATA) public data: eventData) {}

  onNoClick(): void {
    this.dialogRef.close();
  }

}
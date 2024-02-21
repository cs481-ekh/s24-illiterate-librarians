import { Component } from '@angular/core';
import { CalendarOptions } from '@fullcalendar/core'; // useful for typechecking
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Inject } from '@angular/core';


export interface eventData {
  title: string;
  start: Date;
  id: string;
}

@Component({
  selector: 'app-dashboard',
  template: `
  <span class="topcomponents">
    <div class="calendar">
        <full-calendar [options]="calendarOptions"></full-calendar>
    </div>
  </span>
  `,
  styleUrl: './dashboard.component.css'
})
export class DashboardComponent {

  calendarOptions: CalendarOptions = {
    initialView: 'dayGridWeek',
    plugins: [dayGridPlugin, interactionPlugin],
    dateClick: this.handleDateClick.bind(this), 
    eventClick: (info) => {
      this.openDialog(info.event.title, info.event.start, info.event.id);
    },
    events: [
      { title: 'Tutoring', start: '2024-02-20', id: 'a', time: '10:00' }, //we will use a service to pass events when endpoint is ready
    ],
    height: '100%',
  };

  constructor(public dialog: MatDialog) {
  }

  openDialog(title: string, start: Date | null, id: string) {
    const dialogRef = this.dialog.open(SessionModal, {
      width: '500px',
      height: '500px',
      data: {title, start, id},
      position: { top: '7%'},
    });
  }

  handleDateClick(arg: any) {
    console.log("clicked!")
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
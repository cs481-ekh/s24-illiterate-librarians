import { Component, Input } from '@angular/core';
import { Event } from '../../event';

@Component({
  selector: 'app-dashsessions',
  //renders a frame to hold tutoring sessions
  template: `
  <div class="frame">
    <div class="title">
      <h1>Upcoming Tutoring Sessions</h1>
    </div>
    <div id="events">
  <div *ngFor="let event of events">
    <a [routerLink]="['/session']">
      <div class="event">
        <span class="timeTitle">
          <div class="event-title">{{ event.title }}</div>
          <div class="event-time">{{ event.time }}</div>
        </span>
        <span class="event-date">{{ event.start.toLocaleDateString() }}</span>
      </div>
    </a>
  </div>
</div>

  `,
  styleUrl: './dashsessions.component.css'
})
export class DashsessionsComponent {
  @Input() events!: Event[];
}

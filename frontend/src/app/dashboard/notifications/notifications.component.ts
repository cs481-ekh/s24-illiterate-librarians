import { Component } from '@angular/core';

@Component({
  selector: 'app-notifications',
  template: `
    <div class="frame">
      <div class="title">
        <h1>Notifications</h1>
      </div>
      <div id="events">
        <div class="event">
          <span class="timeTitle">
            <div class="event-title">New Tutoring Session</div>
            <div class="event-time>10:00"></div>
          </span>
          <span class="event-date">2021-04-01</span>
        </div>
      </div>
    </div>
  `,
  styleUrl: './notifications.component.css',
})
export class NotificationsComponent {}

import { Component, input } from '@angular/core';

@Component({
  selector: 'app-dashsessions',
  //renders a frame to hold tutoring sessions
  template: `
  <div class="frame">
    <div class="title">
      <h1>Upcoming Tutoring Sessions</h1>
    </div>
      <div id="events">      

        <div class="event">
          <span class="timeTitle">
            <div class="event-title">Weekly Tutoring</div>
            <div class="event-time">10:00</div>
          </span>
            <span class="event-date">2024-02-20</span>
        </div>
          <!-- <ul>
              <li *ngFor="let event of events">
                  <div class="event">
                      <div class="event-title">{{event.title}}</div>
                      <div class="event-time
                      ">{{event.time}}</div> -->
</div>

  `,
  styleUrl: './dashsessions.component.css'
})
export class DashsessionsComponent {
  //pass events in from the dashboard component
}

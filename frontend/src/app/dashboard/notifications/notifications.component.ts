import { Component } from '@angular/core';
import { Notification } from './notification';
import {
  trigger,
  state,
  style,
  animate,
  transition,
} from '@angular/animations';

@Component({
  selector: 'app-notifications',
  animations: [
    //write an animation to slide a div to the right
    trigger('slide', [
      state(
        'in',
        style({
          transform: 'translateX(0px)',
        })
      ),
      state(
        'right',
        style({
          transform: 'translateX(740px)',
        })
      ),
      state(
        'left',
        style({
          transform: 'translateX(-740px)',
        })
      ),
      transition('in => right', [animate('300ms')]),
      transition('in => left', [animate('300ms')]),
    ]),
  ],
  template: `
    <div class="frame">
      <div class="title">
        <h1>Announcements</h1>
        <p class="description">
          Check here for recent updates and upcoming events!
        </p>
      </div>
      <div class = "prev-button">
        <button mat-raised-button	class="arrow" color="primary" (click)="prevNotification()"> <span class="material-symbols-outlined">
arrow_back
</span></button>
      </div>
      <div class="notifications">
        <!-- trigger for animation goes here -->
        <div
          class="offset"
          [@slide]="triggerAnimation ? (direction === 'next' ? 'left' : 'right') : 'in'"
        >
          <div class="notification">
            <span class="notification-heading">
              <div class="notification-title">
                {{
                  notifications[
                    currentNotification == 0
                      ? notifications.length - 1
                      : currentNotification - 1
                  ].title
                }}
              </div>
              <div class="notification-time">
                {{
                  notifications[
                    currentNotification == 0
                      ? notifications.length - 1
                      : currentNotification - 1
                  ].time
                }}
              </div>
            </span>
            <span class="notification-content">
              {{
                notifications[
                  currentNotification == 0
                    ? notifications.length - 1
                    : currentNotification - 1
                ].content
              }}
            </span>
          </div>
          <div class="notification">
            <span class="notification-heading">
              <div class="notification-title">
                {{ notifications[currentNotification].title }}
              </div>
              <div class="notification-time">
                {{ notifications[currentNotification].time }}
              </div>
            </span>
            <span class="notification-content">
              {{ notifications[currentNotification].content }}
            </span>
          </div>
          <div class="notification">
            <span class="notification-heading">
              <div class="notification-title">
                {{
                  notifications[
                    currentNotification == notifications.length - 1
                      ? 0
                      : currentNotification + 1
                  ].title
                }}
              </div>
              <div class="notification-time">
                {{
                  notifications[
                    currentNotification == notifications.length - 1
                      ? 0
                      : currentNotification + 1
                  ].time
                }}
              </div>
            </span>
            <span class="notification-content">
              {{
                notifications[
                  currentNotification == notifications.length - 1
                    ? 0
                    : currentNotification + 1
                ].content
              }}
            </span>
          </div>
        </div>
      </div>
      <div class="next-button">
      <button mat-raised-button	class="arrow" color="primary" id="arrowright"(click)="nextNotification()"> <span class="material-symbols-outlined">
arrow_forward
</span></button>
      </div>
    </div>
  `,
  styleUrl: './notifications.component.css',
})
export class NotificationsComponent {

  direction: 'prev' | 'next' = 'next';
  currentNotification = 0;
  triggerAnimation = false;
  notifications: Notification[] = [
    {
      title: 'Virtual Tutoring Event',
      start: new Date('2021-04-17T10:00:00'),
      content:
        'We are excited to announce that we will be hosting a virtual tutoring event on Saturday, April 17th. This event will be open to all students and will feature a variety of tutoring sessions and workshops. More details to come!',
      time: 'April 17th',
      id: '1',
    },
    {
      title: 'New Tutoring Hours',
      time: 'March 15th',
      content:
        'Starting March 15th, we will be extending our tutoring hours to better accommodate our students. We will now be open from 8:00 AM to 10:00 PM, Monday through Saturday.',
      id: '2',
      start: new Date('2021-03-15T10:00:00'),
    },
    {
      title: 'New Tutoring Services',
      time: 'February 1st',
      content:
        'We are excited to announce that we will be offering new tutoring services starting February 1st. We will now be offering tutoring services for history, science, and foreign languages.',
      id: '3',
      start: new Date('2021-02-01T10:00:00'),
    },
    {
      title: 'New Tutoring Center',
      time: 'January 1st',
      content:
        'We are excited to announce that we will be opening a new tutoring center in the heart of downtown. This new location will feature state-of-the-art facilities and will be open to all students.',
      id: '4',
      start: new Date('2021-01-01T10:00:00'),
    },
    {
      title: 'Accepted into The Literacy Lab',
      time: 'January 1st',
      content:
        'We are excited to welcome you to the Boise State University Literacy Lab. Please let us know if you have any questions or need any assistance.',
      id: '5',
      start: new Date('2021-01-01T10:00:00'),
    },
  ];

  prevNotification() {
    this.direction = 'prev';
    this.triggerAnimation = true;
    setTimeout(() => {
      this.currentNotification =
        (this.currentNotification - 1 + this.notifications.length) %
        this.notifications.length;
      this.triggerAnimation = false;
    }, 300);
  }

  nextNotification() {
    this.direction = 'next';
    this.triggerAnimation = true;
    setTimeout(() => {
      this.currentNotification =
        (this.currentNotification + 1) % this.notifications.length;
      this.triggerAnimation = false;
    }, 300);
  }
}

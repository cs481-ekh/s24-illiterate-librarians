import { Component } from '@angular/core';

@Component({
  selector: 'app-notifications',
  template: `
    <div class="frame">
      <div class="title">
        <h1>Announcements</h1>
        <p class="description">Check here for important updates and events</p>
      </div>
      <div class="notifications">
        <div class="notification">
          <span class="notification-heading">
            <div class="notification-title">Virtual Tutoring Event</div>
            <div class="notification-time">April 17th</div>
          </span>
          <span class="notification-content"
            >We are excited to announce that we will be hosting a virtual
            tutoring event on Saturday, April 17th. This event will be open to
            all students and will feature a variety of tutoring sessions and
            workshops. More details to come!</span
          >
        </div>
        <div class="notification">
          <span class="notification-heading">
            <div class="notification-title">New Tutoring Hours</div>
            <div class="notification-time">March 15th</div>
          </span>
          <span class="notification-content"
            >Starting March 15th, we will be extending our tutoring hours to
            better accommodate our students. We will now be open from 8:00 AM to
            10:00 PM, Monday through Saturday.</span
          >
        </div>
        <div class="notification">
          <span class="notification-heading">
            <div class="notification-title">New Tutoring Services</div>
            <div class="notification-time">February 1st</div>
          </span>
          <span class="notification-content"
            >We are excited to announce that we will be offering new tutoring
            services starting February 1st. We will now be offering tutoring
            services for history, science, and foreign languages.</span
          >
        </div>
      </div>
    </div>
  `,
  styleUrl: './notifications.component.css',
})
export class NotificationsComponent {}

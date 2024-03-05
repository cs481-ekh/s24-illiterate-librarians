import { Injectable } from '@angular/core';
import { Event } from '../event';

@Injectable({
  providedIn: 'root'
})
export class EventsService {

  events: Event[] = [
    {
      title: "Math Tutoring",
      start: new Date("2021-04-01T10:00:00"),
      id: "1",
      time: "10:00 AM"
    },
    {
      title: "English Tutoring",
      start: new Date("2021-04-01T14:00:00"),
      id: "2",
      time: "2:00 PM"
    },
    {
      title: "Science Tutoring",
      start: new Date("2021-04-02T10:00:00"),
      id: "3",
      time: "10:00 AM"
    },
    {
      title: "History Tutoring",
      start: new Date("2021-04-02T14:00:00"),
      id: "4",
      time: "2:00 PM"
    },
    {
      title: "Math Tutoring",
      start: new Date("2021-04-03T10:00:00"),
      id: "5",
      time: "10:00 AM"
    },
    {
      title: "English Tutoring",
      start: new Date("2021-04-03T14:00:00"),
      id: "6",
      time: "2:00 PM"
    },
    {
      title: "Science Tutoring",
      start: new Date("2021-04-04T10:00:00"),
      id: "7",
      time: "10:00 AM"
    },
    {
      title: "History Tutoring",
      start: new Date("2021-04-04T14:00:00"),
      id: "8",
      time: "2:00 PM"
    }
  ];

  getEvents(): Event[] {
    return this.events;
  }

}

import { Component } from '@angular/core';
import { Router, NavigationEnd, Event as RouterEvent } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'frontend';
  showHead = true;

  constructor(private router: Router) { 
    this.router.events.subscribe((event: RouterEvent) => {
      if (event instanceof NavigationEnd) {
        if (event.url === '/login' || event.url === '/register' || event.url === '/application') {
          this.showHead = false;
        } else {
          this.showHead = true;
        }
      }
    });
  }
}
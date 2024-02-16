import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import {
  trigger,
  state,
  style,
  animate,
  transition,
} from '@angular/animations';
import { HostListener } from '@angular/core';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [CommonModule, RouterModule],
  animations: [
    trigger('openClose', [
      state(
        'open',
        style({
          opacity: 1,
          visibility: 'visible',
          backgroundColor: 'white',
        })
      ),
      state(
        'closed',
        style({
          visibility: 'hidden',
          height: '0px',
          opacity: 0,
        })
      ),
      transition('open => closed', [animate('0.2s')]),
      transition('closed => open', [animate('0.2s')]),
    ]),
  ],
  template: `
    <header class="header">
      <a href="https://sdp.boisestate.edu/index.html">
        <img src="../assets/sdp-logo-3.png" alt="SDP Logo" width="100px" />
      </a>
      <a (click)="toggle()">
        <div class="material-symbols-outlined">menu</div>
      </a>

      <div id="links">
        <a><div class="option">Home</div></a>
        <a [routerLink]="['/dashboard']"><div class="option">Dashboard</div></a>
        <a [routerLink]="['/profile']"><div class="option">Profile</div></a>
        <a [routerLink]="['/login']"
          ><div class="option" id="login">Login or Register</div></a
        >
      </div>
      <div id="menu" [@openClose]="openClose()">
        <a
          [routerLink]="['/dashboard']"
          [@openClose]="openClose()"
          (click)="toggle()"
          ><div class="option">Dashboard</div></a
        >
        <a
          [routerLink]="['/profile']"
          [@openClose]="openClose()"
          (click)="toggle()"
          ><div class="option">Profile</div></a
        >
        <a
          [routerLink]="['/login']"
          [@openClose]="openClose()"
          (click)="toggle()"
          ><div class="option" id="login">Login or Register</div></a
        >
      </div>
    </header>
  `,
  styleUrl: './header.component.css',
})
export class HeaderComponent {
  isOpen = false;

  toggle() {
    this.isOpen = !this.isOpen;
  }

  openClose() {
    return this.isOpen ? 'open' : 'closed';
  }

  animationState: string = 'start';

  @HostListener('window:resize', ['$event'])
  onResize() {
    this.checkWindowSize();
  }

  ngOnInit() {
    this.checkWindowSize();
  }

  private checkWindowSize(): void {
    if (window.innerWidth > 960) {
      this.isOpen = false;
    }
  }
}

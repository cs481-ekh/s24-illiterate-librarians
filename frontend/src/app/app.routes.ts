import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { ProfileComponent } from './profile/profile.component';
import { LoginComponent } from './login/login.component';
import { AboutComponent } from './about/about.component';
import { ContactComponent } from './contact/contact.component';

export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
    title: 'Literacy Link',
  },
  {
    path: 'dashboard',
    component: DashboardComponent,
    title: 'Dashboard',
  },
  {
    path: 'profile',
    component: ProfileComponent,
    title: 'Profile',
  },
  {
    path: 'login',
    component: LoginComponent,
    title: 'login',
  },
  {
    path: 'about',
    component: AboutComponent,
    title: 'about',
  },
  {
    path: 'contact',
    component: ContactComponent,
    title: 'contact',
  },
];

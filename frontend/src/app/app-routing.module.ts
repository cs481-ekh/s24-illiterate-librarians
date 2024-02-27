import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { ProfileComponent } from './profile/profile.component';
import { LoginComponent } from './login/login.component';
<<<<<<< HEAD:frontend/src/app/app-routing.module.ts
import { SessionComponent } from './session/session.component';
=======
import { AboutComponent } from './about/about.component';
import { ContactComponent } from './contact/contact.component';
>>>>>>> main:frontend/src/app/app.routes.ts

const routes: Routes = [
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
<<<<<<< HEAD:frontend/src/app/app-routing.module.ts
    path: 'session',
    component: SessionComponent,
    title: 'session',
  }
=======
    path: 'about',
    component: AboutComponent,
    title: 'about',
  },
  {
    path: 'contact',
    component: ContactComponent,
    title: 'contact',
  },
>>>>>>> main:frontend/src/app/app.routes.ts
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

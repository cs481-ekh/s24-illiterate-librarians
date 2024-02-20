import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app.component';
import { provideAnimations } from '@angular/platform-browser/animations';
import { adapterFactory } from 'angular-calendar/date-adapters/date-fns';
import { CalendarModule, DateAdapter } from 'angular-calendar';
import { importProvidersFrom } from '@angular/core';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations'

bootstrapApplication(AppComponent, {
  providers: [provideAnimations(),
    appConfig.providers,
  importProvidersFrom(CalendarModule.forRoot({
    provide: DateAdapter,
    useFactory: adapterFactory,
  }))],
}).catch((err) => console.error(err));
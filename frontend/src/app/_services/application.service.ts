import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { HttpHeaders } from '@angular/common/http';
import { Application } from '../application/application';
import { catchError, throwError } from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  }),
  withCredentials: true,
};

@Injectable({
  providedIn: 'root',
})
export class ApplicationService {
  constructor(private http: HttpClient) {}

  createApplication(application: Application) {
    return this.http
      .post(
        environment.BASE_URL + '/api/application',
        application,
        httpOptions
      )
      .pipe(
        catchError((error: any) => {
          return throwError(error);
        })
      );
  }
}

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable, catchError, throwError } from 'rxjs';
import { NewAccount } from '../../new-account';



const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  withCredentials: true,
};

const url = "http://localhost:8080"

@Injectable({
  providedIn: 'root'
})
export class RegistrationService {
 constructor(private http: HttpClient) { }

  register(user: NewAccount): Observable<any> {
    return this.http
      .post(
        url + "/api/user/create",
        user,
        httpOptions
      ).pipe(
        catchError((error: any) => {
          return throwError(error);
        })
      );
  }
}

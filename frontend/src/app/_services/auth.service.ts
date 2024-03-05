import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { HttpHeaders } from '@angular/common/http';
import { Observable, catchError, throwError, tap } from 'rxjs';
import { JwtHelperService } from '@auth0/angular-jwt';
import moment from 'moment';

interface AuthResult {
  expires_at: number;
  id_token: string;
}

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  withCredentials: true,
};

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  constructor(private http: HttpClient, private jwtHelper: JwtHelperService) {}

  login(username: string, password: string): Observable<any> {
    return this.http
      .post(
        'http://localhost:8080/user/login',
        { username, password },
        httpOptions
      )
      .pipe(
        tap((res: any) => {
          this.setSession(res as AuthResult);
        }),
        catchError((error: any) => {
          return throwError(error);
        })
      );
  }

  private setSession(authResult: AuthResult) {
    const expiresAt = moment().add(authResult.expires_at, 'second');

    localStorage.setItem('id_token', authResult.id_token);
    localStorage.setItem('expires_at', JSON.stringify(expiresAt.valueOf()));
  }

  logout() {
    localStorage.removeItem('id_token');
    localStorage.removeItem('expires_at');
  }

  public isLoggedIn() {
    return moment().isBefore(this.getExpiration());
  }

  isLoggedOut() {
    return !this.isLoggedIn();
  }

  getExpiration() {
    const expiration = localStorage.getItem('expires_at');
    if (expiration) {
      const expiresAt = JSON.parse(expiration);
      return moment(expiresAt);
    }
    return null;
  }

  public isAuthenticated(): boolean {
    const token = localStorage.getItem('id_token');
    return !this.jwtHelper.isTokenExpired(token);
  }
}

import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable, catchError, throwError } from 'rxjs';
import { ProfileData } from '../../profile/parent/update-account';



const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
  withCredentials: true,
};

const url = "http://localhost:8080"

@Injectable({
  providedIn: 'root'
})
export class ProfileService {
    constructor(private http: HttpClient) { }
    
    updateProfile(user: ProfileData): Observable<any> {
        return this.http
            .put(
                url + "/api/user/UpdateUserHandler",
                user,
                httpOptions
            ).pipe(
                catchError((error: any) => {
                    return throwError(error);
                })
            );
    }

    fetchProfile(): Observable<any> {
        return this.http
            .get(
                url + "/api/user/lookup",
                httpOptions
            ).pipe(
                catchError((error: any) => {
                    return throwError(error);
                })
            );
    }
}

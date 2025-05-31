import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { User } from '../models/user';
import { catchError, firstValueFrom, map, Observable, of } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(private http: HttpClient) { }

  login(user: User): Observable<Object> {
    return this.http.post(environment.api + 'login', user, {withCredentials: true}).pipe()
  }

  verifyToken(): Promise<boolean> {
    return firstValueFrom(
      this.http.get<{ valid: boolean }>(
        environment.api + 'validate-token',
        { withCredentials: true }
      ).pipe(
        map(response => response.valid),
        catchError(() => of(false))
      )
    );
  }
}

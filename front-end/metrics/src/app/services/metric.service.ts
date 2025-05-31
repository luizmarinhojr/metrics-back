import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { MetricResponse } from '../models/metric';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class MetricService {

  constructor(private http: HttpClient) { }

  GetAll(): Observable<MetricResponse[]> {
    return this.http.get<MetricResponse[]>(environment.api + 'metrics', {withCredentials: true})
      .pipe()
  }
}

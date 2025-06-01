import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { MetricRequest, MetricResponse } from '../models/metric';
import { map, Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class MetricService {

  constructor(private http: HttpClient) { }

  getAll(): Observable<MetricResponse[]> {
    return this.http.get<MetricResponse[]>(environment.api + 'metrics', {withCredentials: true})
  }

  register(metric: MetricRequest) {
    return this.http.post<{}>(environment.api + 'metric', metric, {withCredentials: true})
  }
}

import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class BrokerService {

  constructor() { }

  getBrokerName(): string | null {
    return localStorage.getItem("broker")
  }
}

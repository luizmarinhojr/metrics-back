import { Component, OnInit } from '@angular/core';
import { HeaderHomeComponent } from '../../components/header-home/header-home.component';
import { TableComponent } from "../../components/table/table.component";
import { MetricResponse } from '../../models/metric';

@Component({
  selector: 'app-home',
  imports: [HeaderHomeComponent, TableComponent],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
  metricsResponse!: MetricResponse[];

  constructor() {}

  ngOnInit(): void {
      this.getMetrics();
  }

  getMetrics() {
    this.metricsResponse = [
      
    ]
  }
}

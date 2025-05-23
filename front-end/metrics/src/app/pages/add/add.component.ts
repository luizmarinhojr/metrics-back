import { Component, OnInit } from '@angular/core';
import { MetricRequest } from '../../models/metric';

@Component({
  selector: 'app-add',
  imports: [],
  templateUrl: './add.component.html',
  styleUrl: './add.component.css'
})
export class AddComponent implements OnInit{
  date!: string;
  metrics!: MetricRequest[];

  constructor() {}

  ngOnInit(): void {
    this.setCurrentDate();
  }

  setCurrentDate(): void {
    const hoje = new Date();
    this.date = hoje.toISOString().substring(0, 10);
  }
}

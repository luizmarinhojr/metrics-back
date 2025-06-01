import { Component, Input } from '@angular/core';
import { MetricResponse } from '../../models/metric';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-table',
  imports: [DatePipe],
  templateUrl: './table.component.html',
  styleUrl: './table.component.css'
})
export class TableComponent {
  @Input() metrics!: MetricResponse[];
}

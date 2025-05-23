import { Component, Input } from '@angular/core';
import { MetricResponse } from '../../models/metric';

@Component({
  selector: 'app-table',
  imports: [],
  templateUrl: './table.component.html',
  styleUrl: './table.component.css'
})
export class TableComponent {
  @Input() metrics!: MetricResponse[];
}

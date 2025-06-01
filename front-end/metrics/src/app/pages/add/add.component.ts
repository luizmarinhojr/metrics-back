import { Component, OnInit } from '@angular/core';
import { MetricRequest } from '../../models/metric';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MetricService } from '../../services/metric.service';
import { Router } from '@angular/router';
import { BrokerService } from '../../services/broker.service';

@Component({
  selector: 'app-add',
  imports: [ReactiveFormsModule],
  templateUrl: './add.component.html',
  styleUrl: './add.component.css'
})
export class AddComponent implements OnInit{
  date!: string;
  metrics!: MetricRequest[];
  form!: FormGroup

  constructor(private fb: FormBuilder, 
    private metricService: MetricService, 
    private router: Router,
    private brokerService: BrokerService
  ) {}

  ngOnInit(): void {
    this.form = this.fb.group({
      date: [this.setCurrentDate(), [Validators.required, Validators.maxLength(10)]],
      broker: [{value: this.brokerService.getBrokerName(), disabled: true}, [Validators.required, Validators.maxLength(30)]],
      receivedLeads: [null, [Validators.required, Validators.max(1500)]],
      calls: [null, [Validators.required, Validators.max(1500)]],
      spontaneous: [null, [Validators.required, Validators.max(1500)]],
      captations: [null, [Validators.required, Validators.max(1500)]],
      visits: [null, [Validators.required, Validators.max(1500)]],
      negotiations: [null, [Validators.required, Validators.max(1500)]],
      proposals: [null, [Validators.required, Validators.max(1500)]],
      sells: [null, [Validators.required, Validators.max(1500)]],
      schedules: [null, [Validators.required, Validators.maxLength(300)]]
    })
  }

  send(): void {
    console.log(this.convertMetricRequest())
    this.metricService.register(this.convertMetricRequest()).subscribe({
      next: (value) => {
        console.log('Metrics registered!');
        this.router.navigate(['/']);
      },
      error: (err) => {
        alert('Error to register Metrics')
        console.error('Error to register metrics:', err);
      }
    })
  }

  convertMetricRequest(): MetricRequest {
    return {
      data: this.form.get('date')?.value,
      leads_recebidos: this.form.get('receivedLeads')?.value,
      ligacoes: this.form.get('calls')?.value,
      espontaneo: this.form.get('spontaneous')?.value,
      captacoes: this.form.get('captations')?.value,
      visitas: this.form.get('visits')?.value,
      negociacoes: this.form.get('negotiations')?.value,
      propostas: this.form.get('proposals')?.value,
      vendas: this.form.get('sells')?.value,
      agendamentos: this.form.get('schedules')?.value
    }
  }

  setCurrentDate(): string {
    const hoje = new Date();
    return hoje.toISOString().substring(0, 10);
  }
}

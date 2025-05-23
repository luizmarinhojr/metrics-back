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
      {
        id: 1,
        data: "string",
        nome: "string",
        leadsRecebidos: 3,
        ligacoes: 5,
        espontaneo: 3,
        captacoes: 2,
        visitas: 1,
        negociacoes: 1,
        propostas: 0,
        vendas: 0,
        observacao: "Nulla interdum, nisl vitae semper hendrerit, felis risus dapibus leo, eget dictum lectus diam sed tellus. Mauris lobortis risus sapien, in"
      },
      {
        id: 2,
        data: "string",
        nome: "string",
        leadsRecebidos: 3,
        ligacoes: 5,
        espontaneo: 3,
        captacoes: 2,
        visitas: 1,
        negociacoes: 1,
        propostas: 0,
        vendas: 0,
        observacao: "Nulla interdum, nisl vitae semper hendrerit, felis risus dapibus leo, eget dictum lectus diam sed tellus. Mauris lobortis risus sapien, in"
      }
    ]
  }
}

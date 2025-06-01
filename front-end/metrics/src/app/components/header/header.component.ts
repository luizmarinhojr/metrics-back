import { Component, OnInit } from '@angular/core';
import { BrokerService } from '../../services/broker.service';
import { LoginService } from '../../services/login.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-header',
  imports: [],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent implements OnInit {
  nome!: string | null

  constructor(private brokerService: BrokerService, private loginService: LoginService, private router: Router) {}

  ngOnInit(): void {
    this.nome = this.brokerService.getBrokerName()
  }

  logout():void {
    console.log("cliquei no logout")
    this.loginService.logout().subscribe({
      next: (value) => {
        localStorage.removeItem('broker');
        this.router.navigate(['login']);
      },
      error: (err) => {
        console.error('Erro no logout:', err);
      }
    })
  }

}

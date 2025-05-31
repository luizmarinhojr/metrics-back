import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { LoginService } from '../../services/login.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signin',
  imports: [ReactiveFormsModule],
  templateUrl: './signin.component.html',
  styleUrl: './signin.component.css'
})
export class SigninComponent implements OnInit {

  form!: FormGroup
  response!: boolean

  constructor(private service: LoginService, private fb: FormBuilder, private router: Router) {
    this.form = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(8), Validators.maxLength(25)]]
    })
  }

  ngOnInit(): void {
    this.verifyToken()
  }

  login() {
    this.response = false
    this.service.login({
      email: this.form.get('email')?.value,
      password: this.form.get('password')?.value
    }).subscribe({
        next: (value) => {
          this.response = true;
          console.log('Login bem-sucedido!');
          this.router.navigate(['/inicio']);
        },
        error: (err) => {
          console.error('Erro no login:', err);
        }
    })
  }

  async verifyToken() {
    if (await this.service.verifyToken()) {
      this.router.navigate(['/inicio']);
    }
  }
}

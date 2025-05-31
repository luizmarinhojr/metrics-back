import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { AddComponent } from './pages/add/add.component';
import { SigninComponent } from './pages/signin/signin.component';
import { authGuard } from './auth/auth.guard';

export const routes: Routes = [
    { path: '', redirectTo: '/inicio', pathMatch: 'full' },
    { path: 'inicio', title: 'Métricas - Início', component: HomeComponent, canActivate: [authGuard] },
    { path: 'adicionar', title: 'Métricas - Adicionar', component: AddComponent, canActivate: [authGuard] },
    { path: 'login', title: 'Métricas - Login', component: SigninComponent }
];

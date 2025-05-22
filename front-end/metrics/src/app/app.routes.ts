import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { AddComponent } from './pages/add/add.component';

export const routes: Routes = [
    { path: '', redirectTo: '/inicio', pathMatch: 'full' },
    { path: 'inicio', title: 'Início', component: HomeComponent },
    { path: 'adicionar', title: 'Adicionar', component: AddComponent }
];

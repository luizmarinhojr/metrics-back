import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { LoginService } from '../services/login.service';

export const authGuard: CanActivateFn = async (route, state) => {
    const loginService = inject(LoginService);
    const router = inject(Router);

    let isLogged = await loginService.verifyToken();

    if (!isLogged) {
      router.navigate(['/login'])
    }

  return isLogged;
};

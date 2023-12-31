import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { DocsComponent } from './components/docs/docs.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { LandingComponent } from './components/landing/landing.component';
import { LogoutComponent } from './components/logout/logout.component';
import { ServerErrorComponent } from './components/error-pages/server-error/server-error.component';
import { NotFoundComponent } from './components/error-pages/not-found/not-found.component';
import { IsAuthenticatedGuard } from './guards';
import { UserProfileComponent } from './components/dashboard/user-profile/user-profile.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { ChatsComponent } from './components/dashboard/chats/chats.component';

const routes: Routes = [
  { path: 'docs', component: DocsComponent },
  { path: '', redirectTo: '/landing', pathMatch: 'full' },
  { path: 'log-in', component: LoginComponent },
  { path: 'sign-up', component: RegisterComponent },
  { path: 'landing', component: LandingComponent },
  { path: 'sign-out', component: LogoutComponent },
  { path: '500', component: ServerErrorComponent },
  { path: '404', component: NotFoundComponent },
  {
    path: 'dashboard', component: DashboardComponent, children: [
      {
        path: 'user-profile', component: UserProfileComponent, canActivate: [IsAuthenticatedGuard]
      },
      {
        path: 'chats', component: ChatsComponent, canActivate: [IsAuthenticatedGuard]
      },

      {
        path: '', redirectTo: 'user-profile', pathMatch: 'full'
      },
    ]
  },
  { path: '**', redirectTo: '/404', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true, enableTracing: false })],
  exports: [RouterModule]
})
export class AppRoutingModule { }

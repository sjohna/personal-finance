import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AccountsComponent } from './components/accounts/accounts.component';
import { CurrenciesComponent } from './components/currencies/currencies.component';

const routes: Routes = [
  {
    path: 'accounts',
    component: AccountsComponent
  },
  {
    path: 'currencies',
    component: CurrenciesComponent
  },
  {
    path: '**',
    pathMatch: 'full',
    redirectTo: 'accounts'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

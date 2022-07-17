import { Component, OnInit } from '@angular/core';
import { firstValueFrom, Observable } from 'rxjs';
import { Account, AccountsService } from 'src/app/services/accounts.service';


@Component({
  selector: 'app-accounts',
  templateUrl: './accounts.component.html'
})
export class AccountsComponent {
  public accounts$: Observable<Account[]>;

  constructor(private accountsService: AccountsService) {
    this.accounts$ = this.accountsService.accounts$;
    this.accountsService.loadAccounts();
  }
}

import { Component, OnInit } from '@angular/core';
import { firstValueFrom, Observable } from 'rxjs';
import { Account, AccountsService } from 'src/app/services/accounts.service';
import {MatDialog} from "@angular/material/dialog";
import {AddEditAccountModalComponent} from "./add-edit-account-modal/add-edit-account-modal.component";


@Component({
  selector: 'app-accounts',
  templateUrl: './accounts.component.html'
})
export class AccountsComponent {
  public accounts$: Observable<Account[]>;

  constructor(private accountsService: AccountsService, public dialog: MatDialog) {
    this.accounts$ = this.accountsService.accounts$;
    this.accountsService.loadAccounts();
  }

  public openAddEditDialog() {
    this.dialog.open(AddEditAccountModalComponent);
  }
}

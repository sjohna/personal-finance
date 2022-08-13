import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { CurrenciesService, Currency } from 'src/app/services/currencies.service';
import {MatDialog} from "@angular/material/dialog";
import {AddEditCurrencyModalComponent} from "./add-edit-currency-modal/add-edit-currency-modal.component";

@Component({
  selector: 'app-currencies',
  templateUrl: './currencies.component.html'
})
export class CurrenciesComponent {
  public currencies$: Observable<Currency[]>;

  constructor(private currenciesService: CurrenciesService, private dialog: MatDialog) {
    this.currencies$ = this.currenciesService.currencies$;
    this.currenciesService.loadCurrencies();
  }

  public openAddEditDialog() {
    this.dialog.open(AddEditCurrencyModalComponent);
  }
}

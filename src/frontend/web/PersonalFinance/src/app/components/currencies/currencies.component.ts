import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { CurrenciesService, Currency } from 'src/app/services/currencies.service';

@Component({
  selector: 'app-currencies',
  templateUrl: './currencies.component.html'
})
export class CurrenciesComponent implements OnInit {
  public currencies$: Observable<Currency[]>;

  constructor(private currenciesService: CurrenciesService) { 
    this.currencies$ = this.currenciesService.currencies$;
    this.currenciesService.loadCurrencies();
  }

  ngOnInit(): void {
  }

}
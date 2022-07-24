import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, shareReplay, firstValueFrom } from 'rxjs';
import { LoadingStatus } from './types';

export  enum DebitCreditType {
  Debit = 0,
  Credit,
}

export interface DebitCredit {
  type: DebitCreditType;
  id: number;
  amount: number;
  currencyId: number;
  time: string;
  accountId: number;
  createdAt: string;
  updatedAt: string;
}

export interface Account {
  id: number;
  name: string;
  description: string;
  createdAt: string;
  updatedAt: string;
  debitsAndCredits?: DebitCredit[];
}

@Injectable({
  providedIn: 'root'
})
export class AccountsService {
  private loadingStatus$$ = new BehaviorSubject<LoadingStatus>(LoadingStatus.NotLoaded);
  public loadingStatus$ = this.loadingStatus$$.pipe(shareReplay(1));

  private accounts$$ = new BehaviorSubject<Account[]>([]);
  public accounts$ = this.accounts$$.pipe(shareReplay(1));

  constructor(private http: HttpClient) { }

  public async loadAccounts() {
    // TODO: handle/throttle redundant requests?
    this.loadingStatus$$.next(LoadingStatus.Loading);

    try {
      const accounts = await firstValueFrom(this.http.get<Account[]>('http://localhost:3000/account'))
      this.accounts$$.next(accounts);
      this.loadingStatus$$.next(LoadingStatus.Loaded);
    } catch {
      this.loadingStatus$$.next(LoadingStatus.Error);
    }
  }


}

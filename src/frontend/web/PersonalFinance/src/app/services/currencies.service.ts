import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, firstValueFrom, shareReplay } from 'rxjs';
import { LoadingStatus } from './types';

export interface Currency {
  id: number;
  name: string;
  abbreviation: string;
  magnitude: number;
  createdAt: string;
  updatedAt: string;
}

@Injectable({
  providedIn: 'root'
})
export class CurrenciesService {
  private currencies$$ = new BehaviorSubject<Currency[]>([]);
  public currencies$ = this.currencies$$.pipe(shareReplay(1));

  private loadingStatus$$ = new BehaviorSubject<LoadingStatus>(LoadingStatus.NotLoaded);
  public loadingStatus$ = this.loadingStatus$$.pipe(shareReplay(1));

  constructor(private http: HttpClient) { }

  public async loadCurrencies() {
    this.loadingStatus$$.next(LoadingStatus.Loading);

    try {
      const currencies = await firstValueFrom(this.http.get<Currency[]>('http://localhost:3000/currency'));
      this.currencies$$.next(currencies);
      this.loadingStatus$$.next(LoadingStatus.Loaded);
    } catch {
      this.loadingStatus$$.next(LoadingStatus.Error);
    }
  }
}

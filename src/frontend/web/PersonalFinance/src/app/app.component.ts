import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { firstValueFrom } from 'rxjs';

enum DebitCreditType {
  Debit = 0,
  Credit,
}

interface Currency {
  id: number;
  name: string;
  abbreviation: string;
  magnitude: number;
  createdAt: string;
  updatedAt: string;
}

interface DebitCredit {
  type: DebitCreditType;
  id: number;
  amount: number;
  currencyId: number;
  time: string;
  accountId: number;
  createdAt: string;
  updatedAt: string;
}

interface Account {
  id: number;
  name: string;
  description: string;
  createdAt: string;
  updatedAt: string;
  debitsAndCredits?: DebitCredit[];
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'PersonalFinance';

  accounts: Account[] = [];

  constructor(private http: HttpClient) {
    this.getAccounts()
  }

  async getAccounts(): Promise<void> {
    this.accounts = await firstValueFrom(this.http.get<Account[]>('http://localhost:3000/account'))
  }
}

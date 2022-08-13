import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AccountsComponent } from './components/accounts/accounts.component';
import { CurrenciesComponent } from './components/currencies/currencies.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatTabsModule } from '@angular/material/tabs';
import { MatButtonModule } from '@angular/material/button';
import { FormatDatePipe } from './pipes/format-date.pipe';
import { AddEditAccountModalComponent } from './components/accounts/add-edit-account-modal/add-edit-account-modal.component';
import { MatDialogModule } from "@angular/material/dialog";
import {ReactiveFormsModule} from "@angular/forms";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import { AddEditCurrencyModalComponent } from './components/currencies/add-edit-currency-modal/add-edit-currency-modal.component';

@NgModule({
  declarations: [
    AppComponent,
    AccountsComponent,
    CurrenciesComponent,
    FormatDatePipe,
    AddEditAccountModalComponent,
    AddEditCurrencyModalComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatTabsModule,
    MatButtonModule,
    MatDialogModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }

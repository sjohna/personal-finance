import {Component, Inject, OnInit} from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {AccountsService} from "../../../services/accounts.service";
import {CurrenciesService} from "../../../services/currencies.service";

interface CurrencyForm {
  name: FormControl<string>;
  abbreviation: FormControl<string>;
  magnitude: FormControl<number>;
}

@Component({
  selector: 'app-add-edit-currency-modal',
  templateUrl: './add-edit-currency-modal.component.html',
})
export class AddEditCurrencyModalComponent {
  public currencyForm: FormGroup<CurrencyForm>;

  constructor(private dialogRef: MatDialogRef<AddEditCurrencyModalComponent>,
              @Inject(MAT_DIALOG_DATA) private data: any,
              private fb: FormBuilder,
              private currencies: CurrenciesService) {
    this.currencyForm = this.fb.group<CurrencyForm>(
      {
        name: this.fb.nonNullable.control('', Validators.required),
        abbreviation: this.fb.nonNullable.control('', Validators.required),
        magnitude: this.fb.nonNullable.control(2, Validators.required),
      },
    );
  }

  public async createCurrency() {
    // TODO: avoid duplicate submissions
    if (!this.currencyForm.valid) {
      return;
    }

    const formValue = this.currencyForm.value;
    await this.currencies.createCurrency(formValue.name!, formValue.abbreviation!, formValue.magnitude!);
    this.dialogRef.close();
  }
}

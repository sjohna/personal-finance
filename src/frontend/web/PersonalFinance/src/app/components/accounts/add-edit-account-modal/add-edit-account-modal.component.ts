import {Component, Inject, OnInit} from '@angular/core';
import {MAT_DIALOG_DATA, MatDialogRef} from "@angular/material/dialog";
import {
  FormBuilder,
  FormControl,
  FormGroup,
  Validators
} from "@angular/forms";
import {AccountsService} from "../../../services/accounts.service";

interface AccountForm {
  name: FormControl<string>;
  description: FormControl<string>;
}

@Component({
  selector: 'app-add-edit-account-modal',
  templateUrl: './add-edit-account-modal.component.html'
})
export class AddEditAccountModalComponent {
  public accountForm: FormGroup<AccountForm>;

  constructor(private dialogRef: MatDialogRef<AddEditAccountModalComponent>,
              @Inject(MAT_DIALOG_DATA) private data: any,
              private fb: FormBuilder,
              private accounts: AccountsService) {
    this.accountForm = this.fb.group<AccountForm>(
      {
        name: this.fb.nonNullable.control('', Validators.required),
        description: this.fb.nonNullable.control('', Validators.required),
      },
    );
  }

  public async createAccount() {
    // TODO: avoid duplicate submissions
    if (!this.accountForm.valid) {
      return;
    }

    const formValue = this.accountForm.value;
    await this.accounts.createAccount(formValue.name!, formValue.description!);
    this.dialogRef.close();
  }
}

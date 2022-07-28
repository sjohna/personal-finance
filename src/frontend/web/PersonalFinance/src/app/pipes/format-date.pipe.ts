import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'formatDate'
})
export class FormatDatePipe implements PipeTransform {

  transform(date: Date | string): string {
    if (date && typeof date === 'object') {
      return date.toLocaleString();
    } else {
      return new Date(date).toLocaleString();
    }
  }

}

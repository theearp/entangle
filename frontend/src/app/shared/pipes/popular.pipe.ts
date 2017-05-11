import { Pipe, PipeTransform } from '@angular/core';

import { Listing } from '../models/listing';
/*
 * 
*/
@Pipe({name: 'popularListings'})
export class PopularListingsPipe implements PipeTransform {
  transform(l: Array<Listing>): Listing[] {
    l.sort(this.popSort)
    return l.slice(1, 10);
  }

  popSort(a: Listing, b: Listing) {
    if (a.Views < b.Views) {
      return -1
    }
    if (a.Views > b.Views) {
      return 1
    }
    return 0
  }
}
import { Component, OnInit } from '@angular/core';

import { ProductService } from '../product.service';
import { Listing } from '../product';

@Component({
  selector: 'listing',
  templateUrl: './listing.component.html',
  styleUrls: ['./listing.component.css'],
  providers: [ProductService]
})
export class ListingComponent implements OnInit {
  listings: Listing[];
  errorMessage: any;
  constructor(private ps: ProductService) { }

  ngOnInit() {
    this.ps.getPopularListings()
    .subscribe(
      data => this.listings = data,
      error => {
        console.log(error);
        this.errorMessage = <any>error;
      });
  }
}

import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params }   from '@angular/router';
import { Location }                 from '@angular/common';

import { Listing } from '../shared/models/listing';
import { ListingService } from '../shared/services/listing.service';
import { ShoppingService } from '../shared/services/shopping.service';

import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'listing-category',
  templateUrl: './listing.component.html',
  styleUrls: ['./listing.component.css'],
  providers: [ListingService]
})

export class ListingCategoryComponent implements OnInit {
  listings: Listing[];
  selectedListing: Listing;
  constructor(
    private ps: ListingService, 
    private ss: ShoppingService,
    private route: ActivatedRoute
    ) {
  }
  
  ngOnInit() {
  }

  updateShoppingCart(listing: Listing) {
    this.ss.addItemToCart(listing);
    console.log(listing);
  }

  showDetails(listing: Listing) {
    this.selectedListing = listing;
  }
}
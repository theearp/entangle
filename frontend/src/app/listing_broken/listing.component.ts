import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import { Listing } from '../shared/models/listing';
import { ListingService } from '../shared/services/listing.service';
import { ShoppingService } from '../shared/services/shopping.service';

@Component({
  selector: 'listing',
  templateUrl: './listing.component.html',
  styleUrls: ['./listing.component.css'],
  providers: [ListingService]
})
export class ListingComponent implements OnInit {
  popular: boolean;
  listings: Listing[];
  constructor(
    private ls: ListingService, 
    private ss: ShoppingService,
    private r: ActivatedRoute
    ) {
  }
  
  ngOnInit() {
    this.r.data.subscribe(v => this.popular = v.popular);
    this.ls.getListings()
    .subscribe(data => {
      if (this.popular) {
        this.listings = data.slice(1, 10);
      } else {
        this.listings = data;
      }
    });
  }

  updateShoppingCart(listing: Listing) {
    this.ss.addItemToCart(listing);
  }

  syncListing(id: string) {
    this.ls.syncListing(id);
  }
}

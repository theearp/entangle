import { Component, OnInit } from '@angular/core';

import { Listing } from '../shared/models/listing';
import { ListingService } from '../shared/services/listing.service';
import { ShoppingService } from '../shared/services/shopping.service';

@Component({
  selector: 'app-listing',
  templateUrl: './listing.component.html',
  styleUrls: ['./listing.component.css']
})
export class ListingComponent implements OnInit {
  listings: Listing[];
  syncStatus: String;
  constructor(
    private listingService: ListingService,
    private shoppingService: ShoppingService) { }

  ngOnInit() {
    this.listingService.getListings()
    .subscribe(data => this.listings = data);
  }

  syncListing(id: string) {
    this.listingService.syncListing(id)
    .subscribe(data => this.syncStatus = data);
  }

  addToCart(l: Listing) {
    this.shoppingService.add(l);
  }
}

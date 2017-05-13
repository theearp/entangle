import { Component, Input, OnInit } from '@angular/core';
import { ActivatedRoute, Params }   from '@angular/router';
import { Location }                 from '@angular/common';
import { ListingService } from '../shared/services/listing.service';
import { ShoppingService } from '../shared/services/shopping.service';
import { Listing } from '../shared/models/listing';

import 'rxjs/add/operator/switchMap';

@Component({
  selector: 'listing-detail',
  templateUrl: './listingDetail.component.html',
  styleUrls: ['./listingDetail.component.css'],
  providers: [ListingService]
})
export class ListingDetailComponent implements OnInit {
  listing: Listing;

  constructor(
    private ls: ListingService,
    private route: ActivatedRoute,
    private location: Location,
    private ss: ShoppingService
  ) {}

  ngOnInit() {
    this.route.params
      .switchMap((params: Params) => this.ls.getListing(params['id']))
      .subscribe(data => this.listing = data);
  }

  updateShoppingCart(listing) {
    this.ss.addItemToCart(listing);
  }

  syncListing(listing) {
    this.ls.syncListing(listing);
  }
}
